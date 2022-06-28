# \SimilarityApi

All URIs are relative to *https://app-us.turnitin.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadSimilarityReportPdf**](SimilarityApi.md#DownloadSimilarityReportPdf) | **Get** /submissions/{id}/similarity/pdf/{pdf_id} | GET download pdf
[**GetSimilarityReportPdfStatus**](SimilarityApi.md#GetSimilarityReportPdfStatus) | **Get** /submissions/{id}/similarity/pdf/{pdf_id}/status | GET pdf download status
[**GetSimilarityReportResults**](SimilarityApi.md#GetSimilarityReportResults) | **Get** /submissions/{id}/similarity | Get Similarity Report Results
[**GetSimilarityReportUrl**](SimilarityApi.md#GetSimilarityReportUrl) | **Post** /submissions/{id}/viewer-url | Returns a URL to access Cloud Viewer
[**RequestSimilarityReport**](SimilarityApi.md#RequestSimilarityReport) | **Put** /submissions/{id}/similarity | Request Similarity Report generation
[**RequestSimilarityReportPdf**](SimilarityApi.md#RequestSimilarityReportPdf) | **Post** /submissions/{id}/similarity/pdf | Request Pdf download and returns the Pdf Id



## DownloadSimilarityReportPdf

> *os.File DownloadSimilarityReportPdf(ctx, id, pdfId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

GET download pdf

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
    pdfId := "pdfId_example" // string | The Pdf ID (returned upon a successful POST to /submissions/{submission_id}/similarity/pdf) 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimilarityApi.DownloadSimilarityReportPdf(context.Background(), id, pdfId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimilarityApi.DownloadSimilarityReportPdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadSimilarityReportPdf`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SimilarityApi.DownloadSimilarityReportPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Submission ID (returned upon a successful POST to /submissions)  | 
**pdfId** | **string** | The Pdf ID (returned upon a successful POST to /submissions/{submission_id}/similarity/pdf)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadSimilarityReportPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 



### Return type

[***os.File**](*os.File.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimilarityReportPdfStatus

> PdfStatusResponse GetSimilarityReportPdfStatus(ctx, id, pdfId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

GET pdf download status

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
    pdfId := "pdfId_example" // string | The Pdf ID (returned upon a successful POST to /submissions/{submission_id}/similarity/pdf) 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimilarityApi.GetSimilarityReportPdfStatus(context.Background(), id, pdfId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimilarityApi.GetSimilarityReportPdfStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSimilarityReportPdfStatus`: PdfStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `SimilarityApi.GetSimilarityReportPdfStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Submission ID (returned upon a successful POST to /submissions)  | 
**pdfId** | **string** | The Pdf ID (returned upon a successful POST to /submissions/{submission_id}/similarity/pdf)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarityReportPdfStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 



### Return type

[**PdfStatusResponse**](PdfStatusResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimilarityReportResults

> SimilarityMetadata GetSimilarityReportResults(ctx, id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

Get Similarity Report Results

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
    resp, r, err := apiClient.SimilarityApi.GetSimilarityReportResults(context.Background(), id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimilarityApi.GetSimilarityReportResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSimilarityReportResults`: SimilarityMetadata
    fmt.Fprintf(os.Stdout, "Response from `SimilarityApi.GetSimilarityReportResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Submission ID (returned upon a successful POST to /submissions)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarityReportResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 


### Return type

[**SimilarityMetadata**](SimilarityMetadata.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimilarityReportUrl

> SimilarityViewerUrlResponse GetSimilarityReportUrl(ctx, id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()

Returns a URL to access Cloud Viewer

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
    data := *openapiclient.NewSimilarityViewerUrlSettings() // SimilarityViewerUrlSettings | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimilarityApi.GetSimilarityReportUrl(context.Background(), id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimilarityApi.GetSimilarityReportUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSimilarityReportUrl`: SimilarityViewerUrlResponse
    fmt.Fprintf(os.Stdout, "Response from `SimilarityApi.GetSimilarityReportUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Submission ID (returned upon a successful POST to /submissions)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarityReportUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 

 **data** | [**SimilarityViewerUrlSettings**](SimilarityViewerUrlSettings.md) |  | 

### Return type

[**SimilarityViewerUrlResponse**](SimilarityViewerUrlResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestSimilarityReport

> SuccessMessage RequestSimilarityReport(ctx, id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()

Request Similarity Report generation

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
    data := *openapiclient.NewSimilarityPutRequest(*openapiclient.NewSimilarityGenerationSettings([]string{"SearchRepositories_example"})) // SimilarityPutRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimilarityApi.RequestSimilarityReport(context.Background(), id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimilarityApi.RequestSimilarityReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestSimilarityReport`: SuccessMessage
    fmt.Fprintf(os.Stdout, "Response from `SimilarityApi.RequestSimilarityReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Submission ID (returned upon a successful POST to /submissions)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestSimilarityReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 

 **data** | [**SimilarityPutRequest**](SimilarityPutRequest.md) |  | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestSimilarityReportPdf

> RequestPdfResponse RequestSimilarityReportPdf(ctx, id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

Request Pdf download and returns the Pdf Id

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
    resp, r, err := apiClient.SimilarityApi.RequestSimilarityReportPdf(context.Background(), id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimilarityApi.RequestSimilarityReportPdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestSimilarityReportPdf`: RequestPdfResponse
    fmt.Fprintf(os.Stdout, "Response from `SimilarityApi.RequestSimilarityReportPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Submission ID (returned upon a successful POST to /submissions)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestSimilarityReportPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 


### Return type

[**RequestPdfResponse**](RequestPdfResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

