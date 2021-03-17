/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package administration

import (
	"github.com/ScottHolden/go-vmware-nsxt/common"
)

type FileThumbprint struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// File name
	Name string `json:"name"`

	// File's SHA1 thumbprint
	Sha1 string `json:"sha1"`

	// File's SHA256 thumbprint
	Sha256 string `json:"sha256"`
}
