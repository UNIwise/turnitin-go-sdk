# Sidebar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMode** | Pointer to **string** | The default mode shown in the sidebar panel. The selected mode must be enabled in the tenant&#39;s license, request parameter, as well as the viewer&#39;s permission. The default mode is similarity.  | [optional] 

## Methods

### NewSidebar

`func NewSidebar() *Sidebar`

NewSidebar instantiates a new Sidebar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSidebarWithDefaults

`func NewSidebarWithDefaults() *Sidebar`

NewSidebarWithDefaults instantiates a new Sidebar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMode

`func (o *Sidebar) GetDefaultMode() string`

GetDefaultMode returns the DefaultMode field if non-nil, zero value otherwise.

### GetDefaultModeOk

`func (o *Sidebar) GetDefaultModeOk() (*string, bool)`

GetDefaultModeOk returns a tuple with the DefaultMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMode

`func (o *Sidebar) SetDefaultMode(v string)`

SetDefaultMode sets DefaultMode field to given value.

### HasDefaultMode

`func (o *Sidebar) HasDefaultMode() bool`

HasDefaultMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


