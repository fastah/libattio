# V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveFrom** | **time.Time** | The point in time at which this value was made \&quot;active\&quot;. &#x60;active_from&#x60; can be considered roughly analogous to &#x60;created_at&#x60;. | 
**ActiveUntil** | **NullableTime** | The point in time at which this value was deactivated. If &#x60;null&#x60;, the value is active. | 
**CreatedByActor** | [**V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor**](V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor.md) |  | 
**Value** | **string** | A raw text field. Values are limited to 10MB. | 
**AttributeType** | **string** | The attribute type of the value. | 

## Methods

### NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12

`func NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12(activeFrom time.Time, activeUntil NullableTime, createdByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor, value string, attributeType string, ) *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12`

NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12 instantiates a new V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12WithDefaults

`func NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12WithDefaults() *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12`

NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12WithDefaults instantiates a new V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveFrom

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) GetActiveFrom() time.Time`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) GetActiveFromOk() (*time.Time, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) SetActiveFrom(v time.Time)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveUntil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) GetActiveUntil() time.Time`

GetActiveUntil returns the ActiveUntil field if non-nil, zero value otherwise.

### GetActiveUntilOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) GetActiveUntilOk() (*time.Time, bool)`

GetActiveUntilOk returns a tuple with the ActiveUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUntil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) SetActiveUntil(v time.Time)`

SetActiveUntil sets ActiveUntil field to given value.


### SetActiveUntilNil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) SetActiveUntilNil(b bool)`

 SetActiveUntilNil sets the value for ActiveUntil to be an explicit nil

### UnsetActiveUntil
`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) UnsetActiveUntil()`

UnsetActiveUntil ensures that no value is present for ActiveUntil, not even an explicit nil
### GetCreatedByActor

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) GetCreatedByActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) GetCreatedByActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) SetCreatedByActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.


### GetValue

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) SetValue(v string)`

SetValue sets Value field to given value.


### GetAttributeType

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf12) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


