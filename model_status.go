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

// checks if the Status type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Status{}

// Status struct for Status
type Status struct {
	Id StatusId `json:"id"`
	// The title of the status
	Title string `json:"title"`
	// Whether or not to archive the status. See our [archiving guide](/docs/archiving-vs-deleting) for more information on archiving.
	IsArchived bool `json:"is_archived"`
	// Whether arriving at this status triggers a celebration effect in the UI
	CelebrationEnabled bool `json:"celebration_enabled"`
	// Target time for a record to spend in given status expressed as a ISO-8601 duration string
	TargetTimeInStatus NullableString `json:"target_time_in_status"`
}

type _Status Status

// NewStatus instantiates a new Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus(id StatusId, title string, isArchived bool, celebrationEnabled bool, targetTimeInStatus NullableString) *Status {
	this := Status{}
	this.Id = id
	this.Title = title
	this.IsArchived = isArchived
	this.CelebrationEnabled = celebrationEnabled
	this.TargetTimeInStatus = targetTimeInStatus
	return &this
}

// NewStatusWithDefaults instantiates a new Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusWithDefaults() *Status {
	this := Status{}
	return &this
}

// GetId returns the Id field value
func (o *Status) GetId() StatusId {
	if o == nil {
		var ret StatusId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Status) GetIdOk() (*StatusId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Status) SetId(v StatusId) {
	o.Id = v
}


// GetTitle returns the Title field value
func (o *Status) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Status) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Status) SetTitle(v string) {
	o.Title = v
}


// GetIsArchived returns the IsArchived field value
func (o *Status) GetIsArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsArchived
}

// GetIsArchivedOk returns a tuple with the IsArchived field value
// and a boolean to check if the value has been set.
func (o *Status) GetIsArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsArchived, true
}

// SetIsArchived sets field value
func (o *Status) SetIsArchived(v bool) {
	o.IsArchived = v
}


// GetCelebrationEnabled returns the CelebrationEnabled field value
func (o *Status) GetCelebrationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CelebrationEnabled
}

// GetCelebrationEnabledOk returns a tuple with the CelebrationEnabled field value
// and a boolean to check if the value has been set.
func (o *Status) GetCelebrationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CelebrationEnabled, true
}

// SetCelebrationEnabled sets field value
func (o *Status) SetCelebrationEnabled(v bool) {
	o.CelebrationEnabled = v
}


// GetTargetTimeInStatus returns the TargetTimeInStatus field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Status) GetTargetTimeInStatus() string {
	if o == nil || o.TargetTimeInStatus.Get() == nil {
		var ret string
		return ret
	}

	return *o.TargetTimeInStatus.Get()
}

// GetTargetTimeInStatusOk returns a tuple with the TargetTimeInStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Status) GetTargetTimeInStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetTimeInStatus.Get(), o.TargetTimeInStatus.IsSet()
}

// SetTargetTimeInStatus sets field value
func (o *Status) SetTargetTimeInStatus(v string) {
	o.TargetTimeInStatus.Set(&v)
}


func (o Status) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Status) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["is_archived"] = o.IsArchived
	toSerialize["celebration_enabled"] = o.CelebrationEnabled
	toSerialize["target_time_in_status"] = o.TargetTimeInStatus.Get()
	return toSerialize, nil
}

func (o *Status) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"title",
		"is_archived",
		"celebration_enabled",
		"target_time_in_status",
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
	varStatus := _Status{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStatus)

	if err != nil {
		return err
	}

	*o = Status(varStatus)

	return err
}

type NullableStatus struct {
	value *Status
	isSet bool
}

func (v NullableStatus) Get() *Status {
	return v.value
}

func (v *NullableStatus) Set(val *Status) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus(val *Status) *NullableStatus {
	return &NullableStatus{value: val, isSet: true}
}

func (v NullableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


