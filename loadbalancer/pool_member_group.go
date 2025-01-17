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

type PoolMemberGroup struct {

	// The list is used to show the customized pool member settings. User can only user pool member action API to update the admin state for a specific IP address.
	CustomizedMembers []PoolMemberSetting `json:"customized_members,omitempty"`

	// Load balancer pool support grouping object as dynamic pool members. The IP list of the grouping object such as NSGroup would be used as pool member IP setting.
	GroupingObject *common.ResourceReference `json:"grouping_object"`

	// Ip revision filter is used to filter IPv4 or IPv6 addresses from the grouping object. If the filter is not specified, both IPv4 and IPv6 addresses would be used as server IPs. The link local and loopback addresses would be always filtered out.
	IpRevisionFilter string `json:"ip_revision_filter,omitempty"`

	// The size is used to define the maximum number of grouping object IP address list. These IP addresses would be used as pool members. If the grouping object includes more than certain number of IP addresses, the redundant parts would be ignored and those IP addresses would not be treated as pool members.
	// In order to destinguish between zero and unspecified value, use pointer
	MaxIpListSize *int64 `json:"max_ip_list_size,omitempty"`

	// If port is specified, all connections will be sent to this port. If unset, the same port the client connected to will be used, it could be overridden by default_pool_member_ports setting in virtual server. The port should not specified for multiple ports case.
	Port int32 `json:"port,omitempty"`
}
