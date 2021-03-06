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

// SimilaritySettingsViewSettings struct for SimilaritySettingsViewSettings
type SimilaritySettingsViewSettings struct {
	// Used to enable save changes in the Viewer and trigger SIMILARITY_UPDATED webhook callback
	SaveChanges *bool `json:"save_changes,omitempty"`
}

// NewSimilaritySettingsViewSettings instantiates a new SimilaritySettingsViewSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimilaritySettingsViewSettings() *SimilaritySettingsViewSettings {
	this := SimilaritySettingsViewSettings{}
	return &this
}

// NewSimilaritySettingsViewSettingsWithDefaults instantiates a new SimilaritySettingsViewSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimilaritySettingsViewSettingsWithDefaults() *SimilaritySettingsViewSettings {
	this := SimilaritySettingsViewSettings{}
	return &this
}

// GetSaveChanges returns the SaveChanges field value if set, zero value otherwise.
func (o *SimilaritySettingsViewSettings) GetSaveChanges() bool {
	if o == nil || o.SaveChanges == nil {
		var ret bool
		return ret
	}
	return *o.SaveChanges
}

// GetSaveChangesOk returns a tuple with the SaveChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilaritySettingsViewSettings) GetSaveChangesOk() (*bool, bool) {
	if o == nil || o.SaveChanges == nil {
		return nil, false
	}
	return o.SaveChanges, true
}

// HasSaveChanges returns a boolean if a field has been set.
func (o *SimilaritySettingsViewSettings) HasSaveChanges() bool {
	if o != nil && o.SaveChanges != nil {
		return true
	}

	return false
}

// SetSaveChanges gets a reference to the given bool and assigns it to the SaveChanges field.
func (o *SimilaritySettingsViewSettings) SetSaveChanges(v bool) {
	o.SaveChanges = &v
}

func (o SimilaritySettingsViewSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SaveChanges != nil {
		toSerialize["save_changes"] = o.SaveChanges
	}
	return json.Marshal(toSerialize)
}

type NullableSimilaritySettingsViewSettings struct {
	value *SimilaritySettingsViewSettings
	isSet bool
}

func (v NullableSimilaritySettingsViewSettings) Get() *SimilaritySettingsViewSettings {
	return v.value
}

func (v *NullableSimilaritySettingsViewSettings) Set(val *SimilaritySettingsViewSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSimilaritySettingsViewSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSimilaritySettingsViewSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimilaritySettingsViewSettings(val *SimilaritySettingsViewSettings) *NullableSimilaritySettingsViewSettings {
	return &NullableSimilaritySettingsViewSettings{value: val, isSet: true}
}

func (v NullableSimilaritySettingsViewSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimilaritySettingsViewSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


