# CommentRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordId** | **string** | The ID of the record the comment belongs to. | 
**ObjectId** | **string** | The ID of the object the record belongs to. | 

## Methods

### NewCommentRecord

`func NewCommentRecord(recordId string, objectId string, ) *CommentRecord`

NewCommentRecord instantiates a new CommentRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentRecordWithDefaults

`func NewCommentRecordWithDefaults() *CommentRecord`

NewCommentRecordWithDefaults instantiates a new CommentRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordId

`func (o *CommentRecord) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *CommentRecord) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *CommentRecord) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.


### GetObjectId

`func (o *CommentRecord) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *CommentRecord) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *CommentRecord) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


