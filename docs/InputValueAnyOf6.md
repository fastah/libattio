# InputValueAnyOf6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### NewInputValueAnyOf6

`func NewInputValueAnyOf6(line1 NullableString, line2 NullableString, line3 NullableString, line4 NullableString, locality NullableString, region NullableString, postcode NullableString, countryCode NullableString, latitude NullableString, longitude NullableString, ) *InputValueAnyOf6`

NewInputValueAnyOf6 instantiates a new InputValueAnyOf6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputValueAnyOf6WithDefaults

`func NewInputValueAnyOf6WithDefaults() *InputValueAnyOf6`

NewInputValueAnyOf6WithDefaults instantiates a new InputValueAnyOf6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine1

`func (o *InputValueAnyOf6) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *InputValueAnyOf6) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *InputValueAnyOf6) SetLine1(v string)`

SetLine1 sets Line1 field to given value.


### SetLine1Nil

`func (o *InputValueAnyOf6) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *InputValueAnyOf6) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *InputValueAnyOf6) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *InputValueAnyOf6) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *InputValueAnyOf6) SetLine2(v string)`

SetLine2 sets Line2 field to given value.


### SetLine2Nil

`func (o *InputValueAnyOf6) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *InputValueAnyOf6) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetLine3

`func (o *InputValueAnyOf6) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *InputValueAnyOf6) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *InputValueAnyOf6) SetLine3(v string)`

SetLine3 sets Line3 field to given value.


### SetLine3Nil

`func (o *InputValueAnyOf6) SetLine3Nil(b bool)`

 SetLine3Nil sets the value for Line3 to be an explicit nil

### UnsetLine3
`func (o *InputValueAnyOf6) UnsetLine3()`

UnsetLine3 ensures that no value is present for Line3, not even an explicit nil
### GetLine4

`func (o *InputValueAnyOf6) GetLine4() string`

GetLine4 returns the Line4 field if non-nil, zero value otherwise.

### GetLine4Ok

`func (o *InputValueAnyOf6) GetLine4Ok() (*string, bool)`

GetLine4Ok returns a tuple with the Line4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine4

`func (o *InputValueAnyOf6) SetLine4(v string)`

SetLine4 sets Line4 field to given value.


### SetLine4Nil

`func (o *InputValueAnyOf6) SetLine4Nil(b bool)`

 SetLine4Nil sets the value for Line4 to be an explicit nil

### UnsetLine4
`func (o *InputValueAnyOf6) UnsetLine4()`

UnsetLine4 ensures that no value is present for Line4, not even an explicit nil
### GetLocality

`func (o *InputValueAnyOf6) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *InputValueAnyOf6) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *InputValueAnyOf6) SetLocality(v string)`

SetLocality sets Locality field to given value.


### SetLocalityNil

`func (o *InputValueAnyOf6) SetLocalityNil(b bool)`

 SetLocalityNil sets the value for Locality to be an explicit nil

### UnsetLocality
`func (o *InputValueAnyOf6) UnsetLocality()`

UnsetLocality ensures that no value is present for Locality, not even an explicit nil
### GetRegion

`func (o *InputValueAnyOf6) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InputValueAnyOf6) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InputValueAnyOf6) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *InputValueAnyOf6) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *InputValueAnyOf6) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetPostcode

`func (o *InputValueAnyOf6) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *InputValueAnyOf6) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *InputValueAnyOf6) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.


### SetPostcodeNil

`func (o *InputValueAnyOf6) SetPostcodeNil(b bool)`

 SetPostcodeNil sets the value for Postcode to be an explicit nil

### UnsetPostcode
`func (o *InputValueAnyOf6) UnsetPostcode()`

UnsetPostcode ensures that no value is present for Postcode, not even an explicit nil
### GetCountryCode

`func (o *InputValueAnyOf6) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *InputValueAnyOf6) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *InputValueAnyOf6) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### SetCountryCodeNil

`func (o *InputValueAnyOf6) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *InputValueAnyOf6) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetLatitude

`func (o *InputValueAnyOf6) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *InputValueAnyOf6) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *InputValueAnyOf6) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.


### SetLatitudeNil

`func (o *InputValueAnyOf6) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *InputValueAnyOf6) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *InputValueAnyOf6) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *InputValueAnyOf6) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *InputValueAnyOf6) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.


### SetLongitudeNil

`func (o *InputValueAnyOf6) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *InputValueAnyOf6) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


