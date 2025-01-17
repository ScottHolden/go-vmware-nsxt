/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

import (
	"github.com/ScottHolden/go-vmware-nsxt/common"
)

type DneKeyPolicyStats struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// The number of bytes decrypted by the key policy. If key policy action is integrity only, it indicates the number of bytes which passed integrity check.
	BytesDecrypted int64 `json:"bytes_decrypted,omitempty"`

	// The number of bytes dropped for the key policy.
	BytesDropped int64 `json:"bytes_dropped,omitempty"`

	// The number of bytes encrypted by the key policy. If key policy action is integrity only, it indicates the number of bytes with the addition of integrity check.
	BytesEncrypted int64 `json:"bytes_encrypted,omitempty"`

	// Key policy identifier of the DNE key policy. This is a globally unique number.
	KeyPolicyIdentifier string `json:"key_policy_identifier,omitempty"`

	// The number of packets decrypted by the key policy. If key policy action is integrity only, it indicates the number of packets which passed integrity check.
	PacketsDecrypted int64 `json:"packets_decrypted,omitempty"`

	// The number of dropped packets for the key policy.
	PacketsDropped int64 `json:"packets_dropped,omitempty"`

	// The number of packets encrypted by the key policy. If key policy action is integrity only, it indicates the number of packets with the addition of integrity check.
	PacketsEncrypted int64 `json:"packets_encrypted,omitempty"`
}
