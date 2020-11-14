package generate

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
	"fmt"
	insecureRand "math/rand"
	"strings"
	"time"

	"github.com/bishopfox/sliver/server/db"
	"github.com/bishopfox/sliver/server/db/models"
)

const (
	// CanaryBucketName - DNS Canary bucket name
	CanaryBucketName = "canaries"
	canaryPrefix     = "can://"
	canarySize       = 6
)

var (
	dnsCharSet = []rune("abcdefghijklmnopqrstuvwxyz0123456789-_")
)

func canarySubDomain() string {
	subdomain := []rune{}
	index := insecureRand.Intn(len(dnsCharSet) - 12) // ensure first char is alphabetic
	subdomain = append(subdomain, dnsCharSet[index])
	for i := 0; i < canarySize; i++ {
		index := insecureRand.Intn(len(dnsCharSet))
		subdomain = append(subdomain, dnsCharSet[index])
	}
	return string(subdomain)
}

// ListCanaries - List of all embedded canaries
func ListCanaries() ([]*models.DNSCanary, error) {
	// TODO
	return []*models.DNSCanary{}, nil
}

// CheckCanary - Check if a canary exists
func CheckCanary(domain string) (*models.DNSCanary, error) {
	// TODO
	return nil, nil
}

// UpdateCanary - Update an existing canary
func UpdateCanary(canary *models.DNSCanary) error {
	// TODO
	return nil
}

// CanaryGenerator - Holds data related to canary generation
type CanaryGenerator struct {
	ImplantName   string
	ParentDomains []string
}

// GenerateCanary - Generate a canary domain and save it to the db
// 				    currently this gets called by template engine
func (g *CanaryGenerator) GenerateCanary() string {

	if len(g.ParentDomains) < 1 {
		buildLog.Warnf("No parent domains")
		return ""
	}

	// Don't need secure random here
	insecureRand.Seed(time.Now().UnixNano())
	index := insecureRand.Intn(len(g.ParentDomains))

	parentDomain := g.ParentDomains[index]
	if strings.HasPrefix(parentDomain, ".") {
		parentDomain = parentDomain[1:]
	}
	if !strings.HasSuffix(parentDomain, ".") {
		parentDomain += "." // Ensure we have the FQDN
	}

	subdomain := canarySubDomain()
	canaryDomain := fmt.Sprintf("%s.%s", subdomain, parentDomain)
	buildLog.Infof("Generated new canary domain %s", canaryDomain)

	canary := &models.DNSCanary{
		ImplantName: g.ImplantName,
		Domain:      canaryDomain,
		Triggered:   false,
		Count:       0,
	}
	dbSession := db.Session()
	dbSession.Create(&canary)

	return fmt.Sprintf("%s%s", canaryPrefix, canaryDomain)
}
