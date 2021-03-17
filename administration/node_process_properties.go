/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package administration

import (
	"github.com/ScottHolden/go-vmware-nsxt/common"
)

type NodeProcessProperties struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// CPU time (user and system) consumed by process in milliseconds
	CpuTime int64 `json:"cpu_time,omitempty"`

	// Resident set size of process in bytes
	MemResident int64 `json:"mem_resident,omitempty"`

	// Virtual memory used by process in bytes
	MemUsed int64 `json:"mem_used,omitempty"`

	// Process id
	Pid int64 `json:"pid,omitempty"`

	// Parent process id
	Ppid int64 `json:"ppid,omitempty"`

	// Process name
	ProcessName string `json:"process_name,omitempty"`

	// Process start time expressed in milliseconds since epoch
	StartTime int64 `json:"start_time,omitempty"`

	// Milliseconds since process started
	Uptime int64 `json:"uptime,omitempty"`
}
