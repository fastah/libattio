# \TranscriptsAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptDelete**](TranscriptsAPI.md#V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptDelete) | **Delete** /v2/meetings/{meeting_id}/call_recordings/{call_recording_id}/transcript | Delete call recording transcript
[**V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet**](TranscriptsAPI.md#V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet) | **Get** /v2/meetings/{meeting_id}/call_recordings/{call_recording_id}/transcript | Get call transcript
[**V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPost**](TranscriptsAPI.md#V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPost) | **Post** /v2/meetings/{meeting_id}/call_recordings/{call_recording_id}/transcript | Create transcript



## V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptDelete

> map[string]interface{} V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptDelete(ctx, meetingId, callRecordingId).Execute()

Delete call recording transcript



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
	resp, r, err := apiClient.TranscriptsAPI.V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptDelete(context.Background(), meetingId, callRecordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranscriptsAPI.V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TranscriptsAPI.V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingId** | **string** |  | 
**callRecordingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptDeleteRequest struct via the builder pattern


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


## V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet

> V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200Response V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet(ctx, meetingId, callRecordingId).Cursor(cursor).Execute()

Get call transcript



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
	cursor := "eyJkZXNjcmlwdGlvbiI6ICJ0aGlzIGlzIGEgY3Vyc29yIn0=.eM56CGbqZ6G1NHiJchTIkH4vKDr" // string |  (optional)

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.TranscriptsAPI.V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet(context.Background(), meetingId, callRecordingId).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranscriptsAPI.V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet`: V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TranscriptsAPI.V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingId** | **string** |  | 
**callRecordingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** |  | 

### Return type

[**V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200Response**](V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPost

> V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPost200Response V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPost(ctx, meetingId, callRecordingId).V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequest(v2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequest).Execute()

Create transcript



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
	v2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequest := *libattio.NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequest(*libattio.NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestData([]libattio.V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestDataTranscriptInner{*libattio.NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestDataTranscriptInner("Hello everyone, welcome to the meeting.", float32(0.5123), float32(3.2123), *libattio.NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestDataTranscriptInnerSpeaker("Simon Mitchell"))})) // V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequest | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.TranscriptsAPI.V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPost(context.Background(), meetingId, callRecordingId).V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequest(v2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranscriptsAPI.V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPost`: V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TranscriptsAPI.V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingId** | **string** |  | 
**callRecordingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequest** | [**V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequest**](V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequest.md) |  | 

### Return type

[**V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPost200Response**](V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

