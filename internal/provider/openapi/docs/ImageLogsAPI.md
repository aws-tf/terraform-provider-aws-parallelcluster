# \ImageLogsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImageLogEvents**](ImageLogsAPI.md#GetImageLogEvents) | **Get** /v3/images/custom/{imageId}/logstreams/{logStreamName} | 
[**GetImageStackEvents**](ImageLogsAPI.md#GetImageStackEvents) | **Get** /v3/images/custom/{imageId}/stackevents | 
[**ListImageLogStreams**](ImageLogsAPI.md#ListImageLogStreams) | **Get** /v3/images/custom/{imageId}/logstreams | 



## GetImageLogEvents

> GetImageLogEventsResponseContent GetImageLogEvents(ctx, imageId, logStreamName).Region(region).NextToken(nextToken).StartFromHead(startFromHead).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()





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
	imageId := "imageId_example" // string | Id of the image.
	logStreamName := "logStreamName_example" // string | Name of the log stream.
	region := "region_example" // string | AWS Region that the operation corresponds to. (optional)
	nextToken := "nextToken_example" // string | Token to use for paginated requests. (optional)
	startFromHead := true // bool | If the value is true, the earliest log events are returned first. If the value is false, the latest log events are returned first. (Defaults to 'false'.) (optional)
	limit := int32(56) // int32 | The maximum number of log events returned. If you don't specify a value, the maximum is as many log events as can fit in a response size of 1 MB, up to 10,000 log events. (optional)
	startTime := time.Now() // time.Time | The start of the time range, expressed in ISO 8601 format (e.g. '2021-01-01T20:00:00Z'). Events with a timestamp equal to this time or later than this time are included. (optional)
	endTime := time.Now() // time.Time | The end of the time range, expressed in ISO 8601 format (e.g. '2021-01-01T20:00:00Z'). Events with a timestamp equal to or later than this time are not included. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageLogsAPI.GetImageLogEvents(context.Background(), imageId, logStreamName).Region(region).NextToken(nextToken).StartFromHead(startFromHead).Limit(limit).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageLogsAPI.GetImageLogEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageLogEvents`: GetImageLogEventsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ImageLogsAPI.GetImageLogEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Id of the image. | 
**logStreamName** | **string** | Name of the log stream. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageLogEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **region** | **string** | AWS Region that the operation corresponds to. | 
 **nextToken** | **string** | Token to use for paginated requests. | 
 **startFromHead** | **bool** | If the value is true, the earliest log events are returned first. If the value is false, the latest log events are returned first. (Defaults to &#39;false&#39;.) | 
 **limit** | **int32** | The maximum number of log events returned. If you don&#39;t specify a value, the maximum is as many log events as can fit in a response size of 1 MB, up to 10,000 log events. | 
 **startTime** | **time.Time** | The start of the time range, expressed in ISO 8601 format (e.g. &#39;2021-01-01T20:00:00Z&#39;). Events with a timestamp equal to this time or later than this time are included. | 
 **endTime** | **time.Time** | The end of the time range, expressed in ISO 8601 format (e.g. &#39;2021-01-01T20:00:00Z&#39;). Events with a timestamp equal to or later than this time are not included. | 

### Return type

[**GetImageLogEventsResponseContent**](GetImageLogEventsResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageStackEvents

> GetImageStackEventsResponseContent GetImageStackEvents(ctx, imageId).Region(region).NextToken(nextToken).Execute()





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
	imageId := "imageId_example" // string | Id of the image.
	region := "region_example" // string | AWS Region that the operation corresponds to. (optional)
	nextToken := "nextToken_example" // string | Token to use for paginated requests. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageLogsAPI.GetImageStackEvents(context.Background(), imageId).Region(region).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageLogsAPI.GetImageStackEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageStackEvents`: GetImageStackEventsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ImageLogsAPI.GetImageStackEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Id of the image. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageStackEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | AWS Region that the operation corresponds to. | 
 **nextToken** | **string** | Token to use for paginated requests. | 

### Return type

[**GetImageStackEventsResponseContent**](GetImageStackEventsResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImageLogStreams

> ListImageLogStreamsResponseContent ListImageLogStreams(ctx, imageId).Region(region).NextToken(nextToken).Execute()





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
	imageId := "imageId_example" // string | Id of the image.
	region := "region_example" // string | Region that the given image belongs to. (optional)
	nextToken := "nextToken_example" // string | Token to use for paginated requests. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageLogsAPI.ListImageLogStreams(context.Background(), imageId).Region(region).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageLogsAPI.ListImageLogStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImageLogStreams`: ListImageLogStreamsResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ImageLogsAPI.ListImageLogStreams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Id of the image. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImageLogStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | Region that the given image belongs to. | 
 **nextToken** | **string** | Token to use for paginated requests. | 

### Return type

[**ListImageLogStreamsResponseContent**](ListImageLogStreamsResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

