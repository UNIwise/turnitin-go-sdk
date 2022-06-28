# SimilarityMetadataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubmissionId** | **string** |  | 
**Status** | **string** | possible values PENDING, COMPLETE | 
**TimeGenerated** | **string** | Time the report finished generating.  If not set the report has not finished generating | 
**TimeRequested** | **string** | Time the report was requested | 
**TopMatches** | [**[]SimilarityMetadataAllOfTopMatches**](SimilarityMetadataAllOfTopMatches.md) | Top matches | 
**TopSourceLargestMatchedWordCount** | **int32** | Largest individual matched word count, 0 if there isn&#39;t a match to this submission. | 

## Methods

### NewSimilarityMetadataAllOf

`func NewSimilarityMetadataAllOf(submissionId string, status string, timeGenerated string, timeRequested string, topMatches []SimilarityMetadataAllOfTopMatches, topSourceLargestMatchedWordCount int32, ) *SimilarityMetadataAllOf`

NewSimilarityMetadataAllOf instantiates a new SimilarityMetadataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimilarityMetadataAllOfWithDefaults

`func NewSimilarityMetadataAllOfWithDefaults() *SimilarityMetadataAllOf`

NewSimilarityMetadataAllOfWithDefaults instantiates a new SimilarityMetadataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubmissionId

`func (o *SimilarityMetadataAllOf) GetSubmissionId() string`

GetSubmissionId returns the SubmissionId field if non-nil, zero value otherwise.

### GetSubmissionIdOk

`func (o *SimilarityMetadataAllOf) GetSubmissionIdOk() (*string, bool)`

GetSubmissionIdOk returns a tuple with the SubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionId

`func (o *SimilarityMetadataAllOf) SetSubmissionId(v string)`

SetSubmissionId sets SubmissionId field to given value.


### GetStatus

`func (o *SimilarityMetadataAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimilarityMetadataAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimilarityMetadataAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimeGenerated

`func (o *SimilarityMetadataAllOf) GetTimeGenerated() string`

GetTimeGenerated returns the TimeGenerated field if non-nil, zero value otherwise.

### GetTimeGeneratedOk

`func (o *SimilarityMetadataAllOf) GetTimeGeneratedOk() (*string, bool)`

GetTimeGeneratedOk returns a tuple with the TimeGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeGenerated

`func (o *SimilarityMetadataAllOf) SetTimeGenerated(v string)`

SetTimeGenerated sets TimeGenerated field to given value.


### GetTimeRequested

`func (o *SimilarityMetadataAllOf) GetTimeRequested() string`

GetTimeRequested returns the TimeRequested field if non-nil, zero value otherwise.

### GetTimeRequestedOk

`func (o *SimilarityMetadataAllOf) GetTimeRequestedOk() (*string, bool)`

GetTimeRequestedOk returns a tuple with the TimeRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRequested

`func (o *SimilarityMetadataAllOf) SetTimeRequested(v string)`

SetTimeRequested sets TimeRequested field to given value.


### GetTopMatches

`func (o *SimilarityMetadataAllOf) GetTopMatches() []SimilarityMetadataAllOfTopMatches`

GetTopMatches returns the TopMatches field if non-nil, zero value otherwise.

### GetTopMatchesOk

`func (o *SimilarityMetadataAllOf) GetTopMatchesOk() (*[]SimilarityMetadataAllOfTopMatches, bool)`

GetTopMatchesOk returns a tuple with the TopMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopMatches

`func (o *SimilarityMetadataAllOf) SetTopMatches(v []SimilarityMetadataAllOfTopMatches)`

SetTopMatches sets TopMatches field to given value.


### GetTopSourceLargestMatchedWordCount

`func (o *SimilarityMetadataAllOf) GetTopSourceLargestMatchedWordCount() int32`

GetTopSourceLargestMatchedWordCount returns the TopSourceLargestMatchedWordCount field if non-nil, zero value otherwise.

### GetTopSourceLargestMatchedWordCountOk

`func (o *SimilarityMetadataAllOf) GetTopSourceLargestMatchedWordCountOk() (*int32, bool)`

GetTopSourceLargestMatchedWordCountOk returns a tuple with the TopSourceLargestMatchedWordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSourceLargestMatchedWordCount

`func (o *SimilarityMetadataAllOf) SetTopSourceLargestMatchedWordCount(v int32)`

SetTopSourceLargestMatchedWordCount sets TopSourceLargestMatchedWordCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


