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

// FeaturesEnabled struct for FeaturesEnabled
type FeaturesEnabled struct {
	Similarity *FeaturesSimilarity `json:"similarity,omitempty"`
	Tenant *FeaturesTenant `json:"tenant,omitempty"`
	ProductName *string `json:"product_name,omitempty"`
	AccessOptions []string `json:"access_options,omitempty"`
}

// NewFeaturesEnabled instantiates a new FeaturesEnabled object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturesEnabled() *FeaturesEnabled {
	this := FeaturesEnabled{}
	return &this
}

// NewFeaturesEnabledWithDefaults instantiates a new FeaturesEnabled object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturesEnabledWithDefaults() *FeaturesEnabled {
	this := FeaturesEnabled{}
	return &this
}

// GetSimilarity returns the Similarity field value if set, zero value otherwise.
func (o *FeaturesEnabled) GetSimilarity() FeaturesSimilarity {
	if o == nil || o.Similarity == nil {
		var ret FeaturesSimilarity
		return ret
	}
	return *o.Similarity
}

// GetSimilarityOk returns a tuple with the Similarity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturesEnabled) GetSimilarityOk() (*FeaturesSimilarity, bool) {
	if o == nil || o.Similarity == nil {
		return nil, false
	}
	return o.Similarity, true
}

// HasSimilarity returns a boolean if a field has been set.
func (o *FeaturesEnabled) HasSimilarity() bool {
	if o != nil && o.Similarity != nil {
		return true
	}

	return false
}

// SetSimilarity gets a reference to the given FeaturesSimilarity and assigns it to the Similarity field.
func (o *FeaturesEnabled) SetSimilarity(v FeaturesSimilarity) {
	o.Similarity = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *FeaturesEnabled) GetTenant() FeaturesTenant {
	if o == nil || o.Tenant == nil {
		var ret FeaturesTenant
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturesEnabled) GetTenantOk() (*FeaturesTenant, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *FeaturesEnabled) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given FeaturesTenant and assigns it to the Tenant field.
func (o *FeaturesEnabled) SetTenant(v FeaturesTenant) {
	o.Tenant = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *FeaturesEnabled) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturesEnabled) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *FeaturesEnabled) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *FeaturesEnabled) SetProductName(v string) {
	o.ProductName = &v
}

// GetAccessOptions returns the AccessOptions field value if set, zero value otherwise.
func (o *FeaturesEnabled) GetAccessOptions() []string {
	if o == nil || o.AccessOptions == nil {
		var ret []string
		return ret
	}
	return o.AccessOptions
}

// GetAccessOptionsOk returns a tuple with the AccessOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturesEnabled) GetAccessOptionsOk() ([]string, bool) {
	if o == nil || o.AccessOptions == nil {
		return nil, false
	}
	return o.AccessOptions, true
}

// HasAccessOptions returns a boolean if a field has been set.
func (o *FeaturesEnabled) HasAccessOptions() bool {
	if o != nil && o.AccessOptions != nil {
		return true
	}

	return false
}

// SetAccessOptions gets a reference to the given []string and assigns it to the AccessOptions field.
func (o *FeaturesEnabled) SetAccessOptions(v []string) {
	o.AccessOptions = v
}

func (o FeaturesEnabled) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Similarity != nil {
		toSerialize["similarity"] = o.Similarity
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if o.ProductName != nil {
		toSerialize["product_name"] = o.ProductName
	}
	if o.AccessOptions != nil {
		toSerialize["access_options"] = o.AccessOptions
	}
	return json.Marshal(toSerialize)
}

type NullableFeaturesEnabled struct {
	value *FeaturesEnabled
	isSet bool
}

func (v NullableFeaturesEnabled) Get() *FeaturesEnabled {
	return v.value
}

func (v *NullableFeaturesEnabled) Set(val *FeaturesEnabled) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturesEnabled) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturesEnabled) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturesEnabled(val *FeaturesEnabled) *NullableFeaturesEnabled {
	return &NullableFeaturesEnabled{value: val, isSet: true}
}

func (v NullableFeaturesEnabled) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeaturesEnabled) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

