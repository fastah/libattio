# \WebhooksAPI

All URIs are relative to *https://api.attio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2WebhooksGet**](WebhooksAPI.md#V2WebhooksGet) | **Get** /v2/webhooks | List webhooks
[**V2WebhooksPost**](WebhooksAPI.md#V2WebhooksPost) | **Post** /v2/webhooks | Create a webhook
[**V2WebhooksWebhookIdDelete**](WebhooksAPI.md#V2WebhooksWebhookIdDelete) | **Delete** /v2/webhooks/{webhook_id} | Delete a webhook
[**V2WebhooksWebhookIdGet**](WebhooksAPI.md#V2WebhooksWebhookIdGet) | **Get** /v2/webhooks/{webhook_id} | Get a webhook
[**V2WebhooksWebhookIdPatch**](WebhooksAPI.md#V2WebhooksWebhookIdPatch) | **Patch** /v2/webhooks/{webhook_id} | Update a webhook



## V2WebhooksGet

> V2WebhooksGet200Response V2WebhooksGet(ctx).Limit(limit).Offset(offset).Execute()

List webhooks



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

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.V2WebhooksGet(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.V2WebhooksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2WebhooksGet`: V2WebhooksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.V2WebhooksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2WebhooksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**V2WebhooksGet200Response**](V2WebhooksGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2WebhooksPost

> V2WebhooksPost200Response V2WebhooksPost(ctx).V2WebhooksPostRequest(v2WebhooksPostRequest).Execute()

Create a webhook



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
	v2WebhooksPostRequest := *libattio.NewV2WebhooksPostRequest(*libattio.NewV2WebhooksPostRequestData("https://example.com/webhook", []libattio.V2WebhooksGet200ResponseDataInnerSubscriptionsInner{*libattio.NewV2WebhooksGet200ResponseDataInnerSubscriptionsInner("note.created", "TODO")})) // V2WebhooksPostRequest | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.V2WebhooksPost(context.Background()).V2WebhooksPostRequest(v2WebhooksPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.V2WebhooksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2WebhooksPost`: V2WebhooksPost200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.V2WebhooksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2WebhooksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2WebhooksPostRequest** | [**V2WebhooksPostRequest**](V2WebhooksPostRequest.md) |  | 

### Return type

[**V2WebhooksPost200Response**](V2WebhooksPost200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2WebhooksWebhookIdDelete

> map[string]interface{} V2WebhooksWebhookIdDelete(ctx, webhookId).Execute()

Delete a webhook



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
	webhookId := "23e42eaf-323a-41da-b5bb-fd67eebda553" // string | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.V2WebhooksWebhookIdDelete(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.V2WebhooksWebhookIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2WebhooksWebhookIdDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.V2WebhooksWebhookIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2WebhooksWebhookIdDeleteRequest struct via the builder pattern


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


## V2WebhooksWebhookIdGet

> V2WebhooksWebhookIdGet200Response V2WebhooksWebhookIdGet(ctx, webhookId).Execute()

Get a webhook



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
	webhookId := "23e42eaf-323a-41da-b5bb-fd67eebda553" // string | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.V2WebhooksWebhookIdGet(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.V2WebhooksWebhookIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2WebhooksWebhookIdGet`: V2WebhooksWebhookIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.V2WebhooksWebhookIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2WebhooksWebhookIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2WebhooksWebhookIdGet200Response**](V2WebhooksWebhookIdGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2WebhooksWebhookIdPatch

> V2WebhooksWebhookIdGet200Response V2WebhooksWebhookIdPatch(ctx, webhookId).V2WebhooksWebhookIdPatchRequest(v2WebhooksWebhookIdPatchRequest).Execute()

Update a webhook



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
	webhookId := "23e42eaf-323a-41da-b5bb-fd67eebda553" // string | 
	v2WebhooksWebhookIdPatchRequest := *libattio.NewV2WebhooksWebhookIdPatchRequest(*libattio.NewV2WebhooksWebhookIdPatchRequestData()) // V2WebhooksWebhookIdPatchRequest | 

	configuration := libattio.NewConfiguration()
	apiClient := libattio.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.V2WebhooksWebhookIdPatch(context.Background(), webhookId).V2WebhooksWebhookIdPatchRequest(v2WebhooksWebhookIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.V2WebhooksWebhookIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2WebhooksWebhookIdPatch`: V2WebhooksWebhookIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.V2WebhooksWebhookIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2WebhooksWebhookIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2WebhooksWebhookIdPatchRequest** | [**V2WebhooksWebhookIdPatchRequest**](V2WebhooksWebhookIdPatchRequest.md) |  | 

### Return type

[**V2WebhooksWebhookIdGet200Response**](V2WebhooksWebhookIdGet200Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

