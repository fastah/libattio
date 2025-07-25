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

// checks if the InputValueAnyOf3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputValueAnyOf3{}

// InputValueAnyOf3 struct for InputValueAnyOf3
type InputValueAnyOf3 struct {
	// A numerical representation of the currency value. A decimal with a max of 4 decimal places.
	CurrencyValue float32 `json:"currency_value"`
}

type _InputValueAnyOf3 InputValueAnyOf3

// NewInputValueAnyOf3 instantiates a new InputValueAnyOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputValueAnyOf3(currencyValue float32) *InputValueAnyOf3 {
	this := InputValueAnyOf3{}
	this.CurrencyValue = currencyValue
	return &this
}

// NewInputValueAnyOf3WithDefaults instantiates a new InputValueAnyOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputValueAnyOf3WithDefaults() *InputValueAnyOf3 {
	this := InputValueAnyOf3{}
	return &this
}

// GetCurrencyValue returns the CurrencyValue field value
func (o *InputValueAnyOf3) GetCurrencyValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrencyValue
}

// GetCurrencyValueOk returns a tuple with the CurrencyValue field value
// and a boolean to check if the value has been set.
func (o *InputValueAnyOf3) GetCurrencyValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyValue, true
}

// SetCurrencyValue sets field value
func (o *InputValueAnyOf3) SetCurrencyValue(v float32) {
	o.CurrencyValue = v
}


func (o InputValueAnyOf3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputValueAnyOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency_value"] = o.CurrencyValue
	return toSerialize, nil
}

func (o *InputValueAnyOf3) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency_value",
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
	varInputValueAnyOf3 := _InputValueAnyOf3{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInputValueAnyOf3)

	if err != nil {
		return err
	}

	*o = InputValueAnyOf3(varInputValueAnyOf3)

	return err
}

type NullableInputValueAnyOf3 struct {
	value *InputValueAnyOf3
	isSet bool
}

func (v NullableInputValueAnyOf3) Get() *InputValueAnyOf3 {
	return v.value
}

func (v *NullableInputValueAnyOf3) Set(val *InputValueAnyOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableInputValueAnyOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableInputValueAnyOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputValueAnyOf3(val *InputValueAnyOf3) *NullableInputValueAnyOf3 {
	return &NullableInputValueAnyOf3{value: val, isSet: true}
}

func (v NullableInputValueAnyOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputValueAnyOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


