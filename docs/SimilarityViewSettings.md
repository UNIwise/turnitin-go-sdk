# SimilarityViewSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeQuotes** | Pointer to **bool** | If set to true, text in quotes will not count as similar content  | [optional] 
**ExcludeBibliography** | Pointer to **bool** | If set to true, text in a bibliography section will not count as similar content  | [optional] 
**ExcludeCitations** | Pointer to **bool** | If set to true, text in citations will not count as similar content  | [optional] 
**ExcludeAbstract** | Pointer to **bool** | If set to true, text in the abstract section of the submission will not count as similar content  | [optional] 
**ExcludeMethods** | Pointer to **bool** | If set to true, text in the method section of the submission will not count as similar content  | [optional] 
**ExcludeSmallMatches** | Pointer to **int32** | If set, similarity matches that match less than the specified amount of words will not count as similar content  | [optional] 
**ExcludeInternet** | Pointer to **bool** | If set to true, text matched to the internet collection will not count as similar content  | [optional] 
**ExcludePublications** | Pointer to **bool** | If set to true, text matched to the publications collection will not count as similar content  | [optional] 
**ExcludeCrossref** | Pointer to **bool** | If set to true, text matched to the Crossref collection will not count as similar content  | [optional] 
**ExcludeCrossrefPostedContent** | Pointer to **bool** | If set to true, text matched to the Crossref Posted Content collection will not count as similar content  | [optional] 
**ExcludeSubmittedWorks** | Pointer to **bool** | If set to true, text matched to the submitted works collection will not count as similar content calculated as if submitted work was not part of the paper  | [optional] 
**ExcludeCustomSections** | Pointer to **bool** | If set to true, text matched to the custom sections defined in the admin settings will not count as similar content calculated as if section was not part of the paper  | [optional] 
**ExcludePreprints** | Pointer to **bool** | If set to true, it will exclude preprints. A preprint is a version of a scholarly or scientific paper that precedes formal peer review and publication in a peer-reviewed scholarly or scientific journal.  | [optional] 

## Methods

### NewSimilarityViewSettings

`func NewSimilarityViewSettings() *SimilarityViewSettings`

NewSimilarityViewSettings instantiates a new SimilarityViewSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimilarityViewSettingsWithDefaults

`func NewSimilarityViewSettingsWithDefaults() *SimilarityViewSettings`

NewSimilarityViewSettingsWithDefaults instantiates a new SimilarityViewSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeQuotes

`func (o *SimilarityViewSettings) GetExcludeQuotes() bool`

GetExcludeQuotes returns the ExcludeQuotes field if non-nil, zero value otherwise.

### GetExcludeQuotesOk

`func (o *SimilarityViewSettings) GetExcludeQuotesOk() (*bool, bool)`

GetExcludeQuotesOk returns a tuple with the ExcludeQuotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeQuotes

`func (o *SimilarityViewSettings) SetExcludeQuotes(v bool)`

SetExcludeQuotes sets ExcludeQuotes field to given value.

### HasExcludeQuotes

`func (o *SimilarityViewSettings) HasExcludeQuotes() bool`

HasExcludeQuotes returns a boolean if a field has been set.

### GetExcludeBibliography

`func (o *SimilarityViewSettings) GetExcludeBibliography() bool`

GetExcludeBibliography returns the ExcludeBibliography field if non-nil, zero value otherwise.

### GetExcludeBibliographyOk

`func (o *SimilarityViewSettings) GetExcludeBibliographyOk() (*bool, bool)`

GetExcludeBibliographyOk returns a tuple with the ExcludeBibliography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBibliography

`func (o *SimilarityViewSettings) SetExcludeBibliography(v bool)`

SetExcludeBibliography sets ExcludeBibliography field to given value.

### HasExcludeBibliography

`func (o *SimilarityViewSettings) HasExcludeBibliography() bool`

HasExcludeBibliography returns a boolean if a field has been set.

### GetExcludeCitations

`func (o *SimilarityViewSettings) GetExcludeCitations() bool`

GetExcludeCitations returns the ExcludeCitations field if non-nil, zero value otherwise.

### GetExcludeCitationsOk

`func (o *SimilarityViewSettings) GetExcludeCitationsOk() (*bool, bool)`

GetExcludeCitationsOk returns a tuple with the ExcludeCitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCitations

`func (o *SimilarityViewSettings) SetExcludeCitations(v bool)`

SetExcludeCitations sets ExcludeCitations field to given value.

### HasExcludeCitations

`func (o *SimilarityViewSettings) HasExcludeCitations() bool`

HasExcludeCitations returns a boolean if a field has been set.

### GetExcludeAbstract

`func (o *SimilarityViewSettings) GetExcludeAbstract() bool`

GetExcludeAbstract returns the ExcludeAbstract field if non-nil, zero value otherwise.

### GetExcludeAbstractOk

`func (o *SimilarityViewSettings) GetExcludeAbstractOk() (*bool, bool)`

GetExcludeAbstractOk returns a tuple with the ExcludeAbstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAbstract

`func (o *SimilarityViewSettings) SetExcludeAbstract(v bool)`

SetExcludeAbstract sets ExcludeAbstract field to given value.

### HasExcludeAbstract

`func (o *SimilarityViewSettings) HasExcludeAbstract() bool`

HasExcludeAbstract returns a boolean if a field has been set.

### GetExcludeMethods

`func (o *SimilarityViewSettings) GetExcludeMethods() bool`

GetExcludeMethods returns the ExcludeMethods field if non-nil, zero value otherwise.

### GetExcludeMethodsOk

`func (o *SimilarityViewSettings) GetExcludeMethodsOk() (*bool, bool)`

GetExcludeMethodsOk returns a tuple with the ExcludeMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeMethods

`func (o *SimilarityViewSettings) SetExcludeMethods(v bool)`

SetExcludeMethods sets ExcludeMethods field to given value.

### HasExcludeMethods

`func (o *SimilarityViewSettings) HasExcludeMethods() bool`

HasExcludeMethods returns a boolean if a field has been set.

### GetExcludeSmallMatches

`func (o *SimilarityViewSettings) GetExcludeSmallMatches() int32`

GetExcludeSmallMatches returns the ExcludeSmallMatches field if non-nil, zero value otherwise.

### GetExcludeSmallMatchesOk

`func (o *SimilarityViewSettings) GetExcludeSmallMatchesOk() (*int32, bool)`

GetExcludeSmallMatchesOk returns a tuple with the ExcludeSmallMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSmallMatches

`func (o *SimilarityViewSettings) SetExcludeSmallMatches(v int32)`

SetExcludeSmallMatches sets ExcludeSmallMatches field to given value.

### HasExcludeSmallMatches

`func (o *SimilarityViewSettings) HasExcludeSmallMatches() bool`

HasExcludeSmallMatches returns a boolean if a field has been set.

### GetExcludeInternet

`func (o *SimilarityViewSettings) GetExcludeInternet() bool`

GetExcludeInternet returns the ExcludeInternet field if non-nil, zero value otherwise.

### GetExcludeInternetOk

`func (o *SimilarityViewSettings) GetExcludeInternetOk() (*bool, bool)`

GetExcludeInternetOk returns a tuple with the ExcludeInternet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeInternet

`func (o *SimilarityViewSettings) SetExcludeInternet(v bool)`

SetExcludeInternet sets ExcludeInternet field to given value.

### HasExcludeInternet

`func (o *SimilarityViewSettings) HasExcludeInternet() bool`

HasExcludeInternet returns a boolean if a field has been set.

### GetExcludePublications

`func (o *SimilarityViewSettings) GetExcludePublications() bool`

GetExcludePublications returns the ExcludePublications field if non-nil, zero value otherwise.

### GetExcludePublicationsOk

`func (o *SimilarityViewSettings) GetExcludePublicationsOk() (*bool, bool)`

GetExcludePublicationsOk returns a tuple with the ExcludePublications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePublications

`func (o *SimilarityViewSettings) SetExcludePublications(v bool)`

SetExcludePublications sets ExcludePublications field to given value.

### HasExcludePublications

`func (o *SimilarityViewSettings) HasExcludePublications() bool`

HasExcludePublications returns a boolean if a field has been set.

### GetExcludeCrossref

`func (o *SimilarityViewSettings) GetExcludeCrossref() bool`

GetExcludeCrossref returns the ExcludeCrossref field if non-nil, zero value otherwise.

### GetExcludeCrossrefOk

`func (o *SimilarityViewSettings) GetExcludeCrossrefOk() (*bool, bool)`

GetExcludeCrossrefOk returns a tuple with the ExcludeCrossref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCrossref

`func (o *SimilarityViewSettings) SetExcludeCrossref(v bool)`

SetExcludeCrossref sets ExcludeCrossref field to given value.

### HasExcludeCrossref

`func (o *SimilarityViewSettings) HasExcludeCrossref() bool`

HasExcludeCrossref returns a boolean if a field has been set.

### GetExcludeCrossrefPostedContent

`func (o *SimilarityViewSettings) GetExcludeCrossrefPostedContent() bool`

GetExcludeCrossrefPostedContent returns the ExcludeCrossrefPostedContent field if non-nil, zero value otherwise.

### GetExcludeCrossrefPostedContentOk

`func (o *SimilarityViewSettings) GetExcludeCrossrefPostedContentOk() (*bool, bool)`

GetExcludeCrossrefPostedContentOk returns a tuple with the ExcludeCrossrefPostedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCrossrefPostedContent

`func (o *SimilarityViewSettings) SetExcludeCrossrefPostedContent(v bool)`

SetExcludeCrossrefPostedContent sets ExcludeCrossrefPostedContent field to given value.

### HasExcludeCrossrefPostedContent

`func (o *SimilarityViewSettings) HasExcludeCrossrefPostedContent() bool`

HasExcludeCrossrefPostedContent returns a boolean if a field has been set.

### GetExcludeSubmittedWorks

`func (o *SimilarityViewSettings) GetExcludeSubmittedWorks() bool`

GetExcludeSubmittedWorks returns the ExcludeSubmittedWorks field if non-nil, zero value otherwise.

### GetExcludeSubmittedWorksOk

`func (o *SimilarityViewSettings) GetExcludeSubmittedWorksOk() (*bool, bool)`

GetExcludeSubmittedWorksOk returns a tuple with the ExcludeSubmittedWorks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSubmittedWorks

`func (o *SimilarityViewSettings) SetExcludeSubmittedWorks(v bool)`

SetExcludeSubmittedWorks sets ExcludeSubmittedWorks field to given value.

### HasExcludeSubmittedWorks

`func (o *SimilarityViewSettings) HasExcludeSubmittedWorks() bool`

HasExcludeSubmittedWorks returns a boolean if a field has been set.

### GetExcludeCustomSections

`func (o *SimilarityViewSettings) GetExcludeCustomSections() bool`

GetExcludeCustomSections returns the ExcludeCustomSections field if non-nil, zero value otherwise.

### GetExcludeCustomSectionsOk

`func (o *SimilarityViewSettings) GetExcludeCustomSectionsOk() (*bool, bool)`

GetExcludeCustomSectionsOk returns a tuple with the ExcludeCustomSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCustomSections

`func (o *SimilarityViewSettings) SetExcludeCustomSections(v bool)`

SetExcludeCustomSections sets ExcludeCustomSections field to given value.

### HasExcludeCustomSections

`func (o *SimilarityViewSettings) HasExcludeCustomSections() bool`

HasExcludeCustomSections returns a boolean if a field has been set.

### GetExcludePreprints

`func (o *SimilarityViewSettings) GetExcludePreprints() bool`

GetExcludePreprints returns the ExcludePreprints field if non-nil, zero value otherwise.

### GetExcludePreprintsOk

`func (o *SimilarityViewSettings) GetExcludePreprintsOk() (*bool, bool)`

GetExcludePreprintsOk returns a tuple with the ExcludePreprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePreprints

`func (o *SimilarityViewSettings) SetExcludePreprints(v bool)`

SetExcludePreprints sets ExcludePreprints field to given value.

### HasExcludePreprints

`func (o *SimilarityViewSettings) HasExcludePreprints() bool`

HasExcludePreprints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


