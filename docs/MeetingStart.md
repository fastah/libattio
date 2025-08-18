# MeetingStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | If a non-all day event, a datetime representing when the meeting starts. Datetimes are formatted as UTC if no timezone is available. If a timezone is available, the datetime will offset using the specified timezone. | 
**Timezone** | **string** | The IANA timezone in which the meeting starts, if available. | 
**Date** | **string** | If an all day event, a date representing when the meeting starts. | 

## Methods

### NewMeetingStart

`func NewMeetingStart(datetime string, timezone string, date string, ) *MeetingStart`

NewMeetingStart instantiates a new MeetingStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingStartWithDefaults

`func NewMeetingStartWithDefaults() *MeetingStart`

NewMeetingStartWithDefaults instantiates a new MeetingStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *MeetingStart) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *MeetingStart) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *MeetingStart) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetTimezone

`func (o *MeetingStart) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MeetingStart) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MeetingStart) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetDate

`func (o *MeetingStart) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MeetingStart) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MeetingStart) SetDate(v string)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


