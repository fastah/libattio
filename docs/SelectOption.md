# SelectOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**SelectOptionId**](SelectOptionId.md) |  | 
**Title** | **string** | The title of the select option | 
**IsArchived** | **bool** | Whether or not to archive the select option. See our [archiving guide](/docs/archiving-vs-deleting) for more information on archiving. | 

## Methods

### NewSelectOption

`func NewSelectOption(id SelectOptionId, title string, isArchived bool, ) *SelectOption`

NewSelectOption instantiates a new SelectOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectOptionWithDefaults

`func NewSelectOptionWithDefaults() *SelectOption`

NewSelectOptionWithDefaults instantiates a new SelectOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SelectOption) GetId() SelectOptionId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SelectOption) GetIdOk() (*SelectOptionId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SelectOption) SetId(v SelectOptionId)`

SetId sets Id field to given value.


### GetTitle

`func (o *SelectOption) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SelectOption) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SelectOption) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIsArchived

`func (o *SelectOption) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *SelectOption) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *SelectOption) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


