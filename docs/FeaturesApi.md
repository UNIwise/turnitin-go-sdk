# \FeaturesApi

All URIs are relative to *https://app-us.turnitin.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FeaturesEnabledGet**](FeaturesApi.md#FeaturesEnabledGet) | **Get** /features-enabled | Get information about what features are allowed by the current license



## FeaturesEnabledGet

> FeaturesEnabled FeaturesEnabledGet(ctx).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

Get information about what features are allowed by the current license

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.FeaturesEnabledGet(context.Background()).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.FeaturesEnabledGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FeaturesEnabledGet`: FeaturesEnabled
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.FeaturesEnabledGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFeaturesEnabledGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 

### Return type

[**FeaturesEnabled**](FeaturesEnabled.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

