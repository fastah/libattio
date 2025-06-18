# Object

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**ObjectId**](ObjectId.md) |  | 
**ApiSlug** | **NullableString** | A unique, human-readable slug to access the object through URLs and API calls. Formatted in snake case. | 
**SingularNoun** | **NullableString** | The singular form of the object&#39;s name. | 
**PluralNoun** | **NullableString** | The plural form of the object&#39;s name. | 
**CreatedAt** | **string** | When the object was created. | 

## Methods

### NewObject

`func NewObject(id ObjectId, apiSlug NullableString, singularNoun NullableString, pluralNoun NullableString, createdAt string, ) *Object`

NewObject instantiates a new Object object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectWithDefaults

`func NewObjectWithDefaults() *Object`

NewObjectWithDefaults instantiates a new Object object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Object) GetId() ObjectId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Object) GetIdOk() (*ObjectId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Object) SetId(v ObjectId)`

SetId sets Id field to given value.


### GetApiSlug

`func (o *Object) GetApiSlug() string`

GetApiSlug returns the ApiSlug field if non-nil, zero value otherwise.

### GetApiSlugOk

`func (o *Object) GetApiSlugOk() (*string, bool)`

GetApiSlugOk returns a tuple with the ApiSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSlug

`func (o *Object) SetApiSlug(v string)`

SetApiSlug sets ApiSlug field to given value.


### SetApiSlugNil

`func (o *Object) SetApiSlugNil(b bool)`

 SetApiSlugNil sets the value for ApiSlug to be an explicit nil

### UnsetApiSlug
`func (o *Object) UnsetApiSlug()`

UnsetApiSlug ensures that no value is present for ApiSlug, not even an explicit nil
### GetSingularNoun

`func (o *Object) GetSingularNoun() string`

GetSingularNoun returns the SingularNoun field if non-nil, zero value otherwise.

### GetSingularNounOk

`func (o *Object) GetSingularNounOk() (*string, bool)`

GetSingularNounOk returns a tuple with the SingularNoun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingularNoun

`func (o *Object) SetSingularNoun(v string)`

SetSingularNoun sets SingularNoun field to given value.


### SetSingularNounNil

`func (o *Object) SetSingularNounNil(b bool)`

 SetSingularNounNil sets the value for SingularNoun to be an explicit nil

### UnsetSingularNoun
`func (o *Object) UnsetSingularNoun()`

UnsetSingularNoun ensures that no value is present for SingularNoun, not even an explicit nil
### GetPluralNoun

`func (o *Object) GetPluralNoun() string`

GetPluralNoun returns the PluralNoun field if non-nil, zero value otherwise.

### GetPluralNounOk

`func (o *Object) GetPluralNounOk() (*string, bool)`

GetPluralNounOk returns a tuple with the PluralNoun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluralNoun

`func (o *Object) SetPluralNoun(v string)`

SetPluralNoun sets PluralNoun field to given value.


### SetPluralNounNil

`func (o *Object) SetPluralNounNil(b bool)`

 SetPluralNounNil sets the value for PluralNoun to be an explicit nil

### UnsetPluralNoun
`func (o *Object) UnsetPluralNoun()`

UnsetPluralNoun ensures that no value is present for PluralNoun, not even an explicit nil
### GetCreatedAt

`func (o *Object) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Object) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Object) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


