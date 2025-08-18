# V2ListsListEntriesEntryIdPatchRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryValues** | **map[string][]interface{}** | An object with an attribute &#x60;api_slug&#x60; or &#x60;attribute_id&#x60; as the key, and a single value (for single-select attributes), or an array of values (for single or multi-select attributes) as the values. For complete documentation on values for all attribute types, please see our [attribute type docs](/docs/attribute-types). | 

## Methods

### NewV2ListsListEntriesEntryIdPatchRequestData

`func NewV2ListsListEntriesEntryIdPatchRequestData(entryValues map[string][]interface{}, ) *V2ListsListEntriesEntryIdPatchRequestData`

NewV2ListsListEntriesEntryIdPatchRequestData instantiates a new V2ListsListEntriesEntryIdPatchRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ListsListEntriesEntryIdPatchRequestDataWithDefaults

`func NewV2ListsListEntriesEntryIdPatchRequestDataWithDefaults() *V2ListsListEntriesEntryIdPatchRequestData`

NewV2ListsListEntriesEntryIdPatchRequestDataWithDefaults instantiates a new V2ListsListEntriesEntryIdPatchRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryValues

`func (o *V2ListsListEntriesEntryIdPatchRequestData) GetEntryValues() map[string][]interface{}`

GetEntryValues returns the EntryValues field if non-nil, zero value otherwise.

### GetEntryValuesOk

`func (o *V2ListsListEntriesEntryIdPatchRequestData) GetEntryValuesOk() (*map[string][]interface{}, bool)`

GetEntryValuesOk returns a tuple with the EntryValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryValues

`func (o *V2ListsListEntriesEntryIdPatchRequestData) SetEntryValues(v map[string][]interface{})`

SetEntryValues sets EntryValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


