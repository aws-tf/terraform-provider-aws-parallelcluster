# \ClusterLogsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterLogEvents**](ClusterLogsAPI.md#GetClusterLogEvents) | **Get** /v3/clusters/{clusterName}/logstreams/{logStreamName} | 
[**GetClusterStackEvents**](ClusterLogsAPI.md#GetClusterStackEvents) | **Get** /v3/clusters/{clusterName}/stackevents | 
[**ListClusterLogStreams**](ClusterLogsAPI.md#ListClusterLogStreams) | **Get** /v3/clusters/{clusterName}/logstreams | 



## GetClusterLogEvents

> GetClusterLogEventsResponseContent GetClusterLogEvents(ctx, clusterName, logStreamName).Region(region).NextToken(nextToken).StartFromHead(startFromHead).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Name of the cluster
	logStreamName := "logStreamName_example" // string | Name of the log stream.
	region := "region_example" // string | AWS Region that the operation corresponds to. (optional)
	nextToken := "nextToken_example" // string | Token to use for paginated requests. (optional)
	startFromHead := true // bool | If the value is true, the earliest log events are returned first. If the value is false, the latest log events are returned first. (Defaults to 'false'.) (optional)
	limit := int32(56) // int32 | The maximum number of log events returned. If you don't specify a value, the maximum is as many log events as can fit in a response size of 1 MB, up to 10,000 log events. (optional)
	startTime := time.Now() // time.Time | The start of the time range, expressed in ISO 8601 format (e.g. '2021-01-01T20:00:00Z'). Events with a timestamp equal to this time or later than this time are included. (optional)
	endTime := time.Now() // time.Time | The end of the time range, expressed in ISO 8601 format (e.g. '2021-01-01T20:00:00Z'). Events with a timestamp equal to or later than this time are not included. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterLogsAPI.GetClusterLogEvents(context.Background(), clusterName, logStreamName).Region(region).NextToken(nextToken).StartFromHead(startFromHead).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterLogsAPI.GetClusterLogEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterLogEvents`: GetClusterLogEventsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ClusterLogsAPI.GetClusterLogEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Name of the cluster | 
**logStreamName** | **string** | Name of the log stream. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterLogEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **region** | **string** | AWS Region that the operation corresponds to. | 
 **nextToken** | **string** | Token to use for paginated requests. | 
 **startFromHead** | **bool** | If the value is true, the earliest log events are returned first. If the value is false, the latest log events are returned first. (Defaults to &#39;false&#39;.) | 
 **limit** | **int32** | The maximum number of log events returned. If you don&#39;t specify a value, the maximum is as many log events as can fit in a response size of 1 MB, up to 10,000 log events. | 
 **startTime** | **time.Time** | The start of the time range, expressed in ISO 8601 format (e.g. &#39;2021-01-01T20:00:00Z&#39;). Events with a timestamp equal to this time or later than this time are included. | 
 **endTime** | **time.Time** | The end of the time range, expressed in ISO 8601 format (e.g. &#39;2021-01-01T20:00:00Z&#39;). Events with a timestamp equal to or later than this time are not included. | 

### Return type

[**GetClusterLogEventsResponseContent**](GetClusterLogEventsResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterStackEvents

> GetClusterStackEventsResponseContent GetClusterStackEvents(ctx, clusterName).Region(region).NextToken(nextToken).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterLogsAPI.GetClusterStackEvents(context.Background(), clusterName).Region(region).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterLogsAPI.GetClusterStackEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterStackEvents`: GetClusterStackEventsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ClusterLogsAPI.GetClusterStackEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Name of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterStackEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | AWS Region that the operation corresponds to. | 
 **nextToken** | **string** | Token to use for paginated requests. | 

### Return type

[**GetClusterStackEventsResponseContent**](GetClusterStackEventsResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterLogStreams

> ListClusterLogStreamsResponseContent ListClusterLogStreams(ctx, clusterName).Region(region).Filters(filters).NextToken(nextToken).Execute()





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
	region := "region_example" // string | Region that the given cluster belongs to. (optional)
	filters := []string{"Inner_example"} // []string | Filter the log streams. Format: 'Name=a,Values=1 Name=b,Values=2,3'. Accepted filters are: private-dns-name - The short form of the private DNS name of the instance (e.g. ip-10-0-0-101). node-type - The node type, the only accepted value for this filter is HeadNode. (optional)
	nextToken := "nextToken_example" // string | Token to use for paginated requests. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterLogsAPI.ListClusterLogStreams(context.Background(), clusterName).Region(region).Filters(filters).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterLogsAPI.ListClusterLogStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterLogStreams`: ListClusterLogStreamsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ClusterLogsAPI.ListClusterLogStreams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Name of the cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterLogStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | Region that the given cluster belongs to. | 
 **filters** | **[]string** | Filter the log streams. Format: &#39;Name&#x3D;a,Values&#x3D;1 Name&#x3D;b,Values&#x3D;2,3&#39;. Accepted filters are: private-dns-name - The short form of the private DNS name of the instance (e.g. ip-10-0-0-101). node-type - The node type, the only accepted value for this filter is HeadNode. | 
 **nextToken** | **string** | Token to use for paginated requests. | 

### Return type

[**ListClusterLogStreamsResponseContent**](ListClusterLogStreamsResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

