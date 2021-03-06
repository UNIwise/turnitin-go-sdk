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

// SimilarityMetadataAllOfTopMatches struct for SimilarityMetadataAllOfTopMatches
type SimilarityMetadataAllOfTopMatches struct {
	// Source name
	Name *string `json:"name,omitempty"`
	// Match percentage
	Percentage *float32 `json:"percentage,omitempty"`
	// Matching submission id
	SubmissionId *string `json:"submission_id,omitempty"`
	// Matching submission source type (INTERNET, PUBLICATION, SUBMITTED_WORK)
	SourceType *string `json:"source_type,omitempty"`
	// number of matching words
	MatchedWordCountTotal *float32 `json:"matched_word_count_total,omitempty"`
	// date match was submitted
	SubmittedDate *string `json:"submitted_date,omitempty"`
	// intitution name for matched SUBMITTED_WORK types
	InstitutionName *string `json:"institution_name,omitempty"`
}

// NewSimilarityMetadataAllOfTopMatches instantiates a new SimilarityMetadataAllOfTopMatches object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimilarityMetadataAllOfTopMatches() *SimilarityMetadataAllOfTopMatches {
	this := SimilarityMetadataAllOfTopMatches{}
	return &this
}

// NewSimilarityMetadataAllOfTopMatchesWithDefaults instantiates a new SimilarityMetadataAllOfTopMatches object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimilarityMetadataAllOfTopMatchesWithDefaults() *SimilarityMetadataAllOfTopMatches {
	this := SimilarityMetadataAllOfTopMatches{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SimilarityMetadataAllOfTopMatches) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityMetadataAllOfTopMatches) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SimilarityMetadataAllOfTopMatches) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SimilarityMetadataAllOfTopMatches) SetName(v string) {
	o.Name = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *SimilarityMetadataAllOfTopMatches) GetPercentage() float32 {
	if o == nil || o.Percentage == nil {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityMetadataAllOfTopMatches) GetPercentageOk() (*float32, bool) {
	if o == nil || o.Percentage == nil {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *SimilarityMetadataAllOfTopMatches) HasPercentage() bool {
	if o != nil && o.Percentage != nil {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *SimilarityMetadataAllOfTopMatches) SetPercentage(v float32) {
	o.Percentage = &v
}

// GetSubmissionId returns the SubmissionId field value if set, zero value otherwise.
func (o *SimilarityMetadataAllOfTopMatches) GetSubmissionId() string {
	if o == nil || o.SubmissionId == nil {
		var ret string
		return ret
	}
	return *o.SubmissionId
}

// GetSubmissionIdOk returns a tuple with the SubmissionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityMetadataAllOfTopMatches) GetSubmissionIdOk() (*string, bool) {
	if o == nil || o.SubmissionId == nil {
		return nil, false
	}
	return o.SubmissionId, true
}

// HasSubmissionId returns a boolean if a field has been set.
func (o *SimilarityMetadataAllOfTopMatches) HasSubmissionId() bool {
	if o != nil && o.SubmissionId != nil {
		return true
	}

	return false
}

// SetSubmissionId gets a reference to the given string and assigns it to the SubmissionId field.
func (o *SimilarityMetadataAllOfTopMatches) SetSubmissionId(v string) {
	o.SubmissionId = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *SimilarityMetadataAllOfTopMatches) GetSourceType() string {
	if o == nil || o.SourceType == nil {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityMetadataAllOfTopMatches) GetSourceTypeOk() (*string, bool) {
	if o == nil || o.SourceType == nil {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *SimilarityMetadataAllOfTopMatches) HasSourceType() bool {
	if o != nil && o.SourceType != nil {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *SimilarityMetadataAllOfTopMatches) SetSourceType(v string) {
	o.SourceType = &v
}

// GetMatchedWordCountTotal returns the MatchedWordCountTotal field value if set, zero value otherwise.
func (o *SimilarityMetadataAllOfTopMatches) GetMatchedWordCountTotal() float32 {
	if o == nil || o.MatchedWordCountTotal == nil {
		var ret float32
		return ret
	}
	return *o.MatchedWordCountTotal
}

// GetMatchedWordCountTotalOk returns a tuple with the MatchedWordCountTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityMetadataAllOfTopMatches) GetMatchedWordCountTotalOk() (*float32, bool) {
	if o == nil || o.MatchedWordCountTotal == nil {
		return nil, false
	}
	return o.MatchedWordCountTotal, true
}

// HasMatchedWordCountTotal returns a boolean if a field has been set.
func (o *SimilarityMetadataAllOfTopMatches) HasMatchedWordCountTotal() bool {
	if o != nil && o.MatchedWordCountTotal != nil {
		return true
	}

	return false
}

// SetMatchedWordCountTotal gets a reference to the given float32 and assigns it to the MatchedWordCountTotal field.
func (o *SimilarityMetadataAllOfTopMatches) SetMatchedWordCountTotal(v float32) {
	o.MatchedWordCountTotal = &v
}

// GetSubmittedDate returns the SubmittedDate field value if set, zero value otherwise.
func (o *SimilarityMetadataAllOfTopMatches) GetSubmittedDate() string {
	if o == nil || o.SubmittedDate == nil {
		var ret string
		return ret
	}
	return *o.SubmittedDate
}

// GetSubmittedDateOk returns a tuple with the SubmittedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityMetadataAllOfTopMatches) GetSubmittedDateOk() (*string, bool) {
	if o == nil || o.SubmittedDate == nil {
		return nil, false
	}
	return o.SubmittedDate, true
}

// HasSubmittedDate returns a boolean if a field has been set.
func (o *SimilarityMetadataAllOfTopMatches) HasSubmittedDate() bool {
	if o != nil && o.SubmittedDate != nil {
		return true
	}

	return false
}

// SetSubmittedDate gets a reference to the given string and assigns it to the SubmittedDate field.
func (o *SimilarityMetadataAllOfTopMatches) SetSubmittedDate(v string) {
	o.SubmittedDate = &v
}

// GetInstitutionName returns the InstitutionName field value if set, zero value otherwise.
func (o *SimilarityMetadataAllOfTopMatches) GetInstitutionName() string {
	if o == nil || o.InstitutionName == nil {
		var ret string
		return ret
	}
	return *o.InstitutionName
}

// GetInstitutionNameOk returns a tuple with the InstitutionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityMetadataAllOfTopMatches) GetInstitutionNameOk() (*string, bool) {
	if o == nil || o.InstitutionName == nil {
		return nil, false
	}
	return o.InstitutionName, true
}

// HasInstitutionName returns a boolean if a field has been set.
func (o *SimilarityMetadataAllOfTopMatches) HasInstitutionName() bool {
	if o != nil && o.InstitutionName != nil {
		return true
	}

	return false
}

// SetInstitutionName gets a reference to the given string and assigns it to the InstitutionName field.
func (o *SimilarityMetadataAllOfTopMatches) SetInstitutionName(v string) {
	o.InstitutionName = &v
}

func (o SimilarityMetadataAllOfTopMatches) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Percentage != nil {
		toSerialize["percentage"] = o.Percentage
	}
	if o.SubmissionId != nil {
		toSerialize["submission_id"] = o.SubmissionId
	}
	if o.SourceType != nil {
		toSerialize["source_type"] = o.SourceType
	}
	if o.MatchedWordCountTotal != nil {
		toSerialize["matched_word_count_total"] = o.MatchedWordCountTotal
	}
	if o.SubmittedDate != nil {
		toSerialize["submitted_date"] = o.SubmittedDate
	}
	if o.InstitutionName != nil {
		toSerialize["institution_name"] = o.InstitutionName
	}
	return json.Marshal(toSerialize)
}

type NullableSimilarityMetadataAllOfTopMatches struct {
	value *SimilarityMetadataAllOfTopMatches
	isSet bool
}

func (v NullableSimilarityMetadataAllOfTopMatches) Get() *SimilarityMetadataAllOfTopMatches {
	return v.value
}

func (v *NullableSimilarityMetadataAllOfTopMatches) Set(val *SimilarityMetadataAllOfTopMatches) {
	v.value = val
	v.isSet = true
}

func (v NullableSimilarityMetadataAllOfTopMatches) IsSet() bool {
	return v.isSet
}

func (v *NullableSimilarityMetadataAllOfTopMatches) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimilarityMetadataAllOfTopMatches(val *SimilarityMetadataAllOfTopMatches) *NullableSimilarityMetadataAllOfTopMatches {
	return &NullableSimilarityMetadataAllOfTopMatches{value: val, isSet: true}
}

func (v NullableSimilarityMetadataAllOfTopMatches) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimilarityMetadataAllOfTopMatches) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


