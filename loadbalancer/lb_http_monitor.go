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

type LbHttpMonitor struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	// Schema for this resource
	Schema string `json:"_schema,omitempty"`

	// Link to this resource
	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int32 `json:"_revision"`

	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`

	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`

	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`

	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`

	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity.
	Protection string `json:"_protection,omitempty"`

	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`

	// Description of this resource
	Description string `json:"description,omitempty"`

	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`

	// Unique identifier of this resource
	Id string `json:"id,omitempty"`

	// Load balancers monitor the health of backend servers to ensure traffic is not black holed. There are two types of healthchecks: active and passive. Passive healthchecks depend on failures in actual client traffic (e.g. RST from server in response to a client connection) to detect that the server or the application is down. In case of active healthchecks, load balancer itself initiates new connections (or sends ICMP ping) to the servers periodically to check their health, completely independent of any data traffic. Currently, active health monitors are supported for HTTP, HTTPS, TCP, UDP and ICMP protocols.
	ResourceType string `json:"resource_type"`

	// Opaque identifiers meaningful to the API user
	Tags []common.Tag `json:"tags,omitempty"`

	// num of consecutive checks must fail before marking it down
	FallCount int64 `json:"fall_count,omitempty"`

	// the frequency at which the system issues the monitor check (in second)
	Interval int64 `json:"interval,omitempty"`

	// If the monitor port is specified, it would override pool member port setting for healthcheck. A port range is not supported.
	MonitorPort string `json:"monitor_port,omitempty"`

	// num of consecutive checks must pass before marking it up
	RiseCount int64 `json:"rise_count,omitempty"`

	// the number of seconds the target has in which to respond to the monitor request
	Timeout int64 `json:"timeout,omitempty"`

	// String to send as part of HTTP health check request body. Valid only for certain HTTP methods like POST.
	RequestBody string `json:"request_body,omitempty"`

	// Array of HTTP request headers
	RequestHeaders []LbHttpRequestHeader `json:"request_headers,omitempty"`

	// the health check method for HTTP monitor type
	RequestMethod string `json:"request_method,omitempty"`

	// URL used for HTTP monitor
	RequestUrl string `json:"request_url,omitempty"`

	// HTTP request version
	RequestVersion string `json:"request_version,omitempty"`

	// If HTTP response body match string (regular expressions not supported) is specified (using LbHttpMonitor.response_body) then the healthcheck HTTP response body is matched against the specified string and server is considered healthy only if there is a match. If the response body string is not specified, HTTP healthcheck is considered successful if the HTTP response status code is 2xx, but it can be configured to accept other status codes as successful.
	ResponseBody string `json:"response_body,omitempty"`

	// The HTTP response status code should be a valid HTTP status code.
	ResponseStatusCodes []int32 `json:"response_status_codes,omitempty"`
}
