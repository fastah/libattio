# CommentEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** | The ID of the entry the comment belongs to. | 
**ListId** | **string** | The ID of the list the entry belongs to. | 

## Methods

### NewCommentEntry

`func NewCommentEntry(entryId string, listId string, ) *CommentEntry`

NewCommentEntry instantiates a new CommentEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentEntryWithDefaults

`func NewCommentEntryWithDefaults() *CommentEntry`

NewCommentEntryWithDefaults instantiates a new CommentEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *CommentEntry) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *CommentEntry) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *CommentEntry) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetListId

`func (o *CommentEntry) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *CommentEntry) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *CommentEntry) SetListId(v string)`

SetListId sets ListId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


