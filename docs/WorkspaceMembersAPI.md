# \WorkspaceMembersAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2WorkspaceMembersGet**](WorkspaceMembersAPI.md#V2WorkspaceMembersGet) | **Get** /v2/workspace_members | List workspace members
[**V2WorkspaceMembersWorkspaceMemberIdGet**](WorkspaceMembersAPI.md#V2WorkspaceMembersWorkspaceMemberIdGet) | **Get** /v2/workspace_members/{workspace_member_id} | Get a workspace member



## V2WorkspaceMembersGet

> V2WorkspaceMembersGet200Response V2WorkspaceMembersGet(ctx).Execute()

List workspace members



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceMembersAPI.V2WorkspaceMembersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMembersAPI.V2WorkspaceMembersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2WorkspaceMembersGet`: V2WorkspaceMembersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMembersAPI.V2WorkspaceMembersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2WorkspaceMembersGetRequest struct via the builder pattern


### Return type

[**V2WorkspaceMembersGet200Response**](V2WorkspaceMembersGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2WorkspaceMembersWorkspaceMemberIdGet

> V2WorkspaceMembersWorkspaceMemberIdGet200Response V2WorkspaceMembersWorkspaceMemberIdGet(ctx, workspaceMemberId).Execute()

Get a workspace member



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceMemberId := "50cf242c-7fa3-4cad-87d0-75b1af71c57b" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceMembersAPI.V2WorkspaceMembersWorkspaceMemberIdGet(context.Background(), workspaceMemberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceMembersAPI.V2WorkspaceMembersWorkspaceMemberIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2WorkspaceMembersWorkspaceMemberIdGet`: V2WorkspaceMembersWorkspaceMemberIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceMembersAPI.V2WorkspaceMembersWorkspaceMemberIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceMemberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2WorkspaceMembersWorkspaceMemberIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2WorkspaceMembersWorkspaceMemberIdGet200Response**](V2WorkspaceMembersWorkspaceMemberIdGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

