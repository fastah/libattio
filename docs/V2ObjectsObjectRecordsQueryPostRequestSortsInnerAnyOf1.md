# V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | **string** | The direction to sort the results by. | 
**Path** | [**[][]V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner**]([]V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner.md) | You may use the &#x60;path&#x60; property to traverse record reference attributes and parent records on list entries. &#x60;path&#x60; accepts an array of tuples where the first element of each tuple is the slug or ID of a list/object, and the second element is the slug or ID of an attribute on that list/object. The first element of the first tuple must correspond to the list or object that you are querying. For example, if you wanted to sort by the name of the parent record (a company) on a list with the slug \&quot;sales\&quot;, you would pass the value &#x60;[[&#39;sales&#39;, &#39;parent_record&#39;], [&#39;companies&#39;, &#39;name&#39;]]&#x60;. | 
**Field** | Pointer to **string** | Which field on the value to sort by e.g. \&quot;last_name\&quot; on a name value. | [optional] 

## Methods

### NewV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1

`func NewV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1(direction string, path [][]V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner, ) *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1`

NewV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1 instantiates a new V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1WithDefaults

`func NewV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1WithDefaults() *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1`

NewV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1WithDefaults instantiates a new V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetPath

`func (o *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1) GetPath() [][]V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1) GetPathOk() (*[][]V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1) SetPath(v [][]V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner)`

SetPath sets Path field to given value.


### GetField

`func (o *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


