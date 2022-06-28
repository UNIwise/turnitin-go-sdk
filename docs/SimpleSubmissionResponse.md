# SimpleSubmissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | the unique ID of the submission | [optional] 
**Owner** | Pointer to **string** | the owner of the submission | [optional] 
**Title** | Pointer to **string** | the title of the submission | [optional] 
**Status** | Pointer to **string** | the current status of the Submission | [optional] 
**CreatedTime** | Pointer to **time.Time** | RFC3339 timestamp of when this submission was initially created. This is the time at which the POST to /submissions was made.  | [optional] 

## Methods

### NewSimpleSubmissionResponse

`func NewSimpleSubmissionResponse() *SimpleSubmissionResponse`

NewSimpleSubmissionResponse instantiates a new SimpleSubmissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleSubmissionResponseWithDefaults

`func NewSimpleSubmissionResponseWithDefaults() *SimpleSubmissionResponse`

NewSimpleSubmissionResponseWithDefaults instantiates a new SimpleSubmissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleSubmissionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleSubmissionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleSubmissionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimpleSubmissionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwner

`func (o *SimpleSubmissionResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SimpleSubmissionResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SimpleSubmissionResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SimpleSubmissionResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTitle

`func (o *SimpleSubmissionResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SimpleSubmissionResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SimpleSubmissionResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SimpleSubmissionResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *SimpleSubmissionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimpleSubmissionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimpleSubmissionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SimpleSubmissionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedTime

`func (o *SimpleSubmissionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SimpleSubmissionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SimpleSubmissionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *SimpleSubmissionResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


