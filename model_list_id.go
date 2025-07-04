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

// checks if the ListId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListId{}

// ListId struct for ListId
type ListId struct {
	// A UUID to identify the workspace this list belongs to.
	WorkspaceId string `json:"workspace_id"`
	// A UUID to identify this list.
	ListId string `json:"list_id"`
}

type _ListId ListId

// NewListId instantiates a new ListId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListId(workspaceId string, listId string) *ListId {
	this := ListId{}
	this.WorkspaceId = workspaceId
	this.ListId = listId
	return &this
}

// NewListIdWithDefaults instantiates a new ListId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIdWithDefaults() *ListId {
	this := ListId{}
	return &this
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *ListId) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *ListId) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *ListId) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}


// GetListId returns the ListId field value
func (o *ListId) GetListId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListId
}

// GetListIdOk returns a tuple with the ListId field value
// and a boolean to check if the value has been set.
func (o *ListId) GetListIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListId, true
}

// SetListId sets field value
func (o *ListId) SetListId(v string) {
	o.ListId = v
}


func (o ListId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workspace_id"] = o.WorkspaceId
	toSerialize["list_id"] = o.ListId
	return toSerialize, nil
}

func (o *ListId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"workspace_id",
		"list_id",
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
	varListId := _ListId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListId)

	if err != nil {
		return err
	}

	*o = ListId(varListId)

	return err
}

type NullableListId struct {
	value *ListId
	isSet bool
}

func (v NullableListId) Get() *ListId {
	return v.value
}

func (v *NullableListId) Set(val *ListId) {
	v.value = val
	v.isSet = true
}

func (v NullableListId) IsSet() bool {
	return v.isSet
}

func (v *NullableListId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListId(val *ListId) *NullableListId {
	return &NullableListId{value: val, isSet: true}
}

func (v NullableListId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


