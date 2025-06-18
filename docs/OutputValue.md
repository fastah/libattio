# OutputValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferencedActorType** | **string** | The type of the referenced actor. [Read more information on actor types here](/docs/actors). | 
**ReferencedActorId** | **string** | The ID of the referenced actor. | 
**AttributeType** | **string** | The attribute type of the value. | 
**Value** | **string** | A timestamp value represents a single, universal moment in time using an ISO 8601 formatted string. This means that a timestamp consists of a date, a time (with nanosecond precision), and a time zone. Attio will coerce timestamps which do not provide full nanosecond precision and UTC is assumed if no time zone is provided. For example, \&quot;2023\&quot;, \&quot;2023-01\&quot;, \&quot;2023-01-02\&quot;, \&quot;2023-01-02T13:00\&quot;, \&quot;2023-01-02T13:00:00\&quot;, and \&quot;2023-01-02T13:00:00.000000000\&quot; will all be coerced to \&quot;2023-01-02T13:00:00.000000000Z\&quot;. Timestamps are always returned in UTC. For example, writing a timestamp value using the string \&quot;2023-01-02T13:00:00.000000000+02:00\&quot; will result in the value \&quot;2023-01-02T11:00:00.000000000Z\&quot; being returned. The maximum date is \&quot;9999-12-31T23:59:59.999999999Z\&quot;. | 
**CurrencyValue** | **float32** | A numerical representation of the currency value. A decimal with a max of 4 decimal places. | 
**CurrencyCode** | Pointer to **string** | The ISO4217 currency code representing the currency that the value is stored in. | [optional] 
**Domain** | **string** |  | 
**RootDomain** | **string** |  | 
**OriginalEmailAddress** | **string** |  | 
**EmailAddress** | **string** |  | 
**EmailDomain** | **string** |  | 
**EmailRootDomain** | **string** |  | 
**EmailLocalSpecifier** | **string** |  | 
**TargetObject** | **string** | A slug identifying the object that the referenced record belongs to. | 
**TargetRecordId** | **string** | A UUID to identify the referenced record. | 
**InteractionType** | **string** | The type of interaction e.g. calendar or email. | 
**InteractedAt** | **time.Time** | When the interaction occurred. | 
**OwnerActor** | [**V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor**](V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor.md) |  | 
**Line1** | **string** | The first line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls. | 
**Line2** | **string** | The second line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls. | 
**Line3** | **string** | The third line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls. | 
**Line4** | **string** | The fourth line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls. | 
**Locality** | **string** | The town, neighborhood or area the location is in. | 
**Region** | **string** | The state, county, province or region that the location is in. | 
**Postcode** | **string** | The postcode or zip code for the location. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.} | 
**CountryCode** | **string** | The ISO 3166-1 alpha-2 country code representing the country that this phone number belongs to. | 
**Latitude** | **string** | The latitude of the location. Validated by the regular expression &#x60;/^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$/&#x60;. Values are stored with up to 9 decimal places of precision. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.} | 
**Longitude** | **string** | The longitude of the location. Validated by the regular expression &#x60;/^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$/&#x60;. Values are stored with up to 9 decimal places of precision. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.} | 
**FirstName** | **string** | The first name. | 
**LastName** | **string** | The last name. | 
**FullName** | **string** | The full name. | 
**OriginalPhoneNumber** | **string** | The raw, original phone number, as inputted. | 
**PhoneNumber** | **string** |  | 
**Status** | [**Status**](Status.md) |  | 
**Option** | [**SelectOption**](SelectOption.md) |  | 

## Methods

### NewOutputValue

`func NewOutputValue(referencedActorType string, referencedActorId string, attributeType string, value string, currencyValue float32, domain string, rootDomain string, originalEmailAddress string, emailAddress string, emailDomain string, emailRootDomain string, emailLocalSpecifier string, targetObject string, targetRecordId string, interactionType string, interactedAt time.Time, ownerActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor, line1 string, line2 string, line3 string, line4 string, locality string, region string, postcode string, countryCode string, latitude string, longitude string, firstName string, lastName string, fullName string, originalPhoneNumber string, phoneNumber string, status Status, option SelectOption, ) *OutputValue`

NewOutputValue instantiates a new OutputValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputValueWithDefaults

`func NewOutputValueWithDefaults() *OutputValue`

NewOutputValueWithDefaults instantiates a new OutputValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferencedActorType

`func (o *OutputValue) GetReferencedActorType() string`

GetReferencedActorType returns the ReferencedActorType field if non-nil, zero value otherwise.

### GetReferencedActorTypeOk

`func (o *OutputValue) GetReferencedActorTypeOk() (*string, bool)`

GetReferencedActorTypeOk returns a tuple with the ReferencedActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedActorType

`func (o *OutputValue) SetReferencedActorType(v string)`

SetReferencedActorType sets ReferencedActorType field to given value.


### GetReferencedActorId

`func (o *OutputValue) GetReferencedActorId() string`

GetReferencedActorId returns the ReferencedActorId field if non-nil, zero value otherwise.

### GetReferencedActorIdOk

`func (o *OutputValue) GetReferencedActorIdOk() (*string, bool)`

GetReferencedActorIdOk returns a tuple with the ReferencedActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedActorId

`func (o *OutputValue) SetReferencedActorId(v string)`

SetReferencedActorId sets ReferencedActorId field to given value.


### GetAttributeType

`func (o *OutputValue) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *OutputValue) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *OutputValue) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetValue

`func (o *OutputValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OutputValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OutputValue) SetValue(v string)`

SetValue sets Value field to given value.


### GetCurrencyValue

`func (o *OutputValue) GetCurrencyValue() float32`

GetCurrencyValue returns the CurrencyValue field if non-nil, zero value otherwise.

### GetCurrencyValueOk

`func (o *OutputValue) GetCurrencyValueOk() (*float32, bool)`

GetCurrencyValueOk returns a tuple with the CurrencyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyValue

`func (o *OutputValue) SetCurrencyValue(v float32)`

SetCurrencyValue sets CurrencyValue field to given value.


### GetCurrencyCode

`func (o *OutputValue) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OutputValue) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OutputValue) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *OutputValue) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDomain

`func (o *OutputValue) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OutputValue) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OutputValue) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetRootDomain

`func (o *OutputValue) GetRootDomain() string`

GetRootDomain returns the RootDomain field if non-nil, zero value otherwise.

### GetRootDomainOk

`func (o *OutputValue) GetRootDomainOk() (*string, bool)`

GetRootDomainOk returns a tuple with the RootDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDomain

`func (o *OutputValue) SetRootDomain(v string)`

SetRootDomain sets RootDomain field to given value.


### GetOriginalEmailAddress

`func (o *OutputValue) GetOriginalEmailAddress() string`

GetOriginalEmailAddress returns the OriginalEmailAddress field if non-nil, zero value otherwise.

### GetOriginalEmailAddressOk

`func (o *OutputValue) GetOriginalEmailAddressOk() (*string, bool)`

GetOriginalEmailAddressOk returns a tuple with the OriginalEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalEmailAddress

`func (o *OutputValue) SetOriginalEmailAddress(v string)`

SetOriginalEmailAddress sets OriginalEmailAddress field to given value.


### GetEmailAddress

`func (o *OutputValue) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *OutputValue) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *OutputValue) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetEmailDomain

`func (o *OutputValue) GetEmailDomain() string`

GetEmailDomain returns the EmailDomain field if non-nil, zero value otherwise.

### GetEmailDomainOk

`func (o *OutputValue) GetEmailDomainOk() (*string, bool)`

GetEmailDomainOk returns a tuple with the EmailDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomain

`func (o *OutputValue) SetEmailDomain(v string)`

SetEmailDomain sets EmailDomain field to given value.


### GetEmailRootDomain

`func (o *OutputValue) GetEmailRootDomain() string`

GetEmailRootDomain returns the EmailRootDomain field if non-nil, zero value otherwise.

### GetEmailRootDomainOk

`func (o *OutputValue) GetEmailRootDomainOk() (*string, bool)`

GetEmailRootDomainOk returns a tuple with the EmailRootDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRootDomain

`func (o *OutputValue) SetEmailRootDomain(v string)`

SetEmailRootDomain sets EmailRootDomain field to given value.


### GetEmailLocalSpecifier

`func (o *OutputValue) GetEmailLocalSpecifier() string`

GetEmailLocalSpecifier returns the EmailLocalSpecifier field if non-nil, zero value otherwise.

### GetEmailLocalSpecifierOk

`func (o *OutputValue) GetEmailLocalSpecifierOk() (*string, bool)`

GetEmailLocalSpecifierOk returns a tuple with the EmailLocalSpecifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailLocalSpecifier

`func (o *OutputValue) SetEmailLocalSpecifier(v string)`

SetEmailLocalSpecifier sets EmailLocalSpecifier field to given value.


### GetTargetObject

`func (o *OutputValue) GetTargetObject() string`

GetTargetObject returns the TargetObject field if non-nil, zero value otherwise.

### GetTargetObjectOk

`func (o *OutputValue) GetTargetObjectOk() (*string, bool)`

GetTargetObjectOk returns a tuple with the TargetObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetObject

`func (o *OutputValue) SetTargetObject(v string)`

SetTargetObject sets TargetObject field to given value.


### GetTargetRecordId

`func (o *OutputValue) GetTargetRecordId() string`

GetTargetRecordId returns the TargetRecordId field if non-nil, zero value otherwise.

### GetTargetRecordIdOk

`func (o *OutputValue) GetTargetRecordIdOk() (*string, bool)`

GetTargetRecordIdOk returns a tuple with the TargetRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRecordId

`func (o *OutputValue) SetTargetRecordId(v string)`

SetTargetRecordId sets TargetRecordId field to given value.


### GetInteractionType

`func (o *OutputValue) GetInteractionType() string`

GetInteractionType returns the InteractionType field if non-nil, zero value otherwise.

### GetInteractionTypeOk

`func (o *OutputValue) GetInteractionTypeOk() (*string, bool)`

GetInteractionTypeOk returns a tuple with the InteractionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionType

`func (o *OutputValue) SetInteractionType(v string)`

SetInteractionType sets InteractionType field to given value.


### GetInteractedAt

`func (o *OutputValue) GetInteractedAt() time.Time`

GetInteractedAt returns the InteractedAt field if non-nil, zero value otherwise.

### GetInteractedAtOk

`func (o *OutputValue) GetInteractedAtOk() (*time.Time, bool)`

GetInteractedAtOk returns a tuple with the InteractedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractedAt

`func (o *OutputValue) SetInteractedAt(v time.Time)`

SetInteractedAt sets InteractedAt field to given value.


### GetOwnerActor

`func (o *OutputValue) GetOwnerActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor`

GetOwnerActor returns the OwnerActor field if non-nil, zero value otherwise.

### GetOwnerActorOk

`func (o *OutputValue) GetOwnerActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor, bool)`

GetOwnerActorOk returns a tuple with the OwnerActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerActor

`func (o *OutputValue) SetOwnerActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor)`

SetOwnerActor sets OwnerActor field to given value.


### GetLine1

`func (o *OutputValue) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *OutputValue) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *OutputValue) SetLine1(v string)`

SetLine1 sets Line1 field to given value.


### GetLine2

`func (o *OutputValue) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *OutputValue) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *OutputValue) SetLine2(v string)`

SetLine2 sets Line2 field to given value.


### GetLine3

`func (o *OutputValue) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *OutputValue) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *OutputValue) SetLine3(v string)`

SetLine3 sets Line3 field to given value.


### GetLine4

`func (o *OutputValue) GetLine4() string`

GetLine4 returns the Line4 field if non-nil, zero value otherwise.

### GetLine4Ok

`func (o *OutputValue) GetLine4Ok() (*string, bool)`

GetLine4Ok returns a tuple with the Line4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine4

`func (o *OutputValue) SetLine4(v string)`

SetLine4 sets Line4 field to given value.


### GetLocality

`func (o *OutputValue) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *OutputValue) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *OutputValue) SetLocality(v string)`

SetLocality sets Locality field to given value.


### GetRegion

`func (o *OutputValue) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OutputValue) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OutputValue) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetPostcode

`func (o *OutputValue) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *OutputValue) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *OutputValue) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.


### GetCountryCode

`func (o *OutputValue) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *OutputValue) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *OutputValue) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetLatitude

`func (o *OutputValue) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *OutputValue) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *OutputValue) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *OutputValue) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *OutputValue) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *OutputValue) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.


### GetFirstName

`func (o *OutputValue) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *OutputValue) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *OutputValue) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *OutputValue) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *OutputValue) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *OutputValue) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetFullName

`func (o *OutputValue) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *OutputValue) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *OutputValue) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetOriginalPhoneNumber

`func (o *OutputValue) GetOriginalPhoneNumber() string`

GetOriginalPhoneNumber returns the OriginalPhoneNumber field if non-nil, zero value otherwise.

### GetOriginalPhoneNumberOk

`func (o *OutputValue) GetOriginalPhoneNumberOk() (*string, bool)`

GetOriginalPhoneNumberOk returns a tuple with the OriginalPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPhoneNumber

`func (o *OutputValue) SetOriginalPhoneNumber(v string)`

SetOriginalPhoneNumber sets OriginalPhoneNumber field to given value.


### GetPhoneNumber

`func (o *OutputValue) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *OutputValue) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *OutputValue) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetStatus

`func (o *OutputValue) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OutputValue) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OutputValue) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetOption

`func (o *OutputValue) GetOption() SelectOption`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *OutputValue) GetOptionOk() (*SelectOption, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *OutputValue) SetOption(v SelectOption)`

SetOption sets Option field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


