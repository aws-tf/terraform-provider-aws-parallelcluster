# \ClusterOperationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCluster**](ClusterOperationsAPI.md#CreateCluster) | **Post** /v3/clusters | 
[**DeleteCluster**](ClusterOperationsAPI.md#DeleteCluster) | **Delete** /v3/clusters/{clusterName} | 
[**DescribeCluster**](ClusterOperationsAPI.md#DescribeCluster) | **Get** /v3/clusters/{clusterName} | 
[**ListClusters**](ClusterOperationsAPI.md#ListClusters) | **Get** /v3/clusters | 
[**UpdateCluster**](ClusterOperationsAPI.md#UpdateCluster) | **Put** /v3/clusters/{clusterName} | 



## CreateCluster

> CreateClusterResponseContent CreateCluster(ctx).CreateClusterRequestContent(createClusterRequestContent).Region(region).SuppressValidators(suppressValidators).ValidationFailureLevel(validationFailureLevel).Dryrun(dryrun).RollbackOnFailure(rollbackOnFailure).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createClusterRequestContent := *openapiclient.NewCreateClusterRequestContent("ClusterName_example", "ClusterConfiguration_example") // CreateClusterRequestContent | 
	region := "region_example" // string | AWS Region that the operation corresponds to. (optional)
	suppressValidators := []string{"Inner_example"} // []string | Identifies one or more config validators to suppress. Format: (ALL|type:[A-Za-z0-9]+) (optional)
	validationFailureLevel := openapiclient.ValidationLevel("INFO") // ValidationLevel | Min validation level that will cause the creation to fail. (Defaults to 'ERROR'.) (optional)
	dryrun := true // bool | Only perform request validation without creating any resource. May be used to validate the cluster configuration. (Defaults to 'false'.) (optional)
	rollbackOnFailure := true // bool | When set it automatically initiates a cluster stack rollback on failures. (Defaults to 'true'.) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterOperationsAPI.CreateCluster(context.Background()).CreateClusterRequestContent(createClusterRequestContent).Region(region).SuppressValidators(suppressValidators).ValidationFailureLevel(validationFailureLevel).Dryrun(dryrun).RollbackOnFailure(rollbackOnFailure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterOperationsAPI.CreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCluster`: CreateClusterResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ClusterOperationsAPI.CreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createClusterRequestContent** | [**CreateClusterRequestContent**](CreateClusterRequestContent.md) |  | 
 **region** | **string** | AWS Region that the operation corresponds to. | 
 **suppressValidators** | **[]string** | Identifies one or more config validators to suppress. Format: (ALL|type:[A-Za-z0-9]+) | 
 **validationFailureLevel** | [**ValidationLevel**](ValidationLevel.md) | Min validation level that will cause the creation to fail. (Defaults to &#39;ERROR&#39;.) | 
 **dryrun** | **bool** | Only perform request validation without creating any resource. May be used to validate the cluster configuration. (Defaults to &#39;false&#39;.) | 
 **rollbackOnFailure** | **bool** | When set it automatically initiates a cluster stack rollback on failures. (Defaults to &#39;true&#39;.) | 

### Return type

[**CreateClusterResponseContent**](CreateClusterResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCluster

> DeleteClusterResponseContent DeleteCluster(ctx, clusterName).Region(region).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Name of the cluster
	region := "region_example" // string | AWS Region that the operation corresponds to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterOperationsAPI.DeleteCluster(context.Background(), clusterName).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterOperationsAPI.DeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCluster`: DeleteClusterResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ClusterOperationsAPI.DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Name of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | AWS Region that the operation corresponds to. | 

### Return type

[**DeleteClusterResponseContent**](DeleteClusterResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeCluster

> DescribeClusterResponseContent DescribeCluster(ctx, clusterName).Region(region).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Name of the cluster
	region := "region_example" // string | AWS Region that the operation corresponds to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterOperationsAPI.DescribeCluster(context.Background(), clusterName).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterOperationsAPI.DescribeCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeCluster`: DescribeClusterResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ClusterOperationsAPI.DescribeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Name of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | AWS Region that the operation corresponds to. | 

### Return type

[**DescribeClusterResponseContent**](DescribeClusterResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusters

> ListClustersResponseContent ListClusters(ctx).Region(region).NextToken(nextToken).ClusterStatus(clusterStatus).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	region := "region_example" // string | List clusters deployed to a given AWS Region. (optional)
	nextToken := "nextToken_example" // string | Token to use for paginated requests. (optional)
	clusterStatus := []openapiclient.ClusterStatusFilteringOption{openapiclient.ClusterStatusFilteringOption("CREATE_IN_PROGRESS")} // []ClusterStatusFilteringOption | Filter by cluster status. (Defaults to all clusters.) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterOperationsAPI.ListClusters(context.Background()).Region(region).NextToken(nextToken).ClusterStatus(clusterStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterOperationsAPI.ListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusters`: ListClustersResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ClusterOperationsAPI.ListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** | List clusters deployed to a given AWS Region. | 
 **nextToken** | **string** | Token to use for paginated requests. | 
 **clusterStatus** | [**[]ClusterStatusFilteringOption**](ClusterStatusFilteringOption.md) | Filter by cluster status. (Defaults to all clusters.) | 

### Return type

[**ListClustersResponseContent**](ListClustersResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCluster

> UpdateClusterResponseContent UpdateCluster(ctx, clusterName).UpdateClusterRequestContent(updateClusterRequestContent).SuppressValidators(suppressValidators).ValidationFailureLevel(validationFailureLevel).Region(region).Dryrun(dryrun).ForceUpdate(forceUpdate).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Name of the cluster
	updateClusterRequestContent := *openapiclient.NewUpdateClusterRequestContent("ClusterConfiguration_example") // UpdateClusterRequestContent | 
	suppressValidators := []string{"Inner_example"} // []string | Identifies one or more config validators to suppress. Format: (ALL|type:[A-Za-z0-9]+) (optional)
	validationFailureLevel := openapiclient.ValidationLevel("INFO") // ValidationLevel | Min validation level that will cause the update to fail. (Defaults to 'ERROR'.) (optional)
	region := "region_example" // string | AWS Region that the operation corresponds to. (optional)
	dryrun := true // bool | Only perform request validation without creating any resource. May be used to validate the cluster configuration and update requirements. (Defaults to 'false'.) (optional)
	forceUpdate := true // bool | Force update by ignoring the update validation errors. (Defaults to 'false'.) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterOperationsAPI.UpdateCluster(context.Background(), clusterName).UpdateClusterRequestContent(updateClusterRequestContent).SuppressValidators(suppressValidators).ValidationFailureLevel(validationFailureLevel).Region(region).Dryrun(dryrun).ForceUpdate(forceUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterOperationsAPI.UpdateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCluster`: UpdateClusterResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ClusterOperationsAPI.UpdateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Name of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateClusterRequestContent** | [**UpdateClusterRequestContent**](UpdateClusterRequestContent.md) |  | 
 **suppressValidators** | **[]string** | Identifies one or more config validators to suppress. Format: (ALL|type:[A-Za-z0-9]+) | 
 **validationFailureLevel** | [**ValidationLevel**](ValidationLevel.md) | Min validation level that will cause the update to fail. (Defaults to &#39;ERROR&#39;.) | 
 **region** | **string** | AWS Region that the operation corresponds to. | 
 **dryrun** | **bool** | Only perform request validation without creating any resource. May be used to validate the cluster configuration and update requirements. (Defaults to &#39;false&#39;.) | 
 **forceUpdate** | **bool** | Force update by ignoring the update validation errors. (Defaults to &#39;false&#39;.) | 

### Return type

[**UpdateClusterResponseContent**](UpdateClusterResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

