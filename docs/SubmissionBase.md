# SubmissionBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **interface{}** | Submission id, optional field | [optional] 
**Owner** | Pointer to **string** | ID of the owning user | [optional] 
**OwnerDefaultPermissionSet** | Pointer to **string** | Default viewer permission set, accepts INSTRUCTOR, LEARNER, EDITOR, USER, APPLICANT, ADMINISTRATOR, UNDEFINED | [optional] 
**Title** | Pointer to **string** | the title of the submission | [optional] 
**Submitter** | Pointer to **string** | (optional) ID of the submitting user, if different from the owning user | [optional] 
**SubmitterDefaultPermissionSet** | Pointer to **string** | Default submitter permission set, accepts INSTRUCTOR, LEARNER, EDITOR, USER, APPLICANT, ADMINISTRATOR, UNDEFINED | [optional] 
**Eula** | Pointer to [**Eula**](Eula.md) |  | [optional] 
**Metadata** | Pointer to [**SubmissionBaseMetadata**](SubmissionBaseMetadata.md) |  | [optional] 
**ExtractTextOnly** | Pointer to **bool** | (optional) indicates if the submission should be treated as a text only submission. A text only submission cannot generate full reports or be viewed in the viewer, but can use the index only endpoint to be indexed | [optional] 

## Methods

### NewSubmissionBase

`func NewSubmissionBase() *SubmissionBase`

NewSubmissionBase instantiates a new SubmissionBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionBaseWithDefaults

`func NewSubmissionBaseWithDefaults() *SubmissionBase`

NewSubmissionBaseWithDefaults instantiates a new SubmissionBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubmissionBase) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubmissionBase) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubmissionBase) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *SubmissionBase) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SubmissionBase) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SubmissionBase) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOwner

`func (o *SubmissionBase) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SubmissionBase) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SubmissionBase) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SubmissionBase) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnerDefaultPermissionSet

`func (o *SubmissionBase) GetOwnerDefaultPermissionSet() string`

GetOwnerDefaultPermissionSet returns the OwnerDefaultPermissionSet field if non-nil, zero value otherwise.

### GetOwnerDefaultPermissionSetOk

`func (o *SubmissionBase) GetOwnerDefaultPermissionSetOk() (*string, bool)`

GetOwnerDefaultPermissionSetOk returns a tuple with the OwnerDefaultPermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerDefaultPermissionSet

`func (o *SubmissionBase) SetOwnerDefaultPermissionSet(v string)`

SetOwnerDefaultPermissionSet sets OwnerDefaultPermissionSet field to given value.

### HasOwnerDefaultPermissionSet

`func (o *SubmissionBase) HasOwnerDefaultPermissionSet() bool`

HasOwnerDefaultPermissionSet returns a boolean if a field has been set.

### GetTitle

`func (o *SubmissionBase) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SubmissionBase) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SubmissionBase) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SubmissionBase) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSubmitter

`func (o *SubmissionBase) GetSubmitter() string`

GetSubmitter returns the Submitter field if non-nil, zero value otherwise.

### GetSubmitterOk

`func (o *SubmissionBase) GetSubmitterOk() (*string, bool)`

GetSubmitterOk returns a tuple with the Submitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitter

`func (o *SubmissionBase) SetSubmitter(v string)`

SetSubmitter sets Submitter field to given value.

### HasSubmitter

`func (o *SubmissionBase) HasSubmitter() bool`

HasSubmitter returns a boolean if a field has been set.

### GetSubmitterDefaultPermissionSet

`func (o *SubmissionBase) GetSubmitterDefaultPermissionSet() string`

GetSubmitterDefaultPermissionSet returns the SubmitterDefaultPermissionSet field if non-nil, zero value otherwise.

### GetSubmitterDefaultPermissionSetOk

`func (o *SubmissionBase) GetSubmitterDefaultPermissionSetOk() (*string, bool)`

GetSubmitterDefaultPermissionSetOk returns a tuple with the SubmitterDefaultPermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterDefaultPermissionSet

`func (o *SubmissionBase) SetSubmitterDefaultPermissionSet(v string)`

SetSubmitterDefaultPermissionSet sets SubmitterDefaultPermissionSet field to given value.

### HasSubmitterDefaultPermissionSet

`func (o *SubmissionBase) HasSubmitterDefaultPermissionSet() bool`

HasSubmitterDefaultPermissionSet returns a boolean if a field has been set.

### GetEula

`func (o *SubmissionBase) GetEula() Eula`

GetEula returns the Eula field if non-nil, zero value otherwise.

### GetEulaOk

`func (o *SubmissionBase) GetEulaOk() (*Eula, bool)`

GetEulaOk returns a tuple with the Eula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEula

`func (o *SubmissionBase) SetEula(v Eula)`

SetEula sets Eula field to given value.

### HasEula

`func (o *SubmissionBase) HasEula() bool`

HasEula returns a boolean if a field has been set.

### GetMetadata

`func (o *SubmissionBase) GetMetadata() SubmissionBaseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubmissionBase) GetMetadataOk() (*SubmissionBaseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubmissionBase) SetMetadata(v SubmissionBaseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubmissionBase) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExtractTextOnly

`func (o *SubmissionBase) GetExtractTextOnly() bool`

GetExtractTextOnly returns the ExtractTextOnly field if non-nil, zero value otherwise.

### GetExtractTextOnlyOk

`func (o *SubmissionBase) GetExtractTextOnlyOk() (*bool, bool)`

GetExtractTextOnlyOk returns a tuple with the ExtractTextOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractTextOnly

`func (o *SubmissionBase) SetExtractTextOnly(v bool)`

SetExtractTextOnly sets ExtractTextOnly field to given value.

### HasExtractTextOnly

`func (o *SubmissionBase) HasExtractTextOnly() bool`

HasExtractTextOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


