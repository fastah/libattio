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
	"fmt"
)

// checks if the V2CommentsPostRequestDataAnyOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2CommentsPostRequestDataAnyOf2{}

// V2CommentsPostRequestDataAnyOf2 struct for V2CommentsPostRequestDataAnyOf2
type V2CommentsPostRequestDataAnyOf2 struct {
	// The format that the comment content is provided in. The `plaintext` format uses the line feed character `\\n` to create new lines within the note content. Rich text formatting and links are not supported.
	Format string `json:"format"`
	// The content of the comment itself. Workspace members can be mentioned using their email address, otherwise email addresses will be presented to users as clickable mailto links.
	Content string `json:"content"`
	Author V2CommentsPostRequestDataAnyOfAuthor `json:"author"`
	// `created_at` will default to the current time. However, if you wish to backdate a comment for migration or other purposes, you can override with a custom `created_at` value. Note that dates before 1970 or in the future are not allowed.
	CreatedAt *string `json:"created_at,omitempty"`
	Entry V2CommentsPostRequestDataAnyOf2Entry `json:"entry"`
	AdditionalProperties map[string]interface{}
}

type _V2CommentsPostRequestDataAnyOf2 V2CommentsPostRequestDataAnyOf2

// NewV2CommentsPostRequestDataAnyOf2 instantiates a new V2CommentsPostRequestDataAnyOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2CommentsPostRequestDataAnyOf2(format string, content string, author V2CommentsPostRequestDataAnyOfAuthor, entry V2CommentsPostRequestDataAnyOf2Entry) *V2CommentsPostRequestDataAnyOf2 {
	this := V2CommentsPostRequestDataAnyOf2{}
	this.Format = format
	this.Content = content
	this.Author = author
	this.Entry = entry
	return &this
}

// NewV2CommentsPostRequestDataAnyOf2WithDefaults instantiates a new V2CommentsPostRequestDataAnyOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2CommentsPostRequestDataAnyOf2WithDefaults() *V2CommentsPostRequestDataAnyOf2 {
	this := V2CommentsPostRequestDataAnyOf2{}
	return &this
}

// GetFormat returns the Format field value
func (o *V2CommentsPostRequestDataAnyOf2) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *V2CommentsPostRequestDataAnyOf2) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *V2CommentsPostRequestDataAnyOf2) SetFormat(v string) {
	o.Format = v
}


// GetContent returns the Content field value
func (o *V2CommentsPostRequestDataAnyOf2) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *V2CommentsPostRequestDataAnyOf2) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *V2CommentsPostRequestDataAnyOf2) SetContent(v string) {
	o.Content = v
}


// GetAuthor returns the Author field value
func (o *V2CommentsPostRequestDataAnyOf2) GetAuthor() V2CommentsPostRequestDataAnyOfAuthor {
	if o == nil {
		var ret V2CommentsPostRequestDataAnyOfAuthor
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *V2CommentsPostRequestDataAnyOf2) GetAuthorOk() (*V2CommentsPostRequestDataAnyOfAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *V2CommentsPostRequestDataAnyOf2) SetAuthor(v V2CommentsPostRequestDataAnyOfAuthor) {
	o.Author = v
}


// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *V2CommentsPostRequestDataAnyOf2) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2CommentsPostRequestDataAnyOf2) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *V2CommentsPostRequestDataAnyOf2) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *V2CommentsPostRequestDataAnyOf2) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetEntry returns the Entry field value
func (o *V2CommentsPostRequestDataAnyOf2) GetEntry() V2CommentsPostRequestDataAnyOf2Entry {
	if o == nil {
		var ret V2CommentsPostRequestDataAnyOf2Entry
		return ret
	}

	return o.Entry
}

// GetEntryOk returns a tuple with the Entry field value
// and a boolean to check if the value has been set.
func (o *V2CommentsPostRequestDataAnyOf2) GetEntryOk() (*V2CommentsPostRequestDataAnyOf2Entry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entry, true
}

// SetEntry sets field value
func (o *V2CommentsPostRequestDataAnyOf2) SetEntry(v V2CommentsPostRequestDataAnyOf2Entry) {
	o.Entry = v
}


func (o V2CommentsPostRequestDataAnyOf2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2CommentsPostRequestDataAnyOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["format"] = o.Format
	toSerialize["content"] = o.Content
	toSerialize["author"] = o.Author
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["entry"] = o.Entry

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V2CommentsPostRequestDataAnyOf2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"format",
		"content",
		"author",
		"entry",
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
	varV2CommentsPostRequestDataAnyOf2 := _V2CommentsPostRequestDataAnyOf2{}

	err = json.Unmarshal(data, &varV2CommentsPostRequestDataAnyOf2)

	if err != nil {
		return err
	}

	*o = V2CommentsPostRequestDataAnyOf2(varV2CommentsPostRequestDataAnyOf2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "format")
		delete(additionalProperties, "content")
		delete(additionalProperties, "author")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "entry")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV2CommentsPostRequestDataAnyOf2 struct {
	value *V2CommentsPostRequestDataAnyOf2
	isSet bool
}

func (v NullableV2CommentsPostRequestDataAnyOf2) Get() *V2CommentsPostRequestDataAnyOf2 {
	return v.value
}

func (v *NullableV2CommentsPostRequestDataAnyOf2) Set(val *V2CommentsPostRequestDataAnyOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableV2CommentsPostRequestDataAnyOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableV2CommentsPostRequestDataAnyOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2CommentsPostRequestDataAnyOf2(val *V2CommentsPostRequestDataAnyOf2) *NullableV2CommentsPostRequestDataAnyOf2 {
	return &NullableV2CommentsPostRequestDataAnyOf2{value: val, isSet: true}
}

func (v NullableV2CommentsPostRequestDataAnyOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2CommentsPostRequestDataAnyOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


