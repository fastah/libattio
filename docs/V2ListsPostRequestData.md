# V2ListsPostRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The human-readable name of the list. | 
**ApiSlug** | **string** | A unique, human-readable slug to access the list through API calls. Should be formatted in snake case. | 
**ParentObject** | **string** | A UUID or slug to identify the allowed object type for records added to this list. | 
**WorkspaceAccess** | **NullableString** | The level of access granted to all members of the workspace for this list. Pass &#x60;null&#x60; to keep the list private and only grant access to specific workspace members. | 
**WorkspaceMemberAccess** | [**[]V2ListsPostRequestDataWorkspaceMemberAccessInner**](V2ListsPostRequestDataWorkspaceMemberAccessInner.md) | The level of access granted to specific workspace members for this list. Pass an empty array to grant access to no workspace members. | 

## Methods

### NewV2ListsPostRequestData

`func NewV2ListsPostRequestData(name string, apiSlug string, parentObject string, workspaceAccess NullableString, workspaceMemberAccess []V2ListsPostRequestDataWorkspaceMemberAccessInner, ) *V2ListsPostRequestData`

NewV2ListsPostRequestData instantiates a new V2ListsPostRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ListsPostRequestDataWithDefaults

`func NewV2ListsPostRequestDataWithDefaults() *V2ListsPostRequestData`

NewV2ListsPostRequestDataWithDefaults instantiates a new V2ListsPostRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2ListsPostRequestData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2ListsPostRequestData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2ListsPostRequestData) SetName(v string)`

SetName sets Name field to given value.


### GetApiSlug

`func (o *V2ListsPostRequestData) GetApiSlug() string`

GetApiSlug returns the ApiSlug field if non-nil, zero value otherwise.

### GetApiSlugOk

`func (o *V2ListsPostRequestData) GetApiSlugOk() (*string, bool)`

GetApiSlugOk returns a tuple with the ApiSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSlug

`func (o *V2ListsPostRequestData) SetApiSlug(v string)`

SetApiSlug sets ApiSlug field to given value.


### GetParentObject

`func (o *V2ListsPostRequestData) GetParentObject() string`

GetParentObject returns the ParentObject field if non-nil, zero value otherwise.

### GetParentObjectOk

`func (o *V2ListsPostRequestData) GetParentObjectOk() (*string, bool)`

GetParentObjectOk returns a tuple with the ParentObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObject

`func (o *V2ListsPostRequestData) SetParentObject(v string)`

SetParentObject sets ParentObject field to given value.


### GetWorkspaceAccess

`func (o *V2ListsPostRequestData) GetWorkspaceAccess() string`

GetWorkspaceAccess returns the WorkspaceAccess field if non-nil, zero value otherwise.

### GetWorkspaceAccessOk

`func (o *V2ListsPostRequestData) GetWorkspaceAccessOk() (*string, bool)`

GetWorkspaceAccessOk returns a tuple with the WorkspaceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceAccess

`func (o *V2ListsPostRequestData) SetWorkspaceAccess(v string)`

SetWorkspaceAccess sets WorkspaceAccess field to given value.


### SetWorkspaceAccessNil

`func (o *V2ListsPostRequestData) SetWorkspaceAccessNil(b bool)`

 SetWorkspaceAccessNil sets the value for WorkspaceAccess to be an explicit nil

### UnsetWorkspaceAccess
`func (o *V2ListsPostRequestData) UnsetWorkspaceAccess()`

UnsetWorkspaceAccess ensures that no value is present for WorkspaceAccess, not even an explicit nil
### GetWorkspaceMemberAccess

`func (o *V2ListsPostRequestData) GetWorkspaceMemberAccess() []V2ListsPostRequestDataWorkspaceMemberAccessInner`

GetWorkspaceMemberAccess returns the WorkspaceMemberAccess field if non-nil, zero value otherwise.

### GetWorkspaceMemberAccessOk

`func (o *V2ListsPostRequestData) GetWorkspaceMemberAccessOk() (*[]V2ListsPostRequestDataWorkspaceMemberAccessInner, bool)`

GetWorkspaceMemberAccessOk returns a tuple with the WorkspaceMemberAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceMemberAccess

`func (o *V2ListsPostRequestData) SetWorkspaceMemberAccess(v []V2ListsPostRequestDataWorkspaceMemberAccessInner)`

SetWorkspaceMemberAccess sets WorkspaceMemberAccess field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


