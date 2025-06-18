# V2WebhooksWebhookIdPatchRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetUrl** | Pointer to **string** | URL where the webhook events will be delivered to. | [optional] 
**Subscriptions** | Pointer to [**[]V2WebhooksGet200ResponseDataInnerSubscriptionsInner**](V2WebhooksGet200ResponseDataInnerSubscriptionsInner.md) | One or more events the webhook is subscribed to. | [optional] 

## Methods

### NewV2WebhooksWebhookIdPatchRequestData

`func NewV2WebhooksWebhookIdPatchRequestData() *V2WebhooksWebhookIdPatchRequestData`

NewV2WebhooksWebhookIdPatchRequestData instantiates a new V2WebhooksWebhookIdPatchRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2WebhooksWebhookIdPatchRequestDataWithDefaults

`func NewV2WebhooksWebhookIdPatchRequestDataWithDefaults() *V2WebhooksWebhookIdPatchRequestData`

NewV2WebhooksWebhookIdPatchRequestDataWithDefaults instantiates a new V2WebhooksWebhookIdPatchRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetUrl

`func (o *V2WebhooksWebhookIdPatchRequestData) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *V2WebhooksWebhookIdPatchRequestData) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *V2WebhooksWebhookIdPatchRequestData) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *V2WebhooksWebhookIdPatchRequestData) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetSubscriptions

`func (o *V2WebhooksWebhookIdPatchRequestData) GetSubscriptions() []V2WebhooksGet200ResponseDataInnerSubscriptionsInner`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *V2WebhooksWebhookIdPatchRequestData) GetSubscriptionsOk() (*[]V2WebhooksGet200ResponseDataInnerSubscriptionsInner, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *V2WebhooksWebhookIdPatchRequestData) SetSubscriptions(v []V2WebhooksGet200ResponseDataInnerSubscriptionsInner)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *V2WebhooksWebhookIdPatchRequestData) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


