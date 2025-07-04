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

// checks if the AttributeRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributeRelationship{}

// AttributeRelationship If this attribute is related to another attribute, this is an object that includes an `id` property that identifies the other attribute. `null` means no relationship exists. See [the help center](https://attio.com/help/reference/managing-your-data/attributes#relationship-attributes) for more details about relationship attributes.
type AttributeRelationship struct {
	Id AttributeId `json:"id"`
}

type _AttributeRelationship AttributeRelationship

// NewAttributeRelationship instantiates a new AttributeRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeRelationship(id AttributeId) *AttributeRelationship {
	this := AttributeRelationship{}
	this.Id = id
	return &this
}

// NewAttributeRelationshipWithDefaults instantiates a new AttributeRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeRelationshipWithDefaults() *AttributeRelationship {
	this := AttributeRelationship{}
	return &this
}

// GetId returns the Id field value
func (o *AttributeRelationship) GetId() AttributeId {
	if o == nil {
		var ret AttributeId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttributeRelationship) GetIdOk() (*AttributeId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttributeRelationship) SetId(v AttributeId) {
	o.Id = v
}


func (o AttributeRelationship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributeRelationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AttributeRelationship) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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
	varAttributeRelationship := _AttributeRelationship{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAttributeRelationship)

	if err != nil {
		return err
	}

	*o = AttributeRelationship(varAttributeRelationship)

	return err
}

type NullableAttributeRelationship struct {
	value *AttributeRelationship
	isSet bool
}

func (v NullableAttributeRelationship) Get() *AttributeRelationship {
	return v.value
}

func (v *NullableAttributeRelationship) Set(val *AttributeRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeRelationship(val *AttributeRelationship) *NullableAttributeRelationship {
	return &NullableAttributeRelationship{value: val, isSet: true}
}

func (v NullableAttributeRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


