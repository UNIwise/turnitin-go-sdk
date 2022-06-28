# WebhookWithSecretAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SigningSecret** | ***os.File** | a sercret used to sign all callbacks for this webhook | 

## Methods

### NewWebhookWithSecretAllOf

`func NewWebhookWithSecretAllOf(signingSecret *os.File, ) *WebhookWithSecretAllOf`

NewWebhookWithSecretAllOf instantiates a new WebhookWithSecretAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithSecretAllOfWithDefaults

`func NewWebhookWithSecretAllOfWithDefaults() *WebhookWithSecretAllOf`

NewWebhookWithSecretAllOfWithDefaults instantiates a new WebhookWithSecretAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigningSecret

`func (o *WebhookWithSecretAllOf) GetSigningSecret() *os.File`

GetSigningSecret returns the SigningSecret field if non-nil, zero value otherwise.

### GetSigningSecretOk

`func (o *WebhookWithSecretAllOf) GetSigningSecretOk() (**os.File, bool)`

GetSigningSecretOk returns a tuple with the SigningSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecret

`func (o *WebhookWithSecretAllOf) SetSigningSecret(v *os.File)`

SetSigningSecret sets SigningSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


