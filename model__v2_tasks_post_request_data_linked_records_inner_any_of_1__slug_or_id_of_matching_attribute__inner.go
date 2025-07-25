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


// V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner struct for V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner
type V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner struct {
	V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf *V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf
	V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf1 *V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf1
	V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf2 *V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf2
	V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf3 *V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf3
	V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf4 *V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf4
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf
	err = json.Unmarshal(data, &dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf);
	if err == nil {
		jsonV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf, _ := json.Marshal(dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf)
		if string(jsonV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf) == "{}" { // empty struct
			dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf = nil
		} else {
			return nil // data stored in dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf, return on the first match
		}
	} else {
		dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf = nil
	}

	// try to unmarshal JSON data into V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf1
	err = json.Unmarshal(data, &dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf1);
	if err == nil {
		jsonV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf1, _ := json.Marshal(dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf1)
		if string(jsonV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf1) == "{}" { // empty struct
			dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf1 = nil
		} else {
			return nil // data stored in dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf1, return on the first match
		}
	} else {
		dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf1 = nil
	}

	// try to unmarshal JSON data into V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf2
	err = json.Unmarshal(data, &dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf2);
	if err == nil {
		jsonV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf2, _ := json.Marshal(dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf2)
		if string(jsonV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf2) == "{}" { // empty struct
			dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf2 = nil
		} else {
			return nil // data stored in dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf2, return on the first match
		}
	} else {
		dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf2 = nil
	}

	// try to unmarshal JSON data into V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf3
	err = json.Unmarshal(data, &dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf3);
	if err == nil {
		jsonV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf3, _ := json.Marshal(dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf3)
		if string(jsonV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf3) == "{}" { // empty struct
			dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf3 = nil
		} else {
			return nil // data stored in dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf3, return on the first match
		}
	} else {
		dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf3 = nil
	}

	// try to unmarshal JSON data into V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf4
	err = json.Unmarshal(data, &dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf4);
	if err == nil {
		jsonV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf4, _ := json.Marshal(dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf4)
		if string(jsonV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf4) == "{}" { // empty struct
			dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf4 = nil
		} else {
			return nil // data stored in dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf4, return on the first match
		}
	} else {
		dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf4 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner) MarshalJSON() ([]byte, error) {
	if src.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf != nil {
		return json.Marshal(&src.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf)
	}

	if src.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf1 != nil {
		return json.Marshal(&src.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf1)
	}

	if src.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf2 != nil {
		return json.Marshal(&src.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf2)
	}

	if src.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf3 != nil {
		return json.Marshal(&src.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf3)
	}

	if src.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf4 != nil {
		return json.Marshal(&src.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf4)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner struct {
	value *V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner
	isSet bool
}

func (v NullableV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner) Get() *V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner {
	return v.value
}

func (v *NullableV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner) Set(val *V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner(val *V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner) *NullableV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner {
	return &NullableV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner{value: val, isSet: true}
}

func (v NullableV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


