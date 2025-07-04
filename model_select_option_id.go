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

// checks if the SelectOptionId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectOptionId{}

// SelectOptionId struct for SelectOptionId
type SelectOptionId struct {
	// The ID of the workspace
	WorkspaceId string `json:"workspace_id"`
	// The ID of the object
	ObjectId string `json:"object_id"`
	// The ID of the attribute
	AttributeId string `json:"attribute_id"`
	// The ID of the select option
	OptionId string `json:"option_id"`
}

type _SelectOptionId SelectOptionId

// NewSelectOptionId instantiates a new SelectOptionId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectOptionId(workspaceId string, objectId string, attributeId string, optionId string) *SelectOptionId {
	this := SelectOptionId{}
	this.WorkspaceId = workspaceId
	this.ObjectId = objectId
	this.AttributeId = attributeId
	this.OptionId = optionId
	return &this
}

// NewSelectOptionIdWithDefaults instantiates a new SelectOptionId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectOptionIdWithDefaults() *SelectOptionId {
	this := SelectOptionId{}
	return &this
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *SelectOptionId) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *SelectOptionId) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *SelectOptionId) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}


// GetObjectId returns the ObjectId field value
func (o *SelectOptionId) GetObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *SelectOptionId) GetObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *SelectOptionId) SetObjectId(v string) {
	o.ObjectId = v
}


// GetAttributeId returns the AttributeId field value
func (o *SelectOptionId) GetAttributeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeId
}

// GetAttributeIdOk returns a tuple with the AttributeId field value
// and a boolean to check if the value has been set.
func (o *SelectOptionId) GetAttributeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeId, true
}

// SetAttributeId sets field value
func (o *SelectOptionId) SetAttributeId(v string) {
	o.AttributeId = v
}


// GetOptionId returns the OptionId field value
func (o *SelectOptionId) GetOptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OptionId
}

// GetOptionIdOk returns a tuple with the OptionId field value
// and a boolean to check if the value has been set.
func (o *SelectOptionId) GetOptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptionId, true
}

// SetOptionId sets field value
func (o *SelectOptionId) SetOptionId(v string) {
	o.OptionId = v
}


func (o SelectOptionId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectOptionId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workspace_id"] = o.WorkspaceId
	toSerialize["object_id"] = o.ObjectId
	toSerialize["attribute_id"] = o.AttributeId
	toSerialize["option_id"] = o.OptionId
	return toSerialize, nil
}

func (o *SelectOptionId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"workspace_id",
		"object_id",
		"attribute_id",
		"option_id",
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
	varSelectOptionId := _SelectOptionId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSelectOptionId)

	if err != nil {
		return err
	}

	*o = SelectOptionId(varSelectOptionId)

	return err
}

type NullableSelectOptionId struct {
	value *SelectOptionId
	isSet bool
}

func (v NullableSelectOptionId) Get() *SelectOptionId {
	return v.value
}

func (v *NullableSelectOptionId) Set(val *SelectOptionId) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectOptionId) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectOptionId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectOptionId(val *SelectOptionId) *NullableSelectOptionId {
	return &NullableSelectOptionId{value: val, isSet: true}
}

func (v NullableSelectOptionId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectOptionId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


