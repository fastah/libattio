# V2TasksPostRequestDataLinkedRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetObject** | **string** | A UUID or slug to identify the object that the referenced record belongs to. | 
**TargetRecordId** | **string** | The ID of the parent record the task refers to. | 
**SlugOrIdOfMatchingAttribute** | [**[]V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner**](V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner.md) | In addition to referencing records directly by record ID, you may also reference by a matching attribute of your choice. For example, if you want to add a reference to the person record with email \&quot;alice@website.com\&quot;, you should pass a value with &#x60;target_object&#x60; set to &#x60;\&quot;people\&quot;&#x60; and &#x60;email_addresses&#x60; set to &#x60;[{email_address:\&quot;alice@website.com\&quot;}]&#x60;. The key should be the slug or ID of the matching attribute you would like to use and the value should be an array containing a single value of the appropriate attribute type (as specified below). Matching on multiple values is not currently supported. Matching attributes must be unique. This process is similar to how you use the &#x60;matching_attribute&#x60; query param in Attio&#39;s [assert endpoints](/rest-api/endpoint-reference/records/assert-a-record). | 

## Methods

### NewV2TasksPostRequestDataLinkedRecordsInner

`func NewV2TasksPostRequestDataLinkedRecordsInner(targetObject string, targetRecordId string, slugOrIdOfMatchingAttribute []V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner, ) *V2TasksPostRequestDataLinkedRecordsInner`

NewV2TasksPostRequestDataLinkedRecordsInner instantiates a new V2TasksPostRequestDataLinkedRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2TasksPostRequestDataLinkedRecordsInnerWithDefaults

`func NewV2TasksPostRequestDataLinkedRecordsInnerWithDefaults() *V2TasksPostRequestDataLinkedRecordsInner`

NewV2TasksPostRequestDataLinkedRecordsInnerWithDefaults instantiates a new V2TasksPostRequestDataLinkedRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetObject

`func (o *V2TasksPostRequestDataLinkedRecordsInner) GetTargetObject() string`

GetTargetObject returns the TargetObject field if non-nil, zero value otherwise.

### GetTargetObjectOk

`func (o *V2TasksPostRequestDataLinkedRecordsInner) GetTargetObjectOk() (*string, bool)`

GetTargetObjectOk returns a tuple with the TargetObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetObject

`func (o *V2TasksPostRequestDataLinkedRecordsInner) SetTargetObject(v string)`

SetTargetObject sets TargetObject field to given value.


### GetTargetRecordId

`func (o *V2TasksPostRequestDataLinkedRecordsInner) GetTargetRecordId() string`

GetTargetRecordId returns the TargetRecordId field if non-nil, zero value otherwise.

### GetTargetRecordIdOk

`func (o *V2TasksPostRequestDataLinkedRecordsInner) GetTargetRecordIdOk() (*string, bool)`

GetTargetRecordIdOk returns a tuple with the TargetRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRecordId

`func (o *V2TasksPostRequestDataLinkedRecordsInner) SetTargetRecordId(v string)`

SetTargetRecordId sets TargetRecordId field to given value.


### GetSlugOrIdOfMatchingAttribute

`func (o *V2TasksPostRequestDataLinkedRecordsInner) GetSlugOrIdOfMatchingAttribute() []V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner`

GetSlugOrIdOfMatchingAttribute returns the SlugOrIdOfMatchingAttribute field if non-nil, zero value otherwise.

### GetSlugOrIdOfMatchingAttributeOk

`func (o *V2TasksPostRequestDataLinkedRecordsInner) GetSlugOrIdOfMatchingAttributeOk() (*[]V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner, bool)`

GetSlugOrIdOfMatchingAttributeOk returns a tuple with the SlugOrIdOfMatchingAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlugOrIdOfMatchingAttribute

`func (o *V2TasksPostRequestDataLinkedRecordsInner) SetSlugOrIdOfMatchingAttribute(v []V2TasksPostRequestDataLinkedRecordsInnerAnyOf1SlugOrIdOfMatchingAttributeInner)`

SetSlugOrIdOfMatchingAttribute sets SlugOrIdOfMatchingAttribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


