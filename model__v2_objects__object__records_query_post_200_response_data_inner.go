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

// checks if the V2ObjectsObjectRecordsQueryPost200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsObjectRecordsQueryPost200ResponseDataInner{}

// V2ObjectsObjectRecordsQueryPost200ResponseDataInner struct for V2ObjectsObjectRecordsQueryPost200ResponseDataInner
type V2ObjectsObjectRecordsQueryPost200ResponseDataInner struct {
	Id V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId `json:"id"`
	// When this record was created.
	CreatedAt string `json:"created_at"`
	// A URL that links directly to the record page in the Attio web application.
	WebUrl string `json:"web_url"`
	// A record type with an attribute `api_slug` as the key, and an array of value objects as the values.
	Values map[string][]V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInner `json:"values"`
}

type _V2ObjectsObjectRecordsQueryPost200ResponseDataInner V2ObjectsObjectRecordsQueryPost200ResponseDataInner

// NewV2ObjectsObjectRecordsQueryPost200ResponseDataInner instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInner(id V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId, createdAt string, webUrl string, values map[string][]V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInner) *V2ObjectsObjectRecordsQueryPost200ResponseDataInner {
	this := V2ObjectsObjectRecordsQueryPost200ResponseDataInner{}
	this.Id = id
	this.CreatedAt = createdAt
	this.WebUrl = webUrl
	this.Values = values
	return &this
}

// NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerWithDefaults instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerWithDefaults() *V2ObjectsObjectRecordsQueryPost200ResponseDataInner {
	this := V2ObjectsObjectRecordsQueryPost200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetId() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId {
	if o == nil {
		var ret V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetIdOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) SetId(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId) {
	o.Id = v
}


// GetCreatedAt returns the CreatedAt field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) SetCreatedAt(v string) {
	o.CreatedAt = v
}


// GetWebUrl returns the WebUrl field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetWebUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetWebUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebUrl, true
}

// SetWebUrl sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) SetWebUrl(v string) {
	o.WebUrl = v
}


// GetValues returns the Values field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetValues() map[string][]V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInner {
	if o == nil {
		var ret map[string][]V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetValuesOk() (map[string][]V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInner, bool) {
	if o == nil {
		return map[string][]V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInner{}, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) SetValues(v map[string][]V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInner) {
	o.Values = v
}


func (o V2ObjectsObjectRecordsQueryPost200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsObjectRecordsQueryPost200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["web_url"] = o.WebUrl
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"web_url",
		"values",
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
	varV2ObjectsObjectRecordsQueryPost200ResponseDataInner := _V2ObjectsObjectRecordsQueryPost200ResponseDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2ObjectsObjectRecordsQueryPost200ResponseDataInner)

	if err != nil {
		return err
	}

	*o = V2ObjectsObjectRecordsQueryPost200ResponseDataInner(varV2ObjectsObjectRecordsQueryPost200ResponseDataInner)

	return err
}

type NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInner struct {
	value *V2ObjectsObjectRecordsQueryPost200ResponseDataInner
	isSet bool
}

func (v NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInner) Get() *V2ObjectsObjectRecordsQueryPost200ResponseDataInner {
	return v.value
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInner) Set(val *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsObjectRecordsQueryPost200ResponseDataInner(val *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInner {
	return &NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInner{value: val, isSet: true}
}

func (v NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


