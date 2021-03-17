/*
 * NSX API
 *
 * VMware NSX REST API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package loadbalancer

import (
	"github.com/ScottHolden/go-vmware-nsxt/common"
)

type LbServerSslProfileListResult struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	// Schema for this resource
	Schema string `json:"_schema,omitempty"`

	// Link to this resource
	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor string `json:"cursor,omitempty"`

	// Count of results found (across all pages), set only on first page
	ResultCount int64 `json:"result_count,omitempty"`

	// If true, results are sorted in ascending order
	SortAscending bool `json:"sort_ascending,omitempty"`

	// Field by which records are sorted
	SortBy string `json:"sort_by,omitempty"`

	// paginated list of load balancer server SSL profiles
	Results []LbServerSslProfile `json:"results"`
}
