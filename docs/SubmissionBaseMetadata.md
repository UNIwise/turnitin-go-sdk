# SubmissionBaseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Submitter** | Pointer to [**Users**](Users.md) |  | [optional] 
**Owners** | Pointer to [**[]Users**](Users.md) |  | [optional] 
**Group** | Pointer to [**Group**](Group.md) |  | [optional] 
**GroupContext** | Pointer to [**GroupContext**](GroupContext.md) |  | [optional] 
**OriginalSubmittedTime** | Pointer to **time.Time** | Optional original submision date | [optional] 
**Custom** | Pointer to **string** | custom metadata | [optional] 

## Methods

### NewSubmissionBaseMetadata

`func NewSubmissionBaseMetadata() *SubmissionBaseMetadata`

NewSubmissionBaseMetadata instantiates a new SubmissionBaseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionBaseMetadataWithDefaults

`func NewSubmissionBaseMetadataWithDefaults() *SubmissionBaseMetadata`

NewSubmissionBaseMetadataWithDefaults instantiates a new SubmissionBaseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubmitter

`func (o *SubmissionBaseMetadata) GetSubmitter() Users`

GetSubmitter returns the Submitter field if non-nil, zero value otherwise.

### GetSubmitterOk

`func (o *SubmissionBaseMetadata) GetSubmitterOk() (*Users, bool)`

GetSubmitterOk returns a tuple with the Submitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitter

`func (o *SubmissionBaseMetadata) SetSubmitter(v Users)`

SetSubmitter sets Submitter field to given value.

### HasSubmitter

`func (o *SubmissionBaseMetadata) HasSubmitter() bool`

HasSubmitter returns a boolean if a field has been set.

### GetOwners

`func (o *SubmissionBaseMetadata) GetOwners() []Users`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *SubmissionBaseMetadata) GetOwnersOk() (*[]Users, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *SubmissionBaseMetadata) SetOwners(v []Users)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *SubmissionBaseMetadata) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetGroup

`func (o *SubmissionBaseMetadata) GetGroup() Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SubmissionBaseMetadata) GetGroupOk() (*Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SubmissionBaseMetadata) SetGroup(v Group)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *SubmissionBaseMetadata) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupContext

`func (o *SubmissionBaseMetadata) GetGroupContext() GroupContext`

GetGroupContext returns the GroupContext field if non-nil, zero value otherwise.

### GetGroupContextOk

`func (o *SubmissionBaseMetadata) GetGroupContextOk() (*GroupContext, bool)`

GetGroupContextOk returns a tuple with the GroupContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupContext

`func (o *SubmissionBaseMetadata) SetGroupContext(v GroupContext)`

SetGroupContext sets GroupContext field to given value.

### HasGroupContext

`func (o *SubmissionBaseMetadata) HasGroupContext() bool`

HasGroupContext returns a boolean if a field has been set.

### GetOriginalSubmittedTime

`func (o *SubmissionBaseMetadata) GetOriginalSubmittedTime() time.Time`

GetOriginalSubmittedTime returns the OriginalSubmittedTime field if non-nil, zero value otherwise.

### GetOriginalSubmittedTimeOk

`func (o *SubmissionBaseMetadata) GetOriginalSubmittedTimeOk() (*time.Time, bool)`

GetOriginalSubmittedTimeOk returns a tuple with the OriginalSubmittedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSubmittedTime

`func (o *SubmissionBaseMetadata) SetOriginalSubmittedTime(v time.Time)`

SetOriginalSubmittedTime sets OriginalSubmittedTime field to given value.

### HasOriginalSubmittedTime

`func (o *SubmissionBaseMetadata) HasOriginalSubmittedTime() bool`

HasOriginalSubmittedTime returns a boolean if a field has been set.

### GetCustom

`func (o *SubmissionBaseMetadata) GetCustom() string`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *SubmissionBaseMetadata) GetCustomOk() (*string, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *SubmissionBaseMetadata) SetCustom(v string)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *SubmissionBaseMetadata) HasCustom() bool`

HasCustom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


