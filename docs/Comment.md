# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**CommentId**](CommentId.md) |  | 
**ThreadId** | **string** | The ID of the thread the comment belongs to. | 
**ContentPlaintext** | **string** | A plaintext representation of the content of the comment. References to workspace members are cast into email addresses, all other stylistic elements are removed. | 
**Entry** | [**CommentEntry**](CommentEntry.md) |  | 
**Record** | [**CommentRecord**](CommentRecord.md) |  | 
**ResolvedAt** | **NullableString** | Whether the comment is resolved. | 
**ResolvedBy** | [**CommentResolvedBy**](CommentResolvedBy.md) |  | 
**CreatedAt** | **string** | When the note was created. | 
**Author** | [**CommentAuthor**](CommentAuthor.md) |  | 

## Methods

### NewComment

`func NewComment(id CommentId, threadId string, contentPlaintext string, entry CommentEntry, record CommentRecord, resolvedAt NullableString, resolvedBy CommentResolvedBy, createdAt string, author CommentAuthor, ) *Comment`

NewComment instantiates a new Comment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithDefaults

`func NewCommentWithDefaults() *Comment`

NewCommentWithDefaults instantiates a new Comment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Comment) GetId() CommentId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Comment) GetIdOk() (*CommentId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Comment) SetId(v CommentId)`

SetId sets Id field to given value.


### GetThreadId

`func (o *Comment) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *Comment) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *Comment) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetContentPlaintext

`func (o *Comment) GetContentPlaintext() string`

GetContentPlaintext returns the ContentPlaintext field if non-nil, zero value otherwise.

### GetContentPlaintextOk

`func (o *Comment) GetContentPlaintextOk() (*string, bool)`

GetContentPlaintextOk returns a tuple with the ContentPlaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPlaintext

`func (o *Comment) SetContentPlaintext(v string)`

SetContentPlaintext sets ContentPlaintext field to given value.


### GetEntry

`func (o *Comment) GetEntry() CommentEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *Comment) GetEntryOk() (*CommentEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *Comment) SetEntry(v CommentEntry)`

SetEntry sets Entry field to given value.


### GetRecord

`func (o *Comment) GetRecord() CommentRecord`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *Comment) GetRecordOk() (*CommentRecord, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *Comment) SetRecord(v CommentRecord)`

SetRecord sets Record field to given value.


### GetResolvedAt

`func (o *Comment) GetResolvedAt() string`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *Comment) GetResolvedAtOk() (*string, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *Comment) SetResolvedAt(v string)`

SetResolvedAt sets ResolvedAt field to given value.


### SetResolvedAtNil

`func (o *Comment) SetResolvedAtNil(b bool)`

 SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil

### UnsetResolvedAt
`func (o *Comment) UnsetResolvedAt()`

UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
### GetResolvedBy

`func (o *Comment) GetResolvedBy() CommentResolvedBy`

GetResolvedBy returns the ResolvedBy field if non-nil, zero value otherwise.

### GetResolvedByOk

`func (o *Comment) GetResolvedByOk() (*CommentResolvedBy, bool)`

GetResolvedByOk returns a tuple with the ResolvedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedBy

`func (o *Comment) SetResolvedBy(v CommentResolvedBy)`

SetResolvedBy sets ResolvedBy field to given value.


### GetCreatedAt

`func (o *Comment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Comment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Comment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetAuthor

`func (o *Comment) GetAuthor() CommentAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Comment) GetAuthorOk() (*CommentAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Comment) SetAuthor(v CommentAuthor)`

SetAuthor sets Author field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


