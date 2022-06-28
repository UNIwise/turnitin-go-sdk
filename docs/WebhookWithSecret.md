# WebhookWithSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SigningSecret** | ***os.File** | a sercret used to sign all callbacks for this webhook | 
**Description** | Pointer to **string** | a human readable description of the webhook | [optional] 
**Url** | **string** | the URL to callback with events, this must be https unless (allow_insecure is true) | 
**AllowInsecure** | Pointer to **bool** | if using an non https url, this must be set to true | [optional] 
**EventTypes** | **[]string** | an array of the types of callbacks that will be sent to this webhook | 

## Methods

### NewWebhookWithSecret

`func NewWebhookWithSecret(signingSecret *os.File, url string, eventTypes []string, ) *WebhookWithSecret`

NewWebhookWithSecret instantiates a new WebhookWithSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithSecretWithDefaults

`func NewWebhookWithSecretWithDefaults() *WebhookWithSecret`

NewWebhookWithSecretWithDefaults instantiates a new WebhookWithSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigningSecret

`func (o *WebhookWithSecret) GetSigningSecret() *os.File`

GetSigningSecret returns the SigningSecret field if non-nil, zero value otherwise.

### GetSigningSecretOk

`func (o *WebhookWithSecret) GetSigningSecretOk() (**os.File, bool)`

GetSigningSecretOk returns a tuple with the SigningSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecret

`func (o *WebhookWithSecret) SetSigningSecret(v *os.File)`

SetSigningSecret sets SigningSecret field to given value.


### GetDescription

`func (o *WebhookWithSecret) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookWithSecret) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookWithSecret) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookWithSecret) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookWithSecret) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookWithSecret) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookWithSecret) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAllowInsecure

`func (o *WebhookWithSecret) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *WebhookWithSecret) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *WebhookWithSecret) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *WebhookWithSecret) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetEventTypes

`func (o *WebhookWithSecret) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *WebhookWithSecret) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *WebhookWithSecret) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


