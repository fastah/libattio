# \RecordsAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2ObjectsObjectRecordsPost**](RecordsAPI.md#V2ObjectsObjectRecordsPost) | **Post** /v2/objects/{object}/records | Create a record
[**V2ObjectsObjectRecordsPut**](RecordsAPI.md#V2ObjectsObjectRecordsPut) | **Put** /v2/objects/{object}/records | Assert a record
[**V2ObjectsObjectRecordsQueryPost**](RecordsAPI.md#V2ObjectsObjectRecordsQueryPost) | **Post** /v2/objects/{object}/records/query | List records
[**V2ObjectsObjectRecordsRecordIdAttributesAttributeValuesGet**](RecordsAPI.md#V2ObjectsObjectRecordsRecordIdAttributesAttributeValuesGet) | **Get** /v2/objects/{object}/records/{record_id}/attributes/{attribute}/values | List record attribute values
[**V2ObjectsObjectRecordsRecordIdDelete**](RecordsAPI.md#V2ObjectsObjectRecordsRecordIdDelete) | **Delete** /v2/objects/{object}/records/{record_id} | Delete a record
[**V2ObjectsObjectRecordsRecordIdEntriesGet**](RecordsAPI.md#V2ObjectsObjectRecordsRecordIdEntriesGet) | **Get** /v2/objects/{object}/records/{record_id}/entries | List record entries
[**V2ObjectsObjectRecordsRecordIdGet**](RecordsAPI.md#V2ObjectsObjectRecordsRecordIdGet) | **Get** /v2/objects/{object}/records/{record_id} | Get a record
[**V2ObjectsObjectRecordsRecordIdPatch**](RecordsAPI.md#V2ObjectsObjectRecordsRecordIdPatch) | **Patch** /v2/objects/{object}/records/{record_id} | Update a record (append multiselect values)
[**V2ObjectsObjectRecordsRecordIdPut**](RecordsAPI.md#V2ObjectsObjectRecordsRecordIdPut) | **Put** /v2/objects/{object}/records/{record_id} | Update a record (overwrite multiselect values)



## V2ObjectsObjectRecordsPost

> V2ObjectsObjectRecordsPost200Response V2ObjectsObjectRecordsPost(ctx, object).V2ObjectsObjectRecordsPostRequest(v2ObjectsObjectRecordsPostRequest).Execute()

Create a record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/fastah/libattio"
)

func main() {
	object := "people" // string | 
	v2ObjectsObjectRecordsPostRequest := *libattio.NewV2ObjectsObjectRecordsPostRequest(*libattio.NewV2ObjectsObjectRecordsPostRequestData(map[string][]interface{}{"key": []interface{}{nil}})) // V2ObjectsObjectRecordsPostRequest | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.V2ObjectsObjectRecordsPost(context.Background(), object).V2ObjectsObjectRecordsPostRequest(v2ObjectsObjectRecordsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.V2ObjectsObjectRecordsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ObjectsObjectRecordsPost`: V2ObjectsObjectRecordsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.V2ObjectsObjectRecordsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**object** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ObjectsObjectRecordsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2ObjectsObjectRecordsPostRequest** | [**V2ObjectsObjectRecordsPostRequest**](V2ObjectsObjectRecordsPostRequest.md) |  | 

### Return type

[**V2ObjectsObjectRecordsPost200Response**](V2ObjectsObjectRecordsPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ObjectsObjectRecordsPut

> V2ObjectsObjectRecordsPut200Response V2ObjectsObjectRecordsPut(ctx, object).MatchingAttribute(matchingAttribute).V2ObjectsObjectRecordsPutRequest(v2ObjectsObjectRecordsPutRequest).Execute()

Assert a record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/fastah/libattio"
)

func main() {
	object := "people" // string | 
	matchingAttribute := "41252299-f8c7-4b5e-99c9-4ff8321d2f96" // string | 
	v2ObjectsObjectRecordsPutRequest := *libattio.NewV2ObjectsObjectRecordsPutRequest(*libattio.NewV2ObjectsObjectRecordsPutRequestData(map[string][]interface{}{"key": []interface{}{nil}})) // V2ObjectsObjectRecordsPutRequest | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.V2ObjectsObjectRecordsPut(context.Background(), object).MatchingAttribute(matchingAttribute).V2ObjectsObjectRecordsPutRequest(v2ObjectsObjectRecordsPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.V2ObjectsObjectRecordsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ObjectsObjectRecordsPut`: V2ObjectsObjectRecordsPut200Response
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.V2ObjectsObjectRecordsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**object** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ObjectsObjectRecordsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **matchingAttribute** | **string** |  | 
 **v2ObjectsObjectRecordsPutRequest** | [**V2ObjectsObjectRecordsPutRequest**](V2ObjectsObjectRecordsPutRequest.md) |  | 

### Return type

[**V2ObjectsObjectRecordsPut200Response**](V2ObjectsObjectRecordsPut200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ObjectsObjectRecordsQueryPost

> V2ObjectsObjectRecordsQueryPost200Response V2ObjectsObjectRecordsQueryPost(ctx, object).V2ObjectsObjectRecordsQueryPostRequest(v2ObjectsObjectRecordsQueryPostRequest).Execute()

List records



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/fastah/libattio"
)

func main() {
	object := "people" // string | 
	v2ObjectsObjectRecordsQueryPostRequest := *libattio.NewV2ObjectsObjectRecordsQueryPostRequest() // V2ObjectsObjectRecordsQueryPostRequest | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.V2ObjectsObjectRecordsQueryPost(context.Background(), object).V2ObjectsObjectRecordsQueryPostRequest(v2ObjectsObjectRecordsQueryPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.V2ObjectsObjectRecordsQueryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ObjectsObjectRecordsQueryPost`: V2ObjectsObjectRecordsQueryPost200Response
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.V2ObjectsObjectRecordsQueryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**object** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ObjectsObjectRecordsQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2ObjectsObjectRecordsQueryPostRequest** | [**V2ObjectsObjectRecordsQueryPostRequest**](V2ObjectsObjectRecordsQueryPostRequest.md) |  | 

### Return type

[**V2ObjectsObjectRecordsQueryPost200Response**](V2ObjectsObjectRecordsQueryPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ObjectsObjectRecordsRecordIdAttributesAttributeValuesGet

> V2ObjectsObjectRecordsRecordIdAttributesAttributeValuesGet200Response V2ObjectsObjectRecordsRecordIdAttributesAttributeValuesGet(ctx, object, recordId, attribute).ShowHistoric(showHistoric).Limit(limit).Offset(offset).Execute()

List record attribute values



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/fastah/libattio"
)

func main() {
	object := "people" // string | 
	recordId := "891dcbfc-9141-415d-9b2a-2238a6cc012d" // string | 
	attribute := "41252299-f8c7-4b5e-99c9-4ff8321d2f96" // string | 
	showHistoric := true // bool |  (optional) (default to false)
	limit := int32(10) // int32 |  (optional)
	offset := int32(5) // int32 |  (optional)

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.V2ObjectsObjectRecordsRecordIdAttributesAttributeValuesGet(context.Background(), object, recordId, attribute).ShowHistoric(showHistoric).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.V2ObjectsObjectRecordsRecordIdAttributesAttributeValuesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ObjectsObjectRecordsRecordIdAttributesAttributeValuesGet`: V2ObjectsObjectRecordsRecordIdAttributesAttributeValuesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.V2ObjectsObjectRecordsRecordIdAttributesAttributeValuesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**object** | **string** |  | 
**recordId** | **string** |  | 
**attribute** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ObjectsObjectRecordsRecordIdAttributesAttributeValuesGetRequest struct via the builder pattern


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


## V2ObjectsObjectRecordsRecordIdDelete

> map[string]interface{} V2ObjectsObjectRecordsRecordIdDelete(ctx, object, recordId).Execute()

Delete a record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/fastah/libattio"
)

func main() {
	object := "people" // string | 
	recordId := "891dcbfc-9141-415d-9b2a-2238a6cc012d" // string | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.V2ObjectsObjectRecordsRecordIdDelete(context.Background(), object, recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.V2ObjectsObjectRecordsRecordIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ObjectsObjectRecordsRecordIdDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.V2ObjectsObjectRecordsRecordIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**object** | **string** |  | 
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ObjectsObjectRecordsRecordIdDeleteRequest struct via the builder pattern


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


## V2ObjectsObjectRecordsRecordIdEntriesGet

> V2ObjectsObjectRecordsRecordIdEntriesGet200Response V2ObjectsObjectRecordsRecordIdEntriesGet(ctx, object, recordId).Limit(limit).Offset(offset).Execute()

List record entries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/fastah/libattio"
)

func main() {
	object := "people" // string | 
	recordId := "891dcbfc-9141-415d-9b2a-2238a6cc012d" // string | 
	limit := int32(10) // int32 |  (optional)
	offset := int32(5) // int32 |  (optional)

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.V2ObjectsObjectRecordsRecordIdEntriesGet(context.Background(), object, recordId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.V2ObjectsObjectRecordsRecordIdEntriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ObjectsObjectRecordsRecordIdEntriesGet`: V2ObjectsObjectRecordsRecordIdEntriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.V2ObjectsObjectRecordsRecordIdEntriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**object** | **string** |  | 
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ObjectsObjectRecordsRecordIdEntriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**V2ObjectsObjectRecordsRecordIdEntriesGet200Response**](V2ObjectsObjectRecordsRecordIdEntriesGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ObjectsObjectRecordsRecordIdGet

> V2ObjectsObjectRecordsPut200Response V2ObjectsObjectRecordsRecordIdGet(ctx, object, recordId).Execute()

Get a record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/fastah/libattio"
)

func main() {
	object := "people" // string | 
	recordId := "891dcbfc-9141-415d-9b2a-2238a6cc012d" // string | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.V2ObjectsObjectRecordsRecordIdGet(context.Background(), object, recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.V2ObjectsObjectRecordsRecordIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ObjectsObjectRecordsRecordIdGet`: V2ObjectsObjectRecordsPut200Response
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.V2ObjectsObjectRecordsRecordIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**object** | **string** |  | 
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ObjectsObjectRecordsRecordIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V2ObjectsObjectRecordsPut200Response**](V2ObjectsObjectRecordsPut200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ObjectsObjectRecordsRecordIdPatch

> V2ObjectsObjectRecordsPut200Response V2ObjectsObjectRecordsRecordIdPatch(ctx, object, recordId).V2ObjectsObjectRecordsRecordIdPatchRequest(v2ObjectsObjectRecordsRecordIdPatchRequest).Execute()

Update a record (append multiselect values)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/fastah/libattio"
)

func main() {
	object := "people" // string | 
	recordId := "891dcbfc-9141-415d-9b2a-2238a6cc012d" // string | 
	v2ObjectsObjectRecordsRecordIdPatchRequest := *libattio.NewV2ObjectsObjectRecordsRecordIdPatchRequest(*libattio.NewV2ObjectsObjectRecordsRecordIdPatchRequestData(map[string][]interface{}{"key": []interface{}{nil}})) // V2ObjectsObjectRecordsRecordIdPatchRequest | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.V2ObjectsObjectRecordsRecordIdPatch(context.Background(), object, recordId).V2ObjectsObjectRecordsRecordIdPatchRequest(v2ObjectsObjectRecordsRecordIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.V2ObjectsObjectRecordsRecordIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ObjectsObjectRecordsRecordIdPatch`: V2ObjectsObjectRecordsPut200Response
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.V2ObjectsObjectRecordsRecordIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**object** | **string** |  | 
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ObjectsObjectRecordsRecordIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2ObjectsObjectRecordsRecordIdPatchRequest** | [**V2ObjectsObjectRecordsRecordIdPatchRequest**](V2ObjectsObjectRecordsRecordIdPatchRequest.md) |  | 

### Return type

[**V2ObjectsObjectRecordsPut200Response**](V2ObjectsObjectRecordsPut200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ObjectsObjectRecordsRecordIdPut

> V2ObjectsObjectRecordsPut200Response V2ObjectsObjectRecordsRecordIdPut(ctx, object, recordId).V2ObjectsObjectRecordsPutRequest(v2ObjectsObjectRecordsPutRequest).Execute()

Update a record (overwrite multiselect values)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/fastah/libattio"
)

func main() {
	object := "people" // string | 
	recordId := "891dcbfc-9141-415d-9b2a-2238a6cc012d" // string | 
	v2ObjectsObjectRecordsPutRequest := *libattio.NewV2ObjectsObjectRecordsPutRequest(*libattio.NewV2ObjectsObjectRecordsPutRequestData(map[string][]interface{}{"key": []interface{}{nil}})) // V2ObjectsObjectRecordsPutRequest | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.V2ObjectsObjectRecordsRecordIdPut(context.Background(), object, recordId).V2ObjectsObjectRecordsPutRequest(v2ObjectsObjectRecordsPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.V2ObjectsObjectRecordsRecordIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ObjectsObjectRecordsRecordIdPut`: V2ObjectsObjectRecordsPut200Response
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.V2ObjectsObjectRecordsRecordIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**object** | **string** |  | 
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ObjectsObjectRecordsRecordIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2ObjectsObjectRecordsPutRequest** | [**V2ObjectsObjectRecordsPutRequest**](V2ObjectsObjectRecordsPutRequest.md) |  | 

### Return type

[**V2ObjectsObjectRecordsPut200Response**](V2ObjectsObjectRecordsPut200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

