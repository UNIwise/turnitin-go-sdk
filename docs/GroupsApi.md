# \GroupsApi

All URIs are relative to *https://app-us.turnitin.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupAttachment**](GroupsApi.md#AddGroupAttachment) | **Post** /groups/{group_id}/attachments | Add attachment to a group. will create a group if it does not exist.
[**DeleteGroupAttachment**](GroupsApi.md#DeleteGroupAttachment) | **Delete** /groups/{group_id}/attachments/{attach_id} | Hard delete group attachment
[**GetGroup**](GroupsApi.md#GetGroup) | **Get** /groups/{group_id} | Get group, group context and group context owners info
[**GetGroupAttachment**](GroupsApi.md#GetGroupAttachment) | **Get** /groups/{group_id}/attachments/{attach_id} | Get group attachment
[**GetGroupAttachments**](GroupsApi.md#GetGroupAttachments) | **Get** /groups/{group_id}/attachments | Get all attachments
[**GroupsGroupIdPut**](GroupsApi.md#GroupsGroupIdPut) | **Put** /groups/{group_id} | upsert group, group context and group context owners info
[**UpdateGroupAttachment**](GroupsApi.md#UpdateGroupAttachment) | **Patch** /groups/{group_id}/attachments/{attach_id} | Patch a group attachment
[**UploadGroupAttachment**](GroupsApi.md#UploadGroupAttachment) | **Put** /groups/{group_id}/attachments/{attach_id}/original | Upload Submitted File



## AddGroupAttachment

> AddGroupAttachmentResponse AddGroupAttachment(ctx, groupId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()

Add attachment to a group. will create a group if it does not exist.

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
    groupId := "groupId_example" // string | group_id
    data := *openapiclient.NewAddGroupAttachmentRequest() // AddGroupAttachmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.AddGroupAttachment(context.Background(), groupId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.AddGroupAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGroupAttachment`: AddGroupAttachmentResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.AddGroupAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | group_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 

 **data** | [**AddGroupAttachmentRequest**](AddGroupAttachmentRequest.md) |  | 

### Return type

[**AddGroupAttachmentResponse**](AddGroupAttachmentResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupAttachment

> DeleteGroupAttachment(ctx, groupId, attachId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

Hard delete group attachment

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
    groupId := "groupId_example" // string | group_id
    attachId := "attachId_example" // string | The attachment ID (returned from a successful group attachment request) 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.DeleteGroupAttachment(context.Background(), groupId, attachId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.DeleteGroupAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | group_id | 
**attachId** | **string** | The attachment ID (returned from a successful group attachment request)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupAttachmentRequest struct via the builder pattern


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


## GetGroup

> AggregateGroup GetGroup(ctx, groupId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

Get group, group context and group context owners info

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
    groupId := "groupId_example" // string | group_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroup(context.Background(), groupId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup`: AggregateGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | group_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 


### Return type

[**AggregateGroup**](AggregateGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupAttachment

> GroupAttachmentResponse GetGroupAttachment(ctx, groupId, attachId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

Get group attachment

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
    groupId := "groupId_example" // string | group_id
    attachId := "attachId_example" // string | The attachment ID (returned from a successful group attachment request) 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupAttachment(context.Background(), groupId, attachId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupAttachment`: GroupAttachmentResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | group_id | 
**attachId** | **string** | The attachment ID (returned from a successful group attachment request)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 



### Return type

[**GroupAttachmentResponse**](GroupAttachmentResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupAttachments

> GetGroupAttachments200Response GetGroupAttachments(ctx, groupId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()

Get all attachments

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
    groupId := "groupId_example" // string | group_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupAttachments(context.Background(), groupId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupAttachments`: GetGroupAttachments200Response
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | group_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 


### Return type

[**GetGroupAttachments200Response**](GetGroupAttachments200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGroupIdPut

> AggregateGroup GroupsGroupIdPut(ctx, groupId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()

upsert group, group context and group context owners info

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
    groupId := "groupId_example" // string | group_id
    data := *openapiclient.NewAggregateGroup() // AggregateGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsGroupIdPut(context.Background(), groupId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsGroupIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGroupIdPut`: AggregateGroup
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsGroupIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | group_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 

 **data** | [**AggregateGroup**](AggregateGroup.md) |  | 

### Return type

[**AggregateGroup**](AggregateGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupAttachment

> GroupAttachmentResponse UpdateGroupAttachment(ctx, groupId, attachId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()

Patch a group attachment

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
    groupId := "groupId_example" // string | group_id
    attachId := "attachId_example" // string | The attachment ID (returned from a successful group attachment request) 
    data := *openapiclient.NewAddGroupAttachmentRequest() // AddGroupAttachmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.UpdateGroupAttachment(context.Background(), groupId, attachId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.UpdateGroupAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroupAttachment`: GroupAttachmentResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.UpdateGroupAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | group_id | 
**attachId** | **string** | The attachment ID (returned from a successful group attachment request)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 


 **data** | [**AddGroupAttachmentRequest**](AddGroupAttachmentRequest.md) |  | 

### Return type

[**GroupAttachmentResponse**](GroupAttachmentResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadGroupAttachment

> SuccessMessage UploadGroupAttachment(ctx, groupId, attachId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).ContentDisposition(contentDisposition).File(file).Execute()

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
    groupId := "groupId_example" // string | The Group ID (required to already exist) 
    attachId := "attachId_example" // string | The attachment ID (returned from a successful group attachment request) 
    contentDisposition := "inline; filename="MyFile.docx"" // string | *must include the \"filename\" parameter, e.g. `inline; filename=\"MyFile.docx\"` 
    file := map[string]interface{}{ ... } // map[string]interface{} | the attachment file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.UploadGroupAttachment(context.Background(), groupId, attachId).XTurnitinIntegrationName(xTurnitinIntegrationName).XTurnitinIntegrationVersion(xTurnitinIntegrationVersion).ContentDisposition(contentDisposition).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.UploadGroupAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadGroupAttachment`: SuccessMessage
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.UploadGroupAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The Group ID (required to already exist)  | 
**attachId** | **string** | The attachment ID (returned from a successful group attachment request)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadGroupAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTurnitinIntegrationName** | **string** | a human readable string representing the type of integration being used | 
 **xTurnitinIntegrationVersion** | **string** | the version of the integration platform being used | 


 **contentDisposition** | **string** | *must include the \&quot;filename\&quot; parameter, e.g. &#x60;inline; filename&#x3D;\&quot;MyFile.docx\&quot;&#x60;  | 
 **file** | **map[string]interface{}** | the attachment file | 

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

