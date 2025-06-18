# V2ObjectsObjectRecordsPostRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | **map[string][]interface{}** | An object with an attribute &#x60;api_slug&#x60; or &#x60;attribute_id&#x60; as the key, and a single value (for single-select attributes), or an array of values (for single or multi-select attributes) as the values. For complete documentation on values for all attribute types, please see our [attribute type docs](/docs/attribute-types). | 

## Methods

### NewV2ObjectsObjectRecordsPostRequestData

`func NewV2ObjectsObjectRecordsPostRequestData(values map[string][]interface{}, ) *V2ObjectsObjectRecordsPostRequestData`

NewV2ObjectsObjectRecordsPostRequestData instantiates a new V2ObjectsObjectRecordsPostRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ObjectsObjectRecordsPostRequestDataWithDefaults

`func NewV2ObjectsObjectRecordsPostRequestDataWithDefaults() *V2ObjectsObjectRecordsPostRequestData`

NewV2ObjectsObjectRecordsPostRequestDataWithDefaults instantiates a new V2ObjectsObjectRecordsPostRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *V2ObjectsObjectRecordsPostRequestData) GetValues() map[string][]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *V2ObjectsObjectRecordsPostRequestData) GetValuesOk() (*map[string][]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *V2ObjectsObjectRecordsPostRequestData) SetValues(v map[string][]interface{})`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


