# V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveFrom** | **time.Time** | The point in time at which this value was made \&quot;active\&quot;. &#x60;active_from&#x60; can be considered roughly analogous to &#x60;created_at&#x60;. | 
**ActiveUntil** | **NullableTime** | The point in time at which this value was deactivated. If &#x60;null&#x60;, the value is active. | 
**CreatedByActor** | [**V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor**](V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor.md) |  | 
**Domain** | **string** |  | 
**RootDomain** | **string** |  | 
**AttributeType** | **string** | The attribute type of the value. | 

## Methods

### NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4

`func NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4(activeFrom time.Time, activeUntil NullableTime, createdByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor, domain string, rootDomain string, attributeType string, ) *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4`

NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4 instantiates a new V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4WithDefaults

`func NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4WithDefaults() *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4`

NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4WithDefaults instantiates a new V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveFrom

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) GetActiveFrom() time.Time`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) GetActiveFromOk() (*time.Time, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) SetActiveFrom(v time.Time)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveUntil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) GetActiveUntil() time.Time`

GetActiveUntil returns the ActiveUntil field if non-nil, zero value otherwise.

### GetActiveUntilOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) GetActiveUntilOk() (*time.Time, bool)`

GetActiveUntilOk returns a tuple with the ActiveUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUntil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) SetActiveUntil(v time.Time)`

SetActiveUntil sets ActiveUntil field to given value.


### SetActiveUntilNil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) SetActiveUntilNil(b bool)`

 SetActiveUntilNil sets the value for ActiveUntil to be an explicit nil

### UnsetActiveUntil
`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) UnsetActiveUntil()`

UnsetActiveUntil ensures that no value is present for ActiveUntil, not even an explicit nil
### GetCreatedByActor

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) GetCreatedByActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) GetCreatedByActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) SetCreatedByActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.


### GetDomain

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetRootDomain

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) GetRootDomain() string`

GetRootDomain returns the RootDomain field if non-nil, zero value otherwise.

### GetRootDomainOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) GetRootDomainOk() (*string, bool)`

GetRootDomainOk returns a tuple with the RootDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDomain

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) SetRootDomain(v string)`

SetRootDomain sets RootDomain field to given value.


### GetAttributeType

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerOneOf4) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


