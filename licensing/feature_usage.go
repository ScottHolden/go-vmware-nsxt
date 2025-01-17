/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package licensing

import (
	"github.com/ScottHolden/go-vmware-nsxt/common"
)

type FeatureUsage struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// Capacity Usage List
	CapacityUsage []CapacityUsage `json:"capacity_usage,omitempty"`

	// name of the feature
	Feature string `json:"feature,omitempty"`
}
