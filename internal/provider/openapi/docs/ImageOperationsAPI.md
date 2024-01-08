# \ImageOperationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildImage**](ImageOperationsAPI.md#BuildImage) | **Post** /v3/images/custom | 
[**DeleteImage**](ImageOperationsAPI.md#DeleteImage) | **Delete** /v3/images/custom/{imageId} | 
[**DescribeImage**](ImageOperationsAPI.md#DescribeImage) | **Get** /v3/images/custom/{imageId} | 
[**ListImages**](ImageOperationsAPI.md#ListImages) | **Get** /v3/images/custom | 
[**ListOfficialImages**](ImageOperationsAPI.md#ListOfficialImages) | **Get** /v3/images/official | 



## BuildImage

> BuildImageResponseContent BuildImage(ctx).BuildImageRequestContent(buildImageRequestContent).SuppressValidators(suppressValidators).ValidationFailureLevel(validationFailureLevel).Dryrun(dryrun).RollbackOnFailure(rollbackOnFailure).Region(region).Execute()





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
	buildImageRequestContent := *openapiclient.NewBuildImageRequestContent("ImageConfiguration_example", "ImageId_example") // BuildImageRequestContent | 
	suppressValidators := []string{"Inner_example"} // []string | Identifies one or more config validators to suppress. Format: (ALL|type:[A-Za-z0-9]+) (optional)
	validationFailureLevel := openapiclient.ValidationLevel("INFO") // ValidationLevel | Min validation level that will cause the creation to fail. (Defaults to 'ERROR'.) (optional)
	dryrun := true // bool | Only perform request validation without creating any resource. It can be used to validate the image configuration. (Defaults to 'false'.) (optional)
	rollbackOnFailure := true // bool | When set, will automatically initiate an image stack rollback on failure. (Defaults to 'false'.) (optional)
	region := "region_example" // string | AWS Region that the operation corresponds to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageOperationsAPI.BuildImage(context.Background()).BuildImageRequestContent(buildImageRequestContent).SuppressValidators(suppressValidators).ValidationFailureLevel(validationFailureLevel).Dryrun(dryrun).RollbackOnFailure(rollbackOnFailure).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageOperationsAPI.BuildImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildImage`: BuildImageResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ImageOperationsAPI.BuildImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuildImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildImageRequestContent** | [**BuildImageRequestContent**](BuildImageRequestContent.md) |  | 
 **suppressValidators** | **[]string** | Identifies one or more config validators to suppress. Format: (ALL|type:[A-Za-z0-9]+) | 
 **validationFailureLevel** | [**ValidationLevel**](ValidationLevel.md) | Min validation level that will cause the creation to fail. (Defaults to &#39;ERROR&#39;.) | 
 **dryrun** | **bool** | Only perform request validation without creating any resource. It can be used to validate the image configuration. (Defaults to &#39;false&#39;.) | 
 **rollbackOnFailure** | **bool** | When set, will automatically initiate an image stack rollback on failure. (Defaults to &#39;false&#39;.) | 
 **region** | **string** | AWS Region that the operation corresponds to. | 

### Return type

[**BuildImageResponseContent**](BuildImageResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImage

> DeleteImageResponseContent DeleteImage(ctx, imageId).Region(region).Force(force).Execute()





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
	force := true // bool | Force deletion in case there are instances using the AMI or in case the AMI is shared. (Defaults to 'false'.) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageOperationsAPI.DeleteImage(context.Background(), imageId).Region(region).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageOperationsAPI.DeleteImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteImage`: DeleteImageResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ImageOperationsAPI.DeleteImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Id of the image. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | AWS Region that the operation corresponds to. | 
 **force** | **bool** | Force deletion in case there are instances using the AMI or in case the AMI is shared. (Defaults to &#39;false&#39;.) | 

### Return type

[**DeleteImageResponseContent**](DeleteImageResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeImage

> DescribeImageResponseContent DescribeImage(ctx, imageId).Region(region).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageOperationsAPI.DescribeImage(context.Background(), imageId).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageOperationsAPI.DescribeImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeImage`: DescribeImageResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ImageOperationsAPI.DescribeImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Id of the image. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | AWS Region that the operation corresponds to. | 

### Return type

[**DescribeImageResponseContent**](DescribeImageResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImages

> ListImagesResponseContent ListImages(ctx).ImageStatus(imageStatus).Region(region).NextToken(nextToken).Execute()





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
	imageStatus := openapiclient.ImageStatusFilteringOption("AVAILABLE") // ImageStatusFilteringOption | Filter images by the status provided.
	region := "region_example" // string | List images built in a given AWS Region. (optional)
	nextToken := "nextToken_example" // string | Token to use for paginated requests. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageOperationsAPI.ListImages(context.Background()).ImageStatus(imageStatus).Region(region).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageOperationsAPI.ListImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImages`: ListImagesResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ImageOperationsAPI.ListImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageStatus** | [**ImageStatusFilteringOption**](ImageStatusFilteringOption.md) | Filter images by the status provided. | 
 **region** | **string** | List images built in a given AWS Region. | 
 **nextToken** | **string** | Token to use for paginated requests. | 

### Return type

[**ListImagesResponseContent**](ListImagesResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOfficialImages

> ListOfficialImagesResponseContent ListOfficialImages(ctx).Region(region).Os(os).Architecture(architecture).Execute()





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
	region := "region_example" // string | AWS Region that the operation corresponds to. (optional)
	os := "os_example" // string | Filter by OS distribution (Default is to not filter.) (optional)
	architecture := "architecture_example" // string | Filter by architecture (Default is to not filter.) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageOperationsAPI.ListOfficialImages(context.Background()).Region(region).Os(os).Architecture(architecture).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageOperationsAPI.ListOfficialImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOfficialImages`: ListOfficialImagesResponseContent
	fmt.Fprintf(os.Stdout, "Response from `ImageOperationsAPI.ListOfficialImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOfficialImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** | AWS Region that the operation corresponds to. | 
 **os** | **string** | Filter by OS distribution (Default is to not filter.) | 
 **architecture** | **string** | Filter by architecture (Default is to not filter.) | 

### Return type

[**ListOfficialImagesResponseContent**](ListOfficialImagesResponseContent.md)

### Authorization

[aws.auth.sigv4](../README.md#aws.auth.sigv4)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

