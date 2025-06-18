# V2ListsListPatchRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The human-readable name of the list. | [optional] 
**ApiSlug** | Pointer to **string** | A unique, human-readable slug to access the list through API calls. Should be formatted in snake case. | [optional] 
**WorkspaceAccess** | Pointer to **NullableString** | The level of access granted to all members of the workspace for this list. Pass &#x60;null&#x60; to keep the list private and only grant access to specific workspace members. | [optional] 
**WorkspaceMemberAccess** | Pointer to [**[]V2ListsPostRequestDataWorkspaceMemberAccessInner**](V2ListsPostRequestDataWorkspaceMemberAccessInner.md) | The level of access granted to specific workspace members for this list. Pass an empty array to grant access to no workspace members. | [optional] 

## Methods

### NewV2ListsListPatchRequestData

`func NewV2ListsListPatchRequestData() *V2ListsListPatchRequestData`

NewV2ListsListPatchRequestData instantiates a new V2ListsListPatchRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ListsListPatchRequestDataWithDefaults

`func NewV2ListsListPatchRequestDataWithDefaults() *V2ListsListPatchRequestData`

NewV2ListsListPatchRequestDataWithDefaults instantiates a new V2ListsListPatchRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2ListsListPatchRequestData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2ListsListPatchRequestData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2ListsListPatchRequestData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2ListsListPatchRequestData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApiSlug

`func (o *V2ListsListPatchRequestData) GetApiSlug() string`

GetApiSlug returns the ApiSlug field if non-nil, zero value otherwise.

### GetApiSlugOk

`func (o *V2ListsListPatchRequestData) GetApiSlugOk() (*string, bool)`

GetApiSlugOk returns a tuple with the ApiSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSlug

`func (o *V2ListsListPatchRequestData) SetApiSlug(v string)`

SetApiSlug sets ApiSlug field to given value.

### HasApiSlug

`func (o *V2ListsListPatchRequestData) HasApiSlug() bool`

HasApiSlug returns a boolean if a field has been set.

### GetWorkspaceAccess

`func (o *V2ListsListPatchRequestData) GetWorkspaceAccess() string`

GetWorkspaceAccess returns the WorkspaceAccess field if non-nil, zero value otherwise.

### GetWorkspaceAccessOk

`func (o *V2ListsListPatchRequestData) GetWorkspaceAccessOk() (*string, bool)`

GetWorkspaceAccessOk returns a tuple with the WorkspaceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceAccess

`func (o *V2ListsListPatchRequestData) SetWorkspaceAccess(v string)`

SetWorkspaceAccess sets WorkspaceAccess field to given value.

### HasWorkspaceAccess

`func (o *V2ListsListPatchRequestData) HasWorkspaceAccess() bool`

HasWorkspaceAccess returns a boolean if a field has been set.

### SetWorkspaceAccessNil

`func (o *V2ListsListPatchRequestData) SetWorkspaceAccessNil(b bool)`

 SetWorkspaceAccessNil sets the value for WorkspaceAccess to be an explicit nil

### UnsetWorkspaceAccess
`func (o *V2ListsListPatchRequestData) UnsetWorkspaceAccess()`

UnsetWorkspaceAccess ensures that no value is present for WorkspaceAccess, not even an explicit nil
### GetWorkspaceMemberAccess

`func (o *V2ListsListPatchRequestData) GetWorkspaceMemberAccess() []V2ListsPostRequestDataWorkspaceMemberAccessInner`

GetWorkspaceMemberAccess returns the WorkspaceMemberAccess field if non-nil, zero value otherwise.

### GetWorkspaceMemberAccessOk

`func (o *V2ListsListPatchRequestData) GetWorkspaceMemberAccessOk() (*[]V2ListsPostRequestDataWorkspaceMemberAccessInner, bool)`

GetWorkspaceMemberAccessOk returns a tuple with the WorkspaceMemberAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceMemberAccess

`func (o *V2ListsListPatchRequestData) SetWorkspaceMemberAccess(v []V2ListsPostRequestDataWorkspaceMemberAccessInner)`

SetWorkspaceMemberAccess sets WorkspaceMemberAccess field to given value.

### HasWorkspaceMemberAccess

`func (o *V2ListsListPatchRequestData) HasWorkspaceMemberAccess() bool`

HasWorkspaceMemberAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


