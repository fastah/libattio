# V2NotesPostRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentObject** | **string** | The ID or slug of the parent object the note belongs to. | 
**ParentRecordId** | **string** | The ID of the parent record the note belongs to. | 
**Title** | **string** | The note title. The title is plaintext only and has no formatting. | 
**Format** | **string** | Specify the format for the note&#39;s content. Choose from: - &#x60;plaintext&#x60;: Standard text format where &#x60;\\n&#x60; signifies a new line. - &#x60;markdown&#x60;: Enables rich text formatting using a subset of Markdown syntax:   - **Headings**: Levels 1-3 (&#x60;#&#x60;, &#x60;##&#x60;, &#x60;###&#x60;).   - **Lists**: Unordered (&#x60;-&#x60;, &#x60;*&#x60;, &#x60;+&#x60;) and ordered (&#x60;1.&#x60;, &#x60;2.&#x60;).   - **Text styles**: Bold (&#x60;**bold**&#x60; or &#x60;__bold__&#x60;), italic (&#x60;*italic*&#x60; or &#x60;_italic_&#x60;), strikethrough (&#x60;~~strikethrough~~&#x60;), and highlight (&#x60;&#x3D;&#x3D;highlighted&#x3D;&#x3D;&#x60;).   - **Links**: Standard Markdown links (&#x60;[link text](https://example.com)&#x60;).    *Note: While the Attio interface supports image embeds, they cannot currently be added or retrieved via the API&#39;s markdown format.* | 
**Content** | **string** | The main content of the note, formatted according to the value provided in the &#x60;format&#x60; field. Use &#x60;\\n&#x60; for line breaks in &#x60;plaintext&#x60;. For &#x60;markdown&#x60;, utilize the supported syntax elements to structure and style your note. | 
**CreatedAt** | Pointer to **string** | &#x60;created_at&#x60; will default to the current time. However, if you wish to backdate a note for migration or other purposes, you can override with a custom &#x60;created_at&#x60; value. Note that dates before 1970 or in the future are not allowed. | [optional] 
**MeetingId** | Pointer to **NullableString** | An optional ID to associate this note with a meeting. If provided, the meeting must exist. Use &#x60;null&#x60; to explicitly set no meeting association. | [optional] 

## Methods

### NewV2NotesPostRequestData

`func NewV2NotesPostRequestData(parentObject string, parentRecordId string, title string, format string, content string, ) *V2NotesPostRequestData`

NewV2NotesPostRequestData instantiates a new V2NotesPostRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2NotesPostRequestDataWithDefaults

`func NewV2NotesPostRequestDataWithDefaults() *V2NotesPostRequestData`

NewV2NotesPostRequestDataWithDefaults instantiates a new V2NotesPostRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentObject

`func (o *V2NotesPostRequestData) GetParentObject() string`

GetParentObject returns the ParentObject field if non-nil, zero value otherwise.

### GetParentObjectOk

`func (o *V2NotesPostRequestData) GetParentObjectOk() (*string, bool)`

GetParentObjectOk returns a tuple with the ParentObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObject

`func (o *V2NotesPostRequestData) SetParentObject(v string)`

SetParentObject sets ParentObject field to given value.


### GetParentRecordId

`func (o *V2NotesPostRequestData) GetParentRecordId() string`

GetParentRecordId returns the ParentRecordId field if non-nil, zero value otherwise.

### GetParentRecordIdOk

`func (o *V2NotesPostRequestData) GetParentRecordIdOk() (*string, bool)`

GetParentRecordIdOk returns a tuple with the ParentRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRecordId

`func (o *V2NotesPostRequestData) SetParentRecordId(v string)`

SetParentRecordId sets ParentRecordId field to given value.


### GetTitle

`func (o *V2NotesPostRequestData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V2NotesPostRequestData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V2NotesPostRequestData) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetFormat

`func (o *V2NotesPostRequestData) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *V2NotesPostRequestData) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *V2NotesPostRequestData) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetContent

`func (o *V2NotesPostRequestData) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *V2NotesPostRequestData) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *V2NotesPostRequestData) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreatedAt

`func (o *V2NotesPostRequestData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2NotesPostRequestData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2NotesPostRequestData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2NotesPostRequestData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMeetingId

`func (o *V2NotesPostRequestData) GetMeetingId() string`

GetMeetingId returns the MeetingId field if non-nil, zero value otherwise.

### GetMeetingIdOk

`func (o *V2NotesPostRequestData) GetMeetingIdOk() (*string, bool)`

GetMeetingIdOk returns a tuple with the MeetingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingId

`func (o *V2NotesPostRequestData) SetMeetingId(v string)`

SetMeetingId sets MeetingId field to given value.

### HasMeetingId

`func (o *V2NotesPostRequestData) HasMeetingId() bool`

HasMeetingId returns a boolean if a field has been set.

### SetMeetingIdNil

`func (o *V2NotesPostRequestData) SetMeetingIdNil(b bool)`

 SetMeetingIdNil sets the value for MeetingId to be an explicit nil

### UnsetMeetingId
`func (o *V2NotesPostRequestData) UnsetMeetingId()`

UnsetMeetingId ensures that no value is present for MeetingId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


