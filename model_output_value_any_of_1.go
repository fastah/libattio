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

// checks if the OutputValueAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputValueAnyOf1{}

// OutputValueAnyOf1 struct for OutputValueAnyOf1
type OutputValueAnyOf1 struct {
	// A boolean representing whether the checkbox is checked or not. The string values 'true' and 'false' are also accepted.
	Value bool `json:"value"`
	// The attribute type of the value.
	AttributeType string `json:"attribute_type"`
}

type _OutputValueAnyOf1 OutputValueAnyOf1

// NewOutputValueAnyOf1 instantiates a new OutputValueAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputValueAnyOf1(value bool, attributeType string) *OutputValueAnyOf1 {
	this := OutputValueAnyOf1{}
	this.Value = value
	this.AttributeType = attributeType
	return &this
}

// NewOutputValueAnyOf1WithDefaults instantiates a new OutputValueAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputValueAnyOf1WithDefaults() *OutputValueAnyOf1 {
	this := OutputValueAnyOf1{}
	return &this
}

// GetValue returns the Value field value
func (o *OutputValueAnyOf1) GetValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OutputValueAnyOf1) GetValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *OutputValueAnyOf1) SetValue(v bool) {
	o.Value = v
}


// GetAttributeType returns the AttributeType field value
func (o *OutputValueAnyOf1) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *OutputValueAnyOf1) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *OutputValueAnyOf1) SetAttributeType(v string) {
	o.AttributeType = v
}


func (o OutputValueAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputValueAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["attribute_type"] = o.AttributeType
	return toSerialize, nil
}

func (o *OutputValueAnyOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"attribute_type",
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
	varOutputValueAnyOf1 := _OutputValueAnyOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOutputValueAnyOf1)

	if err != nil {
		return err
	}

	*o = OutputValueAnyOf1(varOutputValueAnyOf1)

	return err
}

type NullableOutputValueAnyOf1 struct {
	value *OutputValueAnyOf1
	isSet bool
}

func (v NullableOutputValueAnyOf1) Get() *OutputValueAnyOf1 {
	return v.value
}

func (v *NullableOutputValueAnyOf1) Set(val *OutputValueAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputValueAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputValueAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputValueAnyOf1(val *OutputValueAnyOf1) *NullableOutputValueAnyOf1 {
	return &NullableOutputValueAnyOf1{value: val, isSet: true}
}

func (v NullableOutputValueAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputValueAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


