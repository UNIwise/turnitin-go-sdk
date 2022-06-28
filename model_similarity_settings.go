/*
Turnitin Core API

Turnitin Core API (TCA) provides direct API access to the core functionality provided by Turnitin. TCA supports file submission, similarity report generation, group management, and visualization of report matches via Cloud Viewer or PDF download. Below is the full flow to successfully set up an integration scope, an API Key, and make calls to TCA. Integration Scope and API Key management is done via the Admin Console UI by logging in as an admin user. For more details, go to our [developer portal documentation page](https://developers.turnitin.com/docs). ## Integration Scope and API Key Management TCA API calls must provide an API Key for authentication, so you must first have at least one integration scope associated with at least one API Key to use TCA. ### Admin Console UI First, login to Admin Console UI as an *Admin* user with permission to create Integration Scopes, under a tenant that is licensed to use the TCA product Integration Scopes (you can create a new one, or add keys to existing)    * Click `Integrations` in the side bar --> `+ Add Integration` at top the top of the page --> Enter a name --> `Add` Button   API Keys   * Click `Integrations` in the side bar --> `Create API Key` Button next to a given Integration Scope -->   Enter a name --> click `Create and View button`   * Copy/Save the key manually or click save to clipboard button to copy it (this is the only time it will show)  ## TCA Flow    *  Register a webhook   *  Create a submission   *  Upload a file for the submission   *  Wait for the submission upload to process      * If you registered a webhook, a callback will be sent to it when upload is complete      * The status of the *submission* will also update to `COMPLETE`   *  Request a Similarity Report   *  Wait for similarity report to process      * If you registered a webhook, a callback will be sent to it when report is complete      * The status of the *report* will also be updated to `COMPLETE`   *  Request a URL with parameters to view the Similarity Report 

API version: 1.0.249
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SimilaritySettings struct for SimilaritySettings
type SimilaritySettings struct {
	// default similarity mode when viewing a report; set to either match_overview or all_sources
	DefaultMode *string `json:"default_mode,omitempty"`
	Modes *SimilaritySettingsModes `json:"modes,omitempty"`
	ViewSettings *SimilaritySettingsViewSettings `json:"view_settings,omitempty"`
}

// NewSimilaritySettings instantiates a new SimilaritySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimilaritySettings() *SimilaritySettings {
	this := SimilaritySettings{}
	return &this
}

// NewSimilaritySettingsWithDefaults instantiates a new SimilaritySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimilaritySettingsWithDefaults() *SimilaritySettings {
	this := SimilaritySettings{}
	return &this
}

// GetDefaultMode returns the DefaultMode field value if set, zero value otherwise.
func (o *SimilaritySettings) GetDefaultMode() string {
	if o == nil || o.DefaultMode == nil {
		var ret string
		return ret
	}
	return *o.DefaultMode
}

// GetDefaultModeOk returns a tuple with the DefaultMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilaritySettings) GetDefaultModeOk() (*string, bool) {
	if o == nil || o.DefaultMode == nil {
		return nil, false
	}
	return o.DefaultMode, true
}

// HasDefaultMode returns a boolean if a field has been set.
func (o *SimilaritySettings) HasDefaultMode() bool {
	if o != nil && o.DefaultMode != nil {
		return true
	}

	return false
}

// SetDefaultMode gets a reference to the given string and assigns it to the DefaultMode field.
func (o *SimilaritySettings) SetDefaultMode(v string) {
	o.DefaultMode = &v
}

// GetModes returns the Modes field value if set, zero value otherwise.
func (o *SimilaritySettings) GetModes() SimilaritySettingsModes {
	if o == nil || o.Modes == nil {
		var ret SimilaritySettingsModes
		return ret
	}
	return *o.Modes
}

// GetModesOk returns a tuple with the Modes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilaritySettings) GetModesOk() (*SimilaritySettingsModes, bool) {
	if o == nil || o.Modes == nil {
		return nil, false
	}
	return o.Modes, true
}

// HasModes returns a boolean if a field has been set.
func (o *SimilaritySettings) HasModes() bool {
	if o != nil && o.Modes != nil {
		return true
	}

	return false
}

// SetModes gets a reference to the given SimilaritySettingsModes and assigns it to the Modes field.
func (o *SimilaritySettings) SetModes(v SimilaritySettingsModes) {
	o.Modes = &v
}

// GetViewSettings returns the ViewSettings field value if set, zero value otherwise.
func (o *SimilaritySettings) GetViewSettings() SimilaritySettingsViewSettings {
	if o == nil || o.ViewSettings == nil {
		var ret SimilaritySettingsViewSettings
		return ret
	}
	return *o.ViewSettings
}

// GetViewSettingsOk returns a tuple with the ViewSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilaritySettings) GetViewSettingsOk() (*SimilaritySettingsViewSettings, bool) {
	if o == nil || o.ViewSettings == nil {
		return nil, false
	}
	return o.ViewSettings, true
}

// HasViewSettings returns a boolean if a field has been set.
func (o *SimilaritySettings) HasViewSettings() bool {
	if o != nil && o.ViewSettings != nil {
		return true
	}

	return false
}

// SetViewSettings gets a reference to the given SimilaritySettingsViewSettings and assigns it to the ViewSettings field.
func (o *SimilaritySettings) SetViewSettings(v SimilaritySettingsViewSettings) {
	o.ViewSettings = &v
}

func (o SimilaritySettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultMode != nil {
		toSerialize["default_mode"] = o.DefaultMode
	}
	if o.Modes != nil {
		toSerialize["modes"] = o.Modes
	}
	if o.ViewSettings != nil {
		toSerialize["view_settings"] = o.ViewSettings
	}
	return json.Marshal(toSerialize)
}

type NullableSimilaritySettings struct {
	value *SimilaritySettings
	isSet bool
}

func (v NullableSimilaritySettings) Get() *SimilaritySettings {
	return v.value
}

func (v *NullableSimilaritySettings) Set(val *SimilaritySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSimilaritySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSimilaritySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimilaritySettings(val *SimilaritySettings) *NullableSimilaritySettings {
	return &NullableSimilaritySettings{value: val, isSet: true}
}

func (v NullableSimilaritySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimilaritySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


