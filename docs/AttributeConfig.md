# AttributeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | [**AttributeConfigCurrency**](AttributeConfigCurrency.md) |  | 
**RecordReference** | [**AttributeConfigRecordReference**](AttributeConfigRecordReference.md) |  | 

## Methods

### NewAttributeConfig

`func NewAttributeConfig(currency AttributeConfigCurrency, recordReference AttributeConfigRecordReference, ) *AttributeConfig`

NewAttributeConfig instantiates a new AttributeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeConfigWithDefaults

`func NewAttributeConfigWithDefaults() *AttributeConfig`

NewAttributeConfigWithDefaults instantiates a new AttributeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *AttributeConfig) GetCurrency() AttributeConfigCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AttributeConfig) GetCurrencyOk() (*AttributeConfigCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AttributeConfig) SetCurrency(v AttributeConfigCurrency)`

SetCurrency sets Currency field to given value.


### GetRecordReference

`func (o *AttributeConfig) GetRecordReference() AttributeConfigRecordReference`

GetRecordReference returns the RecordReference field if non-nil, zero value otherwise.

### GetRecordReferenceOk

`func (o *AttributeConfig) GetRecordReferenceOk() (*AttributeConfigRecordReference, bool)`

GetRecordReferenceOk returns a tuple with the RecordReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordReference

`func (o *AttributeConfig) SetRecordReference(v AttributeConfigRecordReference)`

SetRecordReference sets RecordReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


