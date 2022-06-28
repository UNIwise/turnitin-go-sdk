# WebhookBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | a human readable description of the webhook | [optional] 
**Url** | **string** | the URL to callback with events, this must be https unless (allow_insecure is true) | 
**AllowInsecure** | Pointer to **bool** | if using an non https url, this must be set to true | [optional] 
**EventTypes** | **[]string** | an array of the types of callbacks that will be sent to this webhook | 

## Methods

### NewWebhookBase

`func NewWebhookBase(url string, eventTypes []string, ) *WebhookBase`

NewWebhookBase instantiates a new WebhookBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookBaseWithDefaults

`func NewWebhookBaseWithDefaults() *WebhookBase`

NewWebhookBaseWithDefaults instantiates a new WebhookBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WebhookBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookBase) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookBase) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookBase) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAllowInsecure

`func (o *WebhookBase) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *WebhookBase) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *WebhookBase) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *WebhookBase) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetEventTypes

`func (o *WebhookBase) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *WebhookBase) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *WebhookBase) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


