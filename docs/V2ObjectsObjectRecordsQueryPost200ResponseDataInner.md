# V2ObjectsObjectRecordsQueryPost200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId**](V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId.md) |  | 
**CreatedAt** | **string** | When this record was created. | 
**WebUrl** | **string** | A URL that links directly to the record page in the Attio web application. | 
**Values** | [**map[string][]V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInner**](array.md) | A record type with an attribute &#x60;api_slug&#x60; as the key, and an array of value objects as the values. | 

## Methods

### NewV2ObjectsObjectRecordsQueryPost200ResponseDataInner

`func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInner(id V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId, createdAt string, webUrl string, values map[string][]V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInner, ) *V2ObjectsObjectRecordsQueryPost200ResponseDataInner`

NewV2ObjectsObjectRecordsQueryPost200ResponseDataInner instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerWithDefaults

`func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerWithDefaults() *V2ObjectsObjectRecordsQueryPost200ResponseDataInner`

NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerWithDefaults instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetId() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetIdOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) SetId(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetWebUrl

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetValues

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetValues() map[string][]V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) GetValuesOk() (*map[string][]V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInner) SetValues(v map[string][]V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInner)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


