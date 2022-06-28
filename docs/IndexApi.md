# \IndexApi

All URIs are relative to *https://app-us.turnitin.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIndexStatus**](IndexApi.md#GetIndexStatus) | **Get** /submissions/{id}/index | Return index state of submission
[**IndexSubmission**](IndexApi.md#IndexSubmission) | **Put** /submissions/{id}/index | Index and return index state of submission



## GetIndexStatus

> IndexStateSettings GetIndexStatus(ctx, id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

Return index state of submission

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
    resp, r, err := apiClient.IndexApi.GetIndexStatus(context.Background(), id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexApi.GetIndexStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndexStatus`: IndexStateSettings
    fmt.Fprintf(os.Stdout, "Response from `IndexApi.GetIndexStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Submission ID (returned upon a successful POST to /submissions)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 


### Return type

[**IndexStateSettings**](IndexStateSettings.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexSubmission

> IndexStateSettings IndexSubmission(ctx, id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

Index and return index state of submission

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
    resp, r, err := apiClient.IndexApi.IndexSubmission(context.Background(), id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexApi.IndexSubmission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IndexSubmission`: IndexStateSettings
    fmt.Fprintf(os.Stdout, "Response from `IndexApi.IndexSubmission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Submission ID (returned upon a successful POST to /submissions)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIndexSubmissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 


### Return type

[**IndexStateSettings**](IndexStateSettings.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

