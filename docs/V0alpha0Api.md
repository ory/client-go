# \V0alpha0Api

All URIs are relative to *https://playground.projects.oryapis.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProject**](V0alpha0Api.md#CreateProject) | **Post** /backoffice/public/projects | Create a Project
[**GetProject**](V0alpha0Api.md#GetProject) | **Get** /backoffice/public/projects/{project_id} | Get a Project
[**GetProjectMembers**](V0alpha0Api.md#GetProjectMembers) | **Get** /backoffice/public/projects/{project_id}/members | Get all members associated with this project.
[**ListProjects**](V0alpha0Api.md#ListProjects) | **Get** /backoffice/public/projects | List All Projects
[**RemoveProjectMember**](V0alpha0Api.md#RemoveProjectMember) | **Delete** /backoffice/public/projects/{project_id}/members/{member_id} | Remove a member associated with this project. This also sets their invite status to &#x60;REMOVED&#x60;.
[**UpdateProject**](V0alpha0Api.md#UpdateProject) | **Put** /backoffice/public/projects/{project_id} | Update an Ory Cloud Project Configuration



## CreateProject

> CreateProjectResponse CreateProject(ctx).CreateProjectBody(createProjectBody).Execute()

Create a Project



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
    createProjectBody := *openapiclient.NewCreateProjectBody() // CreateProjectBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V0alpha0Api.CreateProject(context.Background()).CreateProjectBody(createProjectBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V0alpha0Api.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: CreateProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `V0alpha0Api.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProjectBody** | [**CreateProjectBody**](CreateProjectBody.md) |  | 

### Return type

[**CreateProjectResponse**](CreateProjectResponse.md)

### Authorization

[oryAccessToken](../README.md#oryAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProject

> ProjectNext GetProject(ctx, projectId).Execute()

Get a Project



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
    projectId := "projectId_example" // string | Project ID  The project's ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V0alpha0Api.GetProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V0alpha0Api.GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProject`: ProjectNext
    fmt.Fprintf(os.Stdout, "Response from `V0alpha0Api.GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID  The project&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectNext**](ProjectNext.md)

### Authorization

[oryAccessToken](../README.md#oryAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectMembers

> []CloudAccount GetProjectMembers(ctx, projectId).Execute()

Get all members associated with this project.



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
    projectId := "projectId_example" // string | Project ID  The project's ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V0alpha0Api.GetProjectMembers(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V0alpha0Api.GetProjectMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectMembers`: []CloudAccount
    fmt.Fprintf(os.Stdout, "Response from `V0alpha0Api.GetProjectMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID  The project&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CloudAccount**](CloudAccount.md)

### Authorization

[oryAccessToken](../README.md#oryAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjects

> []ProjectNext ListProjects(ctx).Execute()

List All Projects



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V0alpha0Api.ListProjects(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V0alpha0Api.ListProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProjects`: []ProjectNext
    fmt.Fprintf(os.Stdout, "Response from `V0alpha0Api.ListProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectsRequest struct via the builder pattern


### Return type

[**[]ProjectNext**](ProjectNext.md)

### Authorization

[oryAccessToken](../README.md#oryAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProjectMember

> RemoveProjectMember(ctx, projectId, memberId).Execute()

Remove a member associated with this project. This also sets their invite status to `REMOVED`.



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
    projectId := "projectId_example" // string | Project ID  The project's ID.
    memberId := "memberId_example" // string | Member ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V0alpha0Api.RemoveProjectMember(context.Background(), projectId, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V0alpha0Api.RemoveProjectMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID  The project&#39;s ID. | 
**memberId** | **string** | Member ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProjectMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oryAccessToken](../README.md#oryAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProject

> SuccessfulProjectUpdate UpdateProject(ctx, projectId).UpdateProject(updateProject).Execute()

Update an Ory Cloud Project Configuration



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
    projectId := "projectId_example" // string | Project ID  The project's ID.
    updateProject := *openapiclient.NewUpdateProject("Name_example", *openapiclient.NewProjectServices()) // UpdateProject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V0alpha0Api.UpdateProject(context.Background(), projectId).UpdateProject(updateProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V0alpha0Api.UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProject`: SuccessfulProjectUpdate
    fmt.Fprintf(os.Stdout, "Response from `V0alpha0Api.UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID  The project&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProject** | [**UpdateProject**](UpdateProject.md) |  | 

### Return type

[**SuccessfulProjectUpdate**](SuccessfulProjectUpdate.md)

### Authorization

[oryAccessToken](../README.md#oryAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

