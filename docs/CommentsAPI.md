# \CommentsAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2CommentsCommentIdDelete**](CommentsAPI.md#V2CommentsCommentIdDelete) | **Delete** /v2/comments/{comment_id} | Delete a comment
[**V2CommentsCommentIdGet**](CommentsAPI.md#V2CommentsCommentIdGet) | **Get** /v2/comments/{comment_id} | Get a comment
[**V2CommentsPost**](CommentsAPI.md#V2CommentsPost) | **Post** /v2/comments | Create a comment



## V2CommentsCommentIdDelete

> map[string]interface{} V2CommentsCommentIdDelete(ctx, commentId).Execute()

Delete a comment



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
	commentId := "aa1dc1d9-93ac-4c6c-987e-16b6eea9aab2" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.V2CommentsCommentIdDelete(context.Background(), commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.V2CommentsCommentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2CommentsCommentIdDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.V2CommentsCommentIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CommentsCommentIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CommentsCommentIdGet

> V2CommentsPost200Response V2CommentsCommentIdGet(ctx, commentId).Execute()

Get a comment



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
	commentId := "aa1dc1d9-93ac-4c6c-987e-16b6eea9aab2" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.V2CommentsCommentIdGet(context.Background(), commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.V2CommentsCommentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2CommentsCommentIdGet`: V2CommentsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.V2CommentsCommentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CommentsCommentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2CommentsPost200Response**](V2CommentsPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CommentsPost

> V2CommentsPost200Response V2CommentsPost(ctx).V2CommentsPostRequest(v2CommentsPostRequest).Execute()

Create a comment



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
	v2CommentsPostRequest := *openapiclient.NewV2CommentsPostRequest(*openapiclient.NewV2CommentsPostRequestData("Format_example", "If I put the email address of my colleague on Attio in here, e.g. alice@attio.com, they will be notified. Other emails (e.g. person@example.com) will be turned into clickable links.", *openapiclient.NewV2CommentsPostRequestDataAnyOfAuthor("Type_example", "Id_example"), "aa1dc1d9-93ac-4c6c-987e-16b6eea9aab2", *openapiclient.NewV2CommentsPostRequestDataAnyOf1Record("97052eb9-e65e-443f-a297-f2d9a4a7f795", "bf071e1f-6035-429d-b874-d83ea64ea13b"), *openapiclient.NewV2CommentsPostRequestDataAnyOf2Entry("33ebdbe9-e529-47c9-b894-0ba25e9c15c0", "2e6e29ea-c4e0-4f44-842d-78a891f8c156"))) // V2CommentsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.V2CommentsPost(context.Background()).V2CommentsPostRequest(v2CommentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.V2CommentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2CommentsPost`: V2CommentsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.V2CommentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2CommentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2CommentsPostRequest** | [**V2CommentsPostRequest**](V2CommentsPostRequest.md) |  | 

### Return type

[**V2CommentsPost200Response**](V2CommentsPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

