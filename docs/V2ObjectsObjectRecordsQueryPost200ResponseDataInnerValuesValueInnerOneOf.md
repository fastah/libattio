# V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveFrom** | **time.Time** | The point in time at which this value was made \&quot;active\&quot;. &#x60;active_from&#x60; can be considered roughly analogous to &#x60;created_at&#x60;. | 
**ActiveUntil** | **time.Time** | The point in time at which this value was deactivated. If &#x60;null&#x60;, the value is active. | 
**CreatedByActor** | [**V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor**](V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor.md) |  | 
**ReferencedActorType** | **string** | The type of the referenced actor. [Read more information on actor types here](/docs/actors). | 
**ReferencedActorId** | **NullableString** | The ID of the referenced actor. | 
**AttributeType** | **string** | The attribute type of the value. | 

## Methods

### NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf

`func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf(activeFrom time.Time, activeUntil time.Time, createdByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor, referencedActorType string, referencedActorId NullableString, attributeType string, ) *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf`

NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfWithDefaults

`func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfWithDefaults() *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf`

NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfWithDefaults instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveFrom

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) GetActiveFrom() time.Time`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) GetActiveFromOk() (*time.Time, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) SetActiveFrom(v time.Time)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveUntil

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) GetActiveUntil() time.Time`

GetActiveUntil returns the ActiveUntil field if non-nil, zero value otherwise.

### GetActiveUntilOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) GetActiveUntilOk() (*time.Time, bool)`

GetActiveUntilOk returns a tuple with the ActiveUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUntil

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) SetActiveUntil(v time.Time)`

SetActiveUntil sets ActiveUntil field to given value.


### GetCreatedByActor

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) GetCreatedByActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) GetCreatedByActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) SetCreatedByActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.


### GetReferencedActorType

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) GetReferencedActorType() string`

GetReferencedActorType returns the ReferencedActorType field if non-nil, zero value otherwise.

### GetReferencedActorTypeOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) GetReferencedActorTypeOk() (*string, bool)`

GetReferencedActorTypeOk returns a tuple with the ReferencedActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedActorType

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) SetReferencedActorType(v string)`

SetReferencedActorType sets ReferencedActorType field to given value.


### GetReferencedActorId

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) GetReferencedActorId() string`

GetReferencedActorId returns the ReferencedActorId field if non-nil, zero value otherwise.

### GetReferencedActorIdOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) GetReferencedActorIdOk() (*string, bool)`

GetReferencedActorIdOk returns a tuple with the ReferencedActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedActorId

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) SetReferencedActorId(v string)`

SetReferencedActorId sets ReferencedActorId field to given value.


### SetReferencedActorIdNil

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) SetReferencedActorIdNil(b bool)`

 SetReferencedActorIdNil sets the value for ReferencedActorId to be an explicit nil

### UnsetReferencedActorId
`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) UnsetReferencedActorId()`

UnsetReferencedActorId ensures that no value is present for ReferencedActorId, not even an explicit nil
### GetAttributeType

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


