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

// SimilarityViewSettings struct for SimilarityViewSettings
type SimilarityViewSettings struct {
	// If set to true, text in quotes will not count as similar content 
	ExcludeQuotes *bool `json:"exclude_quotes,omitempty"`
	// If set to true, text in a bibliography section will not count as similar content 
	ExcludeBibliography *bool `json:"exclude_bibliography,omitempty"`
	// If set to true, text in citations will not count as similar content 
	ExcludeCitations *bool `json:"exclude_citations,omitempty"`
	// If set to true, text in the abstract section of the submission will not count as similar content 
	ExcludeAbstract *bool `json:"exclude_abstract,omitempty"`
	// If set to true, text in the method section of the submission will not count as similar content 
	ExcludeMethods *bool `json:"exclude_methods,omitempty"`
	// If set, similarity matches that match less than the specified amount of words will not count as similar content 
	ExcludeSmallMatches *int32 `json:"exclude_small_matches,omitempty"`
	// If set to true, text matched to the internet collection will not count as similar content 
	ExcludeInternet *bool `json:"exclude_internet,omitempty"`
	// If set to true, text matched to the publications collection will not count as similar content 
	ExcludePublications *bool `json:"exclude_publications,omitempty"`
	// If set to true, text matched to the Crossref collection will not count as similar content 
	ExcludeCrossref *bool `json:"exclude_crossref,omitempty"`
	// If set to true, text matched to the Crossref Posted Content collection will not count as similar content 
	ExcludeCrossrefPostedContent *bool `json:"exclude_crossref_posted_content,omitempty"`
	// If set to true, text matched to the submitted works collection will not count as similar content calculated as if submitted work was not part of the paper 
	ExcludeSubmittedWorks *bool `json:"exclude_submitted_works,omitempty"`
	// If set to true, text matched to the custom sections defined in the admin settings will not count as similar content calculated as if section was not part of the paper 
	ExcludeCustomSections *bool `json:"exclude_custom_sections,omitempty"`
	// If set to true, it will exclude preprints. A preprint is a version of a scholarly or scientific paper that precedes formal peer review and publication in a peer-reviewed scholarly or scientific journal. 
	ExcludePreprints *bool `json:"exclude_preprints,omitempty"`
}

// NewSimilarityViewSettings instantiates a new SimilarityViewSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimilarityViewSettings() *SimilarityViewSettings {
	this := SimilarityViewSettings{}
	return &this
}

// NewSimilarityViewSettingsWithDefaults instantiates a new SimilarityViewSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimilarityViewSettingsWithDefaults() *SimilarityViewSettings {
	this := SimilarityViewSettings{}
	return &this
}

// GetExcludeQuotes returns the ExcludeQuotes field value if set, zero value otherwise.
func (o *SimilarityViewSettings) GetExcludeQuotes() bool {
	if o == nil || o.ExcludeQuotes == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeQuotes
}

// GetExcludeQuotesOk returns a tuple with the ExcludeQuotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityViewSettings) GetExcludeQuotesOk() (*bool, bool) {
	if o == nil || o.ExcludeQuotes == nil {
		return nil, false
	}
	return o.ExcludeQuotes, true
}

// HasExcludeQuotes returns a boolean if a field has been set.
func (o *SimilarityViewSettings) HasExcludeQuotes() bool {
	if o != nil && o.ExcludeQuotes != nil {
		return true
	}

	return false
}

// SetExcludeQuotes gets a reference to the given bool and assigns it to the ExcludeQuotes field.
func (o *SimilarityViewSettings) SetExcludeQuotes(v bool) {
	o.ExcludeQuotes = &v
}

// GetExcludeBibliography returns the ExcludeBibliography field value if set, zero value otherwise.
func (o *SimilarityViewSettings) GetExcludeBibliography() bool {
	if o == nil || o.ExcludeBibliography == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeBibliography
}

// GetExcludeBibliographyOk returns a tuple with the ExcludeBibliography field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityViewSettings) GetExcludeBibliographyOk() (*bool, bool) {
	if o == nil || o.ExcludeBibliography == nil {
		return nil, false
	}
	return o.ExcludeBibliography, true
}

// HasExcludeBibliography returns a boolean if a field has been set.
func (o *SimilarityViewSettings) HasExcludeBibliography() bool {
	if o != nil && o.ExcludeBibliography != nil {
		return true
	}

	return false
}

// SetExcludeBibliography gets a reference to the given bool and assigns it to the ExcludeBibliography field.
func (o *SimilarityViewSettings) SetExcludeBibliography(v bool) {
	o.ExcludeBibliography = &v
}

// GetExcludeCitations returns the ExcludeCitations field value if set, zero value otherwise.
func (o *SimilarityViewSettings) GetExcludeCitations() bool {
	if o == nil || o.ExcludeCitations == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeCitations
}

// GetExcludeCitationsOk returns a tuple with the ExcludeCitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityViewSettings) GetExcludeCitationsOk() (*bool, bool) {
	if o == nil || o.ExcludeCitations == nil {
		return nil, false
	}
	return o.ExcludeCitations, true
}

// HasExcludeCitations returns a boolean if a field has been set.
func (o *SimilarityViewSettings) HasExcludeCitations() bool {
	if o != nil && o.ExcludeCitations != nil {
		return true
	}

	return false
}

// SetExcludeCitations gets a reference to the given bool and assigns it to the ExcludeCitations field.
func (o *SimilarityViewSettings) SetExcludeCitations(v bool) {
	o.ExcludeCitations = &v
}

// GetExcludeAbstract returns the ExcludeAbstract field value if set, zero value otherwise.
func (o *SimilarityViewSettings) GetExcludeAbstract() bool {
	if o == nil || o.ExcludeAbstract == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeAbstract
}

// GetExcludeAbstractOk returns a tuple with the ExcludeAbstract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityViewSettings) GetExcludeAbstractOk() (*bool, bool) {
	if o == nil || o.ExcludeAbstract == nil {
		return nil, false
	}
	return o.ExcludeAbstract, true
}

// HasExcludeAbstract returns a boolean if a field has been set.
func (o *SimilarityViewSettings) HasExcludeAbstract() bool {
	if o != nil && o.ExcludeAbstract != nil {
		return true
	}

	return false
}

// SetExcludeAbstract gets a reference to the given bool and assigns it to the ExcludeAbstract field.
func (o *SimilarityViewSettings) SetExcludeAbstract(v bool) {
	o.ExcludeAbstract = &v
}

// GetExcludeMethods returns the ExcludeMethods field value if set, zero value otherwise.
func (o *SimilarityViewSettings) GetExcludeMethods() bool {
	if o == nil || o.ExcludeMethods == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeMethods
}

// GetExcludeMethodsOk returns a tuple with the ExcludeMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityViewSettings) GetExcludeMethodsOk() (*bool, bool) {
	if o == nil || o.ExcludeMethods == nil {
		return nil, false
	}
	return o.ExcludeMethods, true
}

// HasExcludeMethods returns a boolean if a field has been set.
func (o *SimilarityViewSettings) HasExcludeMethods() bool {
	if o != nil && o.ExcludeMethods != nil {
		return true
	}

	return false
}

// SetExcludeMethods gets a reference to the given bool and assigns it to the ExcludeMethods field.
func (o *SimilarityViewSettings) SetExcludeMethods(v bool) {
	o.ExcludeMethods = &v
}

// GetExcludeSmallMatches returns the ExcludeSmallMatches field value if set, zero value otherwise.
func (o *SimilarityViewSettings) GetExcludeSmallMatches() int32 {
	if o == nil || o.ExcludeSmallMatches == nil {
		var ret int32
		return ret
	}
	return *o.ExcludeSmallMatches
}

// GetExcludeSmallMatchesOk returns a tuple with the ExcludeSmallMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityViewSettings) GetExcludeSmallMatchesOk() (*int32, bool) {
	if o == nil || o.ExcludeSmallMatches == nil {
		return nil, false
	}
	return o.ExcludeSmallMatches, true
}

// HasExcludeSmallMatches returns a boolean if a field has been set.
func (o *SimilarityViewSettings) HasExcludeSmallMatches() bool {
	if o != nil && o.ExcludeSmallMatches != nil {
		return true
	}

	return false
}

// SetExcludeSmallMatches gets a reference to the given int32 and assigns it to the ExcludeSmallMatches field.
func (o *SimilarityViewSettings) SetExcludeSmallMatches(v int32) {
	o.ExcludeSmallMatches = &v
}

// GetExcludeInternet returns the ExcludeInternet field value if set, zero value otherwise.
func (o *SimilarityViewSettings) GetExcludeInternet() bool {
	if o == nil || o.ExcludeInternet == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeInternet
}

// GetExcludeInternetOk returns a tuple with the ExcludeInternet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityViewSettings) GetExcludeInternetOk() (*bool, bool) {
	if o == nil || o.ExcludeInternet == nil {
		return nil, false
	}
	return o.ExcludeInternet, true
}

// HasExcludeInternet returns a boolean if a field has been set.
func (o *SimilarityViewSettings) HasExcludeInternet() bool {
	if o != nil && o.ExcludeInternet != nil {
		return true
	}

	return false
}

// SetExcludeInternet gets a reference to the given bool and assigns it to the ExcludeInternet field.
func (o *SimilarityViewSettings) SetExcludeInternet(v bool) {
	o.ExcludeInternet = &v
}

// GetExcludePublications returns the ExcludePublications field value if set, zero value otherwise.
func (o *SimilarityViewSettings) GetExcludePublications() bool {
	if o == nil || o.ExcludePublications == nil {
		var ret bool
		return ret
	}
	return *o.ExcludePublications
}

// GetExcludePublicationsOk returns a tuple with the ExcludePublications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityViewSettings) GetExcludePublicationsOk() (*bool, bool) {
	if o == nil || o.ExcludePublications == nil {
		return nil, false
	}
	return o.ExcludePublications, true
}

// HasExcludePublications returns a boolean if a field has been set.
func (o *SimilarityViewSettings) HasExcludePublications() bool {
	if o != nil && o.ExcludePublications != nil {
		return true
	}

	return false
}

// SetExcludePublications gets a reference to the given bool and assigns it to the ExcludePublications field.
func (o *SimilarityViewSettings) SetExcludePublications(v bool) {
	o.ExcludePublications = &v
}

// GetExcludeCrossref returns the ExcludeCrossref field value if set, zero value otherwise.
func (o *SimilarityViewSettings) GetExcludeCrossref() bool {
	if o == nil || o.ExcludeCrossref == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeCrossref
}

// GetExcludeCrossrefOk returns a tuple with the ExcludeCrossref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityViewSettings) GetExcludeCrossrefOk() (*bool, bool) {
	if o == nil || o.ExcludeCrossref == nil {
		return nil, false
	}
	return o.ExcludeCrossref, true
}

// HasExcludeCrossref returns a boolean if a field has been set.
func (o *SimilarityViewSettings) HasExcludeCrossref() bool {
	if o != nil && o.ExcludeCrossref != nil {
		return true
	}

	return false
}

// SetExcludeCrossref gets a reference to the given bool and assigns it to the ExcludeCrossref field.
func (o *SimilarityViewSettings) SetExcludeCrossref(v bool) {
	o.ExcludeCrossref = &v
}

// GetExcludeCrossrefPostedContent returns the ExcludeCrossrefPostedContent field value if set, zero value otherwise.
func (o *SimilarityViewSettings) GetExcludeCrossrefPostedContent() bool {
	if o == nil || o.ExcludeCrossrefPostedContent == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeCrossrefPostedContent
}

// GetExcludeCrossrefPostedContentOk returns a tuple with the ExcludeCrossrefPostedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityViewSettings) GetExcludeCrossrefPostedContentOk() (*bool, bool) {
	if o == nil || o.ExcludeCrossrefPostedContent == nil {
		return nil, false
	}
	return o.ExcludeCrossrefPostedContent, true
}

// HasExcludeCrossrefPostedContent returns a boolean if a field has been set.
func (o *SimilarityViewSettings) HasExcludeCrossrefPostedContent() bool {
	if o != nil && o.ExcludeCrossrefPostedContent != nil {
		return true
	}

	return false
}

// SetExcludeCrossrefPostedContent gets a reference to the given bool and assigns it to the ExcludeCrossrefPostedContent field.
func (o *SimilarityViewSettings) SetExcludeCrossrefPostedContent(v bool) {
	o.ExcludeCrossrefPostedContent = &v
}

// GetExcludeSubmittedWorks returns the ExcludeSubmittedWorks field value if set, zero value otherwise.
func (o *SimilarityViewSettings) GetExcludeSubmittedWorks() bool {
	if o == nil || o.ExcludeSubmittedWorks == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeSubmittedWorks
}

// GetExcludeSubmittedWorksOk returns a tuple with the ExcludeSubmittedWorks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityViewSettings) GetExcludeSubmittedWorksOk() (*bool, bool) {
	if o == nil || o.ExcludeSubmittedWorks == nil {
		return nil, false
	}
	return o.ExcludeSubmittedWorks, true
}

// HasExcludeSubmittedWorks returns a boolean if a field has been set.
func (o *SimilarityViewSettings) HasExcludeSubmittedWorks() bool {
	if o != nil && o.ExcludeSubmittedWorks != nil {
		return true
	}

	return false
}

// SetExcludeSubmittedWorks gets a reference to the given bool and assigns it to the ExcludeSubmittedWorks field.
func (o *SimilarityViewSettings) SetExcludeSubmittedWorks(v bool) {
	o.ExcludeSubmittedWorks = &v
}

// GetExcludeCustomSections returns the ExcludeCustomSections field value if set, zero value otherwise.
func (o *SimilarityViewSettings) GetExcludeCustomSections() bool {
	if o == nil || o.ExcludeCustomSections == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeCustomSections
}

// GetExcludeCustomSectionsOk returns a tuple with the ExcludeCustomSections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityViewSettings) GetExcludeCustomSectionsOk() (*bool, bool) {
	if o == nil || o.ExcludeCustomSections == nil {
		return nil, false
	}
	return o.ExcludeCustomSections, true
}

// HasExcludeCustomSections returns a boolean if a field has been set.
func (o *SimilarityViewSettings) HasExcludeCustomSections() bool {
	if o != nil && o.ExcludeCustomSections != nil {
		return true
	}

	return false
}

// SetExcludeCustomSections gets a reference to the given bool and assigns it to the ExcludeCustomSections field.
func (o *SimilarityViewSettings) SetExcludeCustomSections(v bool) {
	o.ExcludeCustomSections = &v
}

// GetExcludePreprints returns the ExcludePreprints field value if set, zero value otherwise.
func (o *SimilarityViewSettings) GetExcludePreprints() bool {
	if o == nil || o.ExcludePreprints == nil {
		var ret bool
		return ret
	}
	return *o.ExcludePreprints
}

// GetExcludePreprintsOk returns a tuple with the ExcludePreprints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimilarityViewSettings) GetExcludePreprintsOk() (*bool, bool) {
	if o == nil || o.ExcludePreprints == nil {
		return nil, false
	}
	return o.ExcludePreprints, true
}

// HasExcludePreprints returns a boolean if a field has been set.
func (o *SimilarityViewSettings) HasExcludePreprints() bool {
	if o != nil && o.ExcludePreprints != nil {
		return true
	}

	return false
}

// SetExcludePreprints gets a reference to the given bool and assigns it to the ExcludePreprints field.
func (o *SimilarityViewSettings) SetExcludePreprints(v bool) {
	o.ExcludePreprints = &v
}

func (o SimilarityViewSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExcludeQuotes != nil {
		toSerialize["exclude_quotes"] = o.ExcludeQuotes
	}
	if o.ExcludeBibliography != nil {
		toSerialize["exclude_bibliography"] = o.ExcludeBibliography
	}
	if o.ExcludeCitations != nil {
		toSerialize["exclude_citations"] = o.ExcludeCitations
	}
	if o.ExcludeAbstract != nil {
		toSerialize["exclude_abstract"] = o.ExcludeAbstract
	}
	if o.ExcludeMethods != nil {
		toSerialize["exclude_methods"] = o.ExcludeMethods
	}
	if o.ExcludeSmallMatches != nil {
		toSerialize["exclude_small_matches"] = o.ExcludeSmallMatches
	}
	if o.ExcludeInternet != nil {
		toSerialize["exclude_internet"] = o.ExcludeInternet
	}
	if o.ExcludePublications != nil {
		toSerialize["exclude_publications"] = o.ExcludePublications
	}
	if o.ExcludeCrossref != nil {
		toSerialize["exclude_crossref"] = o.ExcludeCrossref
	}
	if o.ExcludeCrossrefPostedContent != nil {
		toSerialize["exclude_crossref_posted_content"] = o.ExcludeCrossrefPostedContent
	}
	if o.ExcludeSubmittedWorks != nil {
		toSerialize["exclude_submitted_works"] = o.ExcludeSubmittedWorks
	}
	if o.ExcludeCustomSections != nil {
		toSerialize["exclude_custom_sections"] = o.ExcludeCustomSections
	}
	if o.ExcludePreprints != nil {
		toSerialize["exclude_preprints"] = o.ExcludePreprints
	}
	return json.Marshal(toSerialize)
}

type NullableSimilarityViewSettings struct {
	value *SimilarityViewSettings
	isSet bool
}

func (v NullableSimilarityViewSettings) Get() *SimilarityViewSettings {
	return v.value
}

func (v *NullableSimilarityViewSettings) Set(val *SimilarityViewSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSimilarityViewSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSimilarityViewSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimilarityViewSettings(val *SimilarityViewSettings) *NullableSimilarityViewSettings {
	return &NullableSimilarityViewSettings{value: val, isSet: true}
}

func (v NullableSimilarityViewSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimilarityViewSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


