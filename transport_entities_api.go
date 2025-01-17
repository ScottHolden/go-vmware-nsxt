/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package nsxt

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/ScottHolden/go-vmware-nsxt/manager"
)

// Linger please
var (
	_ context.Context
)

type TransportEntitiesApiService service

/* TransportEntitiesApiService Tunnel properties
Tunnel properties
* @param ctx context.Context Authentication Context
@param nodeId ID of transport node
@param tunnelName Tunnel name
@param optional (nil or map[string]interface{}) with one or more of:
    @param "source" (string) Data source type.
@return manager.TunnelProperties*/
func (a *TransportEntitiesApiService) GetTunnel(ctx context.Context, nodeId string, tunnelName string, localVarOptionals map[string]interface{}) (manager.TunnelProperties, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     manager.TunnelProperties
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/transport-nodes/{node-id}/tunnels/{tunnel-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"node-id"+"}", fmt.Sprintf("%v", nodeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tunnel-name"+"}", fmt.Sprintf("%v", tunnelName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["source"], "string", "source"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["source"].(string); localVarOk {
		localVarQueryParams.Add("source", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* TransportEntitiesApiService List of tunnels
List of tunnels
* @param ctx context.Context Authentication Context
@param nodeId ID of transport node
@param optional (nil or map[string]interface{}) with one or more of:
    @param "bfdDiagnosticCode" (string) BFD diagnostic code of Tunnel as defined in RFC 5880
    @param "cursor" (string) Opaque cursor to be used for getting next page of records (supplied by current result page)
    @param "includedFields" (string) Comma separated list of fields that should be included to result of query
    @param "pageSize" (int64) Maximum number of results to return in this page (server may return fewer)
    @param "remoteNodeId" (string)
    @param "sortAscending" (bool)
    @param "sortBy" (string) Field by which records are sorted
    @param "source" (string) Data source type.
    @param "status" (string) Tunnel status
@return manager.TunnelList*/
func (a *TransportEntitiesApiService) QueryTunnels(ctx context.Context, nodeId string, localVarOptionals map[string]interface{}) (manager.TunnelList, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     manager.TunnelList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/transport-nodes/{node-id}/tunnels"
	localVarPath = strings.Replace(localVarPath, "{"+"node-id"+"}", fmt.Sprintf("%v", nodeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["bfdDiagnosticCode"], "string", "bfdDiagnosticCode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["cursor"], "string", "cursor"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includedFields"], "string", "includedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pageSize"], "int64", "pageSize"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["remoteNodeId"], "string", "remoteNodeId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortAscending"], "bool", "sortAscending"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortBy"], "string", "sortBy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["source"], "string", "source"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["status"], "string", "status"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["bfdDiagnosticCode"].(string); localVarOk {
		localVarQueryParams.Add("bfd_diagnostic_code", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["cursor"].(string); localVarOk {
		localVarQueryParams.Add("cursor", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includedFields"].(string); localVarOk {
		localVarQueryParams.Add("included_fields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pageSize"].(int64); localVarOk {
		localVarQueryParams.Add("page_size", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["remoteNodeId"].(string); localVarOk {
		localVarQueryParams.Add("remote_node_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortAscending"].(bool); localVarOk {
		localVarQueryParams.Add("sort_ascending", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortBy"].(string); localVarOk {
		localVarQueryParams.Add("sort_by", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["source"].(string); localVarOk {
		localVarQueryParams.Add("source", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["status"].(string); localVarOk {
		localVarQueryParams.Add("status", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}
