/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package monitoring

import (
	"github.com/ScottHolden/go-vmware-nsxt/common"
)

type Traceflow struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// Traceflow result analysis notes
	Analysis []string `json:"analysis,omitempty"`

	// observation counters
	Counters *TraceflowObservationCounters `json:"counters,omitempty"`

	// The id of the traceflow round
	Id string `json:"id,omitempty"`

	// counters of observations from logical components
	LogicalCounters *TraceflowObservationCounters `json:"logical_counters,omitempty"`

	// id of the source logical port used for injecting the traceflow packet
	LportId string `json:"lport_id,omitempty"`

	// Represents the traceflow operation state
	OperationState string `json:"operation_state,omitempty"`

	// A flag, when set true, indicates some observations were deleted from the result set.
	ResultOverflowed bool `json:"result_overflowed,omitempty"`

	// Maximum time (in ms) the management plane will be waiting for this traceflow round.
	Timeout int64 `json:"timeout,omitempty"`
}
