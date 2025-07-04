/*
Attio API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
Contact: support@attio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libattio

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// WebhooksAPIService WebhooksAPI service
type WebhooksAPIService service

type ApiV2WebhooksGetRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	limit *int32
	offset *int32
}

func (r ApiV2WebhooksGetRequest) Limit(limit int32) ApiV2WebhooksGetRequest {
	r.limit = &limit
	return r
}

func (r ApiV2WebhooksGetRequest) Offset(offset int32) ApiV2WebhooksGetRequest {
	r.offset = &offset
	return r
}

func (r ApiV2WebhooksGetRequest) Execute() (*V2WebhooksGet200Response, *http.Response, error) {
	return r.ApiService.V2WebhooksGetExecute(r)
}

/*
V2WebhooksGet List webhooks

Get all of the webhooks in your workspace.

Required scopes: `webhook:read`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV2WebhooksGetRequest
*/
func (a *WebhooksAPIService) V2WebhooksGet(ctx context.Context) ApiV2WebhooksGetRequest {
	return ApiV2WebhooksGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V2WebhooksGet200Response
func (a *WebhooksAPIService) V2WebhooksGetExecute(r ApiV2WebhooksGetRequest) (*V2WebhooksGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V2WebhooksGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.V2WebhooksGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/webhooks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV2WebhooksPostRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	v2WebhooksPostRequest *V2WebhooksPostRequest
}

func (r ApiV2WebhooksPostRequest) V2WebhooksPostRequest(v2WebhooksPostRequest V2WebhooksPostRequest) ApiV2WebhooksPostRequest {
	r.v2WebhooksPostRequest = &v2WebhooksPostRequest
	return r
}

func (r ApiV2WebhooksPostRequest) Execute() (*V2WebhooksPost200Response, *http.Response, error) {
	return r.ApiService.V2WebhooksPostExecute(r)
}

/*
V2WebhooksPost Create a webhook

Create a webhook and associated subscriptions.

Required scopes: `webhook:read-write`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV2WebhooksPostRequest
*/
func (a *WebhooksAPIService) V2WebhooksPost(ctx context.Context) ApiV2WebhooksPostRequest {
	return ApiV2WebhooksPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V2WebhooksPost200Response
func (a *WebhooksAPIService) V2WebhooksPostExecute(r ApiV2WebhooksPostRequest) (*V2WebhooksPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V2WebhooksPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.V2WebhooksPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/webhooks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.v2WebhooksPostRequest == nil {
		return localVarReturnValue, nil, reportError("v2WebhooksPostRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.v2WebhooksPostRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v V2WebhooksPost400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV2WebhooksWebhookIdDeleteRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookId string
}

func (r ApiV2WebhooksWebhookIdDeleteRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.V2WebhooksWebhookIdDeleteExecute(r)
}

/*
V2WebhooksWebhookIdDelete Delete a webhook

Delete a webhook by ID.

Required scopes: `webhook:read-write`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webhookId
 @return ApiV2WebhooksWebhookIdDeleteRequest
*/
func (a *WebhooksAPIService) V2WebhooksWebhookIdDelete(ctx context.Context, webhookId string) ApiV2WebhooksWebhookIdDeleteRequest {
	return ApiV2WebhooksWebhookIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		webhookId: webhookId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *WebhooksAPIService) V2WebhooksWebhookIdDeleteExecute(r ApiV2WebhooksWebhookIdDeleteRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.V2WebhooksWebhookIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/webhooks/{webhook_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhook_id"+"}", url.PathEscape(parameterValueToString(r.webhookId, "webhookId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v V2WebhooksWebhookIdDelete404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV2WebhooksWebhookIdGetRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookId string
}

func (r ApiV2WebhooksWebhookIdGetRequest) Execute() (*V2WebhooksWebhookIdGet200Response, *http.Response, error) {
	return r.ApiService.V2WebhooksWebhookIdGetExecute(r)
}

/*
V2WebhooksWebhookIdGet Get a webhook

Get a single webhook.

Required scopes: `webhook:read`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webhookId
 @return ApiV2WebhooksWebhookIdGetRequest
*/
func (a *WebhooksAPIService) V2WebhooksWebhookIdGet(ctx context.Context, webhookId string) ApiV2WebhooksWebhookIdGetRequest {
	return ApiV2WebhooksWebhookIdGetRequest{
		ApiService: a,
		ctx: ctx,
		webhookId: webhookId,
	}
}

// Execute executes the request
//  @return V2WebhooksWebhookIdGet200Response
func (a *WebhooksAPIService) V2WebhooksWebhookIdGetExecute(r ApiV2WebhooksWebhookIdGetRequest) (*V2WebhooksWebhookIdGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V2WebhooksWebhookIdGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.V2WebhooksWebhookIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/webhooks/{webhook_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhook_id"+"}", url.PathEscape(parameterValueToString(r.webhookId, "webhookId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v V2WebhooksWebhookIdGet404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV2WebhooksWebhookIdPatchRequest struct {
	ctx context.Context
	ApiService *WebhooksAPIService
	webhookId string
	v2WebhooksWebhookIdPatchRequest *V2WebhooksWebhookIdPatchRequest
}

func (r ApiV2WebhooksWebhookIdPatchRequest) V2WebhooksWebhookIdPatchRequest(v2WebhooksWebhookIdPatchRequest V2WebhooksWebhookIdPatchRequest) ApiV2WebhooksWebhookIdPatchRequest {
	r.v2WebhooksWebhookIdPatchRequest = &v2WebhooksWebhookIdPatchRequest
	return r
}

func (r ApiV2WebhooksWebhookIdPatchRequest) Execute() (*V2WebhooksWebhookIdGet200Response, *http.Response, error) {
	return r.ApiService.V2WebhooksWebhookIdPatchExecute(r)
}

/*
V2WebhooksWebhookIdPatch Update a webhook

Update a webhook and associated subscriptions.

Required scopes: `webhook:read-write`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webhookId
 @return ApiV2WebhooksWebhookIdPatchRequest
*/
func (a *WebhooksAPIService) V2WebhooksWebhookIdPatch(ctx context.Context, webhookId string) ApiV2WebhooksWebhookIdPatchRequest {
	return ApiV2WebhooksWebhookIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		webhookId: webhookId,
	}
}

// Execute executes the request
//  @return V2WebhooksWebhookIdGet200Response
func (a *WebhooksAPIService) V2WebhooksWebhookIdPatchExecute(r ApiV2WebhooksWebhookIdPatchRequest) (*V2WebhooksWebhookIdGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V2WebhooksWebhookIdGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.V2WebhooksWebhookIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/webhooks/{webhook_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhook_id"+"}", url.PathEscape(parameterValueToString(r.webhookId, "webhookId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.v2WebhooksWebhookIdPatchRequest == nil {
		return localVarReturnValue, nil, reportError("v2WebhooksWebhookIdPatchRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.v2WebhooksWebhookIdPatchRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v V2WebhooksWebhookIdGet404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
