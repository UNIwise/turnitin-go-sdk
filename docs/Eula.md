# Eula

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedTimestamp** | Pointer to **time.Time** | The timestamp marking when the EULA was accepted  | [optional] 
**Language** | Pointer to **string** | The language code for which language instance of the EULA version was accepted  | [optional] 
**Version** | Pointer to **string** | The unique name of the EULA Version  | [optional] 

## Methods

### NewEula

`func NewEula() *Eula`

NewEula instantiates a new Eula object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEulaWithDefaults

`func NewEulaWithDefaults() *Eula`

NewEulaWithDefaults instantiates a new Eula object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedTimestamp

`func (o *Eula) GetAcceptedTimestamp() time.Time`

GetAcceptedTimestamp returns the AcceptedTimestamp field if non-nil, zero value otherwise.

### GetAcceptedTimestampOk

`func (o *Eula) GetAcceptedTimestampOk() (*time.Time, bool)`

GetAcceptedTimestampOk returns a tuple with the AcceptedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTimestamp

`func (o *Eula) SetAcceptedTimestamp(v time.Time)`

SetAcceptedTimestamp sets AcceptedTimestamp field to given value.

### HasAcceptedTimestamp

`func (o *Eula) HasAcceptedTimestamp() bool`

HasAcceptedTimestamp returns a boolean if a field has been set.

### GetLanguage

`func (o *Eula) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Eula) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Eula) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Eula) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetVersion

`func (o *Eula) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Eula) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Eula) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Eula) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


