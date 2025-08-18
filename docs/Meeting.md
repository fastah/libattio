# Meeting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**MeetingId**](MeetingId.md) |  | 
**Title** | **string** | The title of the meeting. | 
**Description** | **string** | The description of the meeting. | 
**IsAllDay** | **bool** | Whether or not the meeting is an all day event. All day events may span multiple days. | 
**Start** | [**MeetingStart**](MeetingStart.md) |  | 
**End** | [**MeetingEnd**](MeetingEnd.md) |  | 
**Participants** | [**[]MeetingParticipantsInner**](MeetingParticipantsInner.md) |  | 
**LinkedRecords** | [**[]MeetingLinkedRecordsInner**](MeetingLinkedRecordsInner.md) | A list of records that are linked to the meeting. Participants with matching person records are automatically linked to the meeting but other records may also be linked explicitly. | 
**CreatedAt** | **string** | Timestamp representing when the meeting was created. | 
**CreatedByActor** | [**MeetingCreatedByActor**](MeetingCreatedByActor.md) |  | 

## Methods

### NewMeeting

`func NewMeeting(id MeetingId, title string, description string, isAllDay bool, start MeetingStart, end MeetingEnd, participants []MeetingParticipantsInner, linkedRecords []MeetingLinkedRecordsInner, createdAt string, createdByActor MeetingCreatedByActor, ) *Meeting`

NewMeeting instantiates a new Meeting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingWithDefaults

`func NewMeetingWithDefaults() *Meeting`

NewMeetingWithDefaults instantiates a new Meeting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Meeting) GetId() MeetingId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Meeting) GetIdOk() (*MeetingId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Meeting) SetId(v MeetingId)`

SetId sets Id field to given value.


### GetTitle

`func (o *Meeting) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Meeting) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Meeting) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Meeting) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Meeting) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Meeting) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsAllDay

`func (o *Meeting) GetIsAllDay() bool`

GetIsAllDay returns the IsAllDay field if non-nil, zero value otherwise.

### GetIsAllDayOk

`func (o *Meeting) GetIsAllDayOk() (*bool, bool)`

GetIsAllDayOk returns a tuple with the IsAllDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllDay

`func (o *Meeting) SetIsAllDay(v bool)`

SetIsAllDay sets IsAllDay field to given value.


### GetStart

`func (o *Meeting) GetStart() MeetingStart`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Meeting) GetStartOk() (*MeetingStart, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Meeting) SetStart(v MeetingStart)`

SetStart sets Start field to given value.


### GetEnd

`func (o *Meeting) GetEnd() MeetingEnd`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Meeting) GetEndOk() (*MeetingEnd, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Meeting) SetEnd(v MeetingEnd)`

SetEnd sets End field to given value.


### GetParticipants

`func (o *Meeting) GetParticipants() []MeetingParticipantsInner`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *Meeting) GetParticipantsOk() (*[]MeetingParticipantsInner, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *Meeting) SetParticipants(v []MeetingParticipantsInner)`

SetParticipants sets Participants field to given value.


### GetLinkedRecords

`func (o *Meeting) GetLinkedRecords() []MeetingLinkedRecordsInner`

GetLinkedRecords returns the LinkedRecords field if non-nil, zero value otherwise.

### GetLinkedRecordsOk

`func (o *Meeting) GetLinkedRecordsOk() (*[]MeetingLinkedRecordsInner, bool)`

GetLinkedRecordsOk returns a tuple with the LinkedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedRecords

`func (o *Meeting) SetLinkedRecords(v []MeetingLinkedRecordsInner)`

SetLinkedRecords sets LinkedRecords field to given value.


### GetCreatedAt

`func (o *Meeting) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Meeting) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Meeting) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedByActor

`func (o *Meeting) GetCreatedByActor() MeetingCreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *Meeting) GetCreatedByActorOk() (*MeetingCreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *Meeting) SetCreatedByActor(v MeetingCreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


