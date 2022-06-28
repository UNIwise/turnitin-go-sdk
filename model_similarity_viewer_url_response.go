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

// SimilarityViewerUrlResponse struct for SimilarityViewerUrlResponse
type SimilarityViewerUrlResponse struct {
	// URL to be used to access Cloud Viewer visualization of similarity report matches
	ViewerUrl *string `json:"viewer_url,omitempty"`
}

// NewSimilarityViewerUrlResponse instantiates a new SimilarityViewerUrlResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimilarityViewerUrlResponse() *SimilarityViewerUrlResponse {
	this := SimilarityViewerUrlResponse{}
	return &this
}

// NewSimilarityViewerUrlResponseWithDefaults instantiates a new SimilarityViewerUrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimilarityViewerUrlResponseWithDefaults() *SimilarityViewerUrlResponse {
	this := SimilarityViewerUrlResponse{}
	return &this
}

// GetViewerUrl returns the ViewerUrl field value if set, zero value otherwise.
func (o *SimilarityViewerUrlResponse) GetViewerUrl() string {
	if o == nil || o.ViewerUrl == nil {
		var ret string
		return ret
	}
	return *o.ViewerUrl
}

// GetViewerUrlOk returns a tuple with the ViewerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityViewerUrlResponse) GetViewerUrlOk() (*string, bool) {
	if o == nil || o.ViewerUrl == nil {
		return nil, false
	}
	return o.ViewerUrl, true
}

// HasViewerUrl returns a boolean if a field has been set.
func (o *SimilarityViewerUrlResponse) HasViewerUrl() bool {
	if o != nil && o.ViewerUrl != nil {
		return true
	}

	return false
}

// SetViewerUrl gets a reference to the given string and assigns it to the ViewerUrl field.
func (o *SimilarityViewerUrlResponse) SetViewerUrl(v string) {
	o.ViewerUrl = &v
}

func (o SimilarityViewerUrlResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ViewerUrl != nil {
		toSerialize["viewer_url"] = o.ViewerUrl
	}
	return json.Marshal(toSerialize)
}

type NullableSimilarityViewerUrlResponse struct {
	value *SimilarityViewerUrlResponse
	isSet bool
}

func (v NullableSimilarityViewerUrlResponse) Get() *SimilarityViewerUrlResponse {
	return v.value
}

func (v *NullableSimilarityViewerUrlResponse) Set(val *SimilarityViewerUrlResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSimilarityViewerUrlResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSimilarityViewerUrlResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimilarityViewerUrlResponse(val *SimilarityViewerUrlResponse) *NullableSimilarityViewerUrlResponse {
	return &NullableSimilarityViewerUrlResponse{value: val, isSet: true}
}

func (v NullableSimilarityViewerUrlResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimilarityViewerUrlResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

