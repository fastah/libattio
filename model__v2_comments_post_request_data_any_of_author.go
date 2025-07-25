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

// checks if the V2CommentsPostRequestDataAnyOfAuthor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2CommentsPostRequestDataAnyOfAuthor{}

// V2CommentsPostRequestDataAnyOfAuthor The workspace member who wrote this comment. Note that other types of actors are not currently supported.
type V2CommentsPostRequestDataAnyOfAuthor struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

type _V2CommentsPostRequestDataAnyOfAuthor V2CommentsPostRequestDataAnyOfAuthor

// NewV2CommentsPostRequestDataAnyOfAuthor instantiates a new V2CommentsPostRequestDataAnyOfAuthor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2CommentsPostRequestDataAnyOfAuthor(type_ string, id string) *V2CommentsPostRequestDataAnyOfAuthor {
	this := V2CommentsPostRequestDataAnyOfAuthor{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewV2CommentsPostRequestDataAnyOfAuthorWithDefaults instantiates a new V2CommentsPostRequestDataAnyOfAuthor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2CommentsPostRequestDataAnyOfAuthorWithDefaults() *V2CommentsPostRequestDataAnyOfAuthor {
	this := V2CommentsPostRequestDataAnyOfAuthor{}
	return &this
}

// GetType returns the Type field value
func (o *V2CommentsPostRequestDataAnyOfAuthor) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V2CommentsPostRequestDataAnyOfAuthor) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V2CommentsPostRequestDataAnyOfAuthor) SetType(v string) {
	o.Type = v
}


// GetId returns the Id field value
func (o *V2CommentsPostRequestDataAnyOfAuthor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *V2CommentsPostRequestDataAnyOfAuthor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *V2CommentsPostRequestDataAnyOfAuthor) SetId(v string) {
	o.Id = v
}


func (o V2CommentsPostRequestDataAnyOfAuthor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2CommentsPostRequestDataAnyOfAuthor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *V2CommentsPostRequestDataAnyOfAuthor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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
	varV2CommentsPostRequestDataAnyOfAuthor := _V2CommentsPostRequestDataAnyOfAuthor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2CommentsPostRequestDataAnyOfAuthor)

	if err != nil {
		return err
	}

	*o = V2CommentsPostRequestDataAnyOfAuthor(varV2CommentsPostRequestDataAnyOfAuthor)

	return err
}

type NullableV2CommentsPostRequestDataAnyOfAuthor struct {
	value *V2CommentsPostRequestDataAnyOfAuthor
	isSet bool
}

func (v NullableV2CommentsPostRequestDataAnyOfAuthor) Get() *V2CommentsPostRequestDataAnyOfAuthor {
	return v.value
}

func (v *NullableV2CommentsPostRequestDataAnyOfAuthor) Set(val *V2CommentsPostRequestDataAnyOfAuthor) {
	v.value = val
	v.isSet = true
}

func (v NullableV2CommentsPostRequestDataAnyOfAuthor) IsSet() bool {
	return v.isSet
}

func (v *NullableV2CommentsPostRequestDataAnyOfAuthor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2CommentsPostRequestDataAnyOfAuthor(val *V2CommentsPostRequestDataAnyOfAuthor) *NullableV2CommentsPostRequestDataAnyOfAuthor {
	return &NullableV2CommentsPostRequestDataAnyOfAuthor{value: val, isSet: true}
}

func (v NullableV2CommentsPostRequestDataAnyOfAuthor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2CommentsPostRequestDataAnyOfAuthor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


