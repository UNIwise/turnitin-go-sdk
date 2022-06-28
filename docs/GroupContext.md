# GroupContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | (required)  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owners** | Pointer to [**[]Users**](Users.md) |  | [optional] 

## Methods

### NewGroupContext

`func NewGroupContext() *GroupContext`

NewGroupContext instantiates a new GroupContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupContextWithDefaults

`func NewGroupContextWithDefaults() *GroupContext`

NewGroupContextWithDefaults instantiates a new GroupContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupContext) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupContext) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GroupContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupContext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupContext) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwners

`func (o *GroupContext) GetOwners() []Users`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *GroupContext) GetOwnersOk() (*[]Users, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *GroupContext) SetOwners(v []Users)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *GroupContext) HasOwners() bool`

HasOwners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


