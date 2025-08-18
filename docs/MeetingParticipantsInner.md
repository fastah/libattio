# MeetingParticipantsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of the individual meeting participant. | 
**IsOrganizer** | **bool** | Whether or not the participant is the organizer of the meeting. | 
**EmailAddress** | **NullableString** | The normalized email address of the meeting participant. | 

## Methods

### NewMeetingParticipantsInner

`func NewMeetingParticipantsInner(status string, isOrganizer bool, emailAddress NullableString, ) *MeetingParticipantsInner`

NewMeetingParticipantsInner instantiates a new MeetingParticipantsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingParticipantsInnerWithDefaults

`func NewMeetingParticipantsInnerWithDefaults() *MeetingParticipantsInner`

NewMeetingParticipantsInnerWithDefaults instantiates a new MeetingParticipantsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MeetingParticipantsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MeetingParticipantsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MeetingParticipantsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIsOrganizer

`func (o *MeetingParticipantsInner) GetIsOrganizer() bool`

GetIsOrganizer returns the IsOrganizer field if non-nil, zero value otherwise.

### GetIsOrganizerOk

`func (o *MeetingParticipantsInner) GetIsOrganizerOk() (*bool, bool)`

GetIsOrganizerOk returns a tuple with the IsOrganizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrganizer

`func (o *MeetingParticipantsInner) SetIsOrganizer(v bool)`

SetIsOrganizer sets IsOrganizer field to given value.


### GetEmailAddress

`func (o *MeetingParticipantsInner) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *MeetingParticipantsInner) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *MeetingParticipantsInner) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### SetEmailAddressNil

`func (o *MeetingParticipantsInner) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *MeetingParticipantsInner) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


