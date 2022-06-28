# \WebhookApi

All URIs are relative to *https://app-us.turnitin.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhook**](WebhookApi.md#DeleteWebhook) | **Delete** /webhooks/{id} | Delete Webhook endpoint
[**GetWebhook**](WebhookApi.md#GetWebhook) | **Get** /webhooks/{id} | Get Webhook endpoint info
[**UpdateWebhook**](WebhookApi.md#UpdateWebhook) | **Patch** /webhooks/{id} | Update Webhook endpoint info
[**WebhooksGet**](WebhookApi.md#WebhooksGet) | **Get** /webhooks | List registered Webhook endpoints
[**WebhooksPost**](WebhookApi.md#WebhooksPost) | **Post** /webhooks | Register a Webhook endpoint



## DeleteWebhook

> DeleteWebhook(ctx, id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

Delete Webhook endpoint

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.DeleteWebhook(context.Background(), id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.DeleteWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> Webhook GetWebhook(ctx, id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

Get Webhook endpoint info

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.GetWebhook(context.Background(), id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.GetWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 


### Return type

[**Webhook**](Webhook.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> Webhook UpdateWebhook(ctx, id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()

Update Webhook endpoint info

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
    id := "id_example" // string | 
    data := *openapiclient.NewWebhookPathRequest([]string{"EventTypes_example"}) // WebhookPathRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.UpdateWebhook(context.Background(), id).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.UpdateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 

 **data** | [**WebhookPathRequest**](WebhookPathRequest.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksGet

> []Webhook WebhooksGet(ctx).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

List registered Webhook endpoints

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
    resp, r, err := apiClient.WebhookApi.WebhooksGet(context.Background()).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.WebhooksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhooksGet`: []Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.WebhooksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 

### Return type

[**[]Webhook**](Webhook.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksPost

> Webhook WebhooksPost(ctx).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()

Register a Webhook endpoint

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
    data := *openapiclient.NewWebhookWithSecret("TODO", "Url_example", []string{"EventTypes_example"}) // WebhookWithSecret | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.WebhooksPost(context.Background()).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.WebhooksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhooksPost`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.WebhooksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 
 **data** | [**WebhookWithSecret**](WebhookWithSecret.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

