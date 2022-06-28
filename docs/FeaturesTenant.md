# FeaturesTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequireEula** | Pointer to **bool** | a flag indicating whether this tenant requires EULA checks to use this API | [optional] 

## Methods

### NewFeaturesTenant

`func NewFeaturesTenant() *FeaturesTenant`

NewFeaturesTenant instantiates a new FeaturesTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturesTenantWithDefaults

`func NewFeaturesTenantWithDefaults() *FeaturesTenant`

NewFeaturesTenantWithDefaults instantiates a new FeaturesTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequireEula

`func (o *FeaturesTenant) GetRequireEula() bool`

GetRequireEula returns the RequireEula field if non-nil, zero value otherwise.

### GetRequireEulaOk

`func (o *FeaturesTenant) GetRequireEulaOk() (*bool, bool)`

GetRequireEulaOk returns a tuple with the RequireEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEula

`func (o *FeaturesTenant) SetRequireEula(v bool)`

SetRequireEula sets RequireEula field to given value.

### HasRequireEula

`func (o *FeaturesTenant) HasRequireEula() bool`

HasRequireEula returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


