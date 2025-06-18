# InputValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferencedActorType** | **string** | The type of the referenced actor. Currently, only workspace members can be written into actor reference attributes. [Read more information on actor types here](/docs/actors). | 
**ReferencedActorId** | **string** | The ID of the referenced Actor. | 
**WorkspaceMemberEmailAddress** | **string** | Workspace member actors can be referenced by email address as well as actor ID. | 
**Value** | **string** | A timestamp value represents a single, universal moment in time using an ISO 8601 formatted string. This means that a timestamp consists of a date, a time (with nanosecond precision), and a time zone. Attio will coerce timestamps which do not provide full nanosecond precision and UTC is assumed if no time zone is provided. For example, \&quot;2023\&quot;, \&quot;2023-01\&quot;, \&quot;2023-01-02\&quot;, \&quot;2023-01-02T13:00\&quot;, \&quot;2023-01-02T13:00:00\&quot;, and \&quot;2023-01-02T13:00:00.000000000\&quot; will all be coerced to \&quot;2023-01-02T13:00:00.000000000Z\&quot;. Timestamps are always returned in UTC. For example, writing a timestamp value using the string \&quot;2023-01-02T13:00:00.000000000+02:00\&quot; will result in the value \&quot;2023-01-02T11:00:00.000000000Z\&quot; being returned. The maximum date is \&quot;9999-12-31T23:59:59.999999999Z\&quot;. | 
**CurrencyValue** | **float32** | A numerical representation of the currency value. A decimal with a max of 4 decimal places. | 
**Domain** | Pointer to **string** | The full domain of the website. | [optional] 
**EmailAddress** | Pointer to **string** | An email address string | [optional] 
**TargetObject** | **string** | A UUID or slug to identify the object that the referenced record belongs to. | 
**TargetRecordId** | **string** | A UUID to identify the referenced record. | 
**SlugOrIdOfMatchingAttribute** | [**[]V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner**](V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner.md) | In addition to referencing records directly by record ID, you may also reference by a matching attribute of your choice. For example, if you want to add a reference to the person record with email \&quot;alice@website.com\&quot;, you should pass a value with &#x60;target_object&#x60; set to &#x60;\&quot;people\&quot;&#x60; and &#x60;email_addresses&#x60; set to &#x60;[{email_address:\&quot;alice@website.com\&quot;}]&#x60;. The key should be the slug or ID of the matching attribute you would like to use and the value should be an array containing a single value of the appropriate attribute type (as specified below). Matching on multiple values is not currently supported. Matching attributes must be unique. This process is similar to how you use the &#x60;matching_attribute&#x60; query param in Attio&#39;s [assert endpoints](/rest-api/endpoint-reference/records/assert-a-record). | 
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
**CountryCode** | **string** | The ISO 3166-1 alpha-2 country code representing the country that this phone number belongs to. Optional if &#x60;original_phone_number&#x60; includes a country code prefix. | 
**Latitude** | **string** | The latitude of the location. Validated by the regular expression &#x60;/^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$/&#x60;. Values are stored with up to 9 decimal places of precision. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.} | 
**Longitude** | **string** | The longitude of the location. Validated by the regular expression &#x60;/^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$/&#x60;. Values are stored with up to 9 decimal places of precision. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.} | 
**FirstName** | Pointer to **string** | The first name. | [optional] 
**LastName** | Pointer to **string** | The last name. | [optional] 
**FullName** | Pointer to **string** | The full name. | [optional] 
**OriginalPhoneNumber** | **string** | A phone number which is either a) prefixed with a country code (e.g. &#x60;+44....&#x60;) or b) a local number, where &#x60;country_code&#x60; is specified in addition. | 
**Status** | **string** | The UUID or status title identifying the selected status. | 
**Option** | **string** | The UUID or select option title identifying the selected select option. | 

## Methods

### NewInputValue

`func NewInputValue(referencedActorType string, referencedActorId string, workspaceMemberEmailAddress string, value string, currencyValue float32, targetObject string, targetRecordId string, slugOrIdOfMatchingAttribute []V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner, interactionType string, interactedAt time.Time, ownerActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor, line1 string, line2 string, line3 string, line4 string, locality string, region string, postcode string, countryCode string, latitude string, longitude string, originalPhoneNumber string, status string, option string, ) *InputValue`

NewInputValue instantiates a new InputValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputValueWithDefaults

`func NewInputValueWithDefaults() *InputValue`

NewInputValueWithDefaults instantiates a new InputValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferencedActorType

`func (o *InputValue) GetReferencedActorType() string`

GetReferencedActorType returns the ReferencedActorType field if non-nil, zero value otherwise.

### GetReferencedActorTypeOk

`func (o *InputValue) GetReferencedActorTypeOk() (*string, bool)`

GetReferencedActorTypeOk returns a tuple with the ReferencedActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedActorType

`func (o *InputValue) SetReferencedActorType(v string)`

SetReferencedActorType sets ReferencedActorType field to given value.


### GetReferencedActorId

`func (o *InputValue) GetReferencedActorId() string`

GetReferencedActorId returns the ReferencedActorId field if non-nil, zero value otherwise.

### GetReferencedActorIdOk

`func (o *InputValue) GetReferencedActorIdOk() (*string, bool)`

GetReferencedActorIdOk returns a tuple with the ReferencedActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedActorId

`func (o *InputValue) SetReferencedActorId(v string)`

SetReferencedActorId sets ReferencedActorId field to given value.


### GetWorkspaceMemberEmailAddress

`func (o *InputValue) GetWorkspaceMemberEmailAddress() string`

GetWorkspaceMemberEmailAddress returns the WorkspaceMemberEmailAddress field if non-nil, zero value otherwise.

### GetWorkspaceMemberEmailAddressOk

`func (o *InputValue) GetWorkspaceMemberEmailAddressOk() (*string, bool)`

GetWorkspaceMemberEmailAddressOk returns a tuple with the WorkspaceMemberEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceMemberEmailAddress

`func (o *InputValue) SetWorkspaceMemberEmailAddress(v string)`

SetWorkspaceMemberEmailAddress sets WorkspaceMemberEmailAddress field to given value.


### GetValue

`func (o *InputValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InputValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InputValue) SetValue(v string)`

SetValue sets Value field to given value.


### GetCurrencyValue

`func (o *InputValue) GetCurrencyValue() float32`

GetCurrencyValue returns the CurrencyValue field if non-nil, zero value otherwise.

### GetCurrencyValueOk

`func (o *InputValue) GetCurrencyValueOk() (*float32, bool)`

GetCurrencyValueOk returns a tuple with the CurrencyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyValue

`func (o *InputValue) SetCurrencyValue(v float32)`

SetCurrencyValue sets CurrencyValue field to given value.


### GetDomain

`func (o *InputValue) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *InputValue) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *InputValue) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *InputValue) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEmailAddress

`func (o *InputValue) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *InputValue) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *InputValue) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *InputValue) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetTargetObject

`func (o *InputValue) GetTargetObject() string`

GetTargetObject returns the TargetObject field if non-nil, zero value otherwise.

### GetTargetObjectOk

`func (o *InputValue) GetTargetObjectOk() (*string, bool)`

GetTargetObjectOk returns a tuple with the TargetObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetObject

`func (o *InputValue) SetTargetObject(v string)`

SetTargetObject sets TargetObject field to given value.


### GetTargetRecordId

`func (o *InputValue) GetTargetRecordId() string`

GetTargetRecordId returns the TargetRecordId field if non-nil, zero value otherwise.

### GetTargetRecordIdOk

`func (o *InputValue) GetTargetRecordIdOk() (*string, bool)`

GetTargetRecordIdOk returns a tuple with the TargetRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRecordId

`func (o *InputValue) SetTargetRecordId(v string)`

SetTargetRecordId sets TargetRecordId field to given value.


### GetSlugOrIdOfMatchingAttribute

`func (o *InputValue) GetSlugOrIdOfMatchingAttribute() []V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner`

GetSlugOrIdOfMatchingAttribute returns the SlugOrIdOfMatchingAttribute field if non-nil, zero value otherwise.

### GetSlugOrIdOfMatchingAttributeOk

`func (o *InputValue) GetSlugOrIdOfMatchingAttributeOk() (*[]V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner, bool)`

GetSlugOrIdOfMatchingAttributeOk returns a tuple with the SlugOrIdOfMatchingAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugOrIdOfMatchingAttribute

`func (o *InputValue) SetSlugOrIdOfMatchingAttribute(v []V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner)`

SetSlugOrIdOfMatchingAttribute sets SlugOrIdOfMatchingAttribute field to given value.


### GetInteractionType

`func (o *InputValue) GetInteractionType() string`

GetInteractionType returns the InteractionType field if non-nil, zero value otherwise.

### GetInteractionTypeOk

`func (o *InputValue) GetInteractionTypeOk() (*string, bool)`

GetInteractionTypeOk returns a tuple with the InteractionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionType

`func (o *InputValue) SetInteractionType(v string)`

SetInteractionType sets InteractionType field to given value.


### GetInteractedAt

`func (o *InputValue) GetInteractedAt() time.Time`

GetInteractedAt returns the InteractedAt field if non-nil, zero value otherwise.

### GetInteractedAtOk

`func (o *InputValue) GetInteractedAtOk() (*time.Time, bool)`

GetInteractedAtOk returns a tuple with the InteractedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractedAt

`func (o *InputValue) SetInteractedAt(v time.Time)`

SetInteractedAt sets InteractedAt field to given value.


### GetOwnerActor

`func (o *InputValue) GetOwnerActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor`

GetOwnerActor returns the OwnerActor field if non-nil, zero value otherwise.

### GetOwnerActorOk

`func (o *InputValue) GetOwnerActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor, bool)`

GetOwnerActorOk returns a tuple with the OwnerActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerActor

`func (o *InputValue) SetOwnerActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor)`

SetOwnerActor sets OwnerActor field to given value.


### GetLine1

`func (o *InputValue) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *InputValue) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *InputValue) SetLine1(v string)`

SetLine1 sets Line1 field to given value.


### GetLine2

`func (o *InputValue) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *InputValue) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *InputValue) SetLine2(v string)`

SetLine2 sets Line2 field to given value.


### GetLine3

`func (o *InputValue) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *InputValue) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *InputValue) SetLine3(v string)`

SetLine3 sets Line3 field to given value.


### GetLine4

`func (o *InputValue) GetLine4() string`

GetLine4 returns the Line4 field if non-nil, zero value otherwise.

### GetLine4Ok

`func (o *InputValue) GetLine4Ok() (*string, bool)`

GetLine4Ok returns a tuple with the Line4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine4

`func (o *InputValue) SetLine4(v string)`

SetLine4 sets Line4 field to given value.


### GetLocality

`func (o *InputValue) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *InputValue) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *InputValue) SetLocality(v string)`

SetLocality sets Locality field to given value.


### GetRegion

`func (o *InputValue) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InputValue) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InputValue) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetPostcode

`func (o *InputValue) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *InputValue) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *InputValue) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.


### GetCountryCode

`func (o *InputValue) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *InputValue) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *InputValue) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetLatitude

`func (o *InputValue) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *InputValue) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *InputValue) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *InputValue) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *InputValue) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *InputValue) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.


### GetFirstName

`func (o *InputValue) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InputValue) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InputValue) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *InputValue) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *InputValue) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InputValue) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InputValue) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InputValue) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetFullName

`func (o *InputValue) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *InputValue) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *InputValue) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *InputValue) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetOriginalPhoneNumber

`func (o *InputValue) GetOriginalPhoneNumber() string`

GetOriginalPhoneNumber returns the OriginalPhoneNumber field if non-nil, zero value otherwise.

### GetOriginalPhoneNumberOk

`func (o *InputValue) GetOriginalPhoneNumberOk() (*string, bool)`

GetOriginalPhoneNumberOk returns a tuple with the OriginalPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPhoneNumber

`func (o *InputValue) SetOriginalPhoneNumber(v string)`

SetOriginalPhoneNumber sets OriginalPhoneNumber field to given value.


### GetStatus

`func (o *InputValue) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InputValue) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InputValue) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetOption

`func (o *InputValue) GetOption() string`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *InputValue) GetOptionOk() (*string, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *InputValue) SetOption(v string)`

SetOption sets Option field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


