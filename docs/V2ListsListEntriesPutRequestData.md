# V2ListsListEntriesPutRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentRecordId** | **string** | A UUID identifying the record you want to add to the list. The record will become the &#39;parent&#39; of the created list entry. | 
**ParentObject** | **string** | A UUID or slug identifying the object that the added parent record belongs to. | 
**EntryValues** | **map[string][]interface{}** | An object with an attribute &#x60;api_slug&#x60; or &#x60;attribute_id&#x60; as the key, and a single value (for single-select attributes), or an array of values (for single or multi-select attributes) as the values. For complete documentation on values for all attribute types, please see our [attribute type docs](/docs/attribute-types). | 

## Methods

### NewV2ListsListEntriesPutRequestData

`func NewV2ListsListEntriesPutRequestData(parentRecordId string, parentObject string, entryValues map[string][]interface{}, ) *V2ListsListEntriesPutRequestData`

NewV2ListsListEntriesPutRequestData instantiates a new V2ListsListEntriesPutRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ListsListEntriesPutRequestDataWithDefaults

`func NewV2ListsListEntriesPutRequestDataWithDefaults() *V2ListsListEntriesPutRequestData`

NewV2ListsListEntriesPutRequestDataWithDefaults instantiates a new V2ListsListEntriesPutRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentRecordId

`func (o *V2ListsListEntriesPutRequestData) GetParentRecordId() string`

GetParentRecordId returns the ParentRecordId field if non-nil, zero value otherwise.

### GetParentRecordIdOk

`func (o *V2ListsListEntriesPutRequestData) GetParentRecordIdOk() (*string, bool)`

GetParentRecordIdOk returns a tuple with the ParentRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRecordId

`func (o *V2ListsListEntriesPutRequestData) SetParentRecordId(v string)`

SetParentRecordId sets ParentRecordId field to given value.


### GetParentObject

`func (o *V2ListsListEntriesPutRequestData) GetParentObject() string`

GetParentObject returns the ParentObject field if non-nil, zero value otherwise.

### GetParentObjectOk

`func (o *V2ListsListEntriesPutRequestData) GetParentObjectOk() (*string, bool)`

GetParentObjectOk returns a tuple with the ParentObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObject

`func (o *V2ListsListEntriesPutRequestData) SetParentObject(v string)`

SetParentObject sets ParentObject field to given value.


### GetEntryValues

`func (o *V2ListsListEntriesPutRequestData) GetEntryValues() map[string][]interface{}`

GetEntryValues returns the EntryValues field if non-nil, zero value otherwise.

### GetEntryValuesOk

`func (o *V2ListsListEntriesPutRequestData) GetEntryValuesOk() (*map[string][]interface{}, bool)`

GetEntryValuesOk returns a tuple with the EntryValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryValues

`func (o *V2ListsListEntriesPutRequestData) SetEntryValues(v map[string][]interface{})`

SetEntryValues sets EntryValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


