# V2MeetingsMeetingIdCallRecordingsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoUrl** | **string** | A publicly accessible URL to a video file of the call recording. Attio will download the video from this URL asynchronously.  **Requirements:** - **Protocol:** The URL must use the &#x60;https&#x60; protocol. - **File type:** The file must be a &#x60;.mp4&#x60; file. - **File size:** The file must not exceed 30MB in size. - **Accessibility:** For the request to be accepted, the URL must be publicly accessible. Attio will make a &#x60;HEAD&#x60; request to the URL to verify its accessibility and retrieve file metadata. The response to this request must include a &#x60;Content-Length&#x60; header. | 

## Methods

### NewV2MeetingsMeetingIdCallRecordingsPostRequest

`func NewV2MeetingsMeetingIdCallRecordingsPostRequest(videoUrl string, ) *V2MeetingsMeetingIdCallRecordingsPostRequest`

NewV2MeetingsMeetingIdCallRecordingsPostRequest instantiates a new V2MeetingsMeetingIdCallRecordingsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MeetingsMeetingIdCallRecordingsPostRequestWithDefaults

`func NewV2MeetingsMeetingIdCallRecordingsPostRequestWithDefaults() *V2MeetingsMeetingIdCallRecordingsPostRequest`

NewV2MeetingsMeetingIdCallRecordingsPostRequestWithDefaults instantiates a new V2MeetingsMeetingIdCallRecordingsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideoUrl

`func (o *V2MeetingsMeetingIdCallRecordingsPostRequest) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *V2MeetingsMeetingIdCallRecordingsPostRequest) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *V2MeetingsMeetingIdCallRecordingsPostRequest) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


