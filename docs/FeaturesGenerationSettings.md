# FeaturesGenerationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchRepositories** | Pointer to **[]string** | List of repositories to search | [optional] 
**SubmissionAutoExcludes** | Pointer to **bool** |  | [optional] 

## Methods

### NewFeaturesGenerationSettings

`func NewFeaturesGenerationSettings() *FeaturesGenerationSettings`

NewFeaturesGenerationSettings instantiates a new FeaturesGenerationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturesGenerationSettingsWithDefaults

`func NewFeaturesGenerationSettingsWithDefaults() *FeaturesGenerationSettings`

NewFeaturesGenerationSettingsWithDefaults instantiates a new FeaturesGenerationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchRepositories

`func (o *FeaturesGenerationSettings) GetSearchRepositories() []string`

GetSearchRepositories returns the SearchRepositories field if non-nil, zero value otherwise.

### GetSearchRepositoriesOk

`func (o *FeaturesGenerationSettings) GetSearchRepositoriesOk() (*[]string, bool)`

GetSearchRepositoriesOk returns a tuple with the SearchRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchRepositories

`func (o *FeaturesGenerationSettings) SetSearchRepositories(v []string)`

SetSearchRepositories sets SearchRepositories field to given value.

### HasSearchRepositories

`func (o *FeaturesGenerationSettings) HasSearchRepositories() bool`

HasSearchRepositories returns a boolean if a field has been set.

### GetSubmissionAutoExcludes

`func (o *FeaturesGenerationSettings) GetSubmissionAutoExcludes() bool`

GetSubmissionAutoExcludes returns the SubmissionAutoExcludes field if non-nil, zero value otherwise.

### GetSubmissionAutoExcludesOk

`func (o *FeaturesGenerationSettings) GetSubmissionAutoExcludesOk() (*bool, bool)`

GetSubmissionAutoExcludesOk returns a tuple with the SubmissionAutoExcludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionAutoExcludes

`func (o *FeaturesGenerationSettings) SetSubmissionAutoExcludes(v bool)`

SetSubmissionAutoExcludes sets SubmissionAutoExcludes field to given value.

### HasSubmissionAutoExcludes

`func (o *FeaturesGenerationSettings) HasSubmissionAutoExcludes() bool`

HasSubmissionAutoExcludes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


