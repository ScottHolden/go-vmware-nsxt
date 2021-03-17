/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package administration

import (
	"github.com/ScottHolden/go-vmware-nsxt/manager"
)

type AddManagementNodeSpec struct {

	// The certificate thumbprint of the remote node.
	CertThumbprint string `json:"cert_thumbprint,omitempty"`

	MpaMsgClientInfo *manager.MsgClientInfo `json:"mpa_msg_client_info,omitempty"`

	// The password to be used to authenticate with the remote node.
	Password string `json:"password"`

	// The host address of the remote node to which to send this join request.
	RemoteAddress string `json:"remote_address"`

	// must be set to AddManagementNodeSpec
	Type_ string `json:"type"`

	// The username to be used to authenticate with the remote node.
	UserName string `json:"user_name"`
}
