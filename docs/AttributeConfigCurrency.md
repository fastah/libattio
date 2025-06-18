# AttributeConfigCurrency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultCurrencyCode** | **NullableString** | The ISO4217 code representing the currency that values for this attribute should be stored in. | 
**DisplayType** | **NullableString** | How the currency should be displayed across the app. \&quot;code\&quot; will display the ISO currency code e.g. \&quot;USD\&quot;, \&quot;name\&quot; will display the localized currency name e.g. \&quot;British pound\&quot;, \&quot;narrowSymbol\&quot; will display \&quot;$1\&quot; instead of \&quot;US$1\&quot; and \&quot;symbol\&quot; will display a localized currency symbol such as \&quot;$\&quot;. | 

## Methods

### NewAttributeConfigCurrency

`func NewAttributeConfigCurrency(defaultCurrencyCode NullableString, displayType NullableString, ) *AttributeConfigCurrency`

NewAttributeConfigCurrency instantiates a new AttributeConfigCurrency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeConfigCurrencyWithDefaults

`func NewAttributeConfigCurrencyWithDefaults() *AttributeConfigCurrency`

NewAttributeConfigCurrencyWithDefaults instantiates a new AttributeConfigCurrency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultCurrencyCode

`func (o *AttributeConfigCurrency) GetDefaultCurrencyCode() string`

GetDefaultCurrencyCode returns the DefaultCurrencyCode field if non-nil, zero value otherwise.

### GetDefaultCurrencyCodeOk

`func (o *AttributeConfigCurrency) GetDefaultCurrencyCodeOk() (*string, bool)`

GetDefaultCurrencyCodeOk returns a tuple with the DefaultCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrencyCode

`func (o *AttributeConfigCurrency) SetDefaultCurrencyCode(v string)`

SetDefaultCurrencyCode sets DefaultCurrencyCode field to given value.


### SetDefaultCurrencyCodeNil

`func (o *AttributeConfigCurrency) SetDefaultCurrencyCodeNil(b bool)`

 SetDefaultCurrencyCodeNil sets the value for DefaultCurrencyCode to be an explicit nil

### UnsetDefaultCurrencyCode
`func (o *AttributeConfigCurrency) UnsetDefaultCurrencyCode()`

UnsetDefaultCurrencyCode ensures that no value is present for DefaultCurrencyCode, not even an explicit nil
### GetDisplayType

`func (o *AttributeConfigCurrency) GetDisplayType() string`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *AttributeConfigCurrency) GetDisplayTypeOk() (*string, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *AttributeConfigCurrency) SetDisplayType(v string)`

SetDisplayType sets DisplayType field to given value.


### SetDisplayTypeNil

`func (o *AttributeConfigCurrency) SetDisplayTypeNil(b bool)`

 SetDisplayTypeNil sets the value for DisplayType to be an explicit nil

### UnsetDisplayType
`func (o *AttributeConfigCurrency) UnsetDisplayType()`

UnsetDisplayType ensures that no value is present for DisplayType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


