# V2TasksPostRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | The text content of the task, in the format specified by the &#x60;format&#x60; property. | 
**Format** | **string** | The format of the task content to be created. Rich text formatting, links and @references are not supported. | 
**DeadlineAt** | **NullableString** | The deadline of the task, in ISO 8601 format. | 
**IsCompleted** | **bool** | Whether the task has been completed. | 
**LinkedRecords** | [**[]V2TasksPostRequestDataLinkedRecordsInner**](V2TasksPostRequestDataLinkedRecordsInner.md) | Records linked to the task. Creating record links within task content text is not possible via the API at present. | 
**Assignees** | [**[]V2TasksPostRequestDataAssigneesInner**](V2TasksPostRequestDataAssigneesInner.md) | Workspace members assigned to this task. | 

## Methods

### NewV2TasksPostRequestData

`func NewV2TasksPostRequestData(content string, format string, deadlineAt NullableString, isCompleted bool, linkedRecords []V2TasksPostRequestDataLinkedRecordsInner, assignees []V2TasksPostRequestDataAssigneesInner, ) *V2TasksPostRequestData`

NewV2TasksPostRequestData instantiates a new V2TasksPostRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TasksPostRequestDataWithDefaults

`func NewV2TasksPostRequestDataWithDefaults() *V2TasksPostRequestData`

NewV2TasksPostRequestDataWithDefaults instantiates a new V2TasksPostRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *V2TasksPostRequestData) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *V2TasksPostRequestData) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *V2TasksPostRequestData) SetContent(v string)`

SetContent sets Content field to given value.


### GetFormat

`func (o *V2TasksPostRequestData) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *V2TasksPostRequestData) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *V2TasksPostRequestData) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetDeadlineAt

`func (o *V2TasksPostRequestData) GetDeadlineAt() string`

GetDeadlineAt returns the DeadlineAt field if non-nil, zero value otherwise.

### GetDeadlineAtOk

`func (o *V2TasksPostRequestData) GetDeadlineAtOk() (*string, bool)`

GetDeadlineAtOk returns a tuple with the DeadlineAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlineAt

`func (o *V2TasksPostRequestData) SetDeadlineAt(v string)`

SetDeadlineAt sets DeadlineAt field to given value.


### SetDeadlineAtNil

`func (o *V2TasksPostRequestData) SetDeadlineAtNil(b bool)`

 SetDeadlineAtNil sets the value for DeadlineAt to be an explicit nil

### UnsetDeadlineAt
`func (o *V2TasksPostRequestData) UnsetDeadlineAt()`

UnsetDeadlineAt ensures that no value is present for DeadlineAt, not even an explicit nil
### GetIsCompleted

`func (o *V2TasksPostRequestData) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *V2TasksPostRequestData) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *V2TasksPostRequestData) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.


### GetLinkedRecords

`func (o *V2TasksPostRequestData) GetLinkedRecords() []V2TasksPostRequestDataLinkedRecordsInner`

GetLinkedRecords returns the LinkedRecords field if non-nil, zero value otherwise.

### GetLinkedRecordsOk

`func (o *V2TasksPostRequestData) GetLinkedRecordsOk() (*[]V2TasksPostRequestDataLinkedRecordsInner, bool)`

GetLinkedRecordsOk returns a tuple with the LinkedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedRecords

`func (o *V2TasksPostRequestData) SetLinkedRecords(v []V2TasksPostRequestDataLinkedRecordsInner)`

SetLinkedRecords sets LinkedRecords field to given value.


### GetAssignees

`func (o *V2TasksPostRequestData) GetAssignees() []V2TasksPostRequestDataAssigneesInner`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *V2TasksPostRequestData) GetAssigneesOk() (*[]V2TasksPostRequestDataAssigneesInner, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *V2TasksPostRequestData) SetAssignees(v []V2TasksPostRequestDataAssigneesInner)`

SetAssignees sets Assignees field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


