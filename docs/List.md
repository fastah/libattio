# List

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**ListId**](ListId.md) |  | 
**ApiSlug** | **string** | A human-readable slug for use in URLs and responses. | 
**Name** | **string** | The name of the list, as viewed in the UI. | 
**ParentObject** | **[]string** | A UUID or slug to identify the allowed object type for records added to this list. All new Lists are expected to have exactly one parent object. However, some legacy lists may have multiple allowed parents so the return type of this field is an array. | 
**WorkspaceAccess** | **NullableString** | The level of access granted to all members of the workspace for this list. &#x60;null&#x60; values represent a private list that only grants access to specific workspace members via the &#x60;workspace_member_access&#x60; property. | 
**WorkspaceMemberAccess** | [**[]V2ListsPostRequestDataWorkspaceMemberAccessInner**](V2ListsPostRequestDataWorkspaceMemberAccessInner.md) | The level of access granted to specific workspace members for this list. An empty array represents a list that has granted access to no workspace members. | 
**CreatedByActor** | [**ListCreatedByActor**](ListCreatedByActor.md) |  | 
**CreatedAt** | **string** | When the list was created. | 

## Methods

### NewList

`func NewList(id ListId, apiSlug string, name string, parentObject []string, workspaceAccess NullableString, workspaceMemberAccess []V2ListsPostRequestDataWorkspaceMemberAccessInner, createdByActor ListCreatedByActor, createdAt string, ) *List`

NewList instantiates a new List object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWithDefaults

`func NewListWithDefaults() *List`

NewListWithDefaults instantiates a new List object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *List) GetId() ListId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *List) GetIdOk() (*ListId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *List) SetId(v ListId)`

SetId sets Id field to given value.


### GetApiSlug

`func (o *List) GetApiSlug() string`

GetApiSlug returns the ApiSlug field if non-nil, zero value otherwise.

### GetApiSlugOk

`func (o *List) GetApiSlugOk() (*string, bool)`

GetApiSlugOk returns a tuple with the ApiSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSlug

`func (o *List) SetApiSlug(v string)`

SetApiSlug sets ApiSlug field to given value.


### GetName

`func (o *List) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *List) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *List) SetName(v string)`

SetName sets Name field to given value.


### GetParentObject

`func (o *List) GetParentObject() []string`

GetParentObject returns the ParentObject field if non-nil, zero value otherwise.

### GetParentObjectOk

`func (o *List) GetParentObjectOk() (*[]string, bool)`

GetParentObjectOk returns a tuple with the ParentObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObject

`func (o *List) SetParentObject(v []string)`

SetParentObject sets ParentObject field to given value.


### GetWorkspaceAccess

`func (o *List) GetWorkspaceAccess() string`

GetWorkspaceAccess returns the WorkspaceAccess field if non-nil, zero value otherwise.

### GetWorkspaceAccessOk

`func (o *List) GetWorkspaceAccessOk() (*string, bool)`

GetWorkspaceAccessOk returns a tuple with the WorkspaceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceAccess

`func (o *List) SetWorkspaceAccess(v string)`

SetWorkspaceAccess sets WorkspaceAccess field to given value.


### SetWorkspaceAccessNil

`func (o *List) SetWorkspaceAccessNil(b bool)`

 SetWorkspaceAccessNil sets the value for WorkspaceAccess to be an explicit nil

### UnsetWorkspaceAccess
`func (o *List) UnsetWorkspaceAccess()`

UnsetWorkspaceAccess ensures that no value is present for WorkspaceAccess, not even an explicit nil
### GetWorkspaceMemberAccess

`func (o *List) GetWorkspaceMemberAccess() []V2ListsPostRequestDataWorkspaceMemberAccessInner`

GetWorkspaceMemberAccess returns the WorkspaceMemberAccess field if non-nil, zero value otherwise.

### GetWorkspaceMemberAccessOk

`func (o *List) GetWorkspaceMemberAccessOk() (*[]V2ListsPostRequestDataWorkspaceMemberAccessInner, bool)`

GetWorkspaceMemberAccessOk returns a tuple with the WorkspaceMemberAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceMemberAccess

`func (o *List) SetWorkspaceMemberAccess(v []V2ListsPostRequestDataWorkspaceMemberAccessInner)`

SetWorkspaceMemberAccess sets WorkspaceMemberAccess field to given value.


### GetCreatedByActor

`func (o *List) GetCreatedByActor() ListCreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *List) GetCreatedByActorOk() (*ListCreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *List) SetCreatedByActor(v ListCreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.


### GetCreatedAt

`func (o *List) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *List) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *List) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


