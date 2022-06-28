# \EULAApi

All URIs are relative to *https://app-us.turnitin.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EulaVersionIdAcceptPost**](EULAApi.md#EulaVersionIdAcceptPost) | **Post** /eula/{version_id}/accept | Accepts a particular EULA version on behalf of an external user
[**EulaVersionIdAcceptUserIdGet**](EULAApi.md#EulaVersionIdAcceptUserIdGet) | **Get** /eula/{version_id}/accept/{user_id} | Queries the acceptences of a particular EULA version on behalf of an external user
[**EulaVersionIdGet**](EULAApi.md#EulaVersionIdGet) | **Get** /eula/{version_id} | Gets information about a particular EULA version
[**EulaVersionIdViewGet**](EULAApi.md#EulaVersionIdViewGet) | **Get** /eula/{version_id}/view | Gets the text of a particular EULA version



## EulaVersionIdAcceptPost

> EulaAcceptListItem EulaVersionIdAcceptPost(ctx, versionId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()

Accepts a particular EULA version on behalf of an external user

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
    versionId := "latest" // string | The EULA version ID (or `latest`) 
    data := *openapiclient.NewEulaAcceptRequest() // EulaAcceptRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EULAApi.EulaVersionIdAcceptPost(context.Background(), versionId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EULAApi.EulaVersionIdAcceptPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EulaVersionIdAcceptPost`: EulaAcceptListItem
    fmt.Fprintf(os.Stdout, "Response from `EULAApi.EulaVersionIdAcceptPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | The EULA version ID (or &#x60;latest&#x60;)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEulaVersionIdAcceptPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 

 **data** | [**EulaAcceptRequest**](EulaAcceptRequest.md) |  | 

### Return type

[**EulaAcceptListItem**](EulaAcceptListItem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EulaVersionIdAcceptUserIdGet

> []EulaAcceptListItem EulaVersionIdAcceptUserIdGet(ctx, versionId, userId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

Queries the acceptences of a particular EULA version on behalf of an external user

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
    versionId := "latest" // string | The EULA version ID (or `latest`) 
    userId := "userId_example" // string | The user associated with the EULA status 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EULAApi.EulaVersionIdAcceptUserIdGet(context.Background(), versionId, userId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EULAApi.EulaVersionIdAcceptUserIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EulaVersionIdAcceptUserIdGet`: []EulaAcceptListItem
    fmt.Fprintf(os.Stdout, "Response from `EULAApi.EulaVersionIdAcceptUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | The EULA version ID (or &#x60;latest&#x60;)  | 
**userId** | **string** | The user associated with the EULA status  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEulaVersionIdAcceptUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 



### Return type

[**[]EulaAcceptListItem**](EulaAcceptListItem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EulaVersionIdGet

> EulaVersion EulaVersionIdGet(ctx, versionId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Lang(lang).Execute()

Gets information about a particular EULA version

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
    versionId := "latest" // string | The EULA version ID (or `latest`) 
    lang := "en-US" // string | The desired language of the specified EULA version (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EULAApi.EulaVersionIdGet(context.Background(), versionId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Lang(lang).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EULAApi.EulaVersionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EulaVersionIdGet`: EulaVersion
    fmt.Fprintf(os.Stdout, "Response from `EULAApi.EulaVersionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | The EULA version ID (or &#x60;latest&#x60;)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEulaVersionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 

 **lang** | **string** | The desired language of the specified EULA version | 

### Return type

[**EulaVersion**](EulaVersion.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EulaVersionIdViewGet

> string EulaVersionIdViewGet(ctx, versionId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Lang(lang).Execute()

Gets the text of a particular EULA version

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
    versionId := "latest" // string | The EULA version ID (or `latest`) 
    lang := "en-US" // string | The desired language of the specified EULA version (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EULAApi.EulaVersionIdViewGet(context.Background(), versionId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Lang(lang).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EULAApi.EulaVersionIdViewGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EulaVersionIdViewGet`: string
    fmt.Fprintf(os.Stdout, "Response from `EULAApi.EulaVersionIdViewGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | The EULA version ID (or &#x60;latest&#x60;)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEulaVersionIdViewGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 

 **lang** | **string** | The desired language of the specified EULA version | 

### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain; charset=utf-8, application/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

