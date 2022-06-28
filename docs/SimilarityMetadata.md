# SimilarityMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverallMatchPercentage** | **int32** | Represents the percentage match against all sources | 
**InternetMatchPercentage** | Pointer to **NullableInt32** | Represents the percentage match against internet | [optional] 
**PublicationMatchPercentage** | Pointer to **NullableInt32** | Represents the percentage match against all publications | [optional] 
**SubmittedWorksMatchPercentage** | Pointer to **NullableInt32** | Represents the percentage match against all submitted works | [optional] 
**SubmissionId** | **string** |  | 
**Status** | **string** | possible values PENDING, COMPLETE | 
**TimeGenerated** | **string** | Time the report finished generating.  If not set the report has not finished generating | 
**TimeRequested** | **string** | Time the report was requested | 
**TopMatches** | [**[]SimilarityMetadataAllOfTopMatches**](SimilarityMetadataAllOfTopMatches.md) | Top matches | 
**TopSourceLargestMatchedWordCount** | **int32** | Largest individual matched word count, 0 if there isn&#39;t a match to this submission. | 

## Methods

### NewSimilarityMetadata

`func NewSimilarityMetadata(overallMatchPercentage int32, submissionId string, status string, timeGenerated string, timeRequested string, topMatches []SimilarityMetadataAllOfTopMatches, topSourceLargestMatchedWordCount int32, ) *SimilarityMetadata`

NewSimilarityMetadata instantiates a new SimilarityMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimilarityMetadataWithDefaults

`func NewSimilarityMetadataWithDefaults() *SimilarityMetadata`

NewSimilarityMetadataWithDefaults instantiates a new SimilarityMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverallMatchPercentage

`func (o *SimilarityMetadata) GetOverallMatchPercentage() int32`

GetOverallMatchPercentage returns the OverallMatchPercentage field if non-nil, zero value otherwise.

### GetOverallMatchPercentageOk

`func (o *SimilarityMetadata) GetOverallMatchPercentageOk() (*int32, bool)`

GetOverallMatchPercentageOk returns a tuple with the OverallMatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallMatchPercentage

`func (o *SimilarityMetadata) SetOverallMatchPercentage(v int32)`

SetOverallMatchPercentage sets OverallMatchPercentage field to given value.


### GetInternetMatchPercentage

`func (o *SimilarityMetadata) GetInternetMatchPercentage() int32`

GetInternetMatchPercentage returns the InternetMatchPercentage field if non-nil, zero value otherwise.

### GetInternetMatchPercentageOk

`func (o *SimilarityMetadata) GetInternetMatchPercentageOk() (*int32, bool)`

GetInternetMatchPercentageOk returns a tuple with the InternetMatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetMatchPercentage

`func (o *SimilarityMetadata) SetInternetMatchPercentage(v int32)`

SetInternetMatchPercentage sets InternetMatchPercentage field to given value.

### HasInternetMatchPercentage

`func (o *SimilarityMetadata) HasInternetMatchPercentage() bool`

HasInternetMatchPercentage returns a boolean if a field has been set.

### SetInternetMatchPercentageNil

`func (o *SimilarityMetadata) SetInternetMatchPercentageNil(b bool)`

 SetInternetMatchPercentageNil sets the value for InternetMatchPercentage to be an explicit nil

### UnsetInternetMatchPercentage
`func (o *SimilarityMetadata) UnsetInternetMatchPercentage()`

UnsetInternetMatchPercentage ensures that no value is present for InternetMatchPercentage, not even an explicit nil
### GetPublicationMatchPercentage

`func (o *SimilarityMetadata) GetPublicationMatchPercentage() int32`

GetPublicationMatchPercentage returns the PublicationMatchPercentage field if non-nil, zero value otherwise.

### GetPublicationMatchPercentageOk

`func (o *SimilarityMetadata) GetPublicationMatchPercentageOk() (*int32, bool)`

GetPublicationMatchPercentageOk returns a tuple with the PublicationMatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationMatchPercentage

`func (o *SimilarityMetadata) SetPublicationMatchPercentage(v int32)`

SetPublicationMatchPercentage sets PublicationMatchPercentage field to given value.

### HasPublicationMatchPercentage

`func (o *SimilarityMetadata) HasPublicationMatchPercentage() bool`

HasPublicationMatchPercentage returns a boolean if a field has been set.

### SetPublicationMatchPercentageNil

`func (o *SimilarityMetadata) SetPublicationMatchPercentageNil(b bool)`

 SetPublicationMatchPercentageNil sets the value for PublicationMatchPercentage to be an explicit nil

### UnsetPublicationMatchPercentage
`func (o *SimilarityMetadata) UnsetPublicationMatchPercentage()`

UnsetPublicationMatchPercentage ensures that no value is present for PublicationMatchPercentage, not even an explicit nil
### GetSubmittedWorksMatchPercentage

`func (o *SimilarityMetadata) GetSubmittedWorksMatchPercentage() int32`

GetSubmittedWorksMatchPercentage returns the SubmittedWorksMatchPercentage field if non-nil, zero value otherwise.

### GetSubmittedWorksMatchPercentageOk

`func (o *SimilarityMetadata) GetSubmittedWorksMatchPercentageOk() (*int32, bool)`

GetSubmittedWorksMatchPercentageOk returns a tuple with the SubmittedWorksMatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedWorksMatchPercentage

`func (o *SimilarityMetadata) SetSubmittedWorksMatchPercentage(v int32)`

SetSubmittedWorksMatchPercentage sets SubmittedWorksMatchPercentage field to given value.

### HasSubmittedWorksMatchPercentage

`func (o *SimilarityMetadata) HasSubmittedWorksMatchPercentage() bool`

HasSubmittedWorksMatchPercentage returns a boolean if a field has been set.

### SetSubmittedWorksMatchPercentageNil

`func (o *SimilarityMetadata) SetSubmittedWorksMatchPercentageNil(b bool)`

 SetSubmittedWorksMatchPercentageNil sets the value for SubmittedWorksMatchPercentage to be an explicit nil

### UnsetSubmittedWorksMatchPercentage
`func (o *SimilarityMetadata) UnsetSubmittedWorksMatchPercentage()`

UnsetSubmittedWorksMatchPercentage ensures that no value is present for SubmittedWorksMatchPercentage, not even an explicit nil
### GetSubmissionId

`func (o *SimilarityMetadata) GetSubmissionId() string`

GetSubmissionId returns the SubmissionId field if non-nil, zero value otherwise.

### GetSubmissionIdOk

`func (o *SimilarityMetadata) GetSubmissionIdOk() (*string, bool)`

GetSubmissionIdOk returns a tuple with the SubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionId

`func (o *SimilarityMetadata) SetSubmissionId(v string)`

SetSubmissionId sets SubmissionId field to given value.


### GetStatus

`func (o *SimilarityMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimilarityMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimilarityMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimeGenerated

`func (o *SimilarityMetadata) GetTimeGenerated() string`

GetTimeGenerated returns the TimeGenerated field if non-nil, zero value otherwise.

### GetTimeGeneratedOk

`func (o *SimilarityMetadata) GetTimeGeneratedOk() (*string, bool)`

GetTimeGeneratedOk returns a tuple with the TimeGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeGenerated

`func (o *SimilarityMetadata) SetTimeGenerated(v string)`

SetTimeGenerated sets TimeGenerated field to given value.


### GetTimeRequested

`func (o *SimilarityMetadata) GetTimeRequested() string`

GetTimeRequested returns the TimeRequested field if non-nil, zero value otherwise.

### GetTimeRequestedOk

`func (o *SimilarityMetadata) GetTimeRequestedOk() (*string, bool)`

GetTimeRequestedOk returns a tuple with the TimeRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRequested

`func (o *SimilarityMetadata) SetTimeRequested(v string)`

SetTimeRequested sets TimeRequested field to given value.


### GetTopMatches

`func (o *SimilarityMetadata) GetTopMatches() []SimilarityMetadataAllOfTopMatches`

GetTopMatches returns the TopMatches field if non-nil, zero value otherwise.

### GetTopMatchesOk

`func (o *SimilarityMetadata) GetTopMatchesOk() (*[]SimilarityMetadataAllOfTopMatches, bool)`

GetTopMatchesOk returns a tuple with the TopMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopMatches

`func (o *SimilarityMetadata) SetTopMatches(v []SimilarityMetadataAllOfTopMatches)`

SetTopMatches sets TopMatches field to given value.


### GetTopSourceLargestMatchedWordCount

`func (o *SimilarityMetadata) GetTopSourceLargestMatchedWordCount() int32`

GetTopSourceLargestMatchedWordCount returns the TopSourceLargestMatchedWordCount field if non-nil, zero value otherwise.

### GetTopSourceLargestMatchedWordCountOk

`func (o *SimilarityMetadata) GetTopSourceLargestMatchedWordCountOk() (*int32, bool)`

GetTopSourceLargestMatchedWordCountOk returns a tuple with the TopSourceLargestMatchedWordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSourceLargestMatchedWordCount

`func (o *SimilarityMetadata) SetTopSourceLargestMatchedWordCount(v int32)`

SetTopSourceLargestMatchedWordCount sets TopSourceLargestMatchedWordCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


