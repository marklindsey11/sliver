package models

/*
	Sliver Implant Framework
	Copyright (C) 2020  Bishop Fox

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
	"time"

	"github.com/bishopfox/sliver/protobuf/clientpb"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Website - Colletions of content to serve from HTTP(S)
type Website struct {
	gorm.Model

	ID        uuid.UUID `gorm:"primaryKey;->;<-:create;type:uuid;index:unique;"`
	CreatedAt time.Time `gorm:"->;<-:create;"`

	Name string `gorm:"index:unique"` // Website Name

	WebContents []WebContent
}

// BeforeCreate - GORM hook
func (w *Website) BeforeCreate(tx *gorm.DB) (err error) {
	w.ID, err = uuid.NewV4()
	if err != nil {
		return err
	}
	w.CreatedAt = time.Now()
	return nil
}

// ToProtobuf - Converts to protobuf object
func (w *Website) ToProtobuf() *clientpb.Website {
	return &clientpb.Website{
		Name:     w.Name,
		Contents: map[string]*clientpb.WebContent{},
	}
}

// WebContent - One piece of content mapped to a path
type WebContent struct {
	gorm.Model

	ID uuid.UUID `gorm:"primaryKey;->;<-:create;type:uuid;index:unique;"`

	Path        string `gorm:"index:unique"`
	Size        uint64
	ContentType string
	Content     []byte
}

// BeforeCreate - GORM hook to automatically set values
func (wc *WebContent) BeforeCreate(tx *gorm.DB) (err error) {
	wc.ID, err = uuid.NewV4()
	return err
}

// ToProtobuf - Converts to protobuf object
func (wc *WebContent) ToProtobuf() *clientpb.WebContent {
	return &clientpb.WebContent{
		Path:        wc.Path,
		Size:        wc.Size,
		ContentType: wc.ContentType,
		Content:     wc.Content,
	}
}
