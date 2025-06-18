# V2CommentsPostRequestDataAnyOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** | The format that the comment content is provided in. The &#x60;plaintext&#x60; format uses the line feed character &#x60;\\n&#x60; to create new lines within the note content. Rich text formatting and links are not supported. | 
**Content** | **string** | The content of the comment itself. Workspace members can be mentioned using their email address, otherwise email addresses will be presented to users as clickable mailto links. | 
**Author** | [**V2CommentsPostRequestDataAnyOfAuthor**](V2CommentsPostRequestDataAnyOfAuthor.md) |  | 
**CreatedAt** | Pointer to **string** | &#x60;created_at&#x60; will default to the current time. However, if you wish to backdate a comment for migration or other purposes, you can override with a custom &#x60;created_at&#x60; value. Note that dates before 1970 or in the future are not allowed. | [optional] 
**Entry** | [**V2CommentsPostRequestDataAnyOf2Entry**](V2CommentsPostRequestDataAnyOf2Entry.md) |  | 

## Methods

### NewV2CommentsPostRequestDataAnyOf2

`func NewV2CommentsPostRequestDataAnyOf2(format string, content string, author V2CommentsPostRequestDataAnyOfAuthor, entry V2CommentsPostRequestDataAnyOf2Entry, ) *V2CommentsPostRequestDataAnyOf2`

NewV2CommentsPostRequestDataAnyOf2 instantiates a new V2CommentsPostRequestDataAnyOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2CommentsPostRequestDataAnyOf2WithDefaults

`func NewV2CommentsPostRequestDataAnyOf2WithDefaults() *V2CommentsPostRequestDataAnyOf2`

NewV2CommentsPostRequestDataAnyOf2WithDefaults instantiates a new V2CommentsPostRequestDataAnyOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *V2CommentsPostRequestDataAnyOf2) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *V2CommentsPostRequestDataAnyOf2) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *V2CommentsPostRequestDataAnyOf2) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetContent

`func (o *V2CommentsPostRequestDataAnyOf2) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *V2CommentsPostRequestDataAnyOf2) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *V2CommentsPostRequestDataAnyOf2) SetContent(v string)`

SetContent sets Content field to given value.


### GetAuthor

`func (o *V2CommentsPostRequestDataAnyOf2) GetAuthor() V2CommentsPostRequestDataAnyOfAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *V2CommentsPostRequestDataAnyOf2) GetAuthorOk() (*V2CommentsPostRequestDataAnyOfAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *V2CommentsPostRequestDataAnyOf2) SetAuthor(v V2CommentsPostRequestDataAnyOfAuthor)`

SetAuthor sets Author field to given value.


### GetCreatedAt

`func (o *V2CommentsPostRequestDataAnyOf2) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2CommentsPostRequestDataAnyOf2) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2CommentsPostRequestDataAnyOf2) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2CommentsPostRequestDataAnyOf2) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEntry

`func (o *V2CommentsPostRequestDataAnyOf2) GetEntry() V2CommentsPostRequestDataAnyOf2Entry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *V2CommentsPostRequestDataAnyOf2) GetEntryOk() (*V2CommentsPostRequestDataAnyOf2Entry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *V2CommentsPostRequestDataAnyOf2) SetEntry(v V2CommentsPostRequestDataAnyOf2Entry)`

SetEntry sets Entry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


