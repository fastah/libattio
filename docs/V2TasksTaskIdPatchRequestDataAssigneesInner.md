# V2TasksTaskIdPatchRequestDataAssigneesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferencedActorType** | **string** | The actor type of the task assignee. Only &#x60;workspace-member&#x60; actors can be assigned to tasks. [Read more information on actor types here](/docs/actors). | 
**ReferencedActorId** | **string** | The ID of the actor assigned to this task. | 
**WorkspaceMemberEmailAddress** | **string** | Workspace member actors can be referenced by email address as well as actor ID. | 

## Methods

### NewV2TasksTaskIdPatchRequestDataAssigneesInner

`func NewV2TasksTaskIdPatchRequestDataAssigneesInner(referencedActorType string, referencedActorId string, workspaceMemberEmailAddress string, ) *V2TasksTaskIdPatchRequestDataAssigneesInner`

NewV2TasksTaskIdPatchRequestDataAssigneesInner instantiates a new V2TasksTaskIdPatchRequestDataAssigneesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TasksTaskIdPatchRequestDataAssigneesInnerWithDefaults

`func NewV2TasksTaskIdPatchRequestDataAssigneesInnerWithDefaults() *V2TasksTaskIdPatchRequestDataAssigneesInner`

NewV2TasksTaskIdPatchRequestDataAssigneesInnerWithDefaults instantiates a new V2TasksTaskIdPatchRequestDataAssigneesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferencedActorType

`func (o *V2TasksTaskIdPatchRequestDataAssigneesInner) GetReferencedActorType() string`

GetReferencedActorType returns the ReferencedActorType field if non-nil, zero value otherwise.

### GetReferencedActorTypeOk

`func (o *V2TasksTaskIdPatchRequestDataAssigneesInner) GetReferencedActorTypeOk() (*string, bool)`

GetReferencedActorTypeOk returns a tuple with the ReferencedActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedActorType

`func (o *V2TasksTaskIdPatchRequestDataAssigneesInner) SetReferencedActorType(v string)`

SetReferencedActorType sets ReferencedActorType field to given value.


### GetReferencedActorId

`func (o *V2TasksTaskIdPatchRequestDataAssigneesInner) GetReferencedActorId() string`

GetReferencedActorId returns the ReferencedActorId field if non-nil, zero value otherwise.

### GetReferencedActorIdOk

`func (o *V2TasksTaskIdPatchRequestDataAssigneesInner) GetReferencedActorIdOk() (*string, bool)`

GetReferencedActorIdOk returns a tuple with the ReferencedActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedActorId

`func (o *V2TasksTaskIdPatchRequestDataAssigneesInner) SetReferencedActorId(v string)`

SetReferencedActorId sets ReferencedActorId field to given value.


### GetWorkspaceMemberEmailAddress

`func (o *V2TasksTaskIdPatchRequestDataAssigneesInner) GetWorkspaceMemberEmailAddress() string`

GetWorkspaceMemberEmailAddress returns the WorkspaceMemberEmailAddress field if non-nil, zero value otherwise.

### GetWorkspaceMemberEmailAddressOk

`func (o *V2TasksTaskIdPatchRequestDataAssigneesInner) GetWorkspaceMemberEmailAddressOk() (*string, bool)`

GetWorkspaceMemberEmailAddressOk returns a tuple with the WorkspaceMemberEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceMemberEmailAddress

`func (o *V2TasksTaskIdPatchRequestDataAssigneesInner) SetWorkspaceMemberEmailAddress(v string)`

SetWorkspaceMemberEmailAddress sets WorkspaceMemberEmailAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


