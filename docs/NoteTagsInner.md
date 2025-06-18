# NoteTagsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of entity tagged in the note. Can be either &#39;workspace-member&#39; or &#39;record&#39; | 
**WorkspaceMemberId** | **string** | The ID of the workspace member that is tagged in the note. | 
**Object** | **string** | The slug or ID of the object that the tagged record belongs to. | 
**RecordId** | **string** | The ID of the record that is tagged in the note. | 

## Methods

### NewNoteTagsInner

`func NewNoteTagsInner(type_ string, workspaceMemberId string, object string, recordId string, ) *NoteTagsInner`

NewNoteTagsInner instantiates a new NoteTagsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteTagsInnerWithDefaults

`func NewNoteTagsInnerWithDefaults() *NoteTagsInner`

NewNoteTagsInnerWithDefaults instantiates a new NoteTagsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NoteTagsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NoteTagsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NoteTagsInner) SetType(v string)`

SetType sets Type field to given value.


### GetWorkspaceMemberId

`func (o *NoteTagsInner) GetWorkspaceMemberId() string`

GetWorkspaceMemberId returns the WorkspaceMemberId field if non-nil, zero value otherwise.

### GetWorkspaceMemberIdOk

`func (o *NoteTagsInner) GetWorkspaceMemberIdOk() (*string, bool)`

GetWorkspaceMemberIdOk returns a tuple with the WorkspaceMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceMemberId

`func (o *NoteTagsInner) SetWorkspaceMemberId(v string)`

SetWorkspaceMemberId sets WorkspaceMemberId field to given value.


### GetObject

`func (o *NoteTagsInner) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *NoteTagsInner) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *NoteTagsInner) SetObject(v string)`

SetObject sets Object field to given value.


### GetRecordId

`func (o *NoteTagsInner) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *NoteTagsInner) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *NoteTagsInner) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


