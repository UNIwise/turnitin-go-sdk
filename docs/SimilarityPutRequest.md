# SimilarityPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexingSettings** | Pointer to [**IndexingSettings**](IndexingSettings.md) |  | [optional] 
**GenerationSettings** | [**SimilarityGenerationSettings**](SimilarityGenerationSettings.md) |  | 
**ViewSettings** | Pointer to [**SimilarityViewSettings**](SimilarityViewSettings.md) |  | [optional] 

## Methods

### NewSimilarityPutRequest

`func NewSimilarityPutRequest(generationSettings SimilarityGenerationSettings, ) *SimilarityPutRequest`

NewSimilarityPutRequest instantiates a new SimilarityPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimilarityPutRequestWithDefaults

`func NewSimilarityPutRequestWithDefaults() *SimilarityPutRequest`

NewSimilarityPutRequestWithDefaults instantiates a new SimilarityPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexingSettings

`func (o *SimilarityPutRequest) GetIndexingSettings() IndexingSettings`

GetIndexingSettings returns the IndexingSettings field if non-nil, zero value otherwise.

### GetIndexingSettingsOk

`func (o *SimilarityPutRequest) GetIndexingSettingsOk() (*IndexingSettings, bool)`

GetIndexingSettingsOk returns a tuple with the IndexingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingSettings

`func (o *SimilarityPutRequest) SetIndexingSettings(v IndexingSettings)`

SetIndexingSettings sets IndexingSettings field to given value.

### HasIndexingSettings

`func (o *SimilarityPutRequest) HasIndexingSettings() bool`

HasIndexingSettings returns a boolean if a field has been set.

### GetGenerationSettings

`func (o *SimilarityPutRequest) GetGenerationSettings() SimilarityGenerationSettings`

GetGenerationSettings returns the GenerationSettings field if non-nil, zero value otherwise.

### GetGenerationSettingsOk

`func (o *SimilarityPutRequest) GetGenerationSettingsOk() (*SimilarityGenerationSettings, bool)`

GetGenerationSettingsOk returns a tuple with the GenerationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationSettings

`func (o *SimilarityPutRequest) SetGenerationSettings(v SimilarityGenerationSettings)`

SetGenerationSettings sets GenerationSettings field to given value.


### GetViewSettings

`func (o *SimilarityPutRequest) GetViewSettings() SimilarityViewSettings`

GetViewSettings returns the ViewSettings field if non-nil, zero value otherwise.

### GetViewSettingsOk

`func (o *SimilarityPutRequest) GetViewSettingsOk() (*SimilarityViewSettings, bool)`

GetViewSettingsOk returns a tuple with the ViewSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewSettings

`func (o *SimilarityPutRequest) SetViewSettings(v SimilarityViewSettings)`

SetViewSettings sets ViewSettings field to given value.

### HasViewSettings

`func (o *SimilarityPutRequest) HasViewSettings() bool`

HasViewSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


