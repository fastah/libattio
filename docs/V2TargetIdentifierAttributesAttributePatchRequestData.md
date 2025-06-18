# V2TargetIdentifierAttributesAttributePatchRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The name of the attribute. The title will be visible across Attio&#39;s UI. | [optional] 
**Description** | Pointer to **string** | A text description for the attribute. | [optional] 
**ApiSlug** | Pointer to **string** | A unique, human-readable slug to access the attribute through URLs and API calls. Formatted in snake case. | [optional] 
**IsRequired** | Pointer to **bool** | When &#x60;is_required&#x60; is &#x60;true&#x60;, new records/entries must have a value for this attribute. If &#x60;false&#x60;, values may be &#x60;null&#x60;. This value does not affect existing data and you do not need to backfill &#x60;null&#x60; values if changing &#x60;is_required&#x60; from &#x60;false&#x60; to &#x60;true&#x60;. | [optional] 
**IsUnique** | Pointer to **bool** | Whether or not new values for this attribute must be unique. Uniqueness restrictions are only applied to new data and do not apply retroactively to previously created data. | [optional] 
**DefaultValue** | Pointer to [**V2TargetIdentifierAttributesAttributePatchRequestDataDefaultValue**](V2TargetIdentifierAttributesAttributePatchRequestDataDefaultValue.md) |  | [optional] 
**Config** | Pointer to [**V2TargetIdentifierAttributesAttributePatchRequestDataConfig**](V2TargetIdentifierAttributesAttributePatchRequestDataConfig.md) |  | [optional] 
**IsArchived** | Pointer to **bool** | Whether the attribute has been archived or not. See our [archiving guide](/docs/archiving-vs-deleting) for more information on archiving. | [optional] 

## Methods

### NewV2TargetIdentifierAttributesAttributePatchRequestData

`func NewV2TargetIdentifierAttributesAttributePatchRequestData() *V2TargetIdentifierAttributesAttributePatchRequestData`

NewV2TargetIdentifierAttributesAttributePatchRequestData instantiates a new V2TargetIdentifierAttributesAttributePatchRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TargetIdentifierAttributesAttributePatchRequestDataWithDefaults

`func NewV2TargetIdentifierAttributesAttributePatchRequestDataWithDefaults() *V2TargetIdentifierAttributesAttributePatchRequestData`

NewV2TargetIdentifierAttributesAttributePatchRequestDataWithDefaults instantiates a new V2TargetIdentifierAttributesAttributePatchRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApiSlug

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetApiSlug() string`

GetApiSlug returns the ApiSlug field if non-nil, zero value otherwise.

### GetApiSlugOk

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetApiSlugOk() (*string, bool)`

GetApiSlugOk returns a tuple with the ApiSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSlug

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetApiSlug(v string)`

SetApiSlug sets ApiSlug field to given value.

### HasApiSlug

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasApiSlug() bool`

HasApiSlug returns a boolean if a field has been set.

### GetIsRequired

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetIsUnique

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.

### HasIsUnique

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasIsUnique() bool`

HasIsUnique returns a boolean if a field has been set.

### GetDefaultValue

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetDefaultValue() V2TargetIdentifierAttributesAttributePatchRequestDataDefaultValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetDefaultValueOk() (*V2TargetIdentifierAttributesAttributePatchRequestDataDefaultValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetDefaultValue(v V2TargetIdentifierAttributesAttributePatchRequestDataDefaultValue)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetConfig

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetConfig() V2TargetIdentifierAttributesAttributePatchRequestDataConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetConfigOk() (*V2TargetIdentifierAttributesAttributePatchRequestDataConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetConfig(v V2TargetIdentifierAttributesAttributePatchRequestDataConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetIsArchived

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


