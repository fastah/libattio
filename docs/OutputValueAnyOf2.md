# OutputValueAnyOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyValue** | **float32** | A numerical representation of the currency value. A decimal with a max of 4 decimal places. | 
**CurrencyCode** | Pointer to **NullableString** | The ISO4217 currency code representing the currency that the value is stored in. | [optional] 
**AttributeType** | **string** | The attribute type of the value. | 

## Methods

### NewOutputValueAnyOf2

`func NewOutputValueAnyOf2(currencyValue float32, attributeType string, ) *OutputValueAnyOf2`

NewOutputValueAnyOf2 instantiates a new OutputValueAnyOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputValueAnyOf2WithDefaults

`func NewOutputValueAnyOf2WithDefaults() *OutputValueAnyOf2`

NewOutputValueAnyOf2WithDefaults instantiates a new OutputValueAnyOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyValue

`func (o *OutputValueAnyOf2) GetCurrencyValue() float32`

GetCurrencyValue returns the CurrencyValue field if non-nil, zero value otherwise.

### GetCurrencyValueOk

`func (o *OutputValueAnyOf2) GetCurrencyValueOk() (*float32, bool)`

GetCurrencyValueOk returns a tuple with the CurrencyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyValue

`func (o *OutputValueAnyOf2) SetCurrencyValue(v float32)`

SetCurrencyValue sets CurrencyValue field to given value.


### GetCurrencyCode

`func (o *OutputValueAnyOf2) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OutputValueAnyOf2) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OutputValueAnyOf2) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *OutputValueAnyOf2) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *OutputValueAnyOf2) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *OutputValueAnyOf2) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetAttributeType

`func (o *OutputValueAnyOf2) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *OutputValueAnyOf2) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *OutputValueAnyOf2) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


