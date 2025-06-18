# V2WebhooksGet200ResponseDataInnerSubscriptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | Type of event the webhook is subscribed to. | 
**Filter** | [**NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilter**](V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilter.md) |  | 

## Methods

### NewV2WebhooksGet200ResponseDataInnerSubscriptionsInner

`func NewV2WebhooksGet200ResponseDataInnerSubscriptionsInner(eventType string, filter NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilter, ) *V2WebhooksGet200ResponseDataInnerSubscriptionsInner`

NewV2WebhooksGet200ResponseDataInnerSubscriptionsInner instantiates a new V2WebhooksGet200ResponseDataInnerSubscriptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2WebhooksGet200ResponseDataInnerSubscriptionsInnerWithDefaults

`func NewV2WebhooksGet200ResponseDataInnerSubscriptionsInnerWithDefaults() *V2WebhooksGet200ResponseDataInnerSubscriptionsInner`

NewV2WebhooksGet200ResponseDataInnerSubscriptionsInnerWithDefaults instantiates a new V2WebhooksGet200ResponseDataInnerSubscriptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInner) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetFilter

`func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInner) GetFilter() V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInner) GetFilterOk() (*V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInner) SetFilter(v V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilter)`

SetFilter sets Filter field to given value.


### SetFilterNil

`func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInner) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInner) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


