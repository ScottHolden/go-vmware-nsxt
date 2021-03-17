/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package appdiscovery

import (
	"github.com/ScottHolden/go-vmware-nsxt/common"
)

type AppDiscoveryVmInfo struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// External Id of the VM
	VmExternalId string `json:"vm_external_id,omitempty"`

	// Name of the VM
	VmName string `json:"vm_name,omitempty"`
}
