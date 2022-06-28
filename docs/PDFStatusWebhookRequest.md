# PDFStatusWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The requesting Pdf status | [optional] 
**Id** | Pointer to **string** | the PDF ID | [optional] 
**SubmissionId** | Pointer to **string** | the associated submission ID | [optional] 
**Metadata** | Pointer to [**SubmissionCompleteWebhookRequestAllOfMetadata**](SubmissionCompleteWebhookRequestAllOfMetadata.md) |  | [optional] 

## Methods

### NewPDFStatusWebhookRequest

`func NewPDFStatusWebhookRequest() *PDFStatusWebhookRequest`

NewPDFStatusWebhookRequest instantiates a new PDFStatusWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPDFStatusWebhookRequestWithDefaults

`func NewPDFStatusWebhookRequestWithDefaults() *PDFStatusWebhookRequest`

NewPDFStatusWebhookRequestWithDefaults instantiates a new PDFStatusWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PDFStatusWebhookRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PDFStatusWebhookRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PDFStatusWebhookRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PDFStatusWebhookRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetId

`func (o *PDFStatusWebhookRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PDFStatusWebhookRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PDFStatusWebhookRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PDFStatusWebhookRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubmissionId

`func (o *PDFStatusWebhookRequest) GetSubmissionId() string`

GetSubmissionId returns the SubmissionId field if non-nil, zero value otherwise.

### GetSubmissionIdOk

`func (o *PDFStatusWebhookRequest) GetSubmissionIdOk() (*string, bool)`

GetSubmissionIdOk returns a tuple with the SubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionId

`func (o *PDFStatusWebhookRequest) SetSubmissionId(v string)`

SetSubmissionId sets SubmissionId field to given value.

### HasSubmissionId

`func (o *PDFStatusWebhookRequest) HasSubmissionId() bool`

HasSubmissionId returns a boolean if a field has been set.

### GetMetadata

`func (o *PDFStatusWebhookRequest) GetMetadata() SubmissionCompleteWebhookRequestAllOfMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PDFStatusWebhookRequest) GetMetadataOk() (*SubmissionCompleteWebhookRequestAllOfMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PDFStatusWebhookRequest) SetMetadata(v SubmissionCompleteWebhookRequestAllOfMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PDFStatusWebhookRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


