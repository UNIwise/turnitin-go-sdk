# AddGroupAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Attachment title | [optional] 
**Template** | Pointer to **bool** | template | [optional] [default to false]

## Methods

### NewAddGroupAttachmentRequest

`func NewAddGroupAttachmentRequest() *AddGroupAttachmentRequest`

NewAddGroupAttachmentRequest instantiates a new AddGroupAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroupAttachmentRequestWithDefaults

`func NewAddGroupAttachmentRequestWithDefaults() *AddGroupAttachmentRequest`

NewAddGroupAttachmentRequestWithDefaults instantiates a new AddGroupAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AddGroupAttachmentRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AddGroupAttachmentRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AddGroupAttachmentRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AddGroupAttachmentRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTemplate

`func (o *AddGroupAttachmentRequest) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *AddGroupAttachmentRequest) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *AddGroupAttachmentRequest) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *AddGroupAttachmentRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


