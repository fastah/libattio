# V2ListsListEntriesQueryPost200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**V2ListsListEntriesQueryPost200ResponseDataInnerId**](V2ListsListEntriesQueryPost200ResponseDataInnerId.md) |  | 
**ParentRecordId** | **string** | A UUID identifying the record that is parent of the list entry. | 
**ParentObject** | **string** | A UUID or slug identifying the object that the parent record belongs to. | 
**CreatedAt** | **string** | When this entry was created. | 
**EntryValues** | [**map[string][]V2ObjectsObjectRecordsPut200ResponseDataValuesValueInner**](array.md) | A list of attribute values for the list entry (not attribute values for its parent record). | 

## Methods

### NewV2ListsListEntriesQueryPost200ResponseDataInner

`func NewV2ListsListEntriesQueryPost200ResponseDataInner(id V2ListsListEntriesQueryPost200ResponseDataInnerId, parentRecordId string, parentObject string, createdAt string, entryValues map[string][]V2ObjectsObjectRecordsPut200ResponseDataValuesValueInner, ) *V2ListsListEntriesQueryPost200ResponseDataInner`

NewV2ListsListEntriesQueryPost200ResponseDataInner instantiates a new V2ListsListEntriesQueryPost200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ListsListEntriesQueryPost200ResponseDataInnerWithDefaults

`func NewV2ListsListEntriesQueryPost200ResponseDataInnerWithDefaults() *V2ListsListEntriesQueryPost200ResponseDataInner`

NewV2ListsListEntriesQueryPost200ResponseDataInnerWithDefaults instantiates a new V2ListsListEntriesQueryPost200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2ListsListEntriesQueryPost200ResponseDataInner) GetId() V2ListsListEntriesQueryPost200ResponseDataInnerId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2ListsListEntriesQueryPost200ResponseDataInner) GetIdOk() (*V2ListsListEntriesQueryPost200ResponseDataInnerId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2ListsListEntriesQueryPost200ResponseDataInner) SetId(v V2ListsListEntriesQueryPost200ResponseDataInnerId)`

SetId sets Id field to given value.


### GetParentRecordId

`func (o *V2ListsListEntriesQueryPost200ResponseDataInner) GetParentRecordId() string`

GetParentRecordId returns the ParentRecordId field if non-nil, zero value otherwise.

### GetParentRecordIdOk

`func (o *V2ListsListEntriesQueryPost200ResponseDataInner) GetParentRecordIdOk() (*string, bool)`

GetParentRecordIdOk returns a tuple with the ParentRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRecordId

`func (o *V2ListsListEntriesQueryPost200ResponseDataInner) SetParentRecordId(v string)`

SetParentRecordId sets ParentRecordId field to given value.


### GetParentObject

`func (o *V2ListsListEntriesQueryPost200ResponseDataInner) GetParentObject() string`

GetParentObject returns the ParentObject field if non-nil, zero value otherwise.

### GetParentObjectOk

`func (o *V2ListsListEntriesQueryPost200ResponseDataInner) GetParentObjectOk() (*string, bool)`

GetParentObjectOk returns a tuple with the ParentObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObject

`func (o *V2ListsListEntriesQueryPost200ResponseDataInner) SetParentObject(v string)`

SetParentObject sets ParentObject field to given value.


### GetCreatedAt

`func (o *V2ListsListEntriesQueryPost200ResponseDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2ListsListEntriesQueryPost200ResponseDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2ListsListEntriesQueryPost200ResponseDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEntryValues

`func (o *V2ListsListEntriesQueryPost200ResponseDataInner) GetEntryValues() map[string][]V2ObjectsObjectRecordsPut200ResponseDataValuesValueInner`

GetEntryValues returns the EntryValues field if non-nil, zero value otherwise.

### GetEntryValuesOk

`func (o *V2ListsListEntriesQueryPost200ResponseDataInner) GetEntryValuesOk() (*map[string][]V2ObjectsObjectRecordsPut200ResponseDataValuesValueInner, bool)`

GetEntryValuesOk returns a tuple with the EntryValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryValues

`func (o *V2ListsListEntriesQueryPost200ResponseDataInner) SetEntryValues(v map[string][]V2ObjectsObjectRecordsPut200ResponseDataValuesValueInner)`

SetEntryValues sets EntryValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


