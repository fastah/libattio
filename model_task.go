/*
Attio API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
Contact: support@attio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libattio

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Task type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Task{}

// Task struct for Task
type Task struct {
	Id TaskId `json:"id"`
	// The plaintext representation of the task content. Inline linked records will appear as \"@record name\" and are returned in the `linked_records` property.
	ContentPlaintext string `json:"content_plaintext"`
	// The deadline date of the task. Returned as an ISO 8601 timestamp.
	DeadlineAt NullableString `json:"deadline_at"`
	// Whether the task has been completed.
	IsCompleted bool `json:"is_completed"`
	// Records linked to the task. Creating record links within task content text is not possible via the API at present.
	LinkedRecords []TaskLinkedRecordsInner `json:"linked_records"`
	// Workspace members assigned to this task.
	Assignees []TaskAssigneesInner `json:"assignees"`
	CreatedByActor TaskCreatedByActor `json:"created_by_actor"`
	// When the task was created.
	CreatedAt string `json:"created_at"`
}

type _Task Task

// NewTask instantiates a new Task object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTask(id TaskId, contentPlaintext string, deadlineAt NullableString, isCompleted bool, linkedRecords []TaskLinkedRecordsInner, assignees []TaskAssigneesInner, createdByActor TaskCreatedByActor, createdAt string) *Task {
	this := Task{}
	this.Id = id
	this.ContentPlaintext = contentPlaintext
	this.DeadlineAt = deadlineAt
	this.IsCompleted = isCompleted
	this.LinkedRecords = linkedRecords
	this.Assignees = assignees
	this.CreatedByActor = createdByActor
	this.CreatedAt = createdAt
	return &this
}

// NewTaskWithDefaults instantiates a new Task object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWithDefaults() *Task {
	this := Task{}
	return &this
}

// GetId returns the Id field value
func (o *Task) GetId() TaskId {
	if o == nil {
		var ret TaskId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Task) GetIdOk() (*TaskId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Task) SetId(v TaskId) {
	o.Id = v
}


// GetContentPlaintext returns the ContentPlaintext field value
func (o *Task) GetContentPlaintext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentPlaintext
}

// GetContentPlaintextOk returns a tuple with the ContentPlaintext field value
// and a boolean to check if the value has been set.
func (o *Task) GetContentPlaintextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentPlaintext, true
}

// SetContentPlaintext sets field value
func (o *Task) SetContentPlaintext(v string) {
	o.ContentPlaintext = v
}


// GetDeadlineAt returns the DeadlineAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Task) GetDeadlineAt() string {
	if o == nil || o.DeadlineAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.DeadlineAt.Get()
}

// GetDeadlineAtOk returns a tuple with the DeadlineAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetDeadlineAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeadlineAt.Get(), o.DeadlineAt.IsSet()
}

// SetDeadlineAt sets field value
func (o *Task) SetDeadlineAt(v string) {
	o.DeadlineAt.Set(&v)
}


// GetIsCompleted returns the IsCompleted field value
func (o *Task) GetIsCompleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCompleted
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value
// and a boolean to check if the value has been set.
func (o *Task) GetIsCompletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCompleted, true
}

// SetIsCompleted sets field value
func (o *Task) SetIsCompleted(v bool) {
	o.IsCompleted = v
}


// GetLinkedRecords returns the LinkedRecords field value
func (o *Task) GetLinkedRecords() []TaskLinkedRecordsInner {
	if o == nil {
		var ret []TaskLinkedRecordsInner
		return ret
	}

	return o.LinkedRecords
}

// GetLinkedRecordsOk returns a tuple with the LinkedRecords field value
// and a boolean to check if the value has been set.
func (o *Task) GetLinkedRecordsOk() ([]TaskLinkedRecordsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkedRecords, true
}

// SetLinkedRecords sets field value
func (o *Task) SetLinkedRecords(v []TaskLinkedRecordsInner) {
	o.LinkedRecords = v
}


// GetAssignees returns the Assignees field value
func (o *Task) GetAssignees() []TaskAssigneesInner {
	if o == nil {
		var ret []TaskAssigneesInner
		return ret
	}

	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value
// and a boolean to check if the value has been set.
func (o *Task) GetAssigneesOk() ([]TaskAssigneesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assignees, true
}

// SetAssignees sets field value
func (o *Task) SetAssignees(v []TaskAssigneesInner) {
	o.Assignees = v
}


// GetCreatedByActor returns the CreatedByActor field value
func (o *Task) GetCreatedByActor() TaskCreatedByActor {
	if o == nil {
		var ret TaskCreatedByActor
		return ret
	}

	return o.CreatedByActor
}

// GetCreatedByActorOk returns a tuple with the CreatedByActor field value
// and a boolean to check if the value has been set.
func (o *Task) GetCreatedByActorOk() (*TaskCreatedByActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByActor, true
}

// SetCreatedByActor sets field value
func (o *Task) SetCreatedByActor(v TaskCreatedByActor) {
	o.CreatedByActor = v
}


// GetCreatedAt returns the CreatedAt field value
func (o *Task) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Task) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Task) SetCreatedAt(v string) {
	o.CreatedAt = v
}


func (o Task) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Task) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["content_plaintext"] = o.ContentPlaintext
	toSerialize["deadline_at"] = o.DeadlineAt.Get()
	toSerialize["is_completed"] = o.IsCompleted
	toSerialize["linked_records"] = o.LinkedRecords
	toSerialize["assignees"] = o.Assignees
	toSerialize["created_by_actor"] = o.CreatedByActor
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *Task) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"content_plaintext",
		"deadline_at",
		"is_completed",
		"linked_records",
		"assignees",
		"created_by_actor",
		"created_at",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{} {
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == ""{
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil{
			return err
		}
	}
	varTask := _Task{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTask)

	if err != nil {
		return err
	}

	*o = Task(varTask)

	return err
}

type NullableTask struct {
	value *Task
	isSet bool
}

func (v NullableTask) Get() *Task {
	return v.value
}

func (v *NullableTask) Set(val *Task) {
	v.value = val
	v.isSet = true
}

func (v NullableTask) IsSet() bool {
	return v.isSet
}

func (v *NullableTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTask(val *Task) *NullableTask {
	return &NullableTask{value: val, isSet: true}
}

func (v NullableTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


