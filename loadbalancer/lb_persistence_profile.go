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

type LbPersistenceProfile struct {

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

	// Source-ip persistence ensures all connections from a client (identified by IP address) are sent to the same backend server for a specified period. Cookie persistence allows related client connections, identified by the same cookie in HTTP requests, to be redirected to the same server.
	ResourceType string `json:"resource_type"`

	// Opaque identifiers meaningful to the API user
	Tags []common.Tag `json:"tags,omitempty"`

	// If persistence shared flag is not set in the cookie persistence profile bound to a virtual server, it defaults to cookie persistence that is private to each virtual server and is qualified by the pool. This is accomplished by load balancer inserting a cookie with name in the format &lt;name&gt;.&lt;virtual_server_id&gt;.&lt;pool_id&gt;. If persistence shared flag is set in the cookie persistence profile, in cookie insert mode, cookie persistence could be shared across multiple virtual servers that are bound to the same pools. The cookie name would be changed to &lt;name&gt;.&lt;profile-id&gt;.&lt;pool-id&gt;. If persistence shared flag is not set in the sourceIp persistence profile bound to a virtual server, each virtual server that the profile is bound to maintains its own private persistence table. If persistence shared flag is set in the sourceIp persistence profile, all virtual servers the profile is bound to share the same persistence table.
	PersistenceShared bool `json:"persistence_shared"`
}
