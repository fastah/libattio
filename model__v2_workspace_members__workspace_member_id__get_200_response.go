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

// checks if the V2WorkspaceMembersWorkspaceMemberIdGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2WorkspaceMembersWorkspaceMemberIdGet200Response{}

// V2WorkspaceMembersWorkspaceMemberIdGet200Response Success
type V2WorkspaceMembersWorkspaceMemberIdGet200Response struct {
	Data WorkspaceMember `json:"data"`
}

type _V2WorkspaceMembersWorkspaceMemberIdGet200Response V2WorkspaceMembersWorkspaceMemberIdGet200Response

// NewV2WorkspaceMembersWorkspaceMemberIdGet200Response instantiates a new V2WorkspaceMembersWorkspaceMemberIdGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2WorkspaceMembersWorkspaceMemberIdGet200Response(data WorkspaceMember) *V2WorkspaceMembersWorkspaceMemberIdGet200Response {
	this := V2WorkspaceMembersWorkspaceMemberIdGet200Response{}
	this.Data = data
	return &this
}

// NewV2WorkspaceMembersWorkspaceMemberIdGet200ResponseWithDefaults instantiates a new V2WorkspaceMembersWorkspaceMemberIdGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2WorkspaceMembersWorkspaceMemberIdGet200ResponseWithDefaults() *V2WorkspaceMembersWorkspaceMemberIdGet200Response {
	this := V2WorkspaceMembersWorkspaceMemberIdGet200Response{}
	return &this
}

// GetData returns the Data field value
func (o *V2WorkspaceMembersWorkspaceMemberIdGet200Response) GetData() WorkspaceMember {
	if o == nil {
		var ret WorkspaceMember
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *V2WorkspaceMembersWorkspaceMemberIdGet200Response) GetDataOk() (*WorkspaceMember, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *V2WorkspaceMembersWorkspaceMemberIdGet200Response) SetData(v WorkspaceMember) {
	o.Data = v
}


func (o V2WorkspaceMembersWorkspaceMemberIdGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2WorkspaceMembersWorkspaceMemberIdGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *V2WorkspaceMembersWorkspaceMemberIdGet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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
	varV2WorkspaceMembersWorkspaceMemberIdGet200Response := _V2WorkspaceMembersWorkspaceMemberIdGet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2WorkspaceMembersWorkspaceMemberIdGet200Response)

	if err != nil {
		return err
	}

	*o = V2WorkspaceMembersWorkspaceMemberIdGet200Response(varV2WorkspaceMembersWorkspaceMemberIdGet200Response)

	return err
}

type NullableV2WorkspaceMembersWorkspaceMemberIdGet200Response struct {
	value *V2WorkspaceMembersWorkspaceMemberIdGet200Response
	isSet bool
}

func (v NullableV2WorkspaceMembersWorkspaceMemberIdGet200Response) Get() *V2WorkspaceMembersWorkspaceMemberIdGet200Response {
	return v.value
}

func (v *NullableV2WorkspaceMembersWorkspaceMemberIdGet200Response) Set(val *V2WorkspaceMembersWorkspaceMemberIdGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV2WorkspaceMembersWorkspaceMemberIdGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV2WorkspaceMembersWorkspaceMemberIdGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2WorkspaceMembersWorkspaceMemberIdGet200Response(val *V2WorkspaceMembersWorkspaceMemberIdGet200Response) *NullableV2WorkspaceMembersWorkspaceMemberIdGet200Response {
	return &NullableV2WorkspaceMembersWorkspaceMemberIdGet200Response{value: val, isSet: true}
}

func (v NullableV2WorkspaceMembersWorkspaceMemberIdGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2WorkspaceMembersWorkspaceMemberIdGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


