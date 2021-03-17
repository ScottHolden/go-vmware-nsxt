/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package upgrade

import (
	"github.com/ScottHolden/go-vmware-nsxt/common"
)

type AwsGatewayInstanceStatus struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// Different states of gateway deployment
	DeploymentState string `json:"deployment_state,omitempty"`

	// Status of gateway instance deployment in percentage
	DeploymentStatus int64 `json:"deployment_status,omitempty"`

	// Error code for gateway deployment/undeployment failure
	ErrorCode int64 `json:"error_code,omitempty"`

	// Error message for gateway deployment/undeployment failure
	ErrorMessage string `json:"error_message,omitempty"`

	// Index of HA that indicates whether gateway is primary or secondary. If index is 0, then it is primary gateway. Else secondary gateway.
	GatewayHaIndex int64 `json:"gateway_ha_index,omitempty"`

	// ID of the gateway instance
	GatewayInstanceId string `json:"gateway_instance_id,omitempty"`

	// Name of the gateway instance
	GatewayName string `json:"gateway_name,omitempty"`

	// NSX Node ID of the public cloud gateway
	GatewayNodeId string `json:"gateway_node_id,omitempty"`

	// Gateway instance status
	GatewayStatus string `json:"gateway_status,omitempty"`

	// NSX transport node id of the public cloud gateway
	GatewayTnId string `json:"gateway_tn_id,omitempty"`

	// Flag to identify if this is an active gateway
	IsGatewayActive bool `json:"is_gateway_active,omitempty"`

	// Private IP address of the virtual machine
	PrivateIp string `json:"private_ip,omitempty"`

	// Public IP address of the virtual machine
	PublicIp string `json:"public_ip,omitempty"`
}
