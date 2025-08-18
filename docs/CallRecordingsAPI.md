# \CallRecordingsAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2MeetingsMeetingIdCallRecordingsCallRecordingIdDelete**](CallRecordingsAPI.md#V2MeetingsMeetingIdCallRecordingsCallRecordingIdDelete) | **Delete** /v2/meetings/{meeting_id}/call_recordings/{call_recording_id} | Delete call recording
[**V2MeetingsMeetingIdCallRecordingsCallRecordingIdGet**](CallRecordingsAPI.md#V2MeetingsMeetingIdCallRecordingsCallRecordingIdGet) | **Get** /v2/meetings/{meeting_id}/call_recordings/{call_recording_id} | Get call recording
[**V2MeetingsMeetingIdCallRecordingsGet**](CallRecordingsAPI.md#V2MeetingsMeetingIdCallRecordingsGet) | **Get** /v2/meetings/{meeting_id}/call_recordings | List call recordings
[**V2MeetingsMeetingIdCallRecordingsPost**](CallRecordingsAPI.md#V2MeetingsMeetingIdCallRecordingsPost) | **Post** /v2/meetings/{meeting_id}/call_recordings | Create call recording



## V2MeetingsMeetingIdCallRecordingsCallRecordingIdDelete

> map[string]interface{} V2MeetingsMeetingIdCallRecordingsCallRecordingIdDelete(ctx, meetingId, callRecordingId).Execute()

Delete call recording



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
	callRecordingId := "e8f2a3b7-9b4d-4c5e-8a1f-3d7b2c5e8f9a" // string | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.CallRecordingsAPI.V2MeetingsMeetingIdCallRecordingsCallRecordingIdDelete(context.Background(), meetingId, callRecordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingsAPI.V2MeetingsMeetingIdCallRecordingsCallRecordingIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2MeetingsMeetingIdCallRecordingsCallRecordingIdDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingsAPI.V2MeetingsMeetingIdCallRecordingsCallRecordingIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingId** | **string** |  | 
**callRecordingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MeetingsMeetingIdCallRecordingsCallRecordingIdDeleteRequest struct via the builder pattern


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


## V2MeetingsMeetingIdCallRecordingsCallRecordingIdGet

> V2MeetingsMeetingIdCallRecordingsPost200Response V2MeetingsMeetingIdCallRecordingsCallRecordingIdGet(ctx, meetingId, callRecordingId).Execute()

Get call recording



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
	callRecordingId := "e8f2a3b7-9b4d-4c5e-8a1f-3d7b2c5e8f9a" // string | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.CallRecordingsAPI.V2MeetingsMeetingIdCallRecordingsCallRecordingIdGet(context.Background(), meetingId, callRecordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingsAPI.V2MeetingsMeetingIdCallRecordingsCallRecordingIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2MeetingsMeetingIdCallRecordingsCallRecordingIdGet`: V2MeetingsMeetingIdCallRecordingsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingsAPI.V2MeetingsMeetingIdCallRecordingsCallRecordingIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingId** | **string** |  | 
**callRecordingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MeetingsMeetingIdCallRecordingsCallRecordingIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V2MeetingsMeetingIdCallRecordingsPost200Response**](V2MeetingsMeetingIdCallRecordingsPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MeetingsMeetingIdCallRecordingsGet

> V2MeetingsMeetingIdCallRecordingsGet200Response V2MeetingsMeetingIdCallRecordingsGet(ctx, meetingId).Limit(limit).Cursor(cursor).Execute()

List call recordings



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
	limit := int32(50) // int32 |  (optional)
	cursor := "eyJkZXNjcmlwdGlvbiI6ICJ0aGlzIGlzIGEgY3Vyc29yIn0=.eM56CGbqZ6G1NHiJchTIkH4vKDr" // string |  (optional)

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.CallRecordingsAPI.V2MeetingsMeetingIdCallRecordingsGet(context.Background(), meetingId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingsAPI.V2MeetingsMeetingIdCallRecordingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2MeetingsMeetingIdCallRecordingsGet`: V2MeetingsMeetingIdCallRecordingsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingsAPI.V2MeetingsMeetingIdCallRecordingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MeetingsMeetingIdCallRecordingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **cursor** | **string** |  | 

### Return type

[**V2MeetingsMeetingIdCallRecordingsGet200Response**](V2MeetingsMeetingIdCallRecordingsGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MeetingsMeetingIdCallRecordingsPost

> V2MeetingsMeetingIdCallRecordingsPost200Response V2MeetingsMeetingIdCallRecordingsPost(ctx, meetingId).V2MeetingsMeetingIdCallRecordingsPostRequest(v2MeetingsMeetingIdCallRecordingsPostRequest).Execute()

Create call recording



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
	v2MeetingsMeetingIdCallRecordingsPostRequest := *libattio.NewV2MeetingsMeetingIdCallRecordingsPostRequest("https://example.com/recording.mp4") // V2MeetingsMeetingIdCallRecordingsPostRequest | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.CallRecordingsAPI.V2MeetingsMeetingIdCallRecordingsPost(context.Background(), meetingId).V2MeetingsMeetingIdCallRecordingsPostRequest(v2MeetingsMeetingIdCallRecordingsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallRecordingsAPI.V2MeetingsMeetingIdCallRecordingsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2MeetingsMeetingIdCallRecordingsPost`: V2MeetingsMeetingIdCallRecordingsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `CallRecordingsAPI.V2MeetingsMeetingIdCallRecordingsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MeetingsMeetingIdCallRecordingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2MeetingsMeetingIdCallRecordingsPostRequest** | [**V2MeetingsMeetingIdCallRecordingsPostRequest**](V2MeetingsMeetingIdCallRecordingsPostRequest.md) |  | 

### Return type

[**V2MeetingsMeetingIdCallRecordingsPost200Response**](V2MeetingsMeetingIdCallRecordingsPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

