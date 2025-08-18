# \MeetingsAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2MeetingsGet**](MeetingsAPI.md#V2MeetingsGet) | **Get** /v2/meetings | List meetings
[**V2MeetingsMeetingIdGet**](MeetingsAPI.md#V2MeetingsMeetingIdGet) | **Get** /v2/meetings/{meeting_id} | Get a meeting



## V2MeetingsGet

> V2MeetingsGet200Response V2MeetingsGet(ctx).Limit(limit).Cursor(cursor).LinkedObject(linkedObject).LinkedRecordId(linkedRecordId).Participants(participants).Sort(sort).EndsFrom(endsFrom).StartsBefore(startsBefore).Timezone(timezone).Execute()

List meetings



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
	limit := int32(50) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string |  (optional)
	linkedObject := "linkedObject_example" // string |  (optional)
	linkedRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	participants := "participants_example" // string |  (optional) (default to "")
	sort := "sort_example" // string |  (optional) (default to "start_asc")
	endsFrom := "endsFrom_example" // string |  (optional)
	startsBefore := "startsBefore_example" // string |  (optional)
	timezone := "timezone_example" // string |  (optional) (default to "UTC")

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.MeetingsAPI.V2MeetingsGet(context.Background()).Limit(limit).Cursor(cursor).LinkedObject(linkedObject).LinkedRecordId(linkedRecordId).Participants(participants).Sort(sort).EndsFrom(endsFrom).StartsBefore(startsBefore).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingsAPI.V2MeetingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2MeetingsGet`: V2MeetingsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MeetingsAPI.V2MeetingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2MeetingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** |  | 
 **linkedObject** | **string** |  | 
 **linkedRecordId** | **string** |  | 
 **participants** | **string** |  | [default to &quot;&quot;]
 **sort** | **string** |  | [default to &quot;start_asc&quot;]
 **endsFrom** | **string** |  | 
 **startsBefore** | **string** |  | 
 **timezone** | **string** |  | [default to &quot;UTC&quot;]

### Return type

[**V2MeetingsGet200Response**](V2MeetingsGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MeetingsMeetingIdGet

> V2MeetingsMeetingIdGet200Response V2MeetingsMeetingIdGet(ctx, meetingId).Execute()

Get a meeting



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
	meetingId := "cb59ab17-ad15-460c-a126-0715617c0853" // string | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.MeetingsAPI.V2MeetingsMeetingIdGet(context.Background(), meetingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeetingsAPI.V2MeetingsMeetingIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2MeetingsMeetingIdGet`: V2MeetingsMeetingIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MeetingsAPI.V2MeetingsMeetingIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MeetingsMeetingIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2MeetingsMeetingIdGet200Response**](V2MeetingsMeetingIdGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

