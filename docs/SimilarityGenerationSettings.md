# SimilarityGenerationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchRepositories** | **[]string** | List of repositories to search | 
**SubmissionAutoExcludes** | Pointer to **[]string** | List of submission IDs to exclude from report | [optional] 
**AutoExcludeSelfMatchingScope** | Pointer to **string** | self matching submissions to exclude from report | [optional] 
**Priority** | Pointer to **bool** | Priority level of report generation | [optional] 

## Methods

### NewSimilarityGenerationSettings

`func NewSimilarityGenerationSettings(searchRepositories []string, ) *SimilarityGenerationSettings`

NewSimilarityGenerationSettings instantiates a new SimilarityGenerationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimilarityGenerationSettingsWithDefaults

`func NewSimilarityGenerationSettingsWithDefaults() *SimilarityGenerationSettings`

NewSimilarityGenerationSettingsWithDefaults instantiates a new SimilarityGenerationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchRepositories

`func (o *SimilarityGenerationSettings) GetSearchRepositories() []string`

GetSearchRepositories returns the SearchRepositories field if non-nil, zero value otherwise.

### GetSearchRepositoriesOk

`func (o *SimilarityGenerationSettings) GetSearchRepositoriesOk() (*[]string, bool)`

GetSearchRepositoriesOk returns a tuple with the SearchRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchRepositories

`func (o *SimilarityGenerationSettings) SetSearchRepositories(v []string)`

SetSearchRepositories sets SearchRepositories field to given value.


### GetSubmissionAutoExcludes

`func (o *SimilarityGenerationSettings) GetSubmissionAutoExcludes() []string`

GetSubmissionAutoExcludes returns the SubmissionAutoExcludes field if non-nil, zero value otherwise.

### GetSubmissionAutoExcludesOk

`func (o *SimilarityGenerationSettings) GetSubmissionAutoExcludesOk() (*[]string, bool)`

GetSubmissionAutoExcludesOk returns a tuple with the SubmissionAutoExcludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionAutoExcludes

`func (o *SimilarityGenerationSettings) SetSubmissionAutoExcludes(v []string)`

SetSubmissionAutoExcludes sets SubmissionAutoExcludes field to given value.

### HasSubmissionAutoExcludes

`func (o *SimilarityGenerationSettings) HasSubmissionAutoExcludes() bool`

HasSubmissionAutoExcludes returns a boolean if a field has been set.

### GetAutoExcludeSelfMatchingScope

`func (o *SimilarityGenerationSettings) GetAutoExcludeSelfMatchingScope() string`

GetAutoExcludeSelfMatchingScope returns the AutoExcludeSelfMatchingScope field if non-nil, zero value otherwise.

### GetAutoExcludeSelfMatchingScopeOk

`func (o *SimilarityGenerationSettings) GetAutoExcludeSelfMatchingScopeOk() (*string, bool)`

GetAutoExcludeSelfMatchingScopeOk returns a tuple with the AutoExcludeSelfMatchingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExcludeSelfMatchingScope

`func (o *SimilarityGenerationSettings) SetAutoExcludeSelfMatchingScope(v string)`

SetAutoExcludeSelfMatchingScope sets AutoExcludeSelfMatchingScope field to given value.

### HasAutoExcludeSelfMatchingScope

`func (o *SimilarityGenerationSettings) HasAutoExcludeSelfMatchingScope() bool`

HasAutoExcludeSelfMatchingScope returns a boolean if a field has been set.

### GetPriority

`func (o *SimilarityGenerationSettings) GetPriority() bool`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SimilarityGenerationSettings) GetPriorityOk() (*bool, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SimilarityGenerationSettings) SetPriority(v bool)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SimilarityGenerationSettings) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


