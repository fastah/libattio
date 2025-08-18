# MeetingStartAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | If a non-all day event, a datetime representing when the meeting starts. Datetimes are formatted as UTC if no timezone is available. If a timezone is available, the datetime will offset using the specified timezone. | 
**Timezone** | **NullableString** | The IANA timezone in which the meeting starts, if available. | 

## Methods

### NewMeetingStartAnyOf

`func NewMeetingStartAnyOf(datetime string, timezone NullableString, ) *MeetingStartAnyOf`

NewMeetingStartAnyOf instantiates a new MeetingStartAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingStartAnyOfWithDefaults

`func NewMeetingStartAnyOfWithDefaults() *MeetingStartAnyOf`

NewMeetingStartAnyOfWithDefaults instantiates a new MeetingStartAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *MeetingStartAnyOf) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *MeetingStartAnyOf) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *MeetingStartAnyOf) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetTimezone

`func (o *MeetingStartAnyOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MeetingStartAnyOf) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MeetingStartAnyOf) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### SetTimezoneNil

`func (o *MeetingStartAnyOf) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *MeetingStartAnyOf) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


