# InputValueAnyOf9

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalPhoneNumber** | **string** | A phone number which is either a) prefixed with a country code (e.g. &#x60;+44....&#x60;) or b) a local number, where &#x60;country_code&#x60; is specified in addition. | 
**CountryCode** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code representing the country that this phone number belongs to. Optional if &#x60;original_phone_number&#x60; includes a country code prefix. | [optional] 

## Methods

### NewInputValueAnyOf9

`func NewInputValueAnyOf9(originalPhoneNumber string, ) *InputValueAnyOf9`

NewInputValueAnyOf9 instantiates a new InputValueAnyOf9 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputValueAnyOf9WithDefaults

`func NewInputValueAnyOf9WithDefaults() *InputValueAnyOf9`

NewInputValueAnyOf9WithDefaults instantiates a new InputValueAnyOf9 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalPhoneNumber

`func (o *InputValueAnyOf9) GetOriginalPhoneNumber() string`

GetOriginalPhoneNumber returns the OriginalPhoneNumber field if non-nil, zero value otherwise.

### GetOriginalPhoneNumberOk

`func (o *InputValueAnyOf9) GetOriginalPhoneNumberOk() (*string, bool)`

GetOriginalPhoneNumberOk returns a tuple with the OriginalPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPhoneNumber

`func (o *InputValueAnyOf9) SetOriginalPhoneNumber(v string)`

SetOriginalPhoneNumber sets OriginalPhoneNumber field to given value.


### GetCountryCode

`func (o *InputValueAnyOf9) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *InputValueAnyOf9) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *InputValueAnyOf9) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *InputValueAnyOf9) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *InputValueAnyOf9) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *InputValueAnyOf9) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


