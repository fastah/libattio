# InputValueAnyOf15

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | A timestamp value represents a single, universal moment in time using an ISO 8601 formatted string. This means that a timestamp consists of a date, a time (with nanosecond precision), and a time zone. Attio will coerce timestamps which do not provide full nanosecond precision and UTC is assumed if no time zone is provided. For example, \&quot;2023\&quot;, \&quot;2023-01\&quot;, \&quot;2023-01-02\&quot;, \&quot;2023-01-02T13:00\&quot;, \&quot;2023-01-02T13:00:00\&quot;, and \&quot;2023-01-02T13:00:00.000000000\&quot; will all be coerced to \&quot;2023-01-02T13:00:00.000000000Z\&quot;. Timestamps are always returned in UTC. For example, writing a timestamp value using the string \&quot;2023-01-02T13:00:00.000000000+02:00\&quot; will result in the value \&quot;2023-01-02T11:00:00.000000000Z\&quot; being returned. The maximum date is \&quot;9999-12-31T23:59:59.999999999Z\&quot;. | 

## Methods

### NewInputValueAnyOf15

`func NewInputValueAnyOf15(value string, ) *InputValueAnyOf15`

NewInputValueAnyOf15 instantiates a new InputValueAnyOf15 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputValueAnyOf15WithDefaults

`func NewInputValueAnyOf15WithDefaults() *InputValueAnyOf15`

NewInputValueAnyOf15WithDefaults instantiates a new InputValueAnyOf15 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *InputValueAnyOf15) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InputValueAnyOf15) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InputValueAnyOf15) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


