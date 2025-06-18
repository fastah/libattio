# ThreadId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** | The ID of the workspace the thread belongs to. | 
**ThreadId** | **string** | The ID of the thread. | 

## Methods

### NewThreadId

`func NewThreadId(workspaceId string, threadId string, ) *ThreadId`

NewThreadId instantiates a new ThreadId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadIdWithDefaults

`func NewThreadIdWithDefaults() *ThreadId`

NewThreadIdWithDefaults instantiates a new ThreadId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *ThreadId) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ThreadId) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ThreadId) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetThreadId

`func (o *ThreadId) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *ThreadId) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *ThreadId) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


