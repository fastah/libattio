# ListId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** | A UUID to identify the workspace this list belongs to. | 
**ListId** | **string** | A UUID to identify this list. | 

## Methods

### NewListId

`func NewListId(workspaceId string, listId string, ) *ListId`

NewListId instantiates a new ListId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIdWithDefaults

`func NewListIdWithDefaults() *ListId`

NewListIdWithDefaults instantiates a new ListId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *ListId) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ListId) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ListId) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetListId

`func (o *ListId) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *ListId) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *ListId) SetListId(v string)`

SetListId sets ListId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


