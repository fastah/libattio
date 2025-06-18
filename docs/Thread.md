# Thread

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**ThreadId**](ThreadId.md) |  | 
**Comments** | [**[]Comment**](Comment.md) | An array of comments in the thread, sorted by &#x60;created_at&#x60;. | 
**CreatedAt** | **string** | When the thread was created. | 

## Methods

### NewThread

`func NewThread(id ThreadId, comments []Comment, createdAt string, ) *Thread`

NewThread instantiates a new Thread object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadWithDefaults

`func NewThreadWithDefaults() *Thread`

NewThreadWithDefaults instantiates a new Thread object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Thread) GetId() ThreadId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Thread) GetIdOk() (*ThreadId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Thread) SetId(v ThreadId)`

SetId sets Id field to given value.


### GetComments

`func (o *Thread) GetComments() []Comment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Thread) GetCommentsOk() (*[]Comment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Thread) SetComments(v []Comment)`

SetComments sets Comments field to given value.


### GetCreatedAt

`func (o *Thread) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Thread) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Thread) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


