# AttributeConfigRecordReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedObjectIds** | **[]string** | A list of UUIDs to indicate which objects records are allowed to belong to. Leave empty to to allow records from all object types. | 

## Methods

### NewAttributeConfigRecordReference

`func NewAttributeConfigRecordReference(allowedObjectIds []string, ) *AttributeConfigRecordReference`

NewAttributeConfigRecordReference instantiates a new AttributeConfigRecordReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeConfigRecordReferenceWithDefaults

`func NewAttributeConfigRecordReferenceWithDefaults() *AttributeConfigRecordReference`

NewAttributeConfigRecordReferenceWithDefaults instantiates a new AttributeConfigRecordReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedObjectIds

`func (o *AttributeConfigRecordReference) GetAllowedObjectIds() []string`

GetAllowedObjectIds returns the AllowedObjectIds field if non-nil, zero value otherwise.

### GetAllowedObjectIdsOk

`func (o *AttributeConfigRecordReference) GetAllowedObjectIdsOk() (*[]string, bool)`

GetAllowedObjectIdsOk returns a tuple with the AllowedObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedObjectIds

`func (o *AttributeConfigRecordReference) SetAllowedObjectIds(v []string)`

SetAllowedObjectIds sets AllowedObjectIds field to given value.


### SetAllowedObjectIdsNil

`func (o *AttributeConfigRecordReference) SetAllowedObjectIdsNil(b bool)`

 SetAllowedObjectIdsNil sets the value for AllowedObjectIds to be an explicit nil

### UnsetAllowedObjectIds
`func (o *AttributeConfigRecordReference) UnsetAllowedObjectIds()`

UnsetAllowedObjectIds ensures that no value is present for AllowedObjectIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


