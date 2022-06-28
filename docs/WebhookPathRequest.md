# WebhookPathRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | a human readable description of the webhook | [optional] 
**EventTypes** | **[]string** | an array of the types of callbacks that will be sent to this webhook | 

## Methods

### NewWebhookPathRequest

`func NewWebhookPathRequest(eventTypes []string, ) *WebhookPathRequest`

NewWebhookPathRequest instantiates a new WebhookPathRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPathRequestWithDefaults

`func NewWebhookPathRequestWithDefaults() *WebhookPathRequest`

NewWebhookPathRequestWithDefaults instantiates a new WebhookPathRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WebhookPathRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookPathRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookPathRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookPathRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventTypes

`func (o *WebhookPathRequest) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *WebhookPathRequest) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *WebhookPathRequest) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


