# V2CommentsPostRequestDataAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** | The format that the comment content is provided in. The &#x60;plaintext&#x60; format uses the line feed character &#x60;\\n&#x60; to create new lines within the note content. Rich text formatting and links are not supported. | 
**Content** | **string** | The content of the comment itself. Workspace members can be mentioned using their email address, otherwise email addresses will be presented to users as clickable mailto links. | 
**Author** | [**V2CommentsPostRequestDataAnyOfAuthor**](V2CommentsPostRequestDataAnyOfAuthor.md) |  | 
**CreatedAt** | Pointer to **string** | &#x60;created_at&#x60; will default to the current time. However, if you wish to backdate a comment for migration or other purposes, you can override with a custom &#x60;created_at&#x60; value. Note that dates before 1970 or in the future are not allowed. | [optional] 
**ThreadId** | **string** | If responding to an existing thread, this would be the ID of that thread. | 

## Methods

### NewV2CommentsPostRequestDataAnyOf

`func NewV2CommentsPostRequestDataAnyOf(format string, content string, author V2CommentsPostRequestDataAnyOfAuthor, threadId string, ) *V2CommentsPostRequestDataAnyOf`

NewV2CommentsPostRequestDataAnyOf instantiates a new V2CommentsPostRequestDataAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2CommentsPostRequestDataAnyOfWithDefaults

`func NewV2CommentsPostRequestDataAnyOfWithDefaults() *V2CommentsPostRequestDataAnyOf`

NewV2CommentsPostRequestDataAnyOfWithDefaults instantiates a new V2CommentsPostRequestDataAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *V2CommentsPostRequestDataAnyOf) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *V2CommentsPostRequestDataAnyOf) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *V2CommentsPostRequestDataAnyOf) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetContent

`func (o *V2CommentsPostRequestDataAnyOf) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *V2CommentsPostRequestDataAnyOf) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *V2CommentsPostRequestDataAnyOf) SetContent(v string)`

SetContent sets Content field to given value.


### GetAuthor

`func (o *V2CommentsPostRequestDataAnyOf) GetAuthor() V2CommentsPostRequestDataAnyOfAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *V2CommentsPostRequestDataAnyOf) GetAuthorOk() (*V2CommentsPostRequestDataAnyOfAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *V2CommentsPostRequestDataAnyOf) SetAuthor(v V2CommentsPostRequestDataAnyOfAuthor)`

SetAuthor sets Author field to given value.


### GetCreatedAt

`func (o *V2CommentsPostRequestDataAnyOf) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2CommentsPostRequestDataAnyOf) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2CommentsPostRequestDataAnyOf) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2CommentsPostRequestDataAnyOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetThreadId

`func (o *V2CommentsPostRequestDataAnyOf) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *V2CommentsPostRequestDataAnyOf) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *V2CommentsPostRequestDataAnyOf) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


