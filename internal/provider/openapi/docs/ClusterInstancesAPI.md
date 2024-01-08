# \ClusterInstancesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClusterInstances**](ClusterInstancesAPI.md#DeleteClusterInstances) | **Delete** /v3/clusters/{clusterName}/instances | 
[**DescribeClusterInstances**](ClusterInstancesAPI.md#DescribeClusterInstances) | **Get** /v3/clusters/{clusterName}/instances | 



## DeleteClusterInstances

> DeleteClusterInstances(ctx, clusterName).Region(region).Force(force).Execute()





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
	force := true // bool | Force the deletion also when the cluster with the given name is not found. (Defaults to 'false'.) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClusterInstancesAPI.DeleteClusterInstances(context.Background(), clusterName).Region(region).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterInstancesAPI.DeleteClusterInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Name of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | AWS Region that the operation corresponds to. | 
 **force** | **bool** | Force the deletion also when the cluster with the given name is not found. (Defaults to &#39;false&#39;.) | 

### Return type

 (empty response body)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeClusterInstances

> DescribeClusterInstancesResponseContent DescribeClusterInstances(ctx, clusterName).Region(region).NextToken(nextToken).NodeType(nodeType).QueueName(queueName).Execute()





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
	nextToken := "nextToken_example" // string | Token to use for paginated requests. (optional)
	nodeType := openapiclient.NodeType("HeadNode") // NodeType | Filter the instances by node type. (optional)
	queueName := "queueName_example" // string | Filter the instances by queue name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterInstancesAPI.DescribeClusterInstances(context.Background(), clusterName).Region(region).NextToken(nextToken).NodeType(nodeType).QueueName(queueName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterInstancesAPI.DescribeClusterInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeClusterInstances`: DescribeClusterInstancesResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ClusterInstancesAPI.DescribeClusterInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Name of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeClusterInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | AWS Region that the operation corresponds to. | 
 **nextToken** | **string** | Token to use for paginated requests. | 
 **nodeType** | [**NodeType**](NodeType.md) | Filter the instances by node type. | 
 **queueName** | **string** | Filter the instances by queue name. | 

### Return type

[**DescribeClusterInstancesResponseContent**](DescribeClusterInstancesResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

