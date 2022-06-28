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

// AuthorMetadataOverride struct for AuthorMetadataOverride
type AuthorMetadataOverride struct {
	// Given or first name of submission author
	GivenName *string `json:"given_name,omitempty"`
	// Family or last name of submission author
	FamilyName *string `json:"family_name,omitempty"`
}

// NewAuthorMetadataOverride instantiates a new AuthorMetadataOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorMetadataOverride() *AuthorMetadataOverride {
	this := AuthorMetadataOverride{}
	return &this
}

// NewAuthorMetadataOverrideWithDefaults instantiates a new AuthorMetadataOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorMetadataOverrideWithDefaults() *AuthorMetadataOverride {
	this := AuthorMetadataOverride{}
	return &this
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *AuthorMetadataOverride) GetGivenName() string {
	if o == nil || o.GivenName == nil {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorMetadataOverride) GetGivenNameOk() (*string, bool) {
	if o == nil || o.GivenName == nil {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *AuthorMetadataOverride) HasGivenName() bool {
	if o != nil && o.GivenName != nil {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *AuthorMetadataOverride) SetGivenName(v string) {
	o.GivenName = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *AuthorMetadataOverride) GetFamilyName() string {
	if o == nil || o.FamilyName == nil {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorMetadataOverride) GetFamilyNameOk() (*string, bool) {
	if o == nil || o.FamilyName == nil {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *AuthorMetadataOverride) HasFamilyName() bool {
	if o != nil && o.FamilyName != nil {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *AuthorMetadataOverride) SetFamilyName(v string) {
	o.FamilyName = &v
}

func (o AuthorMetadataOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GivenName != nil {
		toSerialize["given_name"] = o.GivenName
	}
	if o.FamilyName != nil {
		toSerialize["family_name"] = o.FamilyName
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorMetadataOverride struct {
	value *AuthorMetadataOverride
	isSet bool
}

func (v NullableAuthorMetadataOverride) Get() *AuthorMetadataOverride {
	return v.value
}

func (v *NullableAuthorMetadataOverride) Set(val *AuthorMetadataOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorMetadataOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorMetadataOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorMetadataOverride(val *AuthorMetadataOverride) *NullableAuthorMetadataOverride {
	return &NullableAuthorMetadataOverride{value: val, isSet: true}
}

func (v NullableAuthorMetadataOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorMetadataOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

