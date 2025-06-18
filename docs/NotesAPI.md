# \NotesAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2NotesGet**](NotesAPI.md#V2NotesGet) | **Get** /v2/notes | List notes
[**V2NotesNoteIdDelete**](NotesAPI.md#V2NotesNoteIdDelete) | **Delete** /v2/notes/{note_id} | Delete a note
[**V2NotesNoteIdGet**](NotesAPI.md#V2NotesNoteIdGet) | **Get** /v2/notes/{note_id} | Get a note
[**V2NotesPost**](NotesAPI.md#V2NotesPost) | **Post** /v2/notes | Create a note



## V2NotesGet

> V2NotesGet200Response V2NotesGet(ctx).Limit(limit).Offset(offset).ParentObject(parentObject).ParentRecordId(parentRecordId).Execute()

List notes



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
	limit := int32(10) // int32 |  (optional)
	offset := int32(5) // int32 |  (optional)
	parentObject := "people" // string |  (optional)
	parentRecordId := "891dcbfc-9141-415d-9b2a-2238a6cc012d" // string |  (optional)

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.V2NotesGet(context.Background()).Limit(limit).Offset(offset).ParentObject(parentObject).ParentRecordId(parentRecordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.V2NotesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2NotesGet`: V2NotesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.V2NotesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2NotesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **parentObject** | **string** |  | 
 **parentRecordId** | **string** |  | 

### Return type

[**V2NotesGet200Response**](V2NotesGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2NotesNoteIdDelete

> map[string]interface{} V2NotesNoteIdDelete(ctx, noteId).Execute()

Delete a note



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
	noteId := "ff3f3bd4-40f4-4f80-8187-cd02385af424" // string | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.V2NotesNoteIdDelete(context.Background(), noteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.V2NotesNoteIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2NotesNoteIdDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.V2NotesNoteIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2NotesNoteIdDeleteRequest struct via the builder pattern


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


## V2NotesNoteIdGet

> V2NotesPost200Response V2NotesNoteIdGet(ctx, noteId).Execute()

Get a note



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
	noteId := "ff3f3bd4-40f4-4f80-8187-cd02385af424" // string | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.V2NotesNoteIdGet(context.Background(), noteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.V2NotesNoteIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2NotesNoteIdGet`: V2NotesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.V2NotesNoteIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2NotesNoteIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2NotesPost200Response**](V2NotesPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2NotesPost

> V2NotesPost200Response V2NotesPost(ctx).V2NotesPostRequest(v2NotesPostRequest).Execute()

Create a note



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
	v2NotesPostRequest := *libattio.NewV2NotesPostRequest(*libattio.NewV2NotesPostRequestData("people", "891dcbfc-9141-415d-9b2a-2238a6cc012d", "Initial Prospecting Call Summary", "Format_example", "# Meeting Recap: Q4 Planning

**Date:** 2023-10-26
**Attendees:** Alex, Jamie, Casey

## Key Discussion Points

- Reviewed Q3 performance metrics.
- Brainstormed key initiatives for Q4.
- Discussed budget allocation for ==Project Phoenix==.

## Action Items

1. Alex to finalize Q4 roadmap by EOD Friday.
2. Jamie to schedule follow-up with [Marketing Team](https://app.attio.com/teams/marketing).
3. Casey to draft initial budget for ~~Project Chimera~~ (now deferred).

*Next steps: Review draft roadmap next week.*")) // V2NotesPostRequest | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.V2NotesPost(context.Background()).V2NotesPostRequest(v2NotesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.V2NotesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2NotesPost`: V2NotesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.V2NotesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2NotesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2NotesPostRequest** | [**V2NotesPostRequest**](V2NotesPostRequest.md) |  | 

### Return type

[**V2NotesPost200Response**](V2NotesPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

