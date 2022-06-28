# WebhookAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | unique id of webhook | 
**CreatedTime** | Pointer to **time.Time** | RFC3339  timestamp of when this Webhook was initially created. This is the time at which the POST to /webhooks was made.  | [optional] 

## Methods

### NewWebhookAllOf

`func NewWebhookAllOf(id string, ) *WebhookAllOf`

NewWebhookAllOf instantiates a new WebhookAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAllOfWithDefaults

`func NewWebhookAllOfWithDefaults() *WebhookAllOf`

NewWebhookAllOfWithDefaults instantiates a new WebhookAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedTime

`func (o *WebhookAllOf) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *WebhookAllOf) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *WebhookAllOf) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *WebhookAllOf) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


