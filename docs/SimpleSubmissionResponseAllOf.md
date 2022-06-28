# SimpleSubmissionResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | the unique ID of the submission | [optional] 
**Owner** | Pointer to **string** | the owner of the submission | [optional] 
**Title** | Pointer to **string** | the title of the submission | [optional] 
**Status** | Pointer to **string** | the current status of the Submission | [optional] 
**CreatedTime** | Pointer to **time.Time** | RFC3339 timestamp of when this submission was initially created. This is the time at which the POST to /submissions was made.  | [optional] 

## Methods

### NewSimpleSubmissionResponseAllOf

`func NewSimpleSubmissionResponseAllOf() *SimpleSubmissionResponseAllOf`

NewSimpleSubmissionResponseAllOf instantiates a new SimpleSubmissionResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleSubmissionResponseAllOfWithDefaults

`func NewSimpleSubmissionResponseAllOfWithDefaults() *SimpleSubmissionResponseAllOf`

NewSimpleSubmissionResponseAllOfWithDefaults instantiates a new SimpleSubmissionResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleSubmissionResponseAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleSubmissionResponseAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleSubmissionResponseAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimpleSubmissionResponseAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwner

`func (o *SimpleSubmissionResponseAllOf) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SimpleSubmissionResponseAllOf) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SimpleSubmissionResponseAllOf) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SimpleSubmissionResponseAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTitle

`func (o *SimpleSubmissionResponseAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SimpleSubmissionResponseAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SimpleSubmissionResponseAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SimpleSubmissionResponseAllOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *SimpleSubmissionResponseAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimpleSubmissionResponseAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimpleSubmissionResponseAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SimpleSubmissionResponseAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedTime

`func (o *SimpleSubmissionResponseAllOf) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SimpleSubmissionResponseAllOf) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SimpleSubmissionResponseAllOf) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *SimpleSubmissionResponseAllOf) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


