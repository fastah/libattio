# V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveFrom** | **time.Time** | The point in time at which this value was made \&quot;active\&quot;. &#x60;active_from&#x60; can be considered roughly analogous to &#x60;created_at&#x60;. | 
**ActiveUntil** | **NullableTime** | The point in time at which this value was deactivated. If &#x60;null&#x60;, the value is active. | 
**CreatedByActor** | [**V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor**](V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor.md) |  | 
**Line1** | **NullableString** | The first line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls. | 
**Line2** | **NullableString** | The second line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls. | 
**Line3** | **NullableString** | The third line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls. | 
**Line4** | **NullableString** | The fourth line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls. | 
**Locality** | **NullableString** | The town, neighborhood or area the location is in. | 
**Region** | **NullableString** | The state, county, province or region that the location is in. | 
**Postcode** | **NullableString** | The postcode or zip code for the location. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.} | 
**CountryCode** | **NullableString** | The ISO 3166-1 alpha-2 country code for the country this location is in. | 
**Latitude** | **NullableString** | The latitude of the location. Validated by the regular expression &#x60;/^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$/&#x60;. Values are stored with up to 9 decimal places of precision. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.} | 
**Longitude** | **NullableString** | The longitude of the location. Validated by the regular expression &#x60;/^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$/&#x60;. Values are stored with up to 9 decimal places of precision. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.} | 
**AttributeType** | **string** | The attribute type of the value. | 

## Methods

### NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7

`func NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7(activeFrom time.Time, activeUntil NullableTime, createdByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor, line1 NullableString, line2 NullableString, line3 NullableString, line4 NullableString, locality NullableString, region NullableString, postcode NullableString, countryCode NullableString, latitude NullableString, longitude NullableString, attributeType string, ) *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7`

NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7 instantiates a new V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7WithDefaults

`func NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7WithDefaults() *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7`

NewV2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7WithDefaults instantiates a new V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveFrom

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetActiveFrom() time.Time`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetActiveFromOk() (*time.Time, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetActiveFrom(v time.Time)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveUntil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetActiveUntil() time.Time`

GetActiveUntil returns the ActiveUntil field if non-nil, zero value otherwise.

### GetActiveUntilOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetActiveUntilOk() (*time.Time, bool)`

GetActiveUntilOk returns a tuple with the ActiveUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUntil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetActiveUntil(v time.Time)`

SetActiveUntil sets ActiveUntil field to given value.


### SetActiveUntilNil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetActiveUntilNil(b bool)`

 SetActiveUntilNil sets the value for ActiveUntil to be an explicit nil

### UnsetActiveUntil
`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) UnsetActiveUntil()`

UnsetActiveUntil ensures that no value is present for ActiveUntil, not even an explicit nil
### GetCreatedByActor

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetCreatedByActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetCreatedByActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetCreatedByActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf1CreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.


### GetLine1

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetLine1(v string)`

SetLine1 sets Line1 field to given value.


### SetLine1Nil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetLine2(v string)`

SetLine2 sets Line2 field to given value.


### SetLine2Nil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetLine3

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetLine3(v string)`

SetLine3 sets Line3 field to given value.


### SetLine3Nil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetLine3Nil(b bool)`

 SetLine3Nil sets the value for Line3 to be an explicit nil

### UnsetLine3
`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) UnsetLine3()`

UnsetLine3 ensures that no value is present for Line3, not even an explicit nil
### GetLine4

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetLine4() string`

GetLine4 returns the Line4 field if non-nil, zero value otherwise.

### GetLine4Ok

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetLine4Ok() (*string, bool)`

GetLine4Ok returns a tuple with the Line4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine4

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetLine4(v string)`

SetLine4 sets Line4 field to given value.


### SetLine4Nil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetLine4Nil(b bool)`

 SetLine4Nil sets the value for Line4 to be an explicit nil

### UnsetLine4
`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) UnsetLine4()`

UnsetLine4 ensures that no value is present for Line4, not even an explicit nil
### GetLocality

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetLocality(v string)`

SetLocality sets Locality field to given value.


### SetLocalityNil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetLocalityNil(b bool)`

 SetLocalityNil sets the value for Locality to be an explicit nil

### UnsetLocality
`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) UnsetLocality()`

UnsetLocality ensures that no value is present for Locality, not even an explicit nil
### GetRegion

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetPostcode

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.


### SetPostcodeNil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetPostcodeNil(b bool)`

 SetPostcodeNil sets the value for Postcode to be an explicit nil

### UnsetPostcode
`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) UnsetPostcode()`

UnsetPostcode ensures that no value is present for Postcode, not even an explicit nil
### GetCountryCode

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### SetCountryCodeNil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetLatitude

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.


### SetLatitudeNil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.


### SetLongitudeNil

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetAttributeType

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *V2ObjectsObjectRecordsPut200ResponseDataValuesValueInnerAnyOf7) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


