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

// SimilarityCompleteWebhookRequest struct for SimilarityCompleteWebhookRequest
type SimilarityCompleteWebhookRequest struct {
	// Represents the percentage match against all sources
	OverallMatchPercentage int32 `json:"overall_match_percentage"`
	// Represents the percentage match against internet
	InternetMatchPercentage NullableInt32 `json:"internet_match_percentage,omitempty"`
	// Represents the percentage match against all publications
	PublicationMatchPercentage NullableInt32 `json:"publication_match_percentage,omitempty"`
	// Represents the percentage match against all submitted works
	SubmittedWorksMatchPercentage NullableInt32 `json:"submitted_works_match_percentage,omitempty"`
	SubmissionId string `json:"submission_id"`
	// possible values PENDING, COMPLETE
	Status string `json:"status"`
	// Time the report finished generating.  If not set the report has not finished generating
	TimeGenerated string `json:"time_generated"`
	// Time the report was requested
	TimeRequested string `json:"time_requested"`
	// Top matches
	TopMatches []SimilarityMetadataAllOfTopMatches `json:"top_matches"`
	// Largest individual matched word count, 0 if there isn't a match to this submission.
	TopSourceLargestMatchedWordCount int32 `json:"top_source_largest_matched_word_count"`
	Metadata *SubmissionCompleteWebhookRequestAllOfMetadata `json:"metadata,omitempty"`
}

// NewSimilarityCompleteWebhookRequest instantiates a new SimilarityCompleteWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimilarityCompleteWebhookRequest(overallMatchPercentage int32, submissionId string, status string, timeGenerated string, timeRequested string, topMatches []SimilarityMetadataAllOfTopMatches, topSourceLargestMatchedWordCount int32) *SimilarityCompleteWebhookRequest {
	this := SimilarityCompleteWebhookRequest{}
	this.OverallMatchPercentage = overallMatchPercentage
	this.SubmissionId = submissionId
	this.Status = status
	this.TimeGenerated = timeGenerated
	this.TimeRequested = timeRequested
	this.TopMatches = topMatches
	this.TopSourceLargestMatchedWordCount = topSourceLargestMatchedWordCount
	return &this
}

// NewSimilarityCompleteWebhookRequestWithDefaults instantiates a new SimilarityCompleteWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimilarityCompleteWebhookRequestWithDefaults() *SimilarityCompleteWebhookRequest {
	this := SimilarityCompleteWebhookRequest{}
	return &this
}

// GetOverallMatchPercentage returns the OverallMatchPercentage field value
func (o *SimilarityCompleteWebhookRequest) GetOverallMatchPercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OverallMatchPercentage
}

// GetOverallMatchPercentageOk returns a tuple with the OverallMatchPercentage field value
// and a boolean to check if the value has been set.
func (o *SimilarityCompleteWebhookRequest) GetOverallMatchPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverallMatchPercentage, true
}

// SetOverallMatchPercentage sets field value
func (o *SimilarityCompleteWebhookRequest) SetOverallMatchPercentage(v int32) {
	o.OverallMatchPercentage = v
}

// GetInternetMatchPercentage returns the InternetMatchPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimilarityCompleteWebhookRequest) GetInternetMatchPercentage() int32 {
	if o == nil || o.InternetMatchPercentage.Get() == nil {
		var ret int32
		return ret
	}
	return *o.InternetMatchPercentage.Get()
}

// GetInternetMatchPercentageOk returns a tuple with the InternetMatchPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimilarityCompleteWebhookRequest) GetInternetMatchPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InternetMatchPercentage.Get(), o.InternetMatchPercentage.IsSet()
}

// HasInternetMatchPercentage returns a boolean if a field has been set.
func (o *SimilarityCompleteWebhookRequest) HasInternetMatchPercentage() bool {
	if o != nil && o.InternetMatchPercentage.IsSet() {
		return true
	}

	return false
}

// SetInternetMatchPercentage gets a reference to the given NullableInt32 and assigns it to the InternetMatchPercentage field.
func (o *SimilarityCompleteWebhookRequest) SetInternetMatchPercentage(v int32) {
	o.InternetMatchPercentage.Set(&v)
}
// SetInternetMatchPercentageNil sets the value for InternetMatchPercentage to be an explicit nil
func (o *SimilarityCompleteWebhookRequest) SetInternetMatchPercentageNil() {
	o.InternetMatchPercentage.Set(nil)
}

// UnsetInternetMatchPercentage ensures that no value is present for InternetMatchPercentage, not even an explicit nil
func (o *SimilarityCompleteWebhookRequest) UnsetInternetMatchPercentage() {
	o.InternetMatchPercentage.Unset()
}

// GetPublicationMatchPercentage returns the PublicationMatchPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimilarityCompleteWebhookRequest) GetPublicationMatchPercentage() int32 {
	if o == nil || o.PublicationMatchPercentage.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PublicationMatchPercentage.Get()
}

// GetPublicationMatchPercentageOk returns a tuple with the PublicationMatchPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimilarityCompleteWebhookRequest) GetPublicationMatchPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicationMatchPercentage.Get(), o.PublicationMatchPercentage.IsSet()
}

// HasPublicationMatchPercentage returns a boolean if a field has been set.
func (o *SimilarityCompleteWebhookRequest) HasPublicationMatchPercentage() bool {
	if o != nil && o.PublicationMatchPercentage.IsSet() {
		return true
	}

	return false
}

// SetPublicationMatchPercentage gets a reference to the given NullableInt32 and assigns it to the PublicationMatchPercentage field.
func (o *SimilarityCompleteWebhookRequest) SetPublicationMatchPercentage(v int32) {
	o.PublicationMatchPercentage.Set(&v)
}
// SetPublicationMatchPercentageNil sets the value for PublicationMatchPercentage to be an explicit nil
func (o *SimilarityCompleteWebhookRequest) SetPublicationMatchPercentageNil() {
	o.PublicationMatchPercentage.Set(nil)
}

// UnsetPublicationMatchPercentage ensures that no value is present for PublicationMatchPercentage, not even an explicit nil
func (o *SimilarityCompleteWebhookRequest) UnsetPublicationMatchPercentage() {
	o.PublicationMatchPercentage.Unset()
}

// GetSubmittedWorksMatchPercentage returns the SubmittedWorksMatchPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimilarityCompleteWebhookRequest) GetSubmittedWorksMatchPercentage() int32 {
	if o == nil || o.SubmittedWorksMatchPercentage.Get() == nil {
		var ret int32
		return ret
	}
	return *o.SubmittedWorksMatchPercentage.Get()
}

// GetSubmittedWorksMatchPercentageOk returns a tuple with the SubmittedWorksMatchPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimilarityCompleteWebhookRequest) GetSubmittedWorksMatchPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubmittedWorksMatchPercentage.Get(), o.SubmittedWorksMatchPercentage.IsSet()
}

// HasSubmittedWorksMatchPercentage returns a boolean if a field has been set.
func (o *SimilarityCompleteWebhookRequest) HasSubmittedWorksMatchPercentage() bool {
	if o != nil && o.SubmittedWorksMatchPercentage.IsSet() {
		return true
	}

	return false
}

// SetSubmittedWorksMatchPercentage gets a reference to the given NullableInt32 and assigns it to the SubmittedWorksMatchPercentage field.
func (o *SimilarityCompleteWebhookRequest) SetSubmittedWorksMatchPercentage(v int32) {
	o.SubmittedWorksMatchPercentage.Set(&v)
}
// SetSubmittedWorksMatchPercentageNil sets the value for SubmittedWorksMatchPercentage to be an explicit nil
func (o *SimilarityCompleteWebhookRequest) SetSubmittedWorksMatchPercentageNil() {
	o.SubmittedWorksMatchPercentage.Set(nil)
}

// UnsetSubmittedWorksMatchPercentage ensures that no value is present for SubmittedWorksMatchPercentage, not even an explicit nil
func (o *SimilarityCompleteWebhookRequest) UnsetSubmittedWorksMatchPercentage() {
	o.SubmittedWorksMatchPercentage.Unset()
}

// GetSubmissionId returns the SubmissionId field value
func (o *SimilarityCompleteWebhookRequest) GetSubmissionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubmissionId
}

// GetSubmissionIdOk returns a tuple with the SubmissionId field value
// and a boolean to check if the value has been set.
func (o *SimilarityCompleteWebhookRequest) GetSubmissionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubmissionId, true
}

// SetSubmissionId sets field value
func (o *SimilarityCompleteWebhookRequest) SetSubmissionId(v string) {
	o.SubmissionId = v
}

// GetStatus returns the Status field value
func (o *SimilarityCompleteWebhookRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SimilarityCompleteWebhookRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SimilarityCompleteWebhookRequest) SetStatus(v string) {
	o.Status = v
}

// GetTimeGenerated returns the TimeGenerated field value
func (o *SimilarityCompleteWebhookRequest) GetTimeGenerated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeGenerated
}

// GetTimeGeneratedOk returns a tuple with the TimeGenerated field value
// and a boolean to check if the value has been set.
func (o *SimilarityCompleteWebhookRequest) GetTimeGeneratedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeGenerated, true
}

// SetTimeGenerated sets field value
func (o *SimilarityCompleteWebhookRequest) SetTimeGenerated(v string) {
	o.TimeGenerated = v
}

// GetTimeRequested returns the TimeRequested field value
func (o *SimilarityCompleteWebhookRequest) GetTimeRequested() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeRequested
}

// GetTimeRequestedOk returns a tuple with the TimeRequested field value
// and a boolean to check if the value has been set.
func (o *SimilarityCompleteWebhookRequest) GetTimeRequestedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRequested, true
}

// SetTimeRequested sets field value
func (o *SimilarityCompleteWebhookRequest) SetTimeRequested(v string) {
	o.TimeRequested = v
}

// GetTopMatches returns the TopMatches field value
func (o *SimilarityCompleteWebhookRequest) GetTopMatches() []SimilarityMetadataAllOfTopMatches {
	if o == nil {
		var ret []SimilarityMetadataAllOfTopMatches
		return ret
	}

	return o.TopMatches
}

// GetTopMatchesOk returns a tuple with the TopMatches field value
// and a boolean to check if the value has been set.
func (o *SimilarityCompleteWebhookRequest) GetTopMatchesOk() ([]SimilarityMetadataAllOfTopMatches, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopMatches, true
}

// SetTopMatches sets field value
func (o *SimilarityCompleteWebhookRequest) SetTopMatches(v []SimilarityMetadataAllOfTopMatches) {
	o.TopMatches = v
}

// GetTopSourceLargestMatchedWordCount returns the TopSourceLargestMatchedWordCount field value
func (o *SimilarityCompleteWebhookRequest) GetTopSourceLargestMatchedWordCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TopSourceLargestMatchedWordCount
}

// GetTopSourceLargestMatchedWordCountOk returns a tuple with the TopSourceLargestMatchedWordCount field value
// and a boolean to check if the value has been set.
func (o *SimilarityCompleteWebhookRequest) GetTopSourceLargestMatchedWordCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopSourceLargestMatchedWordCount, true
}

// SetTopSourceLargestMatchedWordCount sets field value
func (o *SimilarityCompleteWebhookRequest) SetTopSourceLargestMatchedWordCount(v int32) {
	o.TopSourceLargestMatchedWordCount = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SimilarityCompleteWebhookRequest) GetMetadata() SubmissionCompleteWebhookRequestAllOfMetadata {
	if o == nil || o.Metadata == nil {
		var ret SubmissionCompleteWebhookRequestAllOfMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityCompleteWebhookRequest) GetMetadataOk() (*SubmissionCompleteWebhookRequestAllOfMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SimilarityCompleteWebhookRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given SubmissionCompleteWebhookRequestAllOfMetadata and assigns it to the Metadata field.
func (o *SimilarityCompleteWebhookRequest) SetMetadata(v SubmissionCompleteWebhookRequestAllOfMetadata) {
	o.Metadata = &v
}

func (o SimilarityCompleteWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["overall_match_percentage"] = o.OverallMatchPercentage
	}
	if o.InternetMatchPercentage.IsSet() {
		toSerialize["internet_match_percentage"] = o.InternetMatchPercentage.Get()
	}
	if o.PublicationMatchPercentage.IsSet() {
		toSerialize["publication_match_percentage"] = o.PublicationMatchPercentage.Get()
	}
	if o.SubmittedWorksMatchPercentage.IsSet() {
		toSerialize["submitted_works_match_percentage"] = o.SubmittedWorksMatchPercentage.Get()
	}
	if true {
		toSerialize["submission_id"] = o.SubmissionId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["time_generated"] = o.TimeGenerated
	}
	if true {
		toSerialize["time_requested"] = o.TimeRequested
	}
	if true {
		toSerialize["top_matches"] = o.TopMatches
	}
	if true {
		toSerialize["top_source_largest_matched_word_count"] = o.TopSourceLargestMatchedWordCount
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableSimilarityCompleteWebhookRequest struct {
	value *SimilarityCompleteWebhookRequest
	isSet bool
}

func (v NullableSimilarityCompleteWebhookRequest) Get() *SimilarityCompleteWebhookRequest {
	return v.value
}

func (v *NullableSimilarityCompleteWebhookRequest) Set(val *SimilarityCompleteWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSimilarityCompleteWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSimilarityCompleteWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimilarityCompleteWebhookRequest(val *SimilarityCompleteWebhookRequest) *NullableSimilarityCompleteWebhookRequest {
	return &NullableSimilarityCompleteWebhookRequest{value: val, isSet: true}
}

func (v NullableSimilarityCompleteWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimilarityCompleteWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


