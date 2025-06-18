# V2TargetIdentifierAttributesAttributeStatusesPostRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The Title of the status | 
**CelebrationEnabled** | Pointer to **bool** | Whether arriving at this status triggers a celebration effect | [optional] [default to false]
**TargetTimeInStatus** | Pointer to **NullableString** | Target time for a record to spend in given status expressed as a ISO-8601 duration string | [optional] 

## Methods

### NewV2TargetIdentifierAttributesAttributeStatusesPostRequestData

`func NewV2TargetIdentifierAttributesAttributeStatusesPostRequestData(title string, ) *V2TargetIdentifierAttributesAttributeStatusesPostRequestData`

NewV2TargetIdentifierAttributesAttributeStatusesPostRequestData instantiates a new V2TargetIdentifierAttributesAttributeStatusesPostRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TargetIdentifierAttributesAttributeStatusesPostRequestDataWithDefaults

`func NewV2TargetIdentifierAttributesAttributeStatusesPostRequestDataWithDefaults() *V2TargetIdentifierAttributesAttributeStatusesPostRequestData`

NewV2TargetIdentifierAttributesAttributeStatusesPostRequestDataWithDefaults instantiates a new V2TargetIdentifierAttributesAttributeStatusesPostRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *V2TargetIdentifierAttributesAttributeStatusesPostRequestData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V2TargetIdentifierAttributesAttributeStatusesPostRequestData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V2TargetIdentifierAttributesAttributeStatusesPostRequestData) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCelebrationEnabled

`func (o *V2TargetIdentifierAttributesAttributeStatusesPostRequestData) GetCelebrationEnabled() bool`

GetCelebrationEnabled returns the CelebrationEnabled field if non-nil, zero value otherwise.

### GetCelebrationEnabledOk

`func (o *V2TargetIdentifierAttributesAttributeStatusesPostRequestData) GetCelebrationEnabledOk() (*bool, bool)`

GetCelebrationEnabledOk returns a tuple with the CelebrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCelebrationEnabled

`func (o *V2TargetIdentifierAttributesAttributeStatusesPostRequestData) SetCelebrationEnabled(v bool)`

SetCelebrationEnabled sets CelebrationEnabled field to given value.

### HasCelebrationEnabled

`func (o *V2TargetIdentifierAttributesAttributeStatusesPostRequestData) HasCelebrationEnabled() bool`

HasCelebrationEnabled returns a boolean if a field has been set.

### GetTargetTimeInStatus

`func (o *V2TargetIdentifierAttributesAttributeStatusesPostRequestData) GetTargetTimeInStatus() string`

GetTargetTimeInStatus returns the TargetTimeInStatus field if non-nil, zero value otherwise.

### GetTargetTimeInStatusOk

`func (o *V2TargetIdentifierAttributesAttributeStatusesPostRequestData) GetTargetTimeInStatusOk() (*string, bool)`

GetTargetTimeInStatusOk returns a tuple with the TargetTimeInStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTimeInStatus

`func (o *V2TargetIdentifierAttributesAttributeStatusesPostRequestData) SetTargetTimeInStatus(v string)`

SetTargetTimeInStatus sets TargetTimeInStatus field to given value.

### HasTargetTimeInStatus

`func (o *V2TargetIdentifierAttributesAttributeStatusesPostRequestData) HasTargetTimeInStatus() bool`

HasTargetTimeInStatus returns a boolean if a field has been set.

### SetTargetTimeInStatusNil

`func (o *V2TargetIdentifierAttributesAttributeStatusesPostRequestData) SetTargetTimeInStatusNil(b bool)`

 SetTargetTimeInStatusNil sets the value for TargetTimeInStatus to be an explicit nil

### UnsetTargetTimeInStatus
`func (o *V2TargetIdentifierAttributesAttributeStatusesPostRequestData) UnsetTargetTimeInStatus()`

UnsetTargetTimeInStatus ensures that no value is present for TargetTimeInStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


