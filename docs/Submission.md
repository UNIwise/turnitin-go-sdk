# Submission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **string** | the owner of the submission | [optional] 
**Title** | Pointer to **string** | the title of the submission | [optional] 
**Status** | Pointer to **string** | the current status of the Submission | [optional] 
**Id** | Pointer to **string** | the unique ID of the submission | [optional] 
**ContentType** | Pointer to **string** | the content type of the submission | [optional] 
**PageCount** | Pointer to **int32** | the number of pages in the submission | [optional] 
**WordCount** | Pointer to **int32** | the number of words in the submission | [optional] 
**CharacterCount** | Pointer to **int32** | the number of characters in the submission | [optional] 
**ErrorCode** | Pointer to **string** | an error code representing the type of error encountered (if applicable)  | [optional] 
**CreatedTime** | Pointer to **time.Time** | RFC3339 timestamp of when this submission was initially created. This is the time at which the POST to /submissions was made.  | [optional] 
**Capabilities** | Pointer to **[]string** | Set of capabilities available to the current submission | [optional] 

## Methods

### NewSubmission

`func NewSubmission() *Submission`

NewSubmission instantiates a new Submission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionWithDefaults

`func NewSubmissionWithDefaults() *Submission`

NewSubmissionWithDefaults instantiates a new Submission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *Submission) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Submission) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Submission) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Submission) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTitle

`func (o *Submission) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Submission) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Submission) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Submission) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *Submission) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Submission) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Submission) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Submission) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetId

`func (o *Submission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Submission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Submission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Submission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContentType

`func (o *Submission) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Submission) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *Submission) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *Submission) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetPageCount

`func (o *Submission) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *Submission) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *Submission) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *Submission) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetWordCount

`func (o *Submission) GetWordCount() int32`

GetWordCount returns the WordCount field if non-nil, zero value otherwise.

### GetWordCountOk

`func (o *Submission) GetWordCountOk() (*int32, bool)`

GetWordCountOk returns a tuple with the WordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordCount

`func (o *Submission) SetWordCount(v int32)`

SetWordCount sets WordCount field to given value.

### HasWordCount

`func (o *Submission) HasWordCount() bool`

HasWordCount returns a boolean if a field has been set.

### GetCharacterCount

`func (o *Submission) GetCharacterCount() int32`

GetCharacterCount returns the CharacterCount field if non-nil, zero value otherwise.

### GetCharacterCountOk

`func (o *Submission) GetCharacterCountOk() (*int32, bool)`

GetCharacterCountOk returns a tuple with the CharacterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterCount

`func (o *Submission) SetCharacterCount(v int32)`

SetCharacterCount sets CharacterCount field to given value.

### HasCharacterCount

`func (o *Submission) HasCharacterCount() bool`

HasCharacterCount returns a boolean if a field has been set.

### GetErrorCode

`func (o *Submission) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Submission) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Submission) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Submission) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Submission) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Submission) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Submission) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Submission) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCapabilities

`func (o *Submission) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Submission) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Submission) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Submission) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


