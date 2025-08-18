# V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transcript** | [**[]V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestDataTranscriptInner**](V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestDataTranscriptInner.md) | Array of transcript segments with speech, timing, and speaker information. A maximum of 4000 transcript segments and 50 unique speakers are allowed per transcript. | 

## Methods

### NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestData

`func NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestData(transcript []V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestDataTranscriptInner, ) *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestData`

NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestData instantiates a new V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestDataWithDefaults

`func NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestDataWithDefaults() *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestData`

NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestDataWithDefaults instantiates a new V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTranscript

`func (o *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestData) GetTranscript() []V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestDataTranscriptInner`

GetTranscript returns the Transcript field if non-nil, zero value otherwise.

### GetTranscriptOk

`func (o *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestData) GetTranscriptOk() (*[]V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestDataTranscriptInner, bool)`

GetTranscriptOk returns a tuple with the Transcript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscript

`func (o *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestData) SetTranscript(v []V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptPostRequestDataTranscriptInner)`

SetTranscript sets Transcript field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


