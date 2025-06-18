# V2ListsPostRequestDataWorkspaceMemberAccessInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceMemberId** | **string** | A UUID to identify the workspace member to grant access to. | 
**Level** | **string** | The level of access to the list. | 

## Methods

### NewV2ListsPostRequestDataWorkspaceMemberAccessInner

`func NewV2ListsPostRequestDataWorkspaceMemberAccessInner(workspaceMemberId string, level string, ) *V2ListsPostRequestDataWorkspaceMemberAccessInner`

NewV2ListsPostRequestDataWorkspaceMemberAccessInner instantiates a new V2ListsPostRequestDataWorkspaceMemberAccessInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ListsPostRequestDataWorkspaceMemberAccessInnerWithDefaults

`func NewV2ListsPostRequestDataWorkspaceMemberAccessInnerWithDefaults() *V2ListsPostRequestDataWorkspaceMemberAccessInner`

NewV2ListsPostRequestDataWorkspaceMemberAccessInnerWithDefaults instantiates a new V2ListsPostRequestDataWorkspaceMemberAccessInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceMemberId

`func (o *V2ListsPostRequestDataWorkspaceMemberAccessInner) GetWorkspaceMemberId() string`

GetWorkspaceMemberId returns the WorkspaceMemberId field if non-nil, zero value otherwise.

### GetWorkspaceMemberIdOk

`func (o *V2ListsPostRequestDataWorkspaceMemberAccessInner) GetWorkspaceMemberIdOk() (*string, bool)`

GetWorkspaceMemberIdOk returns a tuple with the WorkspaceMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceMemberId

`func (o *V2ListsPostRequestDataWorkspaceMemberAccessInner) SetWorkspaceMemberId(v string)`

SetWorkspaceMemberId sets WorkspaceMemberId field to given value.


### GetLevel

`func (o *V2ListsPostRequestDataWorkspaceMemberAccessInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *V2ListsPostRequestDataWorkspaceMemberAccessInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *V2ListsPostRequestDataWorkspaceMemberAccessInner) SetLevel(v string)`

SetLevel sets Level field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


