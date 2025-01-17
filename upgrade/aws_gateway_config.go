/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package upgrade

import (
	"github.com/ScottHolden/go-vmware-nsxt/common"
)

type AwsGatewayConfig struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// The ID of the Amazon Machine Image on which this gateway resides
	AmiId string `json:"ami_id,omitempty"`

	// Flag to identify if default quarantine policy is enabled
	DefaultQuarantinePolicyEnabled bool `json:"default_quarantine_policy_enabled,omitempty"`

	// Aws Gateway HA configuration
	GatewayHaConfiguration []AwsGatewayHaConfig `json:"gateway_ha_configuration,omitempty"`

	// Flag to identify if HA is enabled
	IsHaEnabled bool `json:"is_ha_enabled,omitempty"`

	// The key pair name required to authenticate into any instance
	KeyPairName string `json:"key_pair_name,omitempty"`

	// Determines if connection to NSX Manager is via public IP or private IP
	NsxManagerConnection string `json:"nsx_manager_connection,omitempty"`
}
