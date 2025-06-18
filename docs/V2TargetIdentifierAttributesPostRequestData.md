# V2TargetIdentifierAttributesPostRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The name of the attribute. The title will be visible across Attio&#39;s UI. | 
**Description** | **string** | A text description for the attribute. | 
**ApiSlug** | **string** | A unique, human-readable slug to access the attribute through URLs and API calls. Formatted in snake case. | 
**Type** | **string** | The type of the attribute. This value affects the possible &#x60;config&#x60; values. Attributes of type \&quot;status\&quot; are not supported on objects. | 
**IsRequired** | **bool** | When &#x60;is_required&#x60; is &#x60;true&#x60;, new records/entries must have a value for this attribute. If &#x60;false&#x60;, values may be &#x60;null&#x60;. This value does not affect existing data and you do not need to backfill &#x60;null&#x60; values if changing &#x60;is_required&#x60; from &#x60;false&#x60; to &#x60;true&#x60;. | 
**IsUnique** | **bool** | Whether or not new values for this attribute must be unique. Uniqueness restrictions are only applied to new data and do not apply retroactively to previously created data. | 
**IsMultiselect** | **bool** | Whether or not this attribute can have multiple values. Multiselect is only available on some value types. | 
**DefaultValue** | Pointer to [**V2TargetIdentifierAttributesPostRequestDataDefaultValue**](V2TargetIdentifierAttributesPostRequestDataDefaultValue.md) |  | [optional] 
**Config** | [**V2TargetIdentifierAttributesPostRequestDataConfig**](V2TargetIdentifierAttributesPostRequestDataConfig.md) |  | 

## Methods

### NewV2TargetIdentifierAttributesPostRequestData

`func NewV2TargetIdentifierAttributesPostRequestData(title string, description string, apiSlug string, type_ string, isRequired bool, isUnique bool, isMultiselect bool, config V2TargetIdentifierAttributesPostRequestDataConfig, ) *V2TargetIdentifierAttributesPostRequestData`

NewV2TargetIdentifierAttributesPostRequestData instantiates a new V2TargetIdentifierAttributesPostRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TargetIdentifierAttributesPostRequestDataWithDefaults

`func NewV2TargetIdentifierAttributesPostRequestDataWithDefaults() *V2TargetIdentifierAttributesPostRequestData`

NewV2TargetIdentifierAttributesPostRequestDataWithDefaults instantiates a new V2TargetIdentifierAttributesPostRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *V2TargetIdentifierAttributesPostRequestData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V2TargetIdentifierAttributesPostRequestData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V2TargetIdentifierAttributesPostRequestData) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *V2TargetIdentifierAttributesPostRequestData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V2TargetIdentifierAttributesPostRequestData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V2TargetIdentifierAttributesPostRequestData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetApiSlug

`func (o *V2TargetIdentifierAttributesPostRequestData) GetApiSlug() string`

GetApiSlug returns the ApiSlug field if non-nil, zero value otherwise.

### GetApiSlugOk

`func (o *V2TargetIdentifierAttributesPostRequestData) GetApiSlugOk() (*string, bool)`

GetApiSlugOk returns a tuple with the ApiSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSlug

`func (o *V2TargetIdentifierAttributesPostRequestData) SetApiSlug(v string)`

SetApiSlug sets ApiSlug field to given value.


### GetType

`func (o *V2TargetIdentifierAttributesPostRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2TargetIdentifierAttributesPostRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2TargetIdentifierAttributesPostRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetIsRequired

`func (o *V2TargetIdentifierAttributesPostRequestData) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *V2TargetIdentifierAttributesPostRequestData) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *V2TargetIdentifierAttributesPostRequestData) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.


### GetIsUnique

`func (o *V2TargetIdentifierAttributesPostRequestData) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *V2TargetIdentifierAttributesPostRequestData) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *V2TargetIdentifierAttributesPostRequestData) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.


### GetIsMultiselect

`func (o *V2TargetIdentifierAttributesPostRequestData) GetIsMultiselect() bool`

GetIsMultiselect returns the IsMultiselect field if non-nil, zero value otherwise.

### GetIsMultiselectOk

`func (o *V2TargetIdentifierAttributesPostRequestData) GetIsMultiselectOk() (*bool, bool)`

GetIsMultiselectOk returns a tuple with the IsMultiselect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiselect

`func (o *V2TargetIdentifierAttributesPostRequestData) SetIsMultiselect(v bool)`

SetIsMultiselect sets IsMultiselect field to given value.


### GetDefaultValue

`func (o *V2TargetIdentifierAttributesPostRequestData) GetDefaultValue() V2TargetIdentifierAttributesPostRequestDataDefaultValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *V2TargetIdentifierAttributesPostRequestData) GetDefaultValueOk() (*V2TargetIdentifierAttributesPostRequestDataDefaultValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *V2TargetIdentifierAttributesPostRequestData) SetDefaultValue(v V2TargetIdentifierAttributesPostRequestDataDefaultValue)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *V2TargetIdentifierAttributesPostRequestData) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetConfig

`func (o *V2TargetIdentifierAttributesPostRequestData) GetConfig() V2TargetIdentifierAttributesPostRequestDataConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V2TargetIdentifierAttributesPostRequestData) GetConfigOk() (*V2TargetIdentifierAttributesPostRequestDataConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V2TargetIdentifierAttributesPostRequestData) SetConfig(v V2TargetIdentifierAttributesPostRequestDataConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


