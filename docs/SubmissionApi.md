# \SubmissionApi

All URIs are relative to *https://app-us.turnitin.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubmission**](SubmissionApi.md#CreateSubmission) | **Post** /submissions | Create a new Submission
[**DeleteSubmission**](SubmissionApi.md#DeleteSubmission) | **Delete** /submissions/{id} | Deletes a submission and associated similarity report.
[**GetSubmiddionDetails**](SubmissionApi.md#GetSubmiddionDetails) | **Get** /submissions/{id} | Get Submission Details
[**RecoverSubmission**](SubmissionApi.md#RecoverSubmission) | **Put** /submissions/{id}/recover | Recover a soft deleted submission
[**UploadSubmittedFile**](SubmissionApi.md#UploadSubmittedFile) | **Put** /submissions/{id}/original | Upload Submitted File



## CreateSubmission

> SimpleSubmissionResponse CreateSubmission(ctx).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()

Create a new Submission

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xTurnitinIntegrationName := "myintegration" // string | a human readable string representing the type of integration being used
    xTurnitinIntegrationVersion := "v1.0.2" // string | the version of the integration platform being used
    data := *openapiclient.NewSubmissionBase() // SubmissionBase | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubmissionApi.CreateSubmission(context.Background()).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubmissionApi.CreateSubmission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubmission`: SimpleSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SubmissionApi.CreateSubmission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubmissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 
 **data** | [**SubmissionBase**](SubmissionBase.md) |  | 

### Return type

[**SimpleSubmissionResponse**](SimpleSubmissionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubmission

> SuccessMessage DeleteSubmission(ctx, id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Hard(hard).Execute()

Deletes a submission and associated similarity report.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xTurnitinIntegrationName := "myintegration" // string | a human readable string representing the type of integration being used
    xTurnitinIntegrationVersion := "v1.0.2" // string | the version of the integration platform being used
    id := "id_example" // string | The Submission ID (returned upon a successful POST to /submissions) 
    hard := "hard_example" // string | Accepts true or false indicating either hard or soft delete. A soft delete removes the associated submission report/index and changes the saved submission state to DELETED. A hard deletion completely removes the submission information from Panda and TCA and removes the associated report/index, which can not be recovered.  (optional) (default to "false")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubmissionApi.DeleteSubmission(context.Background(), id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Hard(hard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubmissionApi.DeleteSubmission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubmission`: SuccessMessage
    fmt.Fprintf(os.Stdout, "Response from `SubmissionApi.DeleteSubmission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Submission ID (returned upon a successful POST to /submissions)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubmissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 

 **hard** | **string** | Accepts true or false indicating either hard or soft delete. A soft delete removes the associated submission report/index and changes the saved submission state to DELETED. A hard deletion completely removes the submission information from Panda and TCA and removes the associated report/index, which can not be recovered.  | [default to &quot;false&quot;]

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubmiddionDetails

> Submission GetSubmiddionDetails(ctx, id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

Get Submission Details

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xTurnitinIntegrationName := "myintegration" // string | a human readable string representing the type of integration being used
    xTurnitinIntegrationVersion := "v1.0.2" // string | the version of the integration platform being used
    id := "id_example" // string | The Submission ID (returned upon a successful POST to /submissions) 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubmissionApi.GetSubmiddionDetails(context.Background(), id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubmissionApi.GetSubmiddionDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubmiddionDetails`: Submission
    fmt.Fprintf(os.Stdout, "Response from `SubmissionApi.GetSubmiddionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Submission ID (returned upon a successful POST to /submissions)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubmiddionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 


### Return type

[**Submission**](Submission.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoverSubmission

> SimpleSubmissionResponse RecoverSubmission(ctx, id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

Recover a soft deleted submission

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xTurnitinIntegrationName := "myintegration" // string | a human readable string representing the type of integration being used
    xTurnitinIntegrationVersion := "v1.0.2" // string | the version of the integration platform being used
    id := "id_example" // string | The Submission ID (returned upon a successful POST to /submissions) 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubmissionApi.RecoverSubmission(context.Background(), id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubmissionApi.RecoverSubmission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoverSubmission`: SimpleSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SubmissionApi.RecoverSubmission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Submission ID (returned upon a successful POST to /submissions)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecoverSubmissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 


### Return type

[**SimpleSubmissionResponse**](SimpleSubmissionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadSubmittedFile

> SuccessMessage UploadSubmittedFile(ctx, id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).ContentType(contentType).ContentDisposition(contentDisposition).File(file).Execute()

Upload Submitted File

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xTurnitinIntegrationName := "myintegration" // string | a human readable string representing the type of integration being used
    xTurnitinIntegrationVersion := "v1.0.2" // string | the version of the integration platform being used
    id := "id_example" // string | The Submission ID (returned upon a successful POST to /submissions) 
    contentType := "contentType_example" // string | *Must be 'binary/octet-stream' 
    contentDisposition := "inline; filename="MyFile.docx"" // string | *must include the \"filename\" parameter, e.g. `inline; filename=\"MyFile.docx\"`. To support UTF-8 filenames, you must URL encode the header 
    file := map[string]interface{}{ ... } // map[string]interface{} | the user's submitted file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubmissionApi.UploadSubmittedFile(context.Background(), id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).ContentType(contentType).ContentDisposition(contentDisposition).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubmissionApi.UploadSubmittedFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadSubmittedFile`: SuccessMessage
    fmt.Fprintf(os.Stdout, "Response from `SubmissionApi.UploadSubmittedFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Submission ID (returned upon a successful POST to /submissions)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadSubmittedFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 

 **contentType** | **string** | *Must be &#39;binary/octet-stream&#39;  | 
 **contentDisposition** | **string** | *must include the \&quot;filename\&quot; parameter, e.g. &#x60;inline; filename&#x3D;\&quot;MyFile.docx\&quot;&#x60;. To support UTF-8 filenames, you must URL encode the header  | 
 **file** | **map[string]interface{}** | the user&#39;s submitted file | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: binary/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

