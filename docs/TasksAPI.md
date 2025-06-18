# \TasksAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2TasksGet**](TasksAPI.md#V2TasksGet) | **Get** /v2/tasks | List tasks
[**V2TasksPost**](TasksAPI.md#V2TasksPost) | **Post** /v2/tasks | Create a task
[**V2TasksTaskIdDelete**](TasksAPI.md#V2TasksTaskIdDelete) | **Delete** /v2/tasks/{task_id} | Delete a task
[**V2TasksTaskIdGet**](TasksAPI.md#V2TasksTaskIdGet) | **Get** /v2/tasks/{task_id} | Get a task
[**V2TasksTaskIdPatch**](TasksAPI.md#V2TasksTaskIdPatch) | **Patch** /v2/tasks/{task_id} | Update a task



## V2TasksGet

> V2TasksGet200Response V2TasksGet(ctx).Limit(limit).Offset(offset).Sort(sort).LinkedObject(linkedObject).LinkedRecordId(linkedRecordId).Assignee(assignee).IsCompleted(isCompleted).Execute()

List tasks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/fastah/libattio"
)

func main() {
	limit := int32(10) // int32 |  (optional)
	offset := int32(5) // int32 |  (optional)
	sort := "created_at:desc" // string |  (optional)
	linkedObject := "people" // string |  (optional)
	linkedRecordId := "891dcbfc-9141-415d-9b2a-2238a6cc012d" // string |  (optional)
	assignee := "assignee_example" // string |  (optional)
	isCompleted := true // bool |  (optional)

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.V2TasksGet(context.Background()).Limit(limit).Offset(offset).Sort(sort).LinkedObject(linkedObject).LinkedRecordId(linkedRecordId).Assignee(assignee).IsCompleted(isCompleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.V2TasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TasksGet`: V2TasksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.V2TasksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2TasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sort** | **string** |  | 
 **linkedObject** | **string** |  | 
 **linkedRecordId** | **string** |  | 
 **assignee** | **string** |  | 
 **isCompleted** | **bool** |  | 

### Return type

[**V2TasksGet200Response**](V2TasksGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2TasksPost

> V2TasksPost200Response V2TasksPost(ctx).V2TasksPostRequest(v2TasksPostRequest).Execute()

Create a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/fastah/libattio"
)

func main() {
	v2TasksPostRequest := *libattio.NewV2TasksPostRequest(*libattio.NewV2TasksPostRequestData("Follow up on current software solutions", "Format_example", "2023-01-01T15:00:00.000000000Z", false, []libattio.V2TasksPostRequestDataLinkedRecordsInner{*libattio.NewV2TasksPostRequestDataLinkedRecordsInner("people", "891dcbfc-9141-415d-9b2a-2238a6cc012d", []libattio.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner{*libattio.NewV2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner()})}, []libattio.V2TasksPostRequestDataAssigneesInner{*libattio.NewV2TasksPostRequestDataAssigneesInner("workspace-member", "50cf242c-7fa3-4cad-87d0-75b1af71c57b", "alice@attio.com")})) // V2TasksPostRequest | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.V2TasksPost(context.Background()).V2TasksPostRequest(v2TasksPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.V2TasksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TasksPost`: V2TasksPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.V2TasksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2TasksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2TasksPostRequest** | [**V2TasksPostRequest**](V2TasksPostRequest.md) |  | 

### Return type

[**V2TasksPost200Response**](V2TasksPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2TasksTaskIdDelete

> map[string]interface{} V2TasksTaskIdDelete(ctx, taskId).Execute()

Delete a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/fastah/libattio"
)

func main() {
	taskId := "649e34f4-c39a-4f4d-99ef-48a36bef8f04" // string | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.V2TasksTaskIdDelete(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.V2TasksTaskIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TasksTaskIdDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.V2TasksTaskIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2TasksTaskIdDeleteRequest struct via the builder pattern


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


## V2TasksTaskIdGet

> V2TasksPost200Response V2TasksTaskIdGet(ctx, taskId).Execute()

Get a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/fastah/libattio"
)

func main() {
	taskId := "649e34f4-c39a-4f4d-99ef-48a36bef8f04" // string | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.V2TasksTaskIdGet(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.V2TasksTaskIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TasksTaskIdGet`: V2TasksPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.V2TasksTaskIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2TasksTaskIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2TasksPost200Response**](V2TasksPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2TasksTaskIdPatch

> V2TasksPost200Response V2TasksTaskIdPatch(ctx, taskId).V2TasksTaskIdPatchRequest(v2TasksTaskIdPatchRequest).Execute()

Update a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	libattio "github.com/fastah/libattio"
)

func main() {
	taskId := "649e34f4-c39a-4f4d-99ef-48a36bef8f04" // string | 
	v2TasksTaskIdPatchRequest := *libattio.NewV2TasksTaskIdPatchRequest(*libattio.NewV2TasksTaskIdPatchRequestData()) // V2TasksTaskIdPatchRequest | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.V2TasksTaskIdPatch(context.Background(), taskId).V2TasksTaskIdPatchRequest(v2TasksTaskIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.V2TasksTaskIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2TasksTaskIdPatch`: V2TasksPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.V2TasksTaskIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2TasksTaskIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2TasksTaskIdPatchRequest** | [**V2TasksTaskIdPatchRequest**](V2TasksTaskIdPatchRequest.md) |  | 

### Return type

[**V2TasksPost200Response**](V2TasksPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

