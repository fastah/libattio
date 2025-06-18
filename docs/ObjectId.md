# ObjectId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** | A UUID to identify the workspace this object belongs to. | 
**ObjectId** | **string** | A UUID to identify the object. | 

## Methods

### NewObjectId

`func NewObjectId(workspaceId string, objectId string, ) *ObjectId`

NewObjectId instantiates a new ObjectId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectIdWithDefaults

`func NewObjectIdWithDefaults() *ObjectId`

NewObjectIdWithDefaults instantiates a new ObjectId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *ObjectId) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ObjectId) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ObjectId) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetObjectId

`func (o *ObjectId) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ObjectId) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ObjectId) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


