/*
ParallelCluster

ParallelCluster API

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// ImageOperationsAPIService ImageOperationsAPI service
type ImageOperationsAPIService service

type ImageOperationsAPIBuildImageRequest struct {
	ctx context.Context
	ApiService *ImageOperationsAPIService
	buildImageRequestContent *BuildImageRequestContent
	suppressValidators *[]string
	validationFailureLevel *ValidationLevel
	dryrun *bool
	rollbackOnFailure *bool
	region *string
}

func (r ImageOperationsAPIBuildImageRequest) BuildImageRequestContent(buildImageRequestContent BuildImageRequestContent) ImageOperationsAPIBuildImageRequest {
	r.buildImageRequestContent = &buildImageRequestContent
	return r
}

// Identifies one or more config validators to suppress. Format: (ALL|type:[A-Za-z0-9]+)
func (r ImageOperationsAPIBuildImageRequest) SuppressValidators(suppressValidators []string) ImageOperationsAPIBuildImageRequest {
	r.suppressValidators = &suppressValidators
	return r
}

// Min validation level that will cause the creation to fail. (Defaults to &#39;ERROR&#39;.)
func (r ImageOperationsAPIBuildImageRequest) ValidationFailureLevel(validationFailureLevel ValidationLevel) ImageOperationsAPIBuildImageRequest {
	r.validationFailureLevel = &validationFailureLevel
	return r
}

// Only perform request validation without creating any resource. It can be used to validate the image configuration. (Defaults to &#39;false&#39;.)
func (r ImageOperationsAPIBuildImageRequest) Dryrun(dryrun bool) ImageOperationsAPIBuildImageRequest {
	r.dryrun = &dryrun
	return r
}

// When set, will automatically initiate an image stack rollback on failure. (Defaults to &#39;false&#39;.)
func (r ImageOperationsAPIBuildImageRequest) RollbackOnFailure(rollbackOnFailure bool) ImageOperationsAPIBuildImageRequest {
	r.rollbackOnFailure = &rollbackOnFailure
	return r
}

// AWS Region that the operation corresponds to.
func (r ImageOperationsAPIBuildImageRequest) Region(region string) ImageOperationsAPIBuildImageRequest {
	r.region = &region
	return r
}

func (r ImageOperationsAPIBuildImageRequest) Execute() (*BuildImageResponseContent, *http.Response, error) {
	return r.ApiService.BuildImageExecute(r)
}

/*
BuildImage Method for BuildImage

Create a custom ParallelCluster image in a given region.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ImageOperationsAPIBuildImageRequest
*/
func (a *ImageOperationsAPIService) BuildImage(ctx context.Context) ImageOperationsAPIBuildImageRequest {
	return ImageOperationsAPIBuildImageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BuildImageResponseContent
func (a *ImageOperationsAPIService) BuildImageExecute(r ImageOperationsAPIBuildImageRequest) (*BuildImageResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BuildImageResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImageOperationsAPIService.BuildImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/images/custom"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.buildImageRequestContent == nil {
		return localVarReturnValue, nil, reportError("buildImageRequestContent is required and must be specified")
	}

	if r.suppressValidators != nil {
		t := *r.suppressValidators
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "suppressValidators", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "suppressValidators", t, "multi")
		}
	}
	if r.validationFailureLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "validationFailureLevel", r.validationFailureLevel, "")
	}
	if r.dryrun != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dryrun", r.dryrun, "")
	}
	if r.rollbackOnFailure != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rollbackOnFailure", r.rollbackOnFailure, "")
	}
	if r.region != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "region", r.region, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.buildImageRequestContent
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["aws.auth.sigv4"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BuildImageBadRequestExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedClientErrorResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ConflictExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 412 {
			var v DryrunOperationExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v LimitExceededExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServiceExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ImageOperationsAPIDeleteImageRequest struct {
	ctx context.Context
	ApiService *ImageOperationsAPIService
	imageId string
	region *string
	force *bool
}

// AWS Region that the operation corresponds to.
func (r ImageOperationsAPIDeleteImageRequest) Region(region string) ImageOperationsAPIDeleteImageRequest {
	r.region = &region
	return r
}

// Force deletion in case there are instances using the AMI or in case the AMI is shared. (Defaults to &#39;false&#39;.)
func (r ImageOperationsAPIDeleteImageRequest) Force(force bool) ImageOperationsAPIDeleteImageRequest {
	r.force = &force
	return r
}

func (r ImageOperationsAPIDeleteImageRequest) Execute() (*DeleteImageResponseContent, *http.Response, error) {
	return r.ApiService.DeleteImageExecute(r)
}

/*
DeleteImage Method for DeleteImage

Initiate the deletion of the custom ParallelCluster image.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageId Id of the image.
 @return ImageOperationsAPIDeleteImageRequest
*/
func (a *ImageOperationsAPIService) DeleteImage(ctx context.Context, imageId string) ImageOperationsAPIDeleteImageRequest {
	return ImageOperationsAPIDeleteImageRequest{
		ApiService: a,
		ctx: ctx,
		imageId: imageId,
	}
}

// Execute executes the request
//  @return DeleteImageResponseContent
func (a *ImageOperationsAPIService) DeleteImageExecute(r ImageOperationsAPIDeleteImageRequest) (*DeleteImageResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeleteImageResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImageOperationsAPIService.DeleteImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/images/custom/{imageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", url.PathEscape(parameterValueToString(r.imageId, "imageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.region != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "region", r.region, "")
	}
	if r.force != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "force", r.force, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["aws.auth.sigv4"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedClientErrorResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v NotFoundExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v LimitExceededExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServiceExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ImageOperationsAPIDescribeImageRequest struct {
	ctx context.Context
	ApiService *ImageOperationsAPIService
	imageId string
	region *string
}

// AWS Region that the operation corresponds to.
func (r ImageOperationsAPIDescribeImageRequest) Region(region string) ImageOperationsAPIDescribeImageRequest {
	r.region = &region
	return r
}

func (r ImageOperationsAPIDescribeImageRequest) Execute() (*DescribeImageResponseContent, *http.Response, error) {
	return r.ApiService.DescribeImageExecute(r)
}

/*
DescribeImage Method for DescribeImage

Get detailed information about an existing image.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageId Id of the image.
 @return ImageOperationsAPIDescribeImageRequest
*/
func (a *ImageOperationsAPIService) DescribeImage(ctx context.Context, imageId string) ImageOperationsAPIDescribeImageRequest {
	return ImageOperationsAPIDescribeImageRequest{
		ApiService: a,
		ctx: ctx,
		imageId: imageId,
	}
}

// Execute executes the request
//  @return DescribeImageResponseContent
func (a *ImageOperationsAPIService) DescribeImageExecute(r ImageOperationsAPIDescribeImageRequest) (*DescribeImageResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DescribeImageResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImageOperationsAPIService.DescribeImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/images/custom/{imageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", url.PathEscape(parameterValueToString(r.imageId, "imageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.region != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "region", r.region, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["aws.auth.sigv4"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedClientErrorResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v NotFoundExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v LimitExceededExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServiceExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ImageOperationsAPIListImagesRequest struct {
	ctx context.Context
	ApiService *ImageOperationsAPIService
	imageStatus *ImageStatusFilteringOption
	region *string
	nextToken *string
}

// Filter images by the status provided.
func (r ImageOperationsAPIListImagesRequest) ImageStatus(imageStatus ImageStatusFilteringOption) ImageOperationsAPIListImagesRequest {
	r.imageStatus = &imageStatus
	return r
}

// List images built in a given AWS Region.
func (r ImageOperationsAPIListImagesRequest) Region(region string) ImageOperationsAPIListImagesRequest {
	r.region = &region
	return r
}

// Token to use for paginated requests.
func (r ImageOperationsAPIListImagesRequest) NextToken(nextToken string) ImageOperationsAPIListImagesRequest {
	r.nextToken = &nextToken
	return r
}

func (r ImageOperationsAPIListImagesRequest) Execute() (*ListImagesResponseContent, *http.Response, error) {
	return r.ApiService.ListImagesExecute(r)
}

/*
ListImages Method for ListImages

Retrieve the list of existing custom images.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ImageOperationsAPIListImagesRequest
*/
func (a *ImageOperationsAPIService) ListImages(ctx context.Context) ImageOperationsAPIListImagesRequest {
	return ImageOperationsAPIListImagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListImagesResponseContent
func (a *ImageOperationsAPIService) ListImagesExecute(r ImageOperationsAPIListImagesRequest) (*ListImagesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListImagesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImageOperationsAPIService.ListImages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/images/custom"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.imageStatus == nil {
		return localVarReturnValue, nil, reportError("imageStatus is required and must be specified")
	}

	if r.region != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "region", r.region, "")
	}
	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextToken", r.nextToken, "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "imageStatus", r.imageStatus, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["aws.auth.sigv4"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedClientErrorResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v LimitExceededExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServiceExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ImageOperationsAPIListOfficialImagesRequest struct {
	ctx context.Context
	ApiService *ImageOperationsAPIService
	region *string
	os *string
	architecture *string
}

// AWS Region that the operation corresponds to.
func (r ImageOperationsAPIListOfficialImagesRequest) Region(region string) ImageOperationsAPIListOfficialImagesRequest {
	r.region = &region
	return r
}

// Filter by OS distribution (Default is to not filter.)
func (r ImageOperationsAPIListOfficialImagesRequest) Os(os string) ImageOperationsAPIListOfficialImagesRequest {
	r.os = &os
	return r
}

// Filter by architecture (Default is to not filter.)
func (r ImageOperationsAPIListOfficialImagesRequest) Architecture(architecture string) ImageOperationsAPIListOfficialImagesRequest {
	r.architecture = &architecture
	return r
}

func (r ImageOperationsAPIListOfficialImagesRequest) Execute() (*ListOfficialImagesResponseContent, *http.Response, error) {
	return r.ApiService.ListOfficialImagesExecute(r)
}

/*
ListOfficialImages Method for ListOfficialImages

List Official ParallelCluster AMIs.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ImageOperationsAPIListOfficialImagesRequest
*/
func (a *ImageOperationsAPIService) ListOfficialImages(ctx context.Context) ImageOperationsAPIListOfficialImagesRequest {
	return ImageOperationsAPIListOfficialImagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListOfficialImagesResponseContent
func (a *ImageOperationsAPIService) ListOfficialImagesExecute(r ImageOperationsAPIListOfficialImagesRequest) (*ListOfficialImagesResponseContent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListOfficialImagesResponseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImageOperationsAPIService.ListOfficialImages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/images/official"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.region != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "region", r.region, "")
	}
	if r.os != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "os", r.os, "")
	}
	if r.architecture != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "architecture", r.architecture, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["aws.auth.sigv4"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedClientErrorResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v LimitExceededExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InternalServiceExceptionResponseContent
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}