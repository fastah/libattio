# V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveFrom** | **time.Time** | The point in time at which this value was made \&quot;active\&quot;. &#x60;active_from&#x60; can be considered roughly analogous to &#x60;created_at&#x60;. | 
**ActiveUntil** | **NullableTime** | The point in time at which this value was deactivated. If &#x60;null&#x60;, the value is active. | 
**CreatedByActor** | [**V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor**](V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor.md) |  | 
**CurrencyValue** | **float32** | A numerical representation of the currency value. A decimal with a max of 4 decimal places. | 
**CurrencyCode** | Pointer to **NullableString** | The ISO4217 currency code representing the currency that the value is stored in. | [optional] 
**AttributeType** | **string** | The attribute type of the value. | 

## Methods

### NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2

`func NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2(activeFrom time.Time, activeUntil NullableTime, createdByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor, currencyValue float32, attributeType string, ) *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2`

NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2 instantiates a new V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2WithDefaults

`func NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2WithDefaults() *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2`

NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2WithDefaults instantiates a new V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveFrom

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) GetActiveFrom() time.Time`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) GetActiveFromOk() (*time.Time, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) SetActiveFrom(v time.Time)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveUntil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) GetActiveUntil() time.Time`

GetActiveUntil returns the ActiveUntil field if non-nil, zero value otherwise.

### GetActiveUntilOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) GetActiveUntilOk() (*time.Time, bool)`

GetActiveUntilOk returns a tuple with the ActiveUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUntil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) SetActiveUntil(v time.Time)`

SetActiveUntil sets ActiveUntil field to given value.


### SetActiveUntilNil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) SetActiveUntilNil(b bool)`

 SetActiveUntilNil sets the value for ActiveUntil to be an explicit nil

### UnsetActiveUntil
`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) UnsetActiveUntil()`

UnsetActiveUntil ensures that no value is present for ActiveUntil, not even an explicit nil
### GetCreatedByActor

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) GetCreatedByActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) GetCreatedByActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) SetCreatedByActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.


### GetCurrencyValue

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) GetCurrencyValue() float32`

GetCurrencyValue returns the CurrencyValue field if non-nil, zero value otherwise.

### GetCurrencyValueOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) GetCurrencyValueOk() (*float32, bool)`

GetCurrencyValueOk returns a tuple with the CurrencyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyValue

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) SetCurrencyValue(v float32)`

SetCurrencyValue sets CurrencyValue field to given value.


### GetCurrencyCode

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetAttributeType

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf2) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


