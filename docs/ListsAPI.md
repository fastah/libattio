# \ListsAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2ListsGet**](ListsAPI.md#V2ListsGet) | **Get** /v2/lists | List all lists
[**V2ListsListGet**](ListsAPI.md#V2ListsListGet) | **Get** /v2/lists/{list} | Get a list
[**V2ListsListPatch**](ListsAPI.md#V2ListsListPatch) | **Patch** /v2/lists/{list} | Update a list
[**V2ListsPost**](ListsAPI.md#V2ListsPost) | **Post** /v2/lists | Create a list



## V2ListsGet

> V2ListsGet200Response V2ListsGet(ctx).Execute()

List all lists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.V2ListsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.V2ListsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ListsGet`: V2ListsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.V2ListsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2ListsGetRequest struct via the builder pattern


### Return type

[**V2ListsGet200Response**](V2ListsGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ListsListGet

> V2ListsPost200Response V2ListsListGet(ctx, list).Execute()

Get a list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	list := "33ebdbe9-e529-47c9-b894-0ba25e9c15c0" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.V2ListsListGet(context.Background(), list).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.V2ListsListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ListsListGet`: V2ListsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.V2ListsListGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**list** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ListsListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2ListsPost200Response**](V2ListsPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ListsListPatch

> V2ListsPost200Response V2ListsListPatch(ctx, list).V2ListsListPatchRequest(v2ListsListPatchRequest).Execute()

Update a list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	list := "33ebdbe9-e529-47c9-b894-0ba25e9c15c0" // string | 
	v2ListsListPatchRequest := *openapiclient.NewV2ListsListPatchRequest(*openapiclient.NewV2ListsListPatchRequestData()) // V2ListsListPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.V2ListsListPatch(context.Background(), list).V2ListsListPatchRequest(v2ListsListPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.V2ListsListPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ListsListPatch`: V2ListsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.V2ListsListPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**list** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ListsListPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2ListsListPatchRequest** | [**V2ListsListPatchRequest**](V2ListsListPatchRequest.md) |  | 

### Return type

[**V2ListsPost200Response**](V2ListsPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ListsPost

> V2ListsPost200Response V2ListsPost(ctx).V2ListsPostRequest(v2ListsPostRequest).Execute()

Create a list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v2ListsPostRequest := *openapiclient.NewV2ListsPostRequest(*openapiclient.NewV2ListsPostRequestData("Enterprise Sales", "enterprise_sales", "people", "read-and-write", []openapiclient.V2ListsPostRequestDataWorkspaceMemberAccessInner{*openapiclient.NewV2ListsPostRequestDataWorkspaceMemberAccessInner("50cf242c-7fa3-4cad-87d0-75b1af71c57b", "read-and-write")})) // V2ListsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.V2ListsPost(context.Background()).V2ListsPostRequest(v2ListsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.V2ListsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ListsPost`: V2ListsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.V2ListsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ListsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2ListsPostRequest** | [**V2ListsPostRequest**](V2ListsPostRequest.md) |  | 

### Return type

[**V2ListsPost200Response**](V2ListsPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

