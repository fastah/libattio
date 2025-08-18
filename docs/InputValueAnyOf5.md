# InputValueAnyOf5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InteractionType** | **string** | The type of interaction e.g. calendar or email. | 
**InteractedAt** | **time.Time** | When the interaction occurred. | 
**OwnerActor** | [**V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor**](V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor.md) |  | 

## Methods

### NewInputValueAnyOf5

`func NewInputValueAnyOf5(interactionType string, interactedAt time.Time, ownerActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor, ) *InputValueAnyOf5`

NewInputValueAnyOf5 instantiates a new InputValueAnyOf5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputValueAnyOf5WithDefaults

`func NewInputValueAnyOf5WithDefaults() *InputValueAnyOf5`

NewInputValueAnyOf5WithDefaults instantiates a new InputValueAnyOf5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInteractionType

`func (o *InputValueAnyOf5) GetInteractionType() string`

GetInteractionType returns the InteractionType field if non-nil, zero value otherwise.

### GetInteractionTypeOk

`func (o *InputValueAnyOf5) GetInteractionTypeOk() (*string, bool)`

GetInteractionTypeOk returns a tuple with the InteractionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionType

`func (o *InputValueAnyOf5) SetInteractionType(v string)`

SetInteractionType sets InteractionType field to given value.


### GetInteractedAt

`func (o *InputValueAnyOf5) GetInteractedAt() time.Time`

GetInteractedAt returns the InteractedAt field if non-nil, zero value otherwise.

### GetInteractedAtOk

`func (o *InputValueAnyOf5) GetInteractedAtOk() (*time.Time, bool)`

GetInteractedAtOk returns a tuple with the InteractedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractedAt

`func (o *InputValueAnyOf5) SetInteractedAt(v time.Time)`

SetInteractedAt sets InteractedAt field to given value.


### GetOwnerActor

`func (o *InputValueAnyOf5) GetOwnerActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor`

GetOwnerActor returns the OwnerActor field if non-nil, zero value otherwise.

### GetOwnerActorOk

`func (o *InputValueAnyOf5) GetOwnerActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor, bool)`

GetOwnerActorOk returns a tuple with the OwnerActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerActor

`func (o *InputValueAnyOf5) SetOwnerActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor)`

SetOwnerActor sets OwnerActor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


