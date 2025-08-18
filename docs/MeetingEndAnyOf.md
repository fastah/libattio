# MeetingEndAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | A datetime representing when the meeting ends. All day meetings will return a date whereas non-all day meetings will return a datetime. Datetimes do not include timezone information; please refer to &#x60;timezone&#x60; for timezone information. Following iCalendar RFC 5545, the &#x60;end_at&#x60; property is exclusive, meaning that the meeting ends before the specified time, not at it. For example, a one day meeting on June 3rd would have an &#x60;end_at&#x60; of June 4th, not June 3rd; a one hour meeting starting at 14:00 would have an &#x60;end_at&#x60; of 15:00, not 14:00. | 
**Timezone** | **NullableString** | The IANA timezone in which the meeting ends, if available. | 

## Methods

### NewMeetingEndAnyOf

`func NewMeetingEndAnyOf(datetime string, timezone NullableString, ) *MeetingEndAnyOf`

NewMeetingEndAnyOf instantiates a new MeetingEndAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingEndAnyOfWithDefaults

`func NewMeetingEndAnyOfWithDefaults() *MeetingEndAnyOf`

NewMeetingEndAnyOfWithDefaults instantiates a new MeetingEndAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *MeetingEndAnyOf) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *MeetingEndAnyOf) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *MeetingEndAnyOf) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetTimezone

`func (o *MeetingEndAnyOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MeetingEndAnyOf) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MeetingEndAnyOf) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### SetTimezoneNil

`func (o *MeetingEndAnyOf) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *MeetingEndAnyOf) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


