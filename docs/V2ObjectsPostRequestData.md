# V2ObjectsPostRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiSlug** | **string** | A unique, human-readable slug to access the object through URLs and API calls. Should be formatted in snake case. | 
**SingularNoun** | **string** | The singular form of the object&#39;s name. | 
**PluralNoun** | **string** | The plural form of the object&#39;s name. | 

## Methods

### NewV2ObjectsPostRequestData

`func NewV2ObjectsPostRequestData(apiSlug string, singularNoun string, pluralNoun string, ) *V2ObjectsPostRequestData`

NewV2ObjectsPostRequestData instantiates a new V2ObjectsPostRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ObjectsPostRequestDataWithDefaults

`func NewV2ObjectsPostRequestDataWithDefaults() *V2ObjectsPostRequestData`

NewV2ObjectsPostRequestDataWithDefaults instantiates a new V2ObjectsPostRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiSlug

`func (o *V2ObjectsPostRequestData) GetApiSlug() string`

GetApiSlug returns the ApiSlug field if non-nil, zero value otherwise.

### GetApiSlugOk

`func (o *V2ObjectsPostRequestData) GetApiSlugOk() (*string, bool)`

GetApiSlugOk returns a tuple with the ApiSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSlug

`func (o *V2ObjectsPostRequestData) SetApiSlug(v string)`

SetApiSlug sets ApiSlug field to given value.


### GetSingularNoun

`func (o *V2ObjectsPostRequestData) GetSingularNoun() string`

GetSingularNoun returns the SingularNoun field if non-nil, zero value otherwise.

### GetSingularNounOk

`func (o *V2ObjectsPostRequestData) GetSingularNounOk() (*string, bool)`

GetSingularNounOk returns a tuple with the SingularNoun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingularNoun

`func (o *V2ObjectsPostRequestData) SetSingularNoun(v string)`

SetSingularNoun sets SingularNoun field to given value.


### GetPluralNoun

`func (o *V2ObjectsPostRequestData) GetPluralNoun() string`

GetPluralNoun returns the PluralNoun field if non-nil, zero value otherwise.

### GetPluralNounOk

`func (o *V2ObjectsPostRequestData) GetPluralNounOk() (*string, bool)`

GetPluralNounOk returns a tuple with the PluralNoun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluralNoun

`func (o *V2ObjectsPostRequestData) SetPluralNoun(v string)`

SetPluralNoun sets PluralNoun field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


