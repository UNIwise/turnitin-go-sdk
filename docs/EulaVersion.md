# EulaVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The unique name of the EULA Version  | [optional] 
**ValidFrom** | Pointer to **time.Time** | The starting date indicating when acceptence of this EULA is considered valid  | [optional] 
**ValidUntil** | Pointer to **NullableTime** | The ending date indicating when acceptence of this EULA is no longer valid  | [optional] 
**Url** | Pointer to **string** | The url where the corresponding EULA page can be found  | [optional] 
**AvailableLanguages** | Pointer to **[]string** | The languages (instances) of this version. 21 language locales are currently supported.  | [optional] 

## Methods

### NewEulaVersion

`func NewEulaVersion() *EulaVersion`

NewEulaVersion instantiates a new EulaVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEulaVersionWithDefaults

`func NewEulaVersionWithDefaults() *EulaVersion`

NewEulaVersionWithDefaults instantiates a new EulaVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *EulaVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EulaVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EulaVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EulaVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetValidFrom

`func (o *EulaVersion) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *EulaVersion) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *EulaVersion) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *EulaVersion) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidUntil

`func (o *EulaVersion) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *EulaVersion) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *EulaVersion) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *EulaVersion) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### SetValidUntilNil

`func (o *EulaVersion) SetValidUntilNil(b bool)`

 SetValidUntilNil sets the value for ValidUntil to be an explicit nil

### UnsetValidUntil
`func (o *EulaVersion) UnsetValidUntil()`

UnsetValidUntil ensures that no value is present for ValidUntil, not even an explicit nil
### GetUrl

`func (o *EulaVersion) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EulaVersion) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EulaVersion) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EulaVersion) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAvailableLanguages

`func (o *EulaVersion) GetAvailableLanguages() []string`

GetAvailableLanguages returns the AvailableLanguages field if non-nil, zero value otherwise.

### GetAvailableLanguagesOk

`func (o *EulaVersion) GetAvailableLanguagesOk() (*[]string, bool)`

GetAvailableLanguagesOk returns a tuple with the AvailableLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableLanguages

`func (o *EulaVersion) SetAvailableLanguages(v []string)`

SetAvailableLanguages sets AvailableLanguages field to given value.

### HasAvailableLanguages

`func (o *EulaVersion) HasAvailableLanguages() bool`

HasAvailableLanguages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


