# SubmissionSizeError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | an HTTP Response JobStatus Code | [optional] 
**Message** | Pointer to **string** | A message explaining what happened | [optional] 

## Methods

### NewSubmissionSizeError

`func NewSubmissionSizeError() *SubmissionSizeError`

NewSubmissionSizeError instantiates a new SubmissionSizeError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionSizeErrorWithDefaults

`func NewSubmissionSizeErrorWithDefaults() *SubmissionSizeError`

NewSubmissionSizeErrorWithDefaults instantiates a new SubmissionSizeError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SubmissionSizeError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmissionSizeError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmissionSizeError) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubmissionSizeError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *SubmissionSizeError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SubmissionSizeError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SubmissionSizeError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SubmissionSizeError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


