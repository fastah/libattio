# WorkspaceMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**WorkspaceMemberId**](WorkspaceMemberId.md) |  | 
**FirstName** | **string** | The first name of the user. | 
**LastName** | **string** | The last name of the user. | 
**AvatarUrl** | **NullableString** | A URL to the user&#39;s avatar image. | 
**EmailAddress** | **string** | The user&#39;s email address. | 
**CreatedAt** | **string** | When the workspace member was created. | 
**AccessLevel** | **string** | Whether the workspace member is suspended or not and what level of privileges they have inside the workspace. We do not delete workspace members so that you can successfully attribute past actions to suspended workspace members. | 

## Methods

### NewWorkspaceMember

`func NewWorkspaceMember(id WorkspaceMemberId, firstName string, lastName string, avatarUrl NullableString, emailAddress string, createdAt string, accessLevel string, ) *WorkspaceMember`

NewWorkspaceMember instantiates a new WorkspaceMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMemberWithDefaults

`func NewWorkspaceMemberWithDefaults() *WorkspaceMember`

NewWorkspaceMemberWithDefaults instantiates a new WorkspaceMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkspaceMember) GetId() WorkspaceMemberId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceMember) GetIdOk() (*WorkspaceMemberId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceMember) SetId(v WorkspaceMemberId)`

SetId sets Id field to given value.


### GetFirstName

`func (o *WorkspaceMember) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *WorkspaceMember) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *WorkspaceMember) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *WorkspaceMember) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *WorkspaceMember) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *WorkspaceMember) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetAvatarUrl

`func (o *WorkspaceMember) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *WorkspaceMember) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *WorkspaceMember) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### SetAvatarUrlNil

`func (o *WorkspaceMember) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *WorkspaceMember) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetEmailAddress

`func (o *WorkspaceMember) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *WorkspaceMember) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *WorkspaceMember) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetCreatedAt

`func (o *WorkspaceMember) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceMember) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceMember) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetAccessLevel

`func (o *WorkspaceMember) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *WorkspaceMember) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *WorkspaceMember) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


