# SimilaritySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMode** | Pointer to **string** | default similarity mode when viewing a report; set to either match_overview or all_sources | [optional] 
**Modes** | Pointer to [**SimilaritySettingsModes**](SimilaritySettingsModes.md) |  | [optional] 
**ViewSettings** | Pointer to [**SimilaritySettingsViewSettings**](SimilaritySettingsViewSettings.md) |  | [optional] 

## Methods

### NewSimilaritySettings

`func NewSimilaritySettings() *SimilaritySettings`

NewSimilaritySettings instantiates a new SimilaritySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimilaritySettingsWithDefaults

`func NewSimilaritySettingsWithDefaults() *SimilaritySettings`

NewSimilaritySettingsWithDefaults instantiates a new SimilaritySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMode

`func (o *SimilaritySettings) GetDefaultMode() string`

GetDefaultMode returns the DefaultMode field if non-nil, zero value otherwise.

### GetDefaultModeOk

`func (o *SimilaritySettings) GetDefaultModeOk() (*string, bool)`

GetDefaultModeOk returns a tuple with the DefaultMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMode

`func (o *SimilaritySettings) SetDefaultMode(v string)`

SetDefaultMode sets DefaultMode field to given value.

### HasDefaultMode

`func (o *SimilaritySettings) HasDefaultMode() bool`

HasDefaultMode returns a boolean if a field has been set.

### GetModes

`func (o *SimilaritySettings) GetModes() SimilaritySettingsModes`

GetModes returns the Modes field if non-nil, zero value otherwise.

### GetModesOk

`func (o *SimilaritySettings) GetModesOk() (*SimilaritySettingsModes, bool)`

GetModesOk returns a tuple with the Modes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModes

`func (o *SimilaritySettings) SetModes(v SimilaritySettingsModes)`

SetModes sets Modes field to given value.

### HasModes

`func (o *SimilaritySettings) HasModes() bool`

HasModes returns a boolean if a field has been set.

### GetViewSettings

`func (o *SimilaritySettings) GetViewSettings() SimilaritySettingsViewSettings`

GetViewSettings returns the ViewSettings field if non-nil, zero value otherwise.

### GetViewSettingsOk

`func (o *SimilaritySettings) GetViewSettingsOk() (*SimilaritySettingsViewSettings, bool)`

GetViewSettingsOk returns a tuple with the ViewSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewSettings

`func (o *SimilaritySettings) SetViewSettings(v SimilaritySettingsViewSettings)`

SetViewSettings sets ViewSettings field to given value.

### HasViewSettings

`func (o *SimilaritySettings) HasViewSettings() bool`

HasViewSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


