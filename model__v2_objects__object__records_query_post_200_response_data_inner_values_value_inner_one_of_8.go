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
	"time"
	"fmt"
)

// checks if the V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8{}

// V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8 struct for V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8
type V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8 struct {
	// The point in time at which this value was made \"active\". `active_from` can be considered roughly analogous to `created_at`.
	ActiveFrom time.Time `json:"active_from"`
	// The point in time at which this value was deactivated. If `null`, the value is active.
	ActiveUntil NullableTime `json:"active_until"`
	CreatedByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor `json:"created_by_actor"`
	// The first line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.
	Line1 NullableString `json:"line_1"`
	// The second line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.
	Line2 NullableString `json:"line_2"`
	// The third line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.
	Line3 NullableString `json:"line_3"`
	// The fourth line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.
	Line4 NullableString `json:"line_4"`
	// The town, neighborhood or area the location is in.
	Locality NullableString `json:"locality"`
	// The state, county, province or region that the location is in.
	Region NullableString `json:"region"`
	// The postcode or zip code for the location. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.}
	Postcode NullableString `json:"postcode"`
	// The ISO 3166-1 alpha-2 country code for the country this location is in.
	CountryCode NullableString `json:"country_code"`
	// The latitude of the location. Validated by the regular expression `/^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$/`. Values are stored with up to 9 decimal places of precision. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.}
	Latitude NullableString `json:"latitude" validate:"regexp=^[-+]?([1-8]?\\\\d(\\\\.\\\\d+)?|90(\\\\.0+)?)$"`
	// The longitude of the location. Validated by the regular expression `/^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$/`. Values are stored with up to 9 decimal places of precision. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.}
	Longitude NullableString `json:"longitude" validate:"regexp=^[-+]?(180(\\\\.0+)?|((1[0-7]\\\\d)|([1-9]?\\\\d))(\\\\.\\\\d+)?)$"`
	// The attribute type of the value.
	AttributeType string `json:"attribute_type"`
	AdditionalProperties map[string]interface{}
}

type _V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8 V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8

// NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8 instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8(activeFrom time.Time, activeUntil NullableTime, createdByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor, line1 NullableString, line2 NullableString, line3 NullableString, line4 NullableString, locality NullableString, region NullableString, postcode NullableString, countryCode NullableString, latitude NullableString, longitude NullableString, attributeType string) *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8 {
	this := V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8{}
	this.ActiveFrom = activeFrom
	this.ActiveUntil = activeUntil
	this.CreatedByActor = createdByActor
	this.Line1 = line1
	this.Line2 = line2
	this.Line3 = line3
	this.Line4 = line4
	this.Locality = locality
	this.Region = region
	this.Postcode = postcode
	this.CountryCode = countryCode
	this.Latitude = latitude
	this.Longitude = longitude
	this.AttributeType = attributeType
	return &this
}

// NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8WithDefaults instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8WithDefaults() *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8 {
	this := V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8{}
	return &this
}

// GetActiveFrom returns the ActiveFrom field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetActiveFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ActiveFrom
}

// GetActiveFromOk returns a tuple with the ActiveFrom field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetActiveFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveFrom, true
}

// SetActiveFrom sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) SetActiveFrom(v time.Time) {
	o.ActiveFrom = v
}


// GetActiveUntil returns the ActiveUntil field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetActiveUntil() time.Time {
	if o == nil || o.ActiveUntil.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ActiveUntil.Get()
}

// GetActiveUntilOk returns a tuple with the ActiveUntil field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetActiveUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveUntil.Get(), o.ActiveUntil.IsSet()
}

// SetActiveUntil sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) SetActiveUntil(v time.Time) {
	o.ActiveUntil.Set(&v)
}


// GetCreatedByActor returns the CreatedByActor field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetCreatedByActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor {
	if o == nil {
		var ret V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor
		return ret
	}

	return o.CreatedByActor
}

// GetCreatedByActorOk returns a tuple with the CreatedByActor field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetCreatedByActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByActor, true
}

// SetCreatedByActor sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) SetCreatedByActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor) {
	o.CreatedByActor = v
}


// GetLine1 returns the Line1 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetLine1() string {
	if o == nil || o.Line1.Get() == nil {
		var ret string
		return ret
	}

	return *o.Line1.Get()
}

// GetLine1Ok returns a tuple with the Line1 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line1.Get(), o.Line1.IsSet()
}

// SetLine1 sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) SetLine1(v string) {
	o.Line1.Set(&v)
}


// GetLine2 returns the Line2 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetLine2() string {
	if o == nil || o.Line2.Get() == nil {
		var ret string
		return ret
	}

	return *o.Line2.Get()
}

// GetLine2Ok returns a tuple with the Line2 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line2.Get(), o.Line2.IsSet()
}

// SetLine2 sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) SetLine2(v string) {
	o.Line2.Set(&v)
}


// GetLine3 returns the Line3 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetLine3() string {
	if o == nil || o.Line3.Get() == nil {
		var ret string
		return ret
	}

	return *o.Line3.Get()
}

// GetLine3Ok returns a tuple with the Line3 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetLine3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line3.Get(), o.Line3.IsSet()
}

// SetLine3 sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) SetLine3(v string) {
	o.Line3.Set(&v)
}


// GetLine4 returns the Line4 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetLine4() string {
	if o == nil || o.Line4.Get() == nil {
		var ret string
		return ret
	}

	return *o.Line4.Get()
}

// GetLine4Ok returns a tuple with the Line4 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetLine4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line4.Get(), o.Line4.IsSet()
}

// SetLine4 sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) SetLine4(v string) {
	o.Line4.Set(&v)
}


// GetLocality returns the Locality field value
// If the value is explicit nil, the zero value for string will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetLocality() string {
	if o == nil || o.Locality.Get() == nil {
		var ret string
		return ret
	}

	return *o.Locality.Get()
}

// GetLocalityOk returns a tuple with the Locality field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetLocalityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locality.Get(), o.Locality.IsSet()
}

// SetLocality sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) SetLocality(v string) {
	o.Locality.Set(&v)
}


// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for string will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) SetRegion(v string) {
	o.Region.Set(&v)
}


// GetPostcode returns the Postcode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetPostcode() string {
	if o == nil || o.Postcode.Get() == nil {
		var ret string
		return ret
	}

	return *o.Postcode.Get()
}

// GetPostcodeOk returns a tuple with the Postcode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetPostcodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Postcode.Get(), o.Postcode.IsSet()
}

// SetPostcode sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) SetPostcode(v string) {
	o.Postcode.Set(&v)
}


// GetCountryCode returns the CountryCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetCountryCode() string {
	if o == nil || o.CountryCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// SetCountryCode sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}


// GetLatitude returns the Latitude field value
// If the value is explicit nil, the zero value for string will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetLatitude() string {
	if o == nil || o.Latitude.Get() == nil {
		var ret string
		return ret
	}

	return *o.Latitude.Get()
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetLatitudeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latitude.Get(), o.Latitude.IsSet()
}

// SetLatitude sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) SetLatitude(v string) {
	o.Latitude.Set(&v)
}


// GetLongitude returns the Longitude field value
// If the value is explicit nil, the zero value for string will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetLongitude() string {
	if o == nil || o.Longitude.Get() == nil {
		var ret string
		return ret
	}

	return *o.Longitude.Get()
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetLongitudeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Longitude.Get(), o.Longitude.IsSet()
}

// SetLongitude sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) SetLongitude(v string) {
	o.Longitude.Set(&v)
}


// GetAttributeType returns the AttributeType field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) SetAttributeType(v string) {
	o.AttributeType = v
}


func (o V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active_from"] = o.ActiveFrom
	toSerialize["active_until"] = o.ActiveUntil.Get()
	toSerialize["created_by_actor"] = o.CreatedByActor
	toSerialize["line_1"] = o.Line1.Get()
	toSerialize["line_2"] = o.Line2.Get()
	toSerialize["line_3"] = o.Line3.Get()
	toSerialize["line_4"] = o.Line4.Get()
	toSerialize["locality"] = o.Locality.Get()
	toSerialize["region"] = o.Region.Get()
	toSerialize["postcode"] = o.Postcode.Get()
	toSerialize["country_code"] = o.CountryCode.Get()
	toSerialize["latitude"] = o.Latitude.Get()
	toSerialize["longitude"] = o.Longitude.Get()
	toSerialize["attribute_type"] = o.AttributeType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active_from",
		"active_until",
		"created_by_actor",
		"line_1",
		"line_2",
		"line_3",
		"line_4",
		"locality",
		"region",
		"postcode",
		"country_code",
		"latitude",
		"longitude",
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
	varV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8 := _V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8{}

	err = json.Unmarshal(data, &varV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8)

	if err != nil {
		return err
	}

	*o = V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8(varV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active_from")
		delete(additionalProperties, "active_until")
		delete(additionalProperties, "created_by_actor")
		delete(additionalProperties, "line_1")
		delete(additionalProperties, "line_2")
		delete(additionalProperties, "line_3")
		delete(additionalProperties, "line_4")
		delete(additionalProperties, "locality")
		delete(additionalProperties, "region")
		delete(additionalProperties, "postcode")
		delete(additionalProperties, "country_code")
		delete(additionalProperties, "latitude")
		delete(additionalProperties, "longitude")
		delete(additionalProperties, "attribute_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8 struct {
	value *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8
	isSet bool
}

func (v NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) Get() *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8 {
	return v.value
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) Set(val *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8(val *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8 {
	return &NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8{value: val, isSet: true}
}

func (v NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf8) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


