# \ObjectsAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2ObjectsGet**](ObjectsAPI.md#V2ObjectsGet) | **Get** /v2/objects | List objects
[**V2ObjectsObjectGet**](ObjectsAPI.md#V2ObjectsObjectGet) | **Get** /v2/objects/{object} | Get an object
[**V2ObjectsObjectPatch**](ObjectsAPI.md#V2ObjectsObjectPatch) | **Patch** /v2/objects/{object} | Update an object
[**V2ObjectsPost**](ObjectsAPI.md#V2ObjectsPost) | **Post** /v2/objects | Create an object



## V2ObjectsGet

> V2ObjectsGet200Response V2ObjectsGet(ctx).Execute()

List objects



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
	resp, r, err := apiClient.ObjectsAPI.V2ObjectsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAPI.V2ObjectsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ObjectsGet`: V2ObjectsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAPI.V2ObjectsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2ObjectsGetRequest struct via the builder pattern


### Return type

[**V2ObjectsGet200Response**](V2ObjectsGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ObjectsObjectGet

> V2ObjectsPost200Response V2ObjectsObjectGet(ctx, object).Execute()

Get an object



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
	object := "people" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectsAPI.V2ObjectsObjectGet(context.Background(), object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAPI.V2ObjectsObjectGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ObjectsObjectGet`: V2ObjectsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAPI.V2ObjectsObjectGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**object** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ObjectsObjectGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2ObjectsPost200Response**](V2ObjectsPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ObjectsObjectPatch

> V2ObjectsPost200Response V2ObjectsObjectPatch(ctx, object).V2ObjectsObjectPatchRequest(v2ObjectsObjectPatchRequest).Execute()

Update an object



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
	object := "people" // string | 
	v2ObjectsObjectPatchRequest := *openapiclient.NewV2ObjectsObjectPatchRequest(*openapiclient.NewV2ObjectsObjectPatchRequestData()) // V2ObjectsObjectPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectsAPI.V2ObjectsObjectPatch(context.Background(), object).V2ObjectsObjectPatchRequest(v2ObjectsObjectPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAPI.V2ObjectsObjectPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ObjectsObjectPatch`: V2ObjectsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAPI.V2ObjectsObjectPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**object** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ObjectsObjectPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2ObjectsObjectPatchRequest** | [**V2ObjectsObjectPatchRequest**](V2ObjectsObjectPatchRequest.md) |  | 

### Return type

[**V2ObjectsPost200Response**](V2ObjectsPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ObjectsPost

> V2ObjectsPost200Response V2ObjectsPost(ctx).V2ObjectsPostRequest(v2ObjectsPostRequest).Execute()

Create an object



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
	v2ObjectsPostRequest := *openapiclient.NewV2ObjectsPostRequest(*openapiclient.NewV2ObjectsPostRequestData("people", "Person", "People")) // V2ObjectsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectsAPI.V2ObjectsPost(context.Background()).V2ObjectsPostRequest(v2ObjectsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAPI.V2ObjectsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ObjectsPost`: V2ObjectsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAPI.V2ObjectsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ObjectsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2ObjectsPostRequest** | [**V2ObjectsPostRequest**](V2ObjectsPostRequest.md) |  | 

### Return type

[**V2ObjectsPost200Response**](V2ObjectsPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

