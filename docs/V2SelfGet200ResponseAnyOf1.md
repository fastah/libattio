# V2SelfGet200ResponseAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Whether the token is currently active and usable. | 
**Scope** | **string** | A space-separated list of scopes associated with this token | 
**ClientId** | **string** | The app ID of the OAuth application that requested this token | 
**TokenType** | **string** | The type of token, always Bearer for tokens acquired via the OAuth 2.0 flow. | 
**Exp** | **NullableFloat32** | The time at which this token will expire, if set, as a number of seconds since January 1 1970 UTC. | 
**Iat** | **float32** | The time at which this token was issued, as a number of seconds since January 1 1970 UTC. | 
**Sub** | **string** | Since Bearer tokens grant Workspace-level permissions, this property contains the workspace_id. | 
**Aud** | **string** | The intended audience for this token, for Bearer tokens this is the same as the client_id. | 
**Iss** | **string** | The issuer of the token. Always attio.com | 
**AuthorizedByWorkspaceMemberId** | **NullableString** | The ID of the workspace member who authorised this token initially, if known | 
**WorkspaceId** | **string** | The ID of the workspace the token is scoped to. | 
**WorkspaceName** | **string** | The name of the workspace the token is scoped to. | 
**WorkspaceSlug** | **string** | The slug of the workspace the token is scoped to. | 
**WorkspaceLogoUrl** | **NullableString** | The logo URL of the workspace the token is scoped to. | 

## Methods

### NewV2SelfGet200ResponseAnyOf1

`func NewV2SelfGet200ResponseAnyOf1(active bool, scope string, clientId string, tokenType string, exp NullableFloat32, iat float32, sub string, aud string, iss string, authorizedByWorkspaceMemberId NullableString, workspaceId string, workspaceName string, workspaceSlug string, workspaceLogoUrl NullableString, ) *V2SelfGet200ResponseAnyOf1`

NewV2SelfGet200ResponseAnyOf1 instantiates a new V2SelfGet200ResponseAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2SelfGet200ResponseAnyOf1WithDefaults

`func NewV2SelfGet200ResponseAnyOf1WithDefaults() *V2SelfGet200ResponseAnyOf1`

NewV2SelfGet200ResponseAnyOf1WithDefaults instantiates a new V2SelfGet200ResponseAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *V2SelfGet200ResponseAnyOf1) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *V2SelfGet200ResponseAnyOf1) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *V2SelfGet200ResponseAnyOf1) SetActive(v bool)`

SetActive sets Active field to given value.


### GetScope

`func (o *V2SelfGet200ResponseAnyOf1) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *V2SelfGet200ResponseAnyOf1) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *V2SelfGet200ResponseAnyOf1) SetScope(v string)`

SetScope sets Scope field to given value.


### GetClientId

`func (o *V2SelfGet200ResponseAnyOf1) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *V2SelfGet200ResponseAnyOf1) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *V2SelfGet200ResponseAnyOf1) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetTokenType

`func (o *V2SelfGet200ResponseAnyOf1) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *V2SelfGet200ResponseAnyOf1) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *V2SelfGet200ResponseAnyOf1) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExp

`func (o *V2SelfGet200ResponseAnyOf1) GetExp() float32`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *V2SelfGet200ResponseAnyOf1) GetExpOk() (*float32, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *V2SelfGet200ResponseAnyOf1) SetExp(v float32)`

SetExp sets Exp field to given value.


### SetExpNil

`func (o *V2SelfGet200ResponseAnyOf1) SetExpNil(b bool)`

 SetExpNil sets the value for Exp to be an explicit nil

### UnsetExp
`func (o *V2SelfGet200ResponseAnyOf1) UnsetExp()`

UnsetExp ensures that no value is present for Exp, not even an explicit nil
### GetIat

`func (o *V2SelfGet200ResponseAnyOf1) GetIat() float32`

GetIat returns the Iat field if non-nil, zero value otherwise.

### GetIatOk

`func (o *V2SelfGet200ResponseAnyOf1) GetIatOk() (*float32, bool)`

GetIatOk returns a tuple with the Iat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIat

`func (o *V2SelfGet200ResponseAnyOf1) SetIat(v float32)`

SetIat sets Iat field to given value.


### GetSub

`func (o *V2SelfGet200ResponseAnyOf1) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *V2SelfGet200ResponseAnyOf1) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *V2SelfGet200ResponseAnyOf1) SetSub(v string)`

SetSub sets Sub field to given value.


### GetAud

`func (o *V2SelfGet200ResponseAnyOf1) GetAud() string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *V2SelfGet200ResponseAnyOf1) GetAudOk() (*string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *V2SelfGet200ResponseAnyOf1) SetAud(v string)`

SetAud sets Aud field to given value.


### GetIss

`func (o *V2SelfGet200ResponseAnyOf1) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *V2SelfGet200ResponseAnyOf1) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *V2SelfGet200ResponseAnyOf1) SetIss(v string)`

SetIss sets Iss field to given value.


### GetAuthorizedByWorkspaceMemberId

`func (o *V2SelfGet200ResponseAnyOf1) GetAuthorizedByWorkspaceMemberId() string`

GetAuthorizedByWorkspaceMemberId returns the AuthorizedByWorkspaceMemberId field if non-nil, zero value otherwise.

### GetAuthorizedByWorkspaceMemberIdOk

`func (o *V2SelfGet200ResponseAnyOf1) GetAuthorizedByWorkspaceMemberIdOk() (*string, bool)`

GetAuthorizedByWorkspaceMemberIdOk returns a tuple with the AuthorizedByWorkspaceMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedByWorkspaceMemberId

`func (o *V2SelfGet200ResponseAnyOf1) SetAuthorizedByWorkspaceMemberId(v string)`

SetAuthorizedByWorkspaceMemberId sets AuthorizedByWorkspaceMemberId field to given value.


### SetAuthorizedByWorkspaceMemberIdNil

`func (o *V2SelfGet200ResponseAnyOf1) SetAuthorizedByWorkspaceMemberIdNil(b bool)`

 SetAuthorizedByWorkspaceMemberIdNil sets the value for AuthorizedByWorkspaceMemberId to be an explicit nil

### UnsetAuthorizedByWorkspaceMemberId
`func (o *V2SelfGet200ResponseAnyOf1) UnsetAuthorizedByWorkspaceMemberId()`

UnsetAuthorizedByWorkspaceMemberId ensures that no value is present for AuthorizedByWorkspaceMemberId, not even an explicit nil
### GetWorkspaceId

`func (o *V2SelfGet200ResponseAnyOf1) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *V2SelfGet200ResponseAnyOf1) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *V2SelfGet200ResponseAnyOf1) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetWorkspaceName

`func (o *V2SelfGet200ResponseAnyOf1) GetWorkspaceName() string`

GetWorkspaceName returns the WorkspaceName field if non-nil, zero value otherwise.

### GetWorkspaceNameOk

`func (o *V2SelfGet200ResponseAnyOf1) GetWorkspaceNameOk() (*string, bool)`

GetWorkspaceNameOk returns a tuple with the WorkspaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceName

`func (o *V2SelfGet200ResponseAnyOf1) SetWorkspaceName(v string)`

SetWorkspaceName sets WorkspaceName field to given value.


### GetWorkspaceSlug

`func (o *V2SelfGet200ResponseAnyOf1) GetWorkspaceSlug() string`

GetWorkspaceSlug returns the WorkspaceSlug field if non-nil, zero value otherwise.

### GetWorkspaceSlugOk

`func (o *V2SelfGet200ResponseAnyOf1) GetWorkspaceSlugOk() (*string, bool)`

GetWorkspaceSlugOk returns a tuple with the WorkspaceSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSlug

`func (o *V2SelfGet200ResponseAnyOf1) SetWorkspaceSlug(v string)`

SetWorkspaceSlug sets WorkspaceSlug field to given value.


### GetWorkspaceLogoUrl

`func (o *V2SelfGet200ResponseAnyOf1) GetWorkspaceLogoUrl() string`

GetWorkspaceLogoUrl returns the WorkspaceLogoUrl field if non-nil, zero value otherwise.

### GetWorkspaceLogoUrlOk

`func (o *V2SelfGet200ResponseAnyOf1) GetWorkspaceLogoUrlOk() (*string, bool)`

GetWorkspaceLogoUrlOk returns a tuple with the WorkspaceLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceLogoUrl

`func (o *V2SelfGet200ResponseAnyOf1) SetWorkspaceLogoUrl(v string)`

SetWorkspaceLogoUrl sets WorkspaceLogoUrl field to given value.


### SetWorkspaceLogoUrlNil

`func (o *V2SelfGet200ResponseAnyOf1) SetWorkspaceLogoUrlNil(b bool)`

 SetWorkspaceLogoUrlNil sets the value for WorkspaceLogoUrl to be an explicit nil

### UnsetWorkspaceLogoUrl
`func (o *V2SelfGet200ResponseAnyOf1) UnsetWorkspaceLogoUrl()`

UnsetWorkspaceLogoUrl ensures that no value is present for WorkspaceLogoUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


