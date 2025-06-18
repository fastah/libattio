# \AttributesAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2TargetIdentifierAttributesAttributeGet**](AttributesAPI.md#V2TargetIdentifierAttributesAttributeGet) | **Get** /v2/{target}/{identifier}/attributes/{attribute} | Get an attribute
[**V2TargetIdentifierAttributesAttributeOptionsGet**](AttributesAPI.md#V2TargetIdentifierAttributesAttributeOptionsGet) | **Get** /v2/{target}/{identifier}/attributes/{attribute}/options | List select options
[**V2TargetIdentifierAttributesAttributeOptionsOptionPatch**](AttributesAPI.md#V2TargetIdentifierAttributesAttributeOptionsOptionPatch) | **Patch** /v2/{target}/{identifier}/attributes/{attribute}/options/{option} | Update a select option
[**V2TargetIdentifierAttributesAttributeOptionsPost**](AttributesAPI.md#V2TargetIdentifierAttributesAttributeOptionsPost) | **Post** /v2/{target}/{identifier}/attributes/{attribute}/options | Create a select option
[**V2TargetIdentifierAttributesAttributePatch**](AttributesAPI.md#V2TargetIdentifierAttributesAttributePatch) | **Patch** /v2/{target}/{identifier}/attributes/{attribute} | Update an attribute
[**V2TargetIdentifierAttributesAttributeStatusesGet**](AttributesAPI.md#V2TargetIdentifierAttributesAttributeStatusesGet) | **Get** /v2/{target}/{identifier}/attributes/{attribute}/statuses | List statuses
[**V2TargetIdentifierAttributesAttributeStatusesPost**](AttributesAPI.md#V2TargetIdentifierAttributesAttributeStatusesPost) | **Post** /v2/{target}/{identifier}/attributes/{attribute}/statuses | Create a status
[**V2TargetIdentifierAttributesAttributeStatusesStatusPatch**](AttributesAPI.md#V2TargetIdentifierAttributesAttributeStatusesStatusPatch) | **Patch** /v2/{target}/{identifier}/attributes/{attribute}/statuses/{status} | Update a status
[**V2TargetIdentifierAttributesGet**](AttributesAPI.md#V2TargetIdentifierAttributesGet) | **Get** /v2/{target}/{identifier}/attributes | List attributes
[**V2TargetIdentifierAttributesPost**](AttributesAPI.md#V2TargetIdentifierAttributesPost) | **Post** /v2/{target}/{identifier}/attributes | Create an attribute



## V2TargetIdentifierAttributesAttributeGet

> V2TargetIdentifierAttributesPost200Response V2TargetIdentifierAttributesAttributeGet(ctx, target, identifier, attribute).Execute()

Get an attribute



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
	target := "lists" // string | 
	identifier := "33ebdbe9-e529-47c9-b894-0ba25e9c15c0" // string | 
	attribute := "41252299-f8c7-4b5e-99c9-4ff8321d2f96" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributesAPI.V2TargetIdentifierAttributesAttributeGet(context.Background(), target, identifier, attribute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.V2TargetIdentifierAttributesAttributeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TargetIdentifierAttributesAttributeGet`: V2TargetIdentifierAttributesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.V2TargetIdentifierAttributesAttributeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 
**identifier** | **string** |  | 
**attribute** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2TargetIdentifierAttributesAttributeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2TargetIdentifierAttributesPost200Response**](V2TargetIdentifierAttributesPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2TargetIdentifierAttributesAttributeOptionsGet

> V2TargetIdentifierAttributesAttributeOptionsGet200Response V2TargetIdentifierAttributesAttributeOptionsGet(ctx, target, identifier, attribute).ShowArchived(showArchived).Execute()

List select options



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
	target := "lists" // string | 
	identifier := "33ebdbe9-e529-47c9-b894-0ba25e9c15c0" // string | 
	attribute := "41252299-f8c7-4b5e-99c9-4ff8321d2f96" // string | 
	showArchived := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributesAPI.V2TargetIdentifierAttributesAttributeOptionsGet(context.Background(), target, identifier, attribute).ShowArchived(showArchived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.V2TargetIdentifierAttributesAttributeOptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TargetIdentifierAttributesAttributeOptionsGet`: V2TargetIdentifierAttributesAttributeOptionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.V2TargetIdentifierAttributesAttributeOptionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 
**identifier** | **string** |  | 
**attribute** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2TargetIdentifierAttributesAttributeOptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **showArchived** | **bool** |  | 

### Return type

[**V2TargetIdentifierAttributesAttributeOptionsGet200Response**](V2TargetIdentifierAttributesAttributeOptionsGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2TargetIdentifierAttributesAttributeOptionsOptionPatch

> V2TargetIdentifierAttributesAttributeOptionsPost200Response V2TargetIdentifierAttributesAttributeOptionsOptionPatch(ctx, target, identifier, attribute, option).V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest(v2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest).Execute()

Update a select option



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
	target := "lists" // string | 
	identifier := "33ebdbe9-e529-47c9-b894-0ba25e9c15c0" // string | 
	attribute := "41252299-f8c7-4b5e-99c9-4ff8321d2f96" // string | 
	option := "Medium" // string | 
	v2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest := *openapiclient.NewV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest(*openapiclient.NewV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequestData()) // V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributesAPI.V2TargetIdentifierAttributesAttributeOptionsOptionPatch(context.Background(), target, identifier, attribute, option).V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest(v2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.V2TargetIdentifierAttributesAttributeOptionsOptionPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TargetIdentifierAttributesAttributeOptionsOptionPatch`: V2TargetIdentifierAttributesAttributeOptionsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.V2TargetIdentifierAttributesAttributeOptionsOptionPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 
**identifier** | **string** |  | 
**attribute** | **string** |  | 
**option** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **v2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest** | [**V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest**](V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest.md) |  | 

### Return type

[**V2TargetIdentifierAttributesAttributeOptionsPost200Response**](V2TargetIdentifierAttributesAttributeOptionsPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2TargetIdentifierAttributesAttributeOptionsPost

> V2TargetIdentifierAttributesAttributeOptionsPost200Response V2TargetIdentifierAttributesAttributeOptionsPost(ctx, target, identifier, attribute).V2TargetIdentifierAttributesAttributeOptionsPostRequest(v2TargetIdentifierAttributesAttributeOptionsPostRequest).Execute()

Create a select option



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
	target := "lists" // string | 
	identifier := "33ebdbe9-e529-47c9-b894-0ba25e9c15c0" // string | 
	attribute := "41252299-f8c7-4b5e-99c9-4ff8321d2f96" // string | 
	v2TargetIdentifierAttributesAttributeOptionsPostRequest := *openapiclient.NewV2TargetIdentifierAttributesAttributeOptionsPostRequest(*openapiclient.NewV2TargetIdentifierAttributesAttributeOptionsPostRequestData("Medium")) // V2TargetIdentifierAttributesAttributeOptionsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributesAPI.V2TargetIdentifierAttributesAttributeOptionsPost(context.Background(), target, identifier, attribute).V2TargetIdentifierAttributesAttributeOptionsPostRequest(v2TargetIdentifierAttributesAttributeOptionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.V2TargetIdentifierAttributesAttributeOptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TargetIdentifierAttributesAttributeOptionsPost`: V2TargetIdentifierAttributesAttributeOptionsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.V2TargetIdentifierAttributesAttributeOptionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 
**identifier** | **string** |  | 
**attribute** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2TargetIdentifierAttributesAttributeOptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2TargetIdentifierAttributesAttributeOptionsPostRequest** | [**V2TargetIdentifierAttributesAttributeOptionsPostRequest**](V2TargetIdentifierAttributesAttributeOptionsPostRequest.md) |  | 

### Return type

[**V2TargetIdentifierAttributesAttributeOptionsPost200Response**](V2TargetIdentifierAttributesAttributeOptionsPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2TargetIdentifierAttributesAttributePatch

> V2TargetIdentifierAttributesPost200Response V2TargetIdentifierAttributesAttributePatch(ctx, target, identifier, attribute).V2TargetIdentifierAttributesAttributePatchRequest(v2TargetIdentifierAttributesAttributePatchRequest).Execute()

Update an attribute



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
	target := "lists" // string | 
	identifier := "33ebdbe9-e529-47c9-b894-0ba25e9c15c0" // string | 
	attribute := "41252299-f8c7-4b5e-99c9-4ff8321d2f96" // string | 
	v2TargetIdentifierAttributesAttributePatchRequest := *openapiclient.NewV2TargetIdentifierAttributesAttributePatchRequest(*openapiclient.NewV2TargetIdentifierAttributesAttributePatchRequestData()) // V2TargetIdentifierAttributesAttributePatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributesAPI.V2TargetIdentifierAttributesAttributePatch(context.Background(), target, identifier, attribute).V2TargetIdentifierAttributesAttributePatchRequest(v2TargetIdentifierAttributesAttributePatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.V2TargetIdentifierAttributesAttributePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TargetIdentifierAttributesAttributePatch`: V2TargetIdentifierAttributesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.V2TargetIdentifierAttributesAttributePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 
**identifier** | **string** |  | 
**attribute** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2TargetIdentifierAttributesAttributePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2TargetIdentifierAttributesAttributePatchRequest** | [**V2TargetIdentifierAttributesAttributePatchRequest**](V2TargetIdentifierAttributesAttributePatchRequest.md) |  | 

### Return type

[**V2TargetIdentifierAttributesPost200Response**](V2TargetIdentifierAttributesPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2TargetIdentifierAttributesAttributeStatusesGet

> V2TargetIdentifierAttributesAttributeStatusesGet200Response V2TargetIdentifierAttributesAttributeStatusesGet(ctx, target, identifier, attribute).ShowArchived(showArchived).Execute()

List statuses



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
	target := "lists" // string | 
	identifier := "33ebdbe9-e529-47c9-b894-0ba25e9c15c0" // string | 
	attribute := "41252299-f8c7-4b5e-99c9-4ff8321d2f96" // string | 
	showArchived := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributesAPI.V2TargetIdentifierAttributesAttributeStatusesGet(context.Background(), target, identifier, attribute).ShowArchived(showArchived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.V2TargetIdentifierAttributesAttributeStatusesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TargetIdentifierAttributesAttributeStatusesGet`: V2TargetIdentifierAttributesAttributeStatusesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.V2TargetIdentifierAttributesAttributeStatusesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 
**identifier** | **string** |  | 
**attribute** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2TargetIdentifierAttributesAttributeStatusesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **showArchived** | **bool** |  | [default to false]

### Return type

[**V2TargetIdentifierAttributesAttributeStatusesGet200Response**](V2TargetIdentifierAttributesAttributeStatusesGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2TargetIdentifierAttributesAttributeStatusesPost

> V2TargetIdentifierAttributesAttributeStatusesPost200Response V2TargetIdentifierAttributesAttributeStatusesPost(ctx, target, identifier, attribute).V2TargetIdentifierAttributesAttributeStatusesPostRequest(v2TargetIdentifierAttributesAttributeStatusesPostRequest).Execute()

Create a status



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
	target := "lists" // string | 
	identifier := "33ebdbe9-e529-47c9-b894-0ba25e9c15c0" // string | 
	attribute := "41252299-f8c7-4b5e-99c9-4ff8321d2f96" // string | 
	v2TargetIdentifierAttributesAttributeStatusesPostRequest := *openapiclient.NewV2TargetIdentifierAttributesAttributeStatusesPostRequest(*openapiclient.NewV2TargetIdentifierAttributesAttributeStatusesPostRequestData("In Progress")) // V2TargetIdentifierAttributesAttributeStatusesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributesAPI.V2TargetIdentifierAttributesAttributeStatusesPost(context.Background(), target, identifier, attribute).V2TargetIdentifierAttributesAttributeStatusesPostRequest(v2TargetIdentifierAttributesAttributeStatusesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.V2TargetIdentifierAttributesAttributeStatusesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TargetIdentifierAttributesAttributeStatusesPost`: V2TargetIdentifierAttributesAttributeStatusesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.V2TargetIdentifierAttributesAttributeStatusesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 
**identifier** | **string** |  | 
**attribute** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2TargetIdentifierAttributesAttributeStatusesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2TargetIdentifierAttributesAttributeStatusesPostRequest** | [**V2TargetIdentifierAttributesAttributeStatusesPostRequest**](V2TargetIdentifierAttributesAttributeStatusesPostRequest.md) |  | 

### Return type

[**V2TargetIdentifierAttributesAttributeStatusesPost200Response**](V2TargetIdentifierAttributesAttributeStatusesPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2TargetIdentifierAttributesAttributeStatusesStatusPatch

> V2TargetIdentifierAttributesAttributeStatusesPost200Response V2TargetIdentifierAttributesAttributeStatusesStatusPatch(ctx, target, identifier, attribute, status).V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequest(v2TargetIdentifierAttributesAttributeStatusesStatusPatchRequest).Execute()

Update a status



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
	target := "lists" // string | 
	identifier := "33ebdbe9-e529-47c9-b894-0ba25e9c15c0" // string | 
	attribute := "41252299-f8c7-4b5e-99c9-4ff8321d2f96" // string | 
	status := "In Progress" // string | 
	v2TargetIdentifierAttributesAttributeStatusesStatusPatchRequest := *openapiclient.NewV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequest(*openapiclient.NewV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData()) // V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributesAPI.V2TargetIdentifierAttributesAttributeStatusesStatusPatch(context.Background(), target, identifier, attribute, status).V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequest(v2TargetIdentifierAttributesAttributeStatusesStatusPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.V2TargetIdentifierAttributesAttributeStatusesStatusPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TargetIdentifierAttributesAttributeStatusesStatusPatch`: V2TargetIdentifierAttributesAttributeStatusesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.V2TargetIdentifierAttributesAttributeStatusesStatusPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 
**identifier** | **string** |  | 
**attribute** | **string** |  | 
**status** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **v2TargetIdentifierAttributesAttributeStatusesStatusPatchRequest** | [**V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequest**](V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequest.md) |  | 

### Return type

[**V2TargetIdentifierAttributesAttributeStatusesPost200Response**](V2TargetIdentifierAttributesAttributeStatusesPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2TargetIdentifierAttributesGet

> V2TargetIdentifierAttributesGet200Response V2TargetIdentifierAttributesGet(ctx, target, identifier).Limit(limit).Offset(offset).ShowArchived(showArchived).Execute()

List attributes



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
	target := "lists" // string | 
	identifier := "33ebdbe9-e529-47c9-b894-0ba25e9c15c0" // string | 
	limit := int32(10) // int32 |  (optional)
	offset := int32(5) // int32 |  (optional)
	showArchived := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributesAPI.V2TargetIdentifierAttributesGet(context.Background(), target, identifier).Limit(limit).Offset(offset).ShowArchived(showArchived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.V2TargetIdentifierAttributesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TargetIdentifierAttributesGet`: V2TargetIdentifierAttributesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.V2TargetIdentifierAttributesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2TargetIdentifierAttributesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **showArchived** | **bool** |  | 

### Return type

[**V2TargetIdentifierAttributesGet200Response**](V2TargetIdentifierAttributesGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2TargetIdentifierAttributesPost

> V2TargetIdentifierAttributesPost200Response V2TargetIdentifierAttributesPost(ctx, target, identifier).V2TargetIdentifierAttributesPostRequest(v2TargetIdentifierAttributesPostRequest).Execute()

Create an attribute



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
	target := "lists" // string | 
	identifier := "97052eb9-e65e-443f-a297-f2d9a4a7f795" // string | 
	v2TargetIdentifierAttributesPostRequest := *openapiclient.NewV2TargetIdentifierAttributesPostRequest(*openapiclient.NewV2TargetIdentifierAttributesPostRequestData("Your Attribute", "Lorem ipsum", "my-attribute", "text", true, true, true, *openapiclient.NewV2TargetIdentifierAttributesPostRequestDataConfig())) // V2TargetIdentifierAttributesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributesAPI.V2TargetIdentifierAttributesPost(context.Background(), target, identifier).V2TargetIdentifierAttributesPostRequest(v2TargetIdentifierAttributesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributesAPI.V2TargetIdentifierAttributesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TargetIdentifierAttributesPost`: V2TargetIdentifierAttributesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AttributesAPI.V2TargetIdentifierAttributesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**target** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2TargetIdentifierAttributesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2TargetIdentifierAttributesPostRequest** | [**V2TargetIdentifierAttributesPostRequest**](V2TargetIdentifierAttributesPostRequest.md) |  | 

### Return type

[**V2TargetIdentifierAttributesPost200Response**](V2TargetIdentifierAttributesPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

