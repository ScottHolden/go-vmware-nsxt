/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package apiservice

import (
	"github.com/ScottHolden/go-vmware-nsxt/common"
)

type AuthenticationPolicyProperties struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// Once a lockout occurs, the account remains locked out of the API for this time period. Only applies to NSX Manager nodes. Ignored on other node types.
	ApiFailedAuthLockoutPeriod int64 `json:"api_failed_auth_lockout_period,omitempty"`

	// In order to trigger an account lockout, all authentication failures must occur in this time window. If the reset period expires, the failed login count is reset to zero. Only applies to NSX Manager nodes. Ignored on other node types.
	ApiFailedAuthResetPeriod int64 `json:"api_failed_auth_reset_period,omitempty"`

	// Only applies to NSX Manager nodes. Ignored on other node types.
	ApiMaxAuthFailures int64 `json:"api_max_auth_failures,omitempty"`

	// Once a lockout occurs, the account remains locked out of the CLI for this time period.
	CliFailedAuthLockoutPeriod int64 `json:"cli_failed_auth_lockout_period,omitempty"`

	// Number of authentication failures that trigger CLI lockout
	CliMaxAuthFailures int64 `json:"cli_max_auth_failures,omitempty"`

	// Minimum number of characters required in account passwords
	MinimumPasswordLength int64 `json:"minimum_password_length,omitempty"`
}
