# \EntriesAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2ListsListEntriesEntryIdAttributesAttributeValuesGet**](EntriesAPI.md#V2ListsListEntriesEntryIdAttributesAttributeValuesGet) | **Get** /v2/lists/{list}/entries/{entry_id}/attributes/{attribute}/values | List attribute values for a list entry
[**V2ListsListEntriesEntryIdDelete**](EntriesAPI.md#V2ListsListEntriesEntryIdDelete) | **Delete** /v2/lists/{list}/entries/{entry_id} | Delete a list entry
[**V2ListsListEntriesEntryIdGet**](EntriesAPI.md#V2ListsListEntriesEntryIdGet) | **Get** /v2/lists/{list}/entries/{entry_id} | Get a list entry
[**V2ListsListEntriesEntryIdPatch**](EntriesAPI.md#V2ListsListEntriesEntryIdPatch) | **Patch** /v2/lists/{list}/entries/{entry_id} | Update a list entry (append multiselect values)
[**V2ListsListEntriesEntryIdPut**](EntriesAPI.md#V2ListsListEntriesEntryIdPut) | **Put** /v2/lists/{list}/entries/{entry_id} | Update a list entry (overwrite multiselect values)
[**V2ListsListEntriesPost**](EntriesAPI.md#V2ListsListEntriesPost) | **Post** /v2/lists/{list}/entries | Create an entry (add record to list)
[**V2ListsListEntriesPut**](EntriesAPI.md#V2ListsListEntriesPut) | **Put** /v2/lists/{list}/entries | Assert a list entry by parent
[**V2ListsListEntriesQueryPost**](EntriesAPI.md#V2ListsListEntriesQueryPost) | **Post** /v2/lists/{list}/entries/query | List entries



## V2ListsListEntriesEntryIdAttributesAttributeValuesGet

> V2ObjectsObjectRecordsRecordIdAttributesAttributeValuesGet200Response V2ListsListEntriesEntryIdAttributesAttributeValuesGet(ctx, list, entryId, attribute).ShowHistoric(showHistoric).Limit(limit).Offset(offset).Execute()

List attribute values for a list entry



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
	list := "enterprise_sales" // string | 
	entryId := "2e6e29ea-c4e0-4f44-842d-78a891f8c156" // string | 
	attribute := "41252299-f8c7-4b5e-99c9-4ff8321d2f96" // string | 
	showHistoric := true // bool |  (optional) (default to false)
	limit := int32(10) // int32 |  (optional)
	offset := int32(5) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.V2ListsListEntriesEntryIdAttributesAttributeValuesGet(context.Background(), list, entryId, attribute).ShowHistoric(showHistoric).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.V2ListsListEntriesEntryIdAttributesAttributeValuesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ListsListEntriesEntryIdAttributesAttributeValuesGet`: V2ObjectsObjectRecordsRecordIdAttributesAttributeValuesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.V2ListsListEntriesEntryIdAttributesAttributeValuesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**list** | **string** |  | 
**entryId** | **string** |  | 
**attribute** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ListsListEntriesEntryIdAttributesAttributeValuesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **showHistoric** | **bool** |  | [default to false]
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**V2ObjectsObjectRecordsRecordIdAttributesAttributeValuesGet200Response**](V2ObjectsObjectRecordsRecordIdAttributesAttributeValuesGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ListsListEntriesEntryIdDelete

> map[string]interface{} V2ListsListEntriesEntryIdDelete(ctx, list, entryId).Execute()

Delete a list entry



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
	list := "enterprise_sales" // string | 
	entryId := "2e6e29ea-c4e0-4f44-842d-78a891f8c156" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.V2ListsListEntriesEntryIdDelete(context.Background(), list, entryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.V2ListsListEntriesEntryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ListsListEntriesEntryIdDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.V2ListsListEntriesEntryIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**list** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ListsListEntriesEntryIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ListsListEntriesEntryIdGet

> V2ListsListEntriesPut200Response V2ListsListEntriesEntryIdGet(ctx, list, entryId).Execute()

Get a list entry



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
	entryId := "2e6e29ea-c4e0-4f44-842d-78a891f8c156" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.V2ListsListEntriesEntryIdGet(context.Background(), list, entryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.V2ListsListEntriesEntryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ListsListEntriesEntryIdGet`: V2ListsListEntriesPut200Response
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.V2ListsListEntriesEntryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**list** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ListsListEntriesEntryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V2ListsListEntriesPut200Response**](V2ListsListEntriesPut200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ListsListEntriesEntryIdPatch

> V2ListsListEntriesPut200Response V2ListsListEntriesEntryIdPatch(ctx, list, entryId).V2ListsListEntriesEntryIdPutRequest(v2ListsListEntriesEntryIdPutRequest).Execute()

Update a list entry (append multiselect values)



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
	entryId := "2e6e29ea-c4e0-4f44-842d-78a891f8c156" // string | 
	v2ListsListEntriesEntryIdPutRequest := *openapiclient.NewV2ListsListEntriesEntryIdPutRequest(*openapiclient.NewV2ListsListEntriesEntryIdPutRequestData(map[string][]interface{}{"key": []interface{}{nil}})) // V2ListsListEntriesEntryIdPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.V2ListsListEntriesEntryIdPatch(context.Background(), list, entryId).V2ListsListEntriesEntryIdPutRequest(v2ListsListEntriesEntryIdPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.V2ListsListEntriesEntryIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ListsListEntriesEntryIdPatch`: V2ListsListEntriesPut200Response
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.V2ListsListEntriesEntryIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**list** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ListsListEntriesEntryIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2ListsListEntriesEntryIdPutRequest** | [**V2ListsListEntriesEntryIdPutRequest**](V2ListsListEntriesEntryIdPutRequest.md) |  | 

### Return type

[**V2ListsListEntriesPut200Response**](V2ListsListEntriesPut200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ListsListEntriesEntryIdPut

> V2ListsListEntriesPut200Response V2ListsListEntriesEntryIdPut(ctx, list, entryId).V2ListsListEntriesEntryIdPutRequest(v2ListsListEntriesEntryIdPutRequest).Execute()

Update a list entry (overwrite multiselect values)



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
	entryId := "2e6e29ea-c4e0-4f44-842d-78a891f8c156" // string | 
	v2ListsListEntriesEntryIdPutRequest := *openapiclient.NewV2ListsListEntriesEntryIdPutRequest(*openapiclient.NewV2ListsListEntriesEntryIdPutRequestData(map[string][]interface{}{"key": []interface{}{nil}})) // V2ListsListEntriesEntryIdPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.V2ListsListEntriesEntryIdPut(context.Background(), list, entryId).V2ListsListEntriesEntryIdPutRequest(v2ListsListEntriesEntryIdPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.V2ListsListEntriesEntryIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ListsListEntriesEntryIdPut`: V2ListsListEntriesPut200Response
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.V2ListsListEntriesEntryIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**list** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ListsListEntriesEntryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2ListsListEntriesEntryIdPutRequest** | [**V2ListsListEntriesEntryIdPutRequest**](V2ListsListEntriesEntryIdPutRequest.md) |  | 

### Return type

[**V2ListsListEntriesPut200Response**](V2ListsListEntriesPut200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ListsListEntriesPost

> V2ListsListEntriesPut200Response V2ListsListEntriesPost(ctx, list).V2ListsListEntriesPutRequest(v2ListsListEntriesPutRequest).Execute()

Create an entry (add record to list)



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
	v2ListsListEntriesPutRequest := *openapiclient.NewV2ListsListEntriesPutRequest(*openapiclient.NewV2ListsListEntriesPutRequestData("891dcbfc-9141-415d-9b2a-2238a6cc012d", "people", map[string][]interface{}{"key": []interface{}{nil}})) // V2ListsListEntriesPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.V2ListsListEntriesPost(context.Background(), list).V2ListsListEntriesPutRequest(v2ListsListEntriesPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.V2ListsListEntriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ListsListEntriesPost`: V2ListsListEntriesPut200Response
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.V2ListsListEntriesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**list** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ListsListEntriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2ListsListEntriesPutRequest** | [**V2ListsListEntriesPutRequest**](V2ListsListEntriesPutRequest.md) |  | 

### Return type

[**V2ListsListEntriesPut200Response**](V2ListsListEntriesPut200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ListsListEntriesPut

> V2ListsListEntriesPut200Response V2ListsListEntriesPut(ctx, list).V2ListsListEntriesPutRequest(v2ListsListEntriesPutRequest).Execute()

Assert a list entry by parent



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
	v2ListsListEntriesPutRequest := *openapiclient.NewV2ListsListEntriesPutRequest(*openapiclient.NewV2ListsListEntriesPutRequestData("891dcbfc-9141-415d-9b2a-2238a6cc012d", "people", map[string][]interface{}{"key": []interface{}{nil}})) // V2ListsListEntriesPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.V2ListsListEntriesPut(context.Background(), list).V2ListsListEntriesPutRequest(v2ListsListEntriesPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.V2ListsListEntriesPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ListsListEntriesPut`: V2ListsListEntriesPut200Response
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.V2ListsListEntriesPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**list** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ListsListEntriesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2ListsListEntriesPutRequest** | [**V2ListsListEntriesPutRequest**](V2ListsListEntriesPutRequest.md) |  | 

### Return type

[**V2ListsListEntriesPut200Response**](V2ListsListEntriesPut200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ListsListEntriesQueryPost

> V2ListsListEntriesQueryPost200Response V2ListsListEntriesQueryPost(ctx, list).V2ObjectsObjectRecordsQueryPostRequest(v2ObjectsObjectRecordsQueryPostRequest).Execute()

List entries



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
	v2ObjectsObjectRecordsQueryPostRequest := *openapiclient.NewV2ObjectsObjectRecordsQueryPostRequest() // V2ObjectsObjectRecordsQueryPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntriesAPI.V2ListsListEntriesQueryPost(context.Background(), list).V2ObjectsObjectRecordsQueryPostRequest(v2ObjectsObjectRecordsQueryPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntriesAPI.V2ListsListEntriesQueryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ListsListEntriesQueryPost`: V2ListsListEntriesQueryPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EntriesAPI.V2ListsListEntriesQueryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**list** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ListsListEntriesQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2ObjectsObjectRecordsQueryPostRequest** | [**V2ObjectsObjectRecordsQueryPostRequest**](V2ObjectsObjectRecordsQueryPostRequest.md) |  | 

### Return type

[**V2ListsListEntriesQueryPost200Response**](V2ListsListEntriesQueryPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

