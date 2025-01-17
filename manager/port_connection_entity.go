/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

import (
	"github.com/ScottHolden/go-vmware-nsxt/common"
)

// Port Connection Entity
type PortConnectionEntity struct {

	// Resource ID is mapped to this. (ID is Generated for Edge node groups, since resource will be null)
	Id string `json:"id,omitempty"`

	// Resource reference with details of the entity
	Resource *common.ManagedResource `json:"resource,omitempty"`
}
