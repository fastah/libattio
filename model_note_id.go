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

// checks if the NoteId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NoteId{}

// NoteId struct for NoteId
type NoteId struct {
	// The ID of the workspace the note belongs to.
	WorkspaceId string `json:"workspace_id"`
	// The ID of the note.
	NoteId string `json:"note_id"`
}

type _NoteId NoteId

// NewNoteId instantiates a new NoteId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoteId(workspaceId string, noteId string) *NoteId {
	this := NoteId{}
	this.WorkspaceId = workspaceId
	this.NoteId = noteId
	return &this
}

// NewNoteIdWithDefaults instantiates a new NoteId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteIdWithDefaults() *NoteId {
	this := NoteId{}
	return &this
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *NoteId) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *NoteId) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *NoteId) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}


// GetNoteId returns the NoteId field value
func (o *NoteId) GetNoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NoteId
}

// GetNoteIdOk returns a tuple with the NoteId field value
// and a boolean to check if the value has been set.
func (o *NoteId) GetNoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoteId, true
}

// SetNoteId sets field value
func (o *NoteId) SetNoteId(v string) {
	o.NoteId = v
}


func (o NoteId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NoteId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workspace_id"] = o.WorkspaceId
	toSerialize["note_id"] = o.NoteId
	return toSerialize, nil
}

func (o *NoteId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"workspace_id",
		"note_id",
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
	varNoteId := _NoteId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNoteId)

	if err != nil {
		return err
	}

	*o = NoteId(varNoteId)

	return err
}

type NullableNoteId struct {
	value *NoteId
	isSet bool
}

func (v NullableNoteId) Get() *NoteId {
	return v.value
}

func (v *NullableNoteId) Set(val *NoteId) {
	v.value = val
	v.isSet = true
}

func (v NullableNoteId) IsSet() bool {
	return v.isSet
}

func (v *NullableNoteId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoteId(val *NoteId) *NullableNoteId {
	return &NullableNoteId{value: val, isSet: true}
}

func (v NullableNoteId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoteId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


