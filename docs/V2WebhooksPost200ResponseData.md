# V2WebhooksPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetUrl** | **string** | URL where the webhook events will be delivered to. | 
**Subscriptions** | [**[]V2WebhooksGet200ResponseDataInnerSubscriptionsInner**](V2WebhooksGet200ResponseDataInnerSubscriptionsInner.md) | One or more events the webhook is subscribed to. | 
**Id** | [**V2WebhooksGet200ResponseDataInnerId**](V2WebhooksGet200ResponseDataInnerId.md) |  | 
**Status** | **string** | The state of the webhook. Webhooks marked as active and degraded will receive events, inactive ones will not. If a webhook remains in the degraded state for 7 days, it will be marked inactive. | 
**CreatedAt** | **string** | When the webhook was created. | 
**Secret** | **string** | The key which is used to sign the webhook events. This is only shown when setting up the webhook initially. | 

## Methods

### NewV2WebhooksPost200ResponseData

`func NewV2WebhooksPost200ResponseData(targetUrl string, subscriptions []V2WebhooksGet200ResponseDataInnerSubscriptionsInner, id V2WebhooksGet200ResponseDataInnerId, status string, createdAt string, secret string, ) *V2WebhooksPost200ResponseData`

NewV2WebhooksPost200ResponseData instantiates a new V2WebhooksPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2WebhooksPost200ResponseDataWithDefaults

`func NewV2WebhooksPost200ResponseDataWithDefaults() *V2WebhooksPost200ResponseData`

NewV2WebhooksPost200ResponseDataWithDefaults instantiates a new V2WebhooksPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetUrl

`func (o *V2WebhooksPost200ResponseData) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *V2WebhooksPost200ResponseData) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *V2WebhooksPost200ResponseData) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### GetSubscriptions

`func (o *V2WebhooksPost200ResponseData) GetSubscriptions() []V2WebhooksGet200ResponseDataInnerSubscriptionsInner`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *V2WebhooksPost200ResponseData) GetSubscriptionsOk() (*[]V2WebhooksGet200ResponseDataInnerSubscriptionsInner, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *V2WebhooksPost200ResponseData) SetSubscriptions(v []V2WebhooksGet200ResponseDataInnerSubscriptionsInner)`

SetSubscriptions sets Subscriptions field to given value.


### GetId

`func (o *V2WebhooksPost200ResponseData) GetId() V2WebhooksGet200ResponseDataInnerId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2WebhooksPost200ResponseData) GetIdOk() (*V2WebhooksGet200ResponseDataInnerId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2WebhooksPost200ResponseData) SetId(v V2WebhooksGet200ResponseDataInnerId)`

SetId sets Id field to given value.


### GetStatus

`func (o *V2WebhooksPost200ResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2WebhooksPost200ResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2WebhooksPost200ResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *V2WebhooksPost200ResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2WebhooksPost200ResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2WebhooksPost200ResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetSecret

`func (o *V2WebhooksPost200ResponseData) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *V2WebhooksPost200ResponseData) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *V2WebhooksPost200ResponseData) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


