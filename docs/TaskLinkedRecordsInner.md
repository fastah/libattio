# TaskLinkedRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetObjectId** | **string** | The ID of the parent object the task refers to. At present, only &#x60;people&#x60; and &#x60;companies&#x60; are supported. | 
**TargetRecordId** | **string** | The ID of the parent record the task refers to. | 

## Methods

### NewTaskLinkedRecordsInner

`func NewTaskLinkedRecordsInner(targetObjectId string, targetRecordId string, ) *TaskLinkedRecordsInner`

NewTaskLinkedRecordsInner instantiates a new TaskLinkedRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskLinkedRecordsInnerWithDefaults

`func NewTaskLinkedRecordsInnerWithDefaults() *TaskLinkedRecordsInner`

NewTaskLinkedRecordsInnerWithDefaults instantiates a new TaskLinkedRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetObjectId

`func (o *TaskLinkedRecordsInner) GetTargetObjectId() string`

GetTargetObjectId returns the TargetObjectId field if non-nil, zero value otherwise.

### GetTargetObjectIdOk

`func (o *TaskLinkedRecordsInner) GetTargetObjectIdOk() (*string, bool)`

GetTargetObjectIdOk returns a tuple with the TargetObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetObjectId

`func (o *TaskLinkedRecordsInner) SetTargetObjectId(v string)`

SetTargetObjectId sets TargetObjectId field to given value.


### GetTargetRecordId

`func (o *TaskLinkedRecordsInner) GetTargetRecordId() string`

GetTargetRecordId returns the TargetRecordId field if non-nil, zero value otherwise.

### GetTargetRecordIdOk

`func (o *TaskLinkedRecordsInner) GetTargetRecordIdOk() (*string, bool)`

GetTargetRecordIdOk returns a tuple with the TargetRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRecordId

`func (o *TaskLinkedRecordsInner) SetTargetRecordId(v string)`

SetTargetRecordId sets TargetRecordId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


