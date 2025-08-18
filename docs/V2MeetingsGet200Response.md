# V2MeetingsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Meeting**](Meeting.md) |  | 
**Pagination** | [**V2MeetingsGet200ResponsePagination**](V2MeetingsGet200ResponsePagination.md) |  | 

## Methods

### NewV2MeetingsGet200Response

`func NewV2MeetingsGet200Response(data []Meeting, pagination V2MeetingsGet200ResponsePagination, ) *V2MeetingsGet200Response`

NewV2MeetingsGet200Response instantiates a new V2MeetingsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MeetingsGet200ResponseWithDefaults

`func NewV2MeetingsGet200ResponseWithDefaults() *V2MeetingsGet200Response`

NewV2MeetingsGet200ResponseWithDefaults instantiates a new V2MeetingsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V2MeetingsGet200Response) GetData() []Meeting`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V2MeetingsGet200Response) GetDataOk() (*[]Meeting, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V2MeetingsGet200Response) SetData(v []Meeting)`

SetData sets Data field to given value.


### GetPagination

`func (o *V2MeetingsGet200Response) GetPagination() V2MeetingsGet200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V2MeetingsGet200Response) GetPaginationOk() (*V2MeetingsGet200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V2MeetingsGet200Response) SetPagination(v V2MeetingsGet200ResponsePagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


