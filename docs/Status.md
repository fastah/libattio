# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**StatusId**](StatusId.md) |  | 
**Title** | **string** | The title of the status | 
**IsArchived** | **bool** | Whether or not to archive the status. See our [archiving guide](/docs/archiving-vs-deleting) for more information on archiving. | 
**CelebrationEnabled** | **bool** | Whether arriving at this status triggers a celebration effect in the UI | 
**TargetTimeInStatus** | **NullableString** | Target time for a record to spend in given status expressed as a ISO-8601 duration string | 

## Methods

### NewStatus

`func NewStatus(id StatusId, title string, isArchived bool, celebrationEnabled bool, targetTimeInStatus NullableString, ) *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Status) GetId() StatusId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Status) GetIdOk() (*StatusId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Status) SetId(v StatusId)`

SetId sets Id field to given value.


### GetTitle

`func (o *Status) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Status) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Status) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIsArchived

`func (o *Status) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *Status) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *Status) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.


### GetCelebrationEnabled

`func (o *Status) GetCelebrationEnabled() bool`

GetCelebrationEnabled returns the CelebrationEnabled field if non-nil, zero value otherwise.

### GetCelebrationEnabledOk

`func (o *Status) GetCelebrationEnabledOk() (*bool, bool)`

GetCelebrationEnabledOk returns a tuple with the CelebrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCelebrationEnabled

`func (o *Status) SetCelebrationEnabled(v bool)`

SetCelebrationEnabled sets CelebrationEnabled field to given value.


### GetTargetTimeInStatus

`func (o *Status) GetTargetTimeInStatus() string`

GetTargetTimeInStatus returns the TargetTimeInStatus field if non-nil, zero value otherwise.

### GetTargetTimeInStatusOk

`func (o *Status) GetTargetTimeInStatusOk() (*string, bool)`

GetTargetTimeInStatusOk returns a tuple with the TargetTimeInStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTimeInStatus

`func (o *Status) SetTargetTimeInStatus(v string)`

SetTargetTimeInStatus sets TargetTimeInStatus field to given value.


### SetTargetTimeInStatusNil

`func (o *Status) SetTargetTimeInStatusNil(b bool)`

 SetTargetTimeInStatusNil sets the value for TargetTimeInStatus to be an explicit nil

### UnsetTargetTimeInStatus
`func (o *Status) UnsetTargetTimeInStatus()`

UnsetTargetTimeInStatus ensures that no value is present for TargetTimeInStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


