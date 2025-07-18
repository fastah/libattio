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

// checks if the V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest{}

// V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest struct for V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest
type V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest struct {
	Data V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequestData `json:"data"`
}

type _V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest

// NewV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest instantiates a new V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest(data V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequestData) *V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest {
	this := V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest{}
	this.Data = data
	return &this
}

// NewV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequestWithDefaults instantiates a new V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequestWithDefaults() *V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest {
	this := V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest{}
	return &this
}

// GetData returns the Data field value
func (o *V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest) GetData() V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequestData {
	if o == nil {
		var ret V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest) GetDataOk() (*V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest) SetData(v V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequestData) {
	o.Data = v
}


func (o V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest) UnmarshalJSON(data []byte) (err error) {
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
	varV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest := _V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest)

	if err != nil {
		return err
	}

	*o = V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest(varV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest)

	return err
}

type NullableV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest struct {
	value *V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest
	isSet bool
}

func (v NullableV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest) Get() *V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest {
	return v.value
}

func (v *NullableV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest) Set(val *V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest(val *V2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest) *NullableV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest {
	return &NullableV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest{value: val, isSet: true}
}

func (v NullableV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TargetIdentifierAttributesAttributeOptionsOptionPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


