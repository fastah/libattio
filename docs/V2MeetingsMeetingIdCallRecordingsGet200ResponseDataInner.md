# V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInnerId**](V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInnerId.md) |  | 
**Status** | **string** | The status of the call recording. When a call recording is first created, it will have a status of &#x60;PROCESSING&#x60;. Once the recording is ready, it will transition to &#x60;COMPLETED&#x60;. If the recording fails for any reason, the status will be &#x60;FAILED&#x60;. | 
**CreatedByActor** | [**V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInnerCreatedByActor**](V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInnerCreatedByActor.md) |  | 
**CreatedAt** | **string** | The timestamp of when the call recording was created. | 

## Methods

### NewV2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner

`func NewV2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner(id V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInnerId, status string, createdByActor V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInnerCreatedByActor, createdAt string, ) *V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner`

NewV2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner instantiates a new V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MeetingsMeetingIdCallRecordingsGet200ResponseDataInnerWithDefaults

`func NewV2MeetingsMeetingIdCallRecordingsGet200ResponseDataInnerWithDefaults() *V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner`

NewV2MeetingsMeetingIdCallRecordingsGet200ResponseDataInnerWithDefaults instantiates a new V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner) GetId() V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInnerId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner) GetIdOk() (*V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInnerId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner) SetId(v V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInnerId)`

SetId sets Id field to given value.


### GetStatus

`func (o *V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedByActor

`func (o *V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner) GetCreatedByActor() V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInnerCreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner) GetCreatedByActorOk() (*V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInnerCreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner) SetCreatedByActor(v V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInnerCreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.


### GetCreatedAt

`func (o *V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2MeetingsMeetingIdCallRecordingsGet200ResponseDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


