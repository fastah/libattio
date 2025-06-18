# OutputValueAnyOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeType** | **string** | The attribute type of the value. | 
**Value** | **string** | A date represents a single calendar year, month and day, independent of timezone. If hours, months, seconds or timezones are provided, they will be trimmed. For example, \&quot;2023\&quot; and \&quot;2023-01\&quot; will be coerced into \&quot;2023-01-01\&quot;, and \&quot;2023-01-02\&quot;, \&quot;2023-01-02T13:00\&quot;, \&quot;2023-01-02T14:00:00\&quot;, \&quot;2023-01-02T15:00:00.000000000\&quot;, and \&quot;2023-01-02T15:00:00.000000000+02:00\&quot; will all be coerced to \&quot;2023-01-02\&quot;. If a timezone is provided that would result in a different calendar date in UTC, the date will be coerced to UTC and then the timezone component will be trimmed. For example, the value \&quot;2023-01-02T23:00:00-10:00\&quot; will be returned as \&quot;2023-01-03\&quot;. The maximum date is \&quot;9999-12-31\&quot;. | 

## Methods

### NewOutputValueAnyOf3

`func NewOutputValueAnyOf3(attributeType string, value string, ) *OutputValueAnyOf3`

NewOutputValueAnyOf3 instantiates a new OutputValueAnyOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputValueAnyOf3WithDefaults

`func NewOutputValueAnyOf3WithDefaults() *OutputValueAnyOf3`

NewOutputValueAnyOf3WithDefaults instantiates a new OutputValueAnyOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeType

`func (o *OutputValueAnyOf3) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *OutputValueAnyOf3) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *OutputValueAnyOf3) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetValue

`func (o *OutputValueAnyOf3) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OutputValueAnyOf3) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OutputValueAnyOf3) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


