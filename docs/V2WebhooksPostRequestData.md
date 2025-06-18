# V2WebhooksPostRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetUrl** | **string** | URL where the webhook events will be delivered to. | 
**Subscriptions** | [**[]V2WebhooksGet200ResponseDataInnerSubscriptionsInner**](V2WebhooksGet200ResponseDataInnerSubscriptionsInner.md) | One or more events the webhook is subscribed to. | 

## Methods

### NewV2WebhooksPostRequestData

`func NewV2WebhooksPostRequestData(targetUrl string, subscriptions []V2WebhooksGet200ResponseDataInnerSubscriptionsInner, ) *V2WebhooksPostRequestData`

NewV2WebhooksPostRequestData instantiates a new V2WebhooksPostRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2WebhooksPostRequestDataWithDefaults

`func NewV2WebhooksPostRequestDataWithDefaults() *V2WebhooksPostRequestData`

NewV2WebhooksPostRequestDataWithDefaults instantiates a new V2WebhooksPostRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetUrl

`func (o *V2WebhooksPostRequestData) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *V2WebhooksPostRequestData) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *V2WebhooksPostRequestData) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### GetSubscriptions

`func (o *V2WebhooksPostRequestData) GetSubscriptions() []V2WebhooksGet200ResponseDataInnerSubscriptionsInner`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *V2WebhooksPostRequestData) GetSubscriptionsOk() (*[]V2WebhooksGet200ResponseDataInnerSubscriptionsInner, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *V2WebhooksPostRequestData) SetSubscriptions(v []V2WebhooksGet200ResponseDataInnerSubscriptionsInner)`

SetSubscriptions sets Subscriptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


