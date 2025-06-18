# OutputValueAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferencedActorType** | **string** | The type of the referenced actor. [Read more information on actor types here](/docs/actors). | 
**ReferencedActorId** | **NullableString** | The ID of the referenced actor. | 
**AttributeType** | **string** | The attribute type of the value. | 

## Methods

### NewOutputValueAnyOf

`func NewOutputValueAnyOf(referencedActorType string, referencedActorId NullableString, attributeType string, ) *OutputValueAnyOf`

NewOutputValueAnyOf instantiates a new OutputValueAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputValueAnyOfWithDefaults

`func NewOutputValueAnyOfWithDefaults() *OutputValueAnyOf`

NewOutputValueAnyOfWithDefaults instantiates a new OutputValueAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferencedActorType

`func (o *OutputValueAnyOf) GetReferencedActorType() string`

GetReferencedActorType returns the ReferencedActorType field if non-nil, zero value otherwise.

### GetReferencedActorTypeOk

`func (o *OutputValueAnyOf) GetReferencedActorTypeOk() (*string, bool)`

GetReferencedActorTypeOk returns a tuple with the ReferencedActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedActorType

`func (o *OutputValueAnyOf) SetReferencedActorType(v string)`

SetReferencedActorType sets ReferencedActorType field to given value.


### GetReferencedActorId

`func (o *OutputValueAnyOf) GetReferencedActorId() string`

GetReferencedActorId returns the ReferencedActorId field if non-nil, zero value otherwise.

### GetReferencedActorIdOk

`func (o *OutputValueAnyOf) GetReferencedActorIdOk() (*string, bool)`

GetReferencedActorIdOk returns a tuple with the ReferencedActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedActorId

`func (o *OutputValueAnyOf) SetReferencedActorId(v string)`

SetReferencedActorId sets ReferencedActorId field to given value.


### SetReferencedActorIdNil

`func (o *OutputValueAnyOf) SetReferencedActorIdNil(b bool)`

 SetReferencedActorIdNil sets the value for ReferencedActorId to be an explicit nil

### UnsetReferencedActorId
`func (o *OutputValueAnyOf) UnsetReferencedActorId()`

UnsetReferencedActorId ensures that no value is present for ReferencedActorId, not even an explicit nil
### GetAttributeType

`func (o *OutputValueAnyOf) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *OutputValueAnyOf) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *OutputValueAnyOf) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


