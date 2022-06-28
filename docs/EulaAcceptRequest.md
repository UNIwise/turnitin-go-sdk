# EulaAcceptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | The unique id of the user in the external system  | [optional] 
**AcceptedTimestamp** | Pointer to **time.Time** | The timestamp marking when the EULA was accepted  | [optional] 
**Language** | Pointer to **string** | The language code for which language instance of the EULA version was accepted  | [optional] 
**Version** | Pointer to **string** | The unique name of the EULA Version  | [optional] 

## Methods

### NewEulaAcceptRequest

`func NewEulaAcceptRequest() *EulaAcceptRequest`

NewEulaAcceptRequest instantiates a new EulaAcceptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEulaAcceptRequestWithDefaults

`func NewEulaAcceptRequestWithDefaults() *EulaAcceptRequest`

NewEulaAcceptRequestWithDefaults instantiates a new EulaAcceptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *EulaAcceptRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EulaAcceptRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EulaAcceptRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EulaAcceptRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAcceptedTimestamp

`func (o *EulaAcceptRequest) GetAcceptedTimestamp() time.Time`

GetAcceptedTimestamp returns the AcceptedTimestamp field if non-nil, zero value otherwise.

### GetAcceptedTimestampOk

`func (o *EulaAcceptRequest) GetAcceptedTimestampOk() (*time.Time, bool)`

GetAcceptedTimestampOk returns a tuple with the AcceptedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTimestamp

`func (o *EulaAcceptRequest) SetAcceptedTimestamp(v time.Time)`

SetAcceptedTimestamp sets AcceptedTimestamp field to given value.

### HasAcceptedTimestamp

`func (o *EulaAcceptRequest) HasAcceptedTimestamp() bool`

HasAcceptedTimestamp returns a boolean if a field has been set.

### GetLanguage

`func (o *EulaAcceptRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *EulaAcceptRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *EulaAcceptRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *EulaAcceptRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetVersion

`func (o *EulaAcceptRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EulaAcceptRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EulaAcceptRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EulaAcceptRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


