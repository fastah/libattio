# V2ListsListEntriesQueryPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **map[string]interface{}** | An object used to filter results to a subset of results. See the [full guide to filtering and sorting here](/rest-api/how-to/filtering-and-sorting). | [optional] 
**Sorts** | Pointer to [**[]V2ObjectsObjectRecordsQueryPostRequestSortsInner**](V2ObjectsObjectRecordsQueryPostRequestSortsInner.md) | An object used to sort results. See the [full guide to filtering and sorting here](/rest-api/how-to/filtering-and-sorting). | [optional] 
**Limit** | Pointer to **float32** | The maximum number of results to return. Defaults to 500. See the [full guide to pagination here](/rest-api/how-to/pagination). | [optional] 
**Offset** | Pointer to **float32** | The number of results to skip over before returning. Defaults to 0. See the [full guide to pagination here](/rest-api/how-to/pagination). | [optional] 

## Methods

### NewV2ListsListEntriesQueryPostRequest

`func NewV2ListsListEntriesQueryPostRequest() *V2ListsListEntriesQueryPostRequest`

NewV2ListsListEntriesQueryPostRequest instantiates a new V2ListsListEntriesQueryPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ListsListEntriesQueryPostRequestWithDefaults

`func NewV2ListsListEntriesQueryPostRequestWithDefaults() *V2ListsListEntriesQueryPostRequest`

NewV2ListsListEntriesQueryPostRequestWithDefaults instantiates a new V2ListsListEntriesQueryPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *V2ListsListEntriesQueryPostRequest) GetFilter() map[string]interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *V2ListsListEntriesQueryPostRequest) GetFilterOk() (*map[string]interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *V2ListsListEntriesQueryPostRequest) SetFilter(v map[string]interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *V2ListsListEntriesQueryPostRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSorts

`func (o *V2ListsListEntriesQueryPostRequest) GetSorts() []V2ObjectsObjectRecordsQueryPostRequestSortsInner`

GetSorts returns the Sorts field if non-nil, zero value otherwise.

### GetSortsOk

`func (o *V2ListsListEntriesQueryPostRequest) GetSortsOk() (*[]V2ObjectsObjectRecordsQueryPostRequestSortsInner, bool)`

GetSortsOk returns a tuple with the Sorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSorts

`func (o *V2ListsListEntriesQueryPostRequest) SetSorts(v []V2ObjectsObjectRecordsQueryPostRequestSortsInner)`

SetSorts sets Sorts field to given value.

### HasSorts

`func (o *V2ListsListEntriesQueryPostRequest) HasSorts() bool`

HasSorts returns a boolean if a field has been set.

### GetLimit

`func (o *V2ListsListEntriesQueryPostRequest) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *V2ListsListEntriesQueryPostRequest) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *V2ListsListEntriesQueryPostRequest) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *V2ListsListEntriesQueryPostRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *V2ListsListEntriesQueryPostRequest) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *V2ListsListEntriesQueryPostRequest) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *V2ListsListEntriesQueryPostRequest) SetOffset(v float32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *V2ListsListEntriesQueryPostRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


