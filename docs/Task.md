# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**TaskId**](TaskId.md) |  | 
**ContentPlaintext** | **string** | The plaintext representation of the task content. Inline linked records will appear as \&quot;@record name\&quot; and are returned in the &#x60;linked_records&#x60; property. | 
**DeadlineAt** | **NullableString** | The deadline date of the task. Returned as an ISO 8601 timestamp. | 
**IsCompleted** | **bool** | Whether the task has been completed. | 
**LinkedRecords** | [**[]TaskLinkedRecordsInner**](TaskLinkedRecordsInner.md) | Records linked to the task. Creating record links within task content text is not possible via the API at present. | 
**Assignees** | [**[]TaskAssigneesInner**](TaskAssigneesInner.md) | Workspace members assigned to this task. | 
**CreatedByActor** | [**TaskCreatedByActor**](TaskCreatedByActor.md) |  | 
**CreatedAt** | **string** | When the task was created. | 

## Methods

### NewTask

`func NewTask(id TaskId, contentPlaintext string, deadlineAt NullableString, isCompleted bool, linkedRecords []TaskLinkedRecordsInner, assignees []TaskAssigneesInner, createdByActor TaskCreatedByActor, createdAt string, ) *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Task) GetId() TaskId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (*TaskId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Task) SetId(v TaskId)`

SetId sets Id field to given value.


### GetContentPlaintext

`func (o *Task) GetContentPlaintext() string`

GetContentPlaintext returns the ContentPlaintext field if non-nil, zero value otherwise.

### GetContentPlaintextOk

`func (o *Task) GetContentPlaintextOk() (*string, bool)`

GetContentPlaintextOk returns a tuple with the ContentPlaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPlaintext

`func (o *Task) SetContentPlaintext(v string)`

SetContentPlaintext sets ContentPlaintext field to given value.


### GetDeadlineAt

`func (o *Task) GetDeadlineAt() string`

GetDeadlineAt returns the DeadlineAt field if non-nil, zero value otherwise.

### GetDeadlineAtOk

`func (o *Task) GetDeadlineAtOk() (*string, bool)`

GetDeadlineAtOk returns a tuple with the DeadlineAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlineAt

`func (o *Task) SetDeadlineAt(v string)`

SetDeadlineAt sets DeadlineAt field to given value.


### SetDeadlineAtNil

`func (o *Task) SetDeadlineAtNil(b bool)`

 SetDeadlineAtNil sets the value for DeadlineAt to be an explicit nil

### UnsetDeadlineAt
`func (o *Task) UnsetDeadlineAt()`

UnsetDeadlineAt ensures that no value is present for DeadlineAt, not even an explicit nil
### GetIsCompleted

`func (o *Task) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *Task) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *Task) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.


### GetLinkedRecords

`func (o *Task) GetLinkedRecords() []TaskLinkedRecordsInner`

GetLinkedRecords returns the LinkedRecords field if non-nil, zero value otherwise.

### GetLinkedRecordsOk

`func (o *Task) GetLinkedRecordsOk() (*[]TaskLinkedRecordsInner, bool)`

GetLinkedRecordsOk returns a tuple with the LinkedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedRecords

`func (o *Task) SetLinkedRecords(v []TaskLinkedRecordsInner)`

SetLinkedRecords sets LinkedRecords field to given value.


### GetAssignees

`func (o *Task) GetAssignees() []TaskAssigneesInner`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *Task) GetAssigneesOk() (*[]TaskAssigneesInner, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *Task) SetAssignees(v []TaskAssigneesInner)`

SetAssignees sets Assignees field to given value.


### GetCreatedByActor

`func (o *Task) GetCreatedByActor() TaskCreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *Task) GetCreatedByActorOk() (*TaskCreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *Task) SetCreatedByActor(v TaskCreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.


### GetCreatedAt

`func (o *Task) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Task) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Task) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


