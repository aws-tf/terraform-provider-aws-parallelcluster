# \ClusterComputeFleetAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DescribeComputeFleet**](ClusterComputeFleetAPI.md#DescribeComputeFleet) | **Get** /v3/clusters/{clusterName}/computefleet | 
[**UpdateComputeFleet**](ClusterComputeFleetAPI.md#UpdateComputeFleet) | **Patch** /v3/clusters/{clusterName}/computefleet | 



## DescribeComputeFleet

> DescribeComputeFleetResponseContent DescribeComputeFleet(ctx, clusterName).Region(region).Execute()





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
	resp, r, err := apiClient.ClusterComputeFleetAPI.DescribeComputeFleet(context.Background(), clusterName).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterComputeFleetAPI.DescribeComputeFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeComputeFleet`: DescribeComputeFleetResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ClusterComputeFleetAPI.DescribeComputeFleet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Name of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeComputeFleetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | AWS Region that the operation corresponds to. | 

### Return type

[**DescribeComputeFleetResponseContent**](DescribeComputeFleetResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComputeFleet

> UpdateComputeFleetResponseContent UpdateComputeFleet(ctx, clusterName).UpdateComputeFleetRequestContent(updateComputeFleetRequestContent).Region(region).Execute()





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
	updateComputeFleetRequestContent := *openapiclient.NewUpdateComputeFleetRequestContent(openapiclient.RequestedComputeFleetStatus("START_REQUESTED")) // UpdateComputeFleetRequestContent | 
	region := "region_example" // string | AWS Region that the operation corresponds to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterComputeFleetAPI.UpdateComputeFleet(context.Background(), clusterName).UpdateComputeFleetRequestContent(updateComputeFleetRequestContent).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterComputeFleetAPI.UpdateComputeFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateComputeFleet`: UpdateComputeFleetResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ClusterComputeFleetAPI.UpdateComputeFleet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Name of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateComputeFleetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateComputeFleetRequestContent** | [**UpdateComputeFleetRequestContent**](UpdateComputeFleetRequestContent.md) |  | 
 **region** | **string** | AWS Region that the operation corresponds to. | 

### Return type

[**UpdateComputeFleetResponseContent**](UpdateComputeFleetResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

