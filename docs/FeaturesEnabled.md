# FeaturesEnabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Similarity** | Pointer to [**FeaturesSimilarity**](FeaturesSimilarity.md) |  | [optional] 
**Tenant** | Pointer to [**FeaturesTenant**](FeaturesTenant.md) |  | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
**AccessOptions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFeaturesEnabled

`func NewFeaturesEnabled() *FeaturesEnabled`

NewFeaturesEnabled instantiates a new FeaturesEnabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturesEnabledWithDefaults

`func NewFeaturesEnabledWithDefaults() *FeaturesEnabled`

NewFeaturesEnabledWithDefaults instantiates a new FeaturesEnabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSimilarity

`func (o *FeaturesEnabled) GetSimilarity() FeaturesSimilarity`

GetSimilarity returns the Similarity field if non-nil, zero value otherwise.

### GetSimilarityOk

`func (o *FeaturesEnabled) GetSimilarityOk() (*FeaturesSimilarity, bool)`

GetSimilarityOk returns a tuple with the Similarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarity

`func (o *FeaturesEnabled) SetSimilarity(v FeaturesSimilarity)`

SetSimilarity sets Similarity field to given value.

### HasSimilarity

`func (o *FeaturesEnabled) HasSimilarity() bool`

HasSimilarity returns a boolean if a field has been set.

### GetTenant

`func (o *FeaturesEnabled) GetTenant() FeaturesTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *FeaturesEnabled) GetTenantOk() (*FeaturesTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *FeaturesEnabled) SetTenant(v FeaturesTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *FeaturesEnabled) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetProductName

`func (o *FeaturesEnabled) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *FeaturesEnabled) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *FeaturesEnabled) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *FeaturesEnabled) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetAccessOptions

`func (o *FeaturesEnabled) GetAccessOptions() []string`

GetAccessOptions returns the AccessOptions field if non-nil, zero value otherwise.

### GetAccessOptionsOk

`func (o *FeaturesEnabled) GetAccessOptionsOk() (*[]string, bool)`

GetAccessOptionsOk returns a tuple with the AccessOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessOptions

`func (o *FeaturesEnabled) SetAccessOptions(v []string)`

SetAccessOptions sets AccessOptions field to given value.

### HasAccessOptions

`func (o *FeaturesEnabled) HasAccessOptions() bool`

HasAccessOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


