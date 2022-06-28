# AggregateGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | (required)  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**GroupContext** | Pointer to [**GroupContext**](GroupContext.md) |  | [optional] 
**DueDate** | Pointer to **string** | due date for the group | [optional] 
**ReportGeneration** | Pointer to **string** |  | [optional] 

## Methods

### NewAggregateGroup

`func NewAggregateGroup() *AggregateGroup`

NewAggregateGroup instantiates a new AggregateGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateGroupWithDefaults

`func NewAggregateGroupWithDefaults() *AggregateGroup`

NewAggregateGroupWithDefaults instantiates a new AggregateGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AggregateGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AggregateGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AggregateGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AggregateGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AggregateGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AggregateGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AggregateGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AggregateGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AggregateGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AggregateGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AggregateGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AggregateGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGroupContext

`func (o *AggregateGroup) GetGroupContext() GroupContext`

GetGroupContext returns the GroupContext field if non-nil, zero value otherwise.

### GetGroupContextOk

`func (o *AggregateGroup) GetGroupContextOk() (*GroupContext, bool)`

GetGroupContextOk returns a tuple with the GroupContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupContext

`func (o *AggregateGroup) SetGroupContext(v GroupContext)`

SetGroupContext sets GroupContext field to given value.

### HasGroupContext

`func (o *AggregateGroup) HasGroupContext() bool`

HasGroupContext returns a boolean if a field has been set.

### GetDueDate

`func (o *AggregateGroup) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *AggregateGroup) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *AggregateGroup) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *AggregateGroup) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetReportGeneration

`func (o *AggregateGroup) GetReportGeneration() string`

GetReportGeneration returns the ReportGeneration field if non-nil, zero value otherwise.

### GetReportGenerationOk

`func (o *AggregateGroup) GetReportGenerationOk() (*string, bool)`

GetReportGenerationOk returns a tuple with the ReportGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportGeneration

`func (o *AggregateGroup) SetReportGeneration(v string)`

SetReportGeneration sets ReportGeneration field to given value.

### HasReportGeneration

`func (o *AggregateGroup) HasReportGeneration() bool`

HasReportGeneration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


