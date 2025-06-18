# Attribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**AttributeId**](AttributeId.md) |  | 
**Title** | **string** | A title for the attribute, to be displayed across the app. | 
**Description** | **NullableString** | A text description of the attribute. | 
**ApiSlug** | **string** | A unique slug for the attribute for use in API responses and URLs. Formatted in snake case. | 
**Type** | **string** | The type of the attribute. | 
**IsSystemAttribute** | **bool** | &#x60;true&#x60; if this is an Attio system-defined attribute, &#x60;false&#x60; if defined by a user or non-Attio system. | 
**IsWritable** | **bool** | Whether or not this attribute can be written to. Can only be false when &#x60;is_system_attribute&#x60; is &#x60;true&#x60; (user-defined attributes are always writeable). If &#x60;false&#x60;, this usually means the attribute is enriched by Attio. | 
**IsRequired** | **bool** | When &#x60;is_required&#x60; is &#x60;true&#x60;, new records/entries must have a value for this attribute. If &#x60;false&#x60;, values may be &#x60;null&#x60;. This value does not affect existing data and you do not need to backfill &#x60;null&#x60; values if changing &#x60;is_required&#x60; from &#x60;false&#x60; to &#x60;true&#x60;. | 
**IsUnique** | **bool** | Whether or not new values for this attribute must be unique. Uniqueness restrictions are only applied to new data and do not apply retroactively to previously created data. | 
**IsMultiselect** | **bool** | Whether or not this attribute can have multiple values. Multiselect is only available on some value types. | 
**IsDefaultValueEnabled** | **bool** | Whether this attribute has a default value enabled. Must be &#x60;true&#x60; when &#x60;is_required&#x60; is &#x60;true&#x60;. | 
**IsArchived** | **bool** | Whether this attribute has been archived. Archived attributes are hidden from most UI, but can be restored either over the API or in workspace settings. See the [guide on archiving and deleting](/docs/archiving-vs-deleting)for more information. | 
**DefaultValue** | [**NullableAttributeDefaultValue**](AttributeDefaultValue.md) |  | 
**Relationship** | [**AttributeRelationship**](AttributeRelationship.md) |  | 
**CreatedAt** | **string** | When this attribute was created. | 
**Config** | [**AttributeConfig**](AttributeConfig.md) |  | 

## Methods

### NewAttribute

`func NewAttribute(id AttributeId, title string, description NullableString, apiSlug string, type_ string, isSystemAttribute bool, isWritable bool, isRequired bool, isUnique bool, isMultiselect bool, isDefaultValueEnabled bool, isArchived bool, defaultValue NullableAttributeDefaultValue, relationship AttributeRelationship, createdAt string, config AttributeConfig, ) *Attribute`

NewAttribute instantiates a new Attribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeWithDefaults

`func NewAttributeWithDefaults() *Attribute`

NewAttributeWithDefaults instantiates a new Attribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Attribute) GetId() AttributeId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Attribute) GetIdOk() (*AttributeId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Attribute) SetId(v AttributeId)`

SetId sets Id field to given value.


### GetTitle

`func (o *Attribute) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Attribute) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Attribute) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Attribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Attribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Attribute) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Attribute) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Attribute) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetApiSlug

`func (o *Attribute) GetApiSlug() string`

GetApiSlug returns the ApiSlug field if non-nil, zero value otherwise.

### GetApiSlugOk

`func (o *Attribute) GetApiSlugOk() (*string, bool)`

GetApiSlugOk returns a tuple with the ApiSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSlug

`func (o *Attribute) SetApiSlug(v string)`

SetApiSlug sets ApiSlug field to given value.


### GetType

`func (o *Attribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Attribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Attribute) SetType(v string)`

SetType sets Type field to given value.


### GetIsSystemAttribute

`func (o *Attribute) GetIsSystemAttribute() bool`

GetIsSystemAttribute returns the IsSystemAttribute field if non-nil, zero value otherwise.

### GetIsSystemAttributeOk

`func (o *Attribute) GetIsSystemAttributeOk() (*bool, bool)`

GetIsSystemAttributeOk returns a tuple with the IsSystemAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemAttribute

`func (o *Attribute) SetIsSystemAttribute(v bool)`

SetIsSystemAttribute sets IsSystemAttribute field to given value.


### GetIsWritable

`func (o *Attribute) GetIsWritable() bool`

GetIsWritable returns the IsWritable field if non-nil, zero value otherwise.

### GetIsWritableOk

`func (o *Attribute) GetIsWritableOk() (*bool, bool)`

GetIsWritableOk returns a tuple with the IsWritable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWritable

`func (o *Attribute) SetIsWritable(v bool)`

SetIsWritable sets IsWritable field to given value.


### GetIsRequired

`func (o *Attribute) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *Attribute) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *Attribute) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.


### GetIsUnique

`func (o *Attribute) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *Attribute) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *Attribute) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.


### GetIsMultiselect

`func (o *Attribute) GetIsMultiselect() bool`

GetIsMultiselect returns the IsMultiselect field if non-nil, zero value otherwise.

### GetIsMultiselectOk

`func (o *Attribute) GetIsMultiselectOk() (*bool, bool)`

GetIsMultiselectOk returns a tuple with the IsMultiselect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiselect

`func (o *Attribute) SetIsMultiselect(v bool)`

SetIsMultiselect sets IsMultiselect field to given value.


### GetIsDefaultValueEnabled

`func (o *Attribute) GetIsDefaultValueEnabled() bool`

GetIsDefaultValueEnabled returns the IsDefaultValueEnabled field if non-nil, zero value otherwise.

### GetIsDefaultValueEnabledOk

`func (o *Attribute) GetIsDefaultValueEnabledOk() (*bool, bool)`

GetIsDefaultValueEnabledOk returns a tuple with the IsDefaultValueEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultValueEnabled

`func (o *Attribute) SetIsDefaultValueEnabled(v bool)`

SetIsDefaultValueEnabled sets IsDefaultValueEnabled field to given value.


### GetIsArchived

`func (o *Attribute) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *Attribute) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *Attribute) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.


### GetDefaultValue

`func (o *Attribute) GetDefaultValue() AttributeDefaultValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *Attribute) GetDefaultValueOk() (*AttributeDefaultValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *Attribute) SetDefaultValue(v AttributeDefaultValue)`

SetDefaultValue sets DefaultValue field to given value.


### SetDefaultValueNil

`func (o *Attribute) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *Attribute) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetRelationship

`func (o *Attribute) GetRelationship() AttributeRelationship`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *Attribute) GetRelationshipOk() (*AttributeRelationship, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *Attribute) SetRelationship(v AttributeRelationship)`

SetRelationship sets Relationship field to given value.


### GetCreatedAt

`func (o *Attribute) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Attribute) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Attribute) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetConfig

`func (o *Attribute) GetConfig() AttributeConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Attribute) GetConfigOk() (*AttributeConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Attribute) SetConfig(v AttributeConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


