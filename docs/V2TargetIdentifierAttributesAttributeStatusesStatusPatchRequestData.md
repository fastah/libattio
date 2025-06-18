# V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The Title of the status | [optional] 
**CelebrationEnabled** | Pointer to **bool** | Whether arriving at this status triggers a celebration effect | [optional] [default to false]
**TargetTimeInStatus** | Pointer to **NullableString** | Target time for a record to spend in given status expressed as a ISO-8601 duration string | [optional] 
**IsArchived** | Pointer to **bool** | Whether or not to archive the status. See our [archiving guide](/docs/archiving-vs-deleting) for more information on archiving. | [optional] 

## Methods

### NewV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData

`func NewV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData() *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData`

NewV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData instantiates a new V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestDataWithDefaults

`func NewV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestDataWithDefaults() *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData`

NewV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestDataWithDefaults instantiates a new V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCelebrationEnabled

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetCelebrationEnabled() bool`

GetCelebrationEnabled returns the CelebrationEnabled field if non-nil, zero value otherwise.

### GetCelebrationEnabledOk

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetCelebrationEnabledOk() (*bool, bool)`

GetCelebrationEnabledOk returns a tuple with the CelebrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCelebrationEnabled

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) SetCelebrationEnabled(v bool)`

SetCelebrationEnabled sets CelebrationEnabled field to given value.

### HasCelebrationEnabled

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) HasCelebrationEnabled() bool`

HasCelebrationEnabled returns a boolean if a field has been set.

### GetTargetTimeInStatus

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetTargetTimeInStatus() string`

GetTargetTimeInStatus returns the TargetTimeInStatus field if non-nil, zero value otherwise.

### GetTargetTimeInStatusOk

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetTargetTimeInStatusOk() (*string, bool)`

GetTargetTimeInStatusOk returns a tuple with the TargetTimeInStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTimeInStatus

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) SetTargetTimeInStatus(v string)`

SetTargetTimeInStatus sets TargetTimeInStatus field to given value.

### HasTargetTimeInStatus

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) HasTargetTimeInStatus() bool`

HasTargetTimeInStatus returns a boolean if a field has been set.

### SetTargetTimeInStatusNil

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) SetTargetTimeInStatusNil(b bool)`

 SetTargetTimeInStatusNil sets the value for TargetTimeInStatus to be an explicit nil

### UnsetTargetTimeInStatus
`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) UnsetTargetTimeInStatus()`

UnsetTargetTimeInStatus ensures that no value is present for TargetTimeInStatus, not even an explicit nil
### GetIsArchived

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


