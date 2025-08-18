# Note

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**NoteId**](NoteId.md) |  | 
**ParentObject** | **string** | The slug or ID of the parent object the note belongs to. | 
**ParentRecordId** | **string** | The ID of the parent record the note belongs to. | 
**Title** | **string** | The note title. The title is plaintext only and has no formatting. | 
**MeetingId** | **NullableString** | The ID of the meeting associated with this note, or null if no meeting is associated. | 
**ContentPlaintext** | **string** | The plaintext representation of the note content. The line feed character &#x60;\\n&#x60; represents new lines within the note content. | 
**ContentMarkdown** | **string** | The markdown representation of the note content. Supports a subset of markdown features including: - Headings (levels 1-3 only with &#x60;#&#x60;, &#x60;##&#x60;, &#x60;###&#x60;) - Unordered lists (&#x60;-&#x60;, &#x60;*&#x60;, &#x60;+&#x60;) - Ordered lists (&#x60;1.&#x60;, &#x60;2.&#x60;, etc.) - Text styling: &#x60;**bold**&#x60;, &#x60;*italic*&#x60;, &#x60;~~strikethrough~~&#x60;, &#x60;&#x3D;&#x3D;highlighted&#x3D;&#x3D;&#x60; - Links: &#x60;[link text](https://example.com)&#x60;  Note that note images are not returned as part of the markdown API representation. | 
**Tags** | [**[]NoteTagsInner**](NoteTagsInner.md) | An array of records or workspace members that are @-tagged in the note content. | 
**CreatedByActor** | [**NoteCreatedByActor**](NoteCreatedByActor.md) |  | 
**CreatedAt** | **string** | When the note was created. | 

## Methods

### NewNote

`func NewNote(id NoteId, parentObject string, parentRecordId string, title string, meetingId NullableString, contentPlaintext string, contentMarkdown string, tags []NoteTagsInner, createdByActor NoteCreatedByActor, createdAt string, ) *Note`

NewNote instantiates a new Note object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteWithDefaults

`func NewNoteWithDefaults() *Note`

NewNoteWithDefaults instantiates a new Note object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Note) GetId() NoteId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Note) GetIdOk() (*NoteId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Note) SetId(v NoteId)`

SetId sets Id field to given value.


### GetParentObject

`func (o *Note) GetParentObject() string`

GetParentObject returns the ParentObject field if non-nil, zero value otherwise.

### GetParentObjectOk

`func (o *Note) GetParentObjectOk() (*string, bool)`

GetParentObjectOk returns a tuple with the ParentObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObject

`func (o *Note) SetParentObject(v string)`

SetParentObject sets ParentObject field to given value.


### GetParentRecordId

`func (o *Note) GetParentRecordId() string`

GetParentRecordId returns the ParentRecordId field if non-nil, zero value otherwise.

### GetParentRecordIdOk

`func (o *Note) GetParentRecordIdOk() (*string, bool)`

GetParentRecordIdOk returns a tuple with the ParentRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRecordId

`func (o *Note) SetParentRecordId(v string)`

SetParentRecordId sets ParentRecordId field to given value.


### GetTitle

`func (o *Note) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Note) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Note) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMeetingId

`func (o *Note) GetMeetingId() string`

GetMeetingId returns the MeetingId field if non-nil, zero value otherwise.

### GetMeetingIdOk

`func (o *Note) GetMeetingIdOk() (*string, bool)`

GetMeetingIdOk returns a tuple with the MeetingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingId

`func (o *Note) SetMeetingId(v string)`

SetMeetingId sets MeetingId field to given value.


### SetMeetingIdNil

`func (o *Note) SetMeetingIdNil(b bool)`

 SetMeetingIdNil sets the value for MeetingId to be an explicit nil

### UnsetMeetingId
`func (o *Note) UnsetMeetingId()`

UnsetMeetingId ensures that no value is present for MeetingId, not even an explicit nil
### GetContentPlaintext

`func (o *Note) GetContentPlaintext() string`

GetContentPlaintext returns the ContentPlaintext field if non-nil, zero value otherwise.

### GetContentPlaintextOk

`func (o *Note) GetContentPlaintextOk() (*string, bool)`

GetContentPlaintextOk returns a tuple with the ContentPlaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPlaintext

`func (o *Note) SetContentPlaintext(v string)`

SetContentPlaintext sets ContentPlaintext field to given value.


### GetContentMarkdown

`func (o *Note) GetContentMarkdown() string`

GetContentMarkdown returns the ContentMarkdown field if non-nil, zero value otherwise.

### GetContentMarkdownOk

`func (o *Note) GetContentMarkdownOk() (*string, bool)`

GetContentMarkdownOk returns a tuple with the ContentMarkdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentMarkdown

`func (o *Note) SetContentMarkdown(v string)`

SetContentMarkdown sets ContentMarkdown field to given value.


### GetTags

`func (o *Note) GetTags() []NoteTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Note) GetTagsOk() (*[]NoteTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Note) SetTags(v []NoteTagsInner)`

SetTags sets Tags field to given value.


### GetCreatedByActor

`func (o *Note) GetCreatedByActor() NoteCreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *Note) GetCreatedByActorOk() (*NoteCreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *Note) SetCreatedByActor(v NoteCreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.


### GetCreatedAt

`func (o *Note) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Note) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Note) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


