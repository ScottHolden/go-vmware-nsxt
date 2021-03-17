/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

import (
	"github.com/ScottHolden/go-vmware-nsxt/common"
)

type VtepListResult struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor string `json:"cursor,omitempty"`

	// Count of results found (across all pages), set only on first page
	ResultCount int64 `json:"result_count,omitempty"`

	SortAscending bool `json:"sort_ascending,omitempty"`

	// Field by which records are sorted
	SortBy string `json:"sort_by,omitempty"`

	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`

	// The id of the logical Switch
	LogicalSwitchId string `json:"logical_switch_id,omitempty"`

	Results []VtepTableEntry `json:"results,omitempty"`

	// Transport node identifier
	TransportNodeId string `json:"transport_node_id,omitempty"`
}
