# AttributeId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** | A UUID representing the workspace this attribute belongs to. | 
**ObjectId** | **string** | A UUID to identify the object or list that this attribute belongs to | 
**AttributeId** | **string** | A UUID to identify this attribute. | 

## Methods

### NewAttributeId

`func NewAttributeId(workspaceId string, objectId string, attributeId string, ) *AttributeId`

NewAttributeId instantiates a new AttributeId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeIdWithDefaults

`func NewAttributeIdWithDefaults() *AttributeId`

NewAttributeIdWithDefaults instantiates a new AttributeId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *AttributeId) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *AttributeId) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *AttributeId) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetObjectId

`func (o *AttributeId) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *AttributeId) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *AttributeId) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetAttributeId

`func (o *AttributeId) GetAttributeId() string`

GetAttributeId returns the AttributeId field if non-nil, zero value otherwise.

### GetAttributeIdOk

`func (o *AttributeId) GetAttributeIdOk() (*string, bool)`

GetAttributeIdOk returns a tuple with the AttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeId

`func (o *AttributeId) SetAttributeId(v string)`

SetAttributeId sets AttributeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


