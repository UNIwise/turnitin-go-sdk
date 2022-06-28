# FeaturesSimilarityViewSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeQuotes** | Pointer to **bool** | If set to true, text in quotes will not count as similar content  | [optional] 
**ExcludeBibliography** | Pointer to **bool** | If set to true, text in a bibliography section will not count as similar content  | [optional] 
**ExcludeCitations** | Pointer to **bool** | If set to true, text in citations will not count as similar content  | [optional] 
**ExcludeAbstract** | Pointer to **bool** | If set to true, text in the abstract section of the submission will not count as similar content  | [optional] 
**ExcludeMethods** | Pointer to **bool** | If set to true, text in the method section of the submission will not count as similar content  | [optional] 
**ExcludeSmallMatches** | Pointer to **bool** | If set, similarity matches that match less than the specified amount of words will not count as similar content  | [optional] 
**ExcludeInternet** | Pointer to **bool** | If set to true, text matched to the internet collection will not count as similar content  | [optional] 
**ExcludePublications** | Pointer to **bool** | If set to true, text matched to the publications collection will not count as similar content  | [optional] 
**ExcludeCrossref** | Pointer to **bool** | If set to true, text matched to the Crossref collection will not count as similar content  | [optional] 
**ExcludeCrossrefPostedContent** | Pointer to **bool** | If set to true, text matched to the Crossref Posted Content collection will not count as similar content  | [optional] 
**ExcludeSubmittedWorks** | Pointer to **bool** | If set to true, text matched to the submitted works collection will not count as similar content calculated as if submitted work was not part of the paper  | [optional] 
**ExcludeCustomSections** | Pointer to **bool** | If set to true, text matched to the custom sections defined in the admin settings will not count as similar content calculated as if section was not part of the paper  | [optional] 
**ExcludePreprints** | Pointer to **bool** | If set to true, it will exclude preprints. A preprint is a version of a scholarly or scientific paper that precedes formal peer review and publication in a peer-reviewed scholarly or scientific journal.  | [optional] 

## Methods

### NewFeaturesSimilarityViewSettings

`func NewFeaturesSimilarityViewSettings() *FeaturesSimilarityViewSettings`

NewFeaturesSimilarityViewSettings instantiates a new FeaturesSimilarityViewSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturesSimilarityViewSettingsWithDefaults

`func NewFeaturesSimilarityViewSettingsWithDefaults() *FeaturesSimilarityViewSettings`

NewFeaturesSimilarityViewSettingsWithDefaults instantiates a new FeaturesSimilarityViewSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeQuotes

`func (o *FeaturesSimilarityViewSettings) GetExcludeQuotes() bool`

GetExcludeQuotes returns the ExcludeQuotes field if non-nil, zero value otherwise.

### GetExcludeQuotesOk

`func (o *FeaturesSimilarityViewSettings) GetExcludeQuotesOk() (*bool, bool)`

GetExcludeQuotesOk returns a tuple with the ExcludeQuotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeQuotes

`func (o *FeaturesSimilarityViewSettings) SetExcludeQuotes(v bool)`

SetExcludeQuotes sets ExcludeQuotes field to given value.

### HasExcludeQuotes

`func (o *FeaturesSimilarityViewSettings) HasExcludeQuotes() bool`

HasExcludeQuotes returns a boolean if a field has been set.

### GetExcludeBibliography

`func (o *FeaturesSimilarityViewSettings) GetExcludeBibliography() bool`

GetExcludeBibliography returns the ExcludeBibliography field if non-nil, zero value otherwise.

### GetExcludeBibliographyOk

`func (o *FeaturesSimilarityViewSettings) GetExcludeBibliographyOk() (*bool, bool)`

GetExcludeBibliographyOk returns a tuple with the ExcludeBibliography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBibliography

`func (o *FeaturesSimilarityViewSettings) SetExcludeBibliography(v bool)`

SetExcludeBibliography sets ExcludeBibliography field to given value.

### HasExcludeBibliography

`func (o *FeaturesSimilarityViewSettings) HasExcludeBibliography() bool`

HasExcludeBibliography returns a boolean if a field has been set.

### GetExcludeCitations

`func (o *FeaturesSimilarityViewSettings) GetExcludeCitations() bool`

GetExcludeCitations returns the ExcludeCitations field if non-nil, zero value otherwise.

### GetExcludeCitationsOk

`func (o *FeaturesSimilarityViewSettings) GetExcludeCitationsOk() (*bool, bool)`

GetExcludeCitationsOk returns a tuple with the ExcludeCitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCitations

`func (o *FeaturesSimilarityViewSettings) SetExcludeCitations(v bool)`

SetExcludeCitations sets ExcludeCitations field to given value.

### HasExcludeCitations

`func (o *FeaturesSimilarityViewSettings) HasExcludeCitations() bool`

HasExcludeCitations returns a boolean if a field has been set.

### GetExcludeAbstract

`func (o *FeaturesSimilarityViewSettings) GetExcludeAbstract() bool`

GetExcludeAbstract returns the ExcludeAbstract field if non-nil, zero value otherwise.

### GetExcludeAbstractOk

`func (o *FeaturesSimilarityViewSettings) GetExcludeAbstractOk() (*bool, bool)`

GetExcludeAbstractOk returns a tuple with the ExcludeAbstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAbstract

`func (o *FeaturesSimilarityViewSettings) SetExcludeAbstract(v bool)`

SetExcludeAbstract sets ExcludeAbstract field to given value.

### HasExcludeAbstract

`func (o *FeaturesSimilarityViewSettings) HasExcludeAbstract() bool`

HasExcludeAbstract returns a boolean if a field has been set.

### GetExcludeMethods

`func (o *FeaturesSimilarityViewSettings) GetExcludeMethods() bool`

GetExcludeMethods returns the ExcludeMethods field if non-nil, zero value otherwise.

### GetExcludeMethodsOk

`func (o *FeaturesSimilarityViewSettings) GetExcludeMethodsOk() (*bool, bool)`

GetExcludeMethodsOk returns a tuple with the ExcludeMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeMethods

`func (o *FeaturesSimilarityViewSettings) SetExcludeMethods(v bool)`

SetExcludeMethods sets ExcludeMethods field to given value.

### HasExcludeMethods

`func (o *FeaturesSimilarityViewSettings) HasExcludeMethods() bool`

HasExcludeMethods returns a boolean if a field has been set.

### GetExcludeSmallMatches

`func (o *FeaturesSimilarityViewSettings) GetExcludeSmallMatches() bool`

GetExcludeSmallMatches returns the ExcludeSmallMatches field if non-nil, zero value otherwise.

### GetExcludeSmallMatchesOk

`func (o *FeaturesSimilarityViewSettings) GetExcludeSmallMatchesOk() (*bool, bool)`

GetExcludeSmallMatchesOk returns a tuple with the ExcludeSmallMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSmallMatches

`func (o *FeaturesSimilarityViewSettings) SetExcludeSmallMatches(v bool)`

SetExcludeSmallMatches sets ExcludeSmallMatches field to given value.

### HasExcludeSmallMatches

`func (o *FeaturesSimilarityViewSettings) HasExcludeSmallMatches() bool`

HasExcludeSmallMatches returns a boolean if a field has been set.

### GetExcludeInternet

`func (o *FeaturesSimilarityViewSettings) GetExcludeInternet() bool`

GetExcludeInternet returns the ExcludeInternet field if non-nil, zero value otherwise.

### GetExcludeInternetOk

`func (o *FeaturesSimilarityViewSettings) GetExcludeInternetOk() (*bool, bool)`

GetExcludeInternetOk returns a tuple with the ExcludeInternet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeInternet

`func (o *FeaturesSimilarityViewSettings) SetExcludeInternet(v bool)`

SetExcludeInternet sets ExcludeInternet field to given value.

### HasExcludeInternet

`func (o *FeaturesSimilarityViewSettings) HasExcludeInternet() bool`

HasExcludeInternet returns a boolean if a field has been set.

### GetExcludePublications

`func (o *FeaturesSimilarityViewSettings) GetExcludePublications() bool`

GetExcludePublications returns the ExcludePublications field if non-nil, zero value otherwise.

### GetExcludePublicationsOk

`func (o *FeaturesSimilarityViewSettings) GetExcludePublicationsOk() (*bool, bool)`

GetExcludePublicationsOk returns a tuple with the ExcludePublications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePublications

`func (o *FeaturesSimilarityViewSettings) SetExcludePublications(v bool)`

SetExcludePublications sets ExcludePublications field to given value.

### HasExcludePublications

`func (o *FeaturesSimilarityViewSettings) HasExcludePublications() bool`

HasExcludePublications returns a boolean if a field has been set.

### GetExcludeCrossref

`func (o *FeaturesSimilarityViewSettings) GetExcludeCrossref() bool`

GetExcludeCrossref returns the ExcludeCrossref field if non-nil, zero value otherwise.

### GetExcludeCrossrefOk

`func (o *FeaturesSimilarityViewSettings) GetExcludeCrossrefOk() (*bool, bool)`

GetExcludeCrossrefOk returns a tuple with the ExcludeCrossref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCrossref

`func (o *FeaturesSimilarityViewSettings) SetExcludeCrossref(v bool)`

SetExcludeCrossref sets ExcludeCrossref field to given value.

### HasExcludeCrossref

`func (o *FeaturesSimilarityViewSettings) HasExcludeCrossref() bool`

HasExcludeCrossref returns a boolean if a field has been set.

### GetExcludeCrossrefPostedContent

`func (o *FeaturesSimilarityViewSettings) GetExcludeCrossrefPostedContent() bool`

GetExcludeCrossrefPostedContent returns the ExcludeCrossrefPostedContent field if non-nil, zero value otherwise.

### GetExcludeCrossrefPostedContentOk

`func (o *FeaturesSimilarityViewSettings) GetExcludeCrossrefPostedContentOk() (*bool, bool)`

GetExcludeCrossrefPostedContentOk returns a tuple with the ExcludeCrossrefPostedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCrossrefPostedContent

`func (o *FeaturesSimilarityViewSettings) SetExcludeCrossrefPostedContent(v bool)`

SetExcludeCrossrefPostedContent sets ExcludeCrossrefPostedContent field to given value.

### HasExcludeCrossrefPostedContent

`func (o *FeaturesSimilarityViewSettings) HasExcludeCrossrefPostedContent() bool`

HasExcludeCrossrefPostedContent returns a boolean if a field has been set.

### GetExcludeSubmittedWorks

`func (o *FeaturesSimilarityViewSettings) GetExcludeSubmittedWorks() bool`

GetExcludeSubmittedWorks returns the ExcludeSubmittedWorks field if non-nil, zero value otherwise.

### GetExcludeSubmittedWorksOk

`func (o *FeaturesSimilarityViewSettings) GetExcludeSubmittedWorksOk() (*bool, bool)`

GetExcludeSubmittedWorksOk returns a tuple with the ExcludeSubmittedWorks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSubmittedWorks

`func (o *FeaturesSimilarityViewSettings) SetExcludeSubmittedWorks(v bool)`

SetExcludeSubmittedWorks sets ExcludeSubmittedWorks field to given value.

### HasExcludeSubmittedWorks

`func (o *FeaturesSimilarityViewSettings) HasExcludeSubmittedWorks() bool`

HasExcludeSubmittedWorks returns a boolean if a field has been set.

### GetExcludeCustomSections

`func (o *FeaturesSimilarityViewSettings) GetExcludeCustomSections() bool`

GetExcludeCustomSections returns the ExcludeCustomSections field if non-nil, zero value otherwise.

### GetExcludeCustomSectionsOk

`func (o *FeaturesSimilarityViewSettings) GetExcludeCustomSectionsOk() (*bool, bool)`

GetExcludeCustomSectionsOk returns a tuple with the ExcludeCustomSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCustomSections

`func (o *FeaturesSimilarityViewSettings) SetExcludeCustomSections(v bool)`

SetExcludeCustomSections sets ExcludeCustomSections field to given value.

### HasExcludeCustomSections

`func (o *FeaturesSimilarityViewSettings) HasExcludeCustomSections() bool`

HasExcludeCustomSections returns a boolean if a field has been set.

### GetExcludePreprints

`func (o *FeaturesSimilarityViewSettings) GetExcludePreprints() bool`

GetExcludePreprints returns the ExcludePreprints field if non-nil, zero value otherwise.

### GetExcludePreprintsOk

`func (o *FeaturesSimilarityViewSettings) GetExcludePreprintsOk() (*bool, bool)`

GetExcludePreprintsOk returns a tuple with the ExcludePreprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePreprints

`func (o *FeaturesSimilarityViewSettings) SetExcludePreprints(v bool)`

SetExcludePreprints sets ExcludePreprints field to given value.

### HasExcludePreprints

`func (o *FeaturesSimilarityViewSettings) HasExcludePreprints() bool`

HasExcludePreprints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


