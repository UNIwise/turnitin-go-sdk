# IndexingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddToIndex** | Pointer to **bool** | If set, the submission will be added to all valid node groups for future matching | [optional] 

## Methods

### NewIndexingSettings

`func NewIndexingSettings() *IndexingSettings`

NewIndexingSettings instantiates a new IndexingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexingSettingsWithDefaults

`func NewIndexingSettingsWithDefaults() *IndexingSettings`

NewIndexingSettingsWithDefaults instantiates a new IndexingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddToIndex

`func (o *IndexingSettings) GetAddToIndex() bool`

GetAddToIndex returns the AddToIndex field if non-nil, zero value otherwise.

### GetAddToIndexOk

`func (o *IndexingSettings) GetAddToIndexOk() (*bool, bool)`

GetAddToIndexOk returns a tuple with the AddToIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToIndex

`func (o *IndexingSettings) SetAddToIndex(v bool)`

SetAddToIndex sets AddToIndex field to given value.

### HasAddToIndex

`func (o *IndexingSettings) HasAddToIndex() bool`

HasAddToIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


