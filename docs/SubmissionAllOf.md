# SubmissionAllOf

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

### NewSubmissionAllOf

`func NewSubmissionAllOf() *SubmissionAllOf`

NewSubmissionAllOf instantiates a new SubmissionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionAllOfWithDefaults

`func NewSubmissionAllOfWithDefaults() *SubmissionAllOf`

NewSubmissionAllOfWithDefaults instantiates a new SubmissionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *SubmissionAllOf) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SubmissionAllOf) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SubmissionAllOf) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SubmissionAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTitle

`func (o *SubmissionAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SubmissionAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SubmissionAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SubmissionAllOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *SubmissionAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmissionAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmissionAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubmissionAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetId

`func (o *SubmissionAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubmissionAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubmissionAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubmissionAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContentType

`func (o *SubmissionAllOf) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *SubmissionAllOf) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *SubmissionAllOf) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *SubmissionAllOf) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetPageCount

`func (o *SubmissionAllOf) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *SubmissionAllOf) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *SubmissionAllOf) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *SubmissionAllOf) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetWordCount

`func (o *SubmissionAllOf) GetWordCount() int32`

GetWordCount returns the WordCount field if non-nil, zero value otherwise.

### GetWordCountOk

`func (o *SubmissionAllOf) GetWordCountOk() (*int32, bool)`

GetWordCountOk returns a tuple with the WordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordCount

`func (o *SubmissionAllOf) SetWordCount(v int32)`

SetWordCount sets WordCount field to given value.

### HasWordCount

`func (o *SubmissionAllOf) HasWordCount() bool`

HasWordCount returns a boolean if a field has been set.

### GetCharacterCount

`func (o *SubmissionAllOf) GetCharacterCount() int32`

GetCharacterCount returns the CharacterCount field if non-nil, zero value otherwise.

### GetCharacterCountOk

`func (o *SubmissionAllOf) GetCharacterCountOk() (*int32, bool)`

GetCharacterCountOk returns a tuple with the CharacterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterCount

`func (o *SubmissionAllOf) SetCharacterCount(v int32)`

SetCharacterCount sets CharacterCount field to given value.

### HasCharacterCount

`func (o *SubmissionAllOf) HasCharacterCount() bool`

HasCharacterCount returns a boolean if a field has been set.

### GetErrorCode

`func (o *SubmissionAllOf) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SubmissionAllOf) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *SubmissionAllOf) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *SubmissionAllOf) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetCreatedTime

`func (o *SubmissionAllOf) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SubmissionAllOf) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SubmissionAllOf) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *SubmissionAllOf) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCapabilities

`func (o *SubmissionAllOf) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *SubmissionAllOf) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *SubmissionAllOf) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *SubmissionAllOf) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


