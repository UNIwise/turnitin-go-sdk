# RateLimitError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | an HTTP Response JobStatus Code | [optional] 
**Message** | Pointer to **string** | A message explaining what happened | [optional] 

## Methods

### NewRateLimitError

`func NewRateLimitError() *RateLimitError`

NewRateLimitError instantiates a new RateLimitError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitErrorWithDefaults

`func NewRateLimitErrorWithDefaults() *RateLimitError`

NewRateLimitErrorWithDefaults instantiates a new RateLimitError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RateLimitError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RateLimitError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RateLimitError) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RateLimitError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *RateLimitError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RateLimitError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RateLimitError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RateLimitError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


