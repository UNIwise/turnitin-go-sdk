# SimilarityCompleteWebhookRequest

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
**Metadata** | Pointer to [**SubmissionCompleteWebhookRequestAllOfMetadata**](SubmissionCompleteWebhookRequestAllOfMetadata.md) |  | [optional] 

## Methods

### NewSimilarityCompleteWebhookRequest

`func NewSimilarityCompleteWebhookRequest(overallMatchPercentage int32, submissionId string, status string, timeGenerated string, timeRequested string, topMatches []SimilarityMetadataAllOfTopMatches, topSourceLargestMatchedWordCount int32, ) *SimilarityCompleteWebhookRequest`

NewSimilarityCompleteWebhookRequest instantiates a new SimilarityCompleteWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimilarityCompleteWebhookRequestWithDefaults

`func NewSimilarityCompleteWebhookRequestWithDefaults() *SimilarityCompleteWebhookRequest`

NewSimilarityCompleteWebhookRequestWithDefaults instantiates a new SimilarityCompleteWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverallMatchPercentage

`func (o *SimilarityCompleteWebhookRequest) GetOverallMatchPercentage() int32`

GetOverallMatchPercentage returns the OverallMatchPercentage field if non-nil, zero value otherwise.

### GetOverallMatchPercentageOk

`func (o *SimilarityCompleteWebhookRequest) GetOverallMatchPercentageOk() (*int32, bool)`

GetOverallMatchPercentageOk returns a tuple with the OverallMatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallMatchPercentage

`func (o *SimilarityCompleteWebhookRequest) SetOverallMatchPercentage(v int32)`

SetOverallMatchPercentage sets OverallMatchPercentage field to given value.


### GetInternetMatchPercentage

`func (o *SimilarityCompleteWebhookRequest) GetInternetMatchPercentage() int32`

GetInternetMatchPercentage returns the InternetMatchPercentage field if non-nil, zero value otherwise.

### GetInternetMatchPercentageOk

`func (o *SimilarityCompleteWebhookRequest) GetInternetMatchPercentageOk() (*int32, bool)`

GetInternetMatchPercentageOk returns a tuple with the InternetMatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetMatchPercentage

`func (o *SimilarityCompleteWebhookRequest) SetInternetMatchPercentage(v int32)`

SetInternetMatchPercentage sets InternetMatchPercentage field to given value.

### HasInternetMatchPercentage

`func (o *SimilarityCompleteWebhookRequest) HasInternetMatchPercentage() bool`

HasInternetMatchPercentage returns a boolean if a field has been set.

### SetInternetMatchPercentageNil

`func (o *SimilarityCompleteWebhookRequest) SetInternetMatchPercentageNil(b bool)`

 SetInternetMatchPercentageNil sets the value for InternetMatchPercentage to be an explicit nil

### UnsetInternetMatchPercentage
`func (o *SimilarityCompleteWebhookRequest) UnsetInternetMatchPercentage()`

UnsetInternetMatchPercentage ensures that no value is present for InternetMatchPercentage, not even an explicit nil
### GetPublicationMatchPercentage

`func (o *SimilarityCompleteWebhookRequest) GetPublicationMatchPercentage() int32`

GetPublicationMatchPercentage returns the PublicationMatchPercentage field if non-nil, zero value otherwise.

### GetPublicationMatchPercentageOk

`func (o *SimilarityCompleteWebhookRequest) GetPublicationMatchPercentageOk() (*int32, bool)`

GetPublicationMatchPercentageOk returns a tuple with the PublicationMatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationMatchPercentage

`func (o *SimilarityCompleteWebhookRequest) SetPublicationMatchPercentage(v int32)`

SetPublicationMatchPercentage sets PublicationMatchPercentage field to given value.

### HasPublicationMatchPercentage

`func (o *SimilarityCompleteWebhookRequest) HasPublicationMatchPercentage() bool`

HasPublicationMatchPercentage returns a boolean if a field has been set.

### SetPublicationMatchPercentageNil

`func (o *SimilarityCompleteWebhookRequest) SetPublicationMatchPercentageNil(b bool)`

 SetPublicationMatchPercentageNil sets the value for PublicationMatchPercentage to be an explicit nil

### UnsetPublicationMatchPercentage
`func (o *SimilarityCompleteWebhookRequest) UnsetPublicationMatchPercentage()`

UnsetPublicationMatchPercentage ensures that no value is present for PublicationMatchPercentage, not even an explicit nil
### GetSubmittedWorksMatchPercentage

`func (o *SimilarityCompleteWebhookRequest) GetSubmittedWorksMatchPercentage() int32`

GetSubmittedWorksMatchPercentage returns the SubmittedWorksMatchPercentage field if non-nil, zero value otherwise.

### GetSubmittedWorksMatchPercentageOk

`func (o *SimilarityCompleteWebhookRequest) GetSubmittedWorksMatchPercentageOk() (*int32, bool)`

GetSubmittedWorksMatchPercentageOk returns a tuple with the SubmittedWorksMatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedWorksMatchPercentage

`func (o *SimilarityCompleteWebhookRequest) SetSubmittedWorksMatchPercentage(v int32)`

SetSubmittedWorksMatchPercentage sets SubmittedWorksMatchPercentage field to given value.

### HasSubmittedWorksMatchPercentage

`func (o *SimilarityCompleteWebhookRequest) HasSubmittedWorksMatchPercentage() bool`

HasSubmittedWorksMatchPercentage returns a boolean if a field has been set.

### SetSubmittedWorksMatchPercentageNil

`func (o *SimilarityCompleteWebhookRequest) SetSubmittedWorksMatchPercentageNil(b bool)`

 SetSubmittedWorksMatchPercentageNil sets the value for SubmittedWorksMatchPercentage to be an explicit nil

### UnsetSubmittedWorksMatchPercentage
`func (o *SimilarityCompleteWebhookRequest) UnsetSubmittedWorksMatchPercentage()`

UnsetSubmittedWorksMatchPercentage ensures that no value is present for SubmittedWorksMatchPercentage, not even an explicit nil
### GetSubmissionId

`func (o *SimilarityCompleteWebhookRequest) GetSubmissionId() string`

GetSubmissionId returns the SubmissionId field if non-nil, zero value otherwise.

### GetSubmissionIdOk

`func (o *SimilarityCompleteWebhookRequest) GetSubmissionIdOk() (*string, bool)`

GetSubmissionIdOk returns a tuple with the SubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionId

`func (o *SimilarityCompleteWebhookRequest) SetSubmissionId(v string)`

SetSubmissionId sets SubmissionId field to given value.


### GetStatus

`func (o *SimilarityCompleteWebhookRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimilarityCompleteWebhookRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimilarityCompleteWebhookRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimeGenerated

`func (o *SimilarityCompleteWebhookRequest) GetTimeGenerated() string`

GetTimeGenerated returns the TimeGenerated field if non-nil, zero value otherwise.

### GetTimeGeneratedOk

`func (o *SimilarityCompleteWebhookRequest) GetTimeGeneratedOk() (*string, bool)`

GetTimeGeneratedOk returns a tuple with the TimeGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeGenerated

`func (o *SimilarityCompleteWebhookRequest) SetTimeGenerated(v string)`

SetTimeGenerated sets TimeGenerated field to given value.


### GetTimeRequested

`func (o *SimilarityCompleteWebhookRequest) GetTimeRequested() string`

GetTimeRequested returns the TimeRequested field if non-nil, zero value otherwise.

### GetTimeRequestedOk

`func (o *SimilarityCompleteWebhookRequest) GetTimeRequestedOk() (*string, bool)`

GetTimeRequestedOk returns a tuple with the TimeRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRequested

`func (o *SimilarityCompleteWebhookRequest) SetTimeRequested(v string)`

SetTimeRequested sets TimeRequested field to given value.


### GetTopMatches

`func (o *SimilarityCompleteWebhookRequest) GetTopMatches() []SimilarityMetadataAllOfTopMatches`

GetTopMatches returns the TopMatches field if non-nil, zero value otherwise.

### GetTopMatchesOk

`func (o *SimilarityCompleteWebhookRequest) GetTopMatchesOk() (*[]SimilarityMetadataAllOfTopMatches, bool)`

GetTopMatchesOk returns a tuple with the TopMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopMatches

`func (o *SimilarityCompleteWebhookRequest) SetTopMatches(v []SimilarityMetadataAllOfTopMatches)`

SetTopMatches sets TopMatches field to given value.


### GetTopSourceLargestMatchedWordCount

`func (o *SimilarityCompleteWebhookRequest) GetTopSourceLargestMatchedWordCount() int32`

GetTopSourceLargestMatchedWordCount returns the TopSourceLargestMatchedWordCount field if non-nil, zero value otherwise.

### GetTopSourceLargestMatchedWordCountOk

`func (o *SimilarityCompleteWebhookRequest) GetTopSourceLargestMatchedWordCountOk() (*int32, bool)`

GetTopSourceLargestMatchedWordCountOk returns a tuple with the TopSourceLargestMatchedWordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSourceLargestMatchedWordCount

`func (o *SimilarityCompleteWebhookRequest) SetTopSourceLargestMatchedWordCount(v int32)`

SetTopSourceLargestMatchedWordCount sets TopSourceLargestMatchedWordCount field to given value.


### GetMetadata

`func (o *SimilarityCompleteWebhookRequest) GetMetadata() SubmissionCompleteWebhookRequestAllOfMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SimilarityCompleteWebhookRequest) GetMetadataOk() (*SubmissionCompleteWebhookRequestAllOfMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SimilarityCompleteWebhookRequest) SetMetadata(v SubmissionCompleteWebhookRequestAllOfMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SimilarityCompleteWebhookRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


