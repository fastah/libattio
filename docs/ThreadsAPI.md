# \ThreadsAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2ThreadsGet**](ThreadsAPI.md#V2ThreadsGet) | **Get** /v2/threads | List threads
[**V2ThreadsThreadIdGet**](ThreadsAPI.md#V2ThreadsThreadIdGet) | **Get** /v2/threads/{thread_id} | Get a thread



## V2ThreadsGet

> V2ThreadsGet200Response V2ThreadsGet(ctx).RecordId(recordId).Object(object).EntryId(entryId).List(list).Limit(limit).Offset(offset).Execute()

List threads



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
	recordId := "891dcbfc-9141-415d-9b2a-2238a6cc012d" // string |  (optional)
	object := "people" // string |  (optional)
	entryId := "2e6e29ea-c4e0-4f44-842d-78a891f8c156" // string |  (optional)
	list := "33ebdbe9-e529-47c9-b894-0ba25e9c15c0" // string |  (optional)
	limit := int32(10) // int32 |  (optional)
	offset := int32(5) // int32 |  (optional)

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.V2ThreadsGet(context.Background()).RecordId(recordId).Object(object).EntryId(entryId).List(list).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.V2ThreadsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ThreadsGet`: V2ThreadsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.V2ThreadsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ThreadsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordId** | **string** |  | 
 **object** | **string** |  | 
 **entryId** | **string** |  | 
 **list** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**V2ThreadsGet200Response**](V2ThreadsGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ThreadsThreadIdGet

> V2ThreadsThreadIdGet200Response V2ThreadsThreadIdGet(ctx, threadId).Execute()

Get a thread



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
	threadId := "a649e4d9-435c-43fb-83ba-847b4876f27a" // string | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreadsAPI.V2ThreadsThreadIdGet(context.Background(), threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreadsAPI.V2ThreadsThreadIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ThreadsThreadIdGet`: V2ThreadsThreadIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ThreadsAPI.V2ThreadsThreadIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ThreadsThreadIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2ThreadsThreadIdGet200Response**](V2ThreadsThreadIdGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

