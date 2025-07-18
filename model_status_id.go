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

// checks if the StatusId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusId{}

// StatusId struct for StatusId
type StatusId struct {
	// The ID of the workspace
	WorkspaceId string `json:"workspace_id"`
	// The ID of the object
	ObjectId string `json:"object_id"`
	// The ID of the attribute
	AttributeId string `json:"attribute_id"`
	// The ID of the status
	StatusId string `json:"status_id"`
}

type _StatusId StatusId

// NewStatusId instantiates a new StatusId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusId(workspaceId string, objectId string, attributeId string, statusId string) *StatusId {
	this := StatusId{}
	this.WorkspaceId = workspaceId
	this.ObjectId = objectId
	this.AttributeId = attributeId
	this.StatusId = statusId
	return &this
}

// NewStatusIdWithDefaults instantiates a new StatusId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusIdWithDefaults() *StatusId {
	this := StatusId{}
	return &this
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *StatusId) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *StatusId) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *StatusId) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}


// GetObjectId returns the ObjectId field value
func (o *StatusId) GetObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *StatusId) GetObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *StatusId) SetObjectId(v string) {
	o.ObjectId = v
}


// GetAttributeId returns the AttributeId field value
func (o *StatusId) GetAttributeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeId
}

// GetAttributeIdOk returns a tuple with the AttributeId field value
// and a boolean to check if the value has been set.
func (o *StatusId) GetAttributeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeId, true
}

// SetAttributeId sets field value
func (o *StatusId) SetAttributeId(v string) {
	o.AttributeId = v
}


// GetStatusId returns the StatusId field value
func (o *StatusId) GetStatusId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusId
}

// GetStatusIdOk returns a tuple with the StatusId field value
// and a boolean to check if the value has been set.
func (o *StatusId) GetStatusIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusId, true
}

// SetStatusId sets field value
func (o *StatusId) SetStatusId(v string) {
	o.StatusId = v
}


func (o StatusId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workspace_id"] = o.WorkspaceId
	toSerialize["object_id"] = o.ObjectId
	toSerialize["attribute_id"] = o.AttributeId
	toSerialize["status_id"] = o.StatusId
	return toSerialize, nil
}

func (o *StatusId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"workspace_id",
		"object_id",
		"attribute_id",
		"status_id",
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
	varStatusId := _StatusId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStatusId)

	if err != nil {
		return err
	}

	*o = StatusId(varStatusId)

	return err
}

type NullableStatusId struct {
	value *StatusId
	isSet bool
}

func (v NullableStatusId) Get() *StatusId {
	return v.value
}

func (v *NullableStatusId) Set(val *StatusId) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusId) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusId(val *StatusId) *NullableStatusId {
	return &NullableStatusId{value: val, isSet: true}
}

func (v NullableStatusId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


