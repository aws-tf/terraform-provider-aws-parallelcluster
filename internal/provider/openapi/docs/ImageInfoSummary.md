# ImageInfoSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | **string** | Id of the image. | 
**Ec2AmiInfo** | Pointer to [**Ec2AmiInfoSummary**](Ec2AmiInfoSummary.md) |  | [optional] 
**Region** | **string** | AWS region where the image is built. | 
**Version** | **string** | ParallelCluster version used to build the image. | 
**CloudformationStackArn** | Pointer to **string** | ARN of the main CloudFormation stack. | [optional] 
**ImageBuildStatus** | [**ImageBuildStatus**](ImageBuildStatus.md) |  | 
**CloudformationStackStatus** | Pointer to [**CloudFormationStackStatus**](CloudFormationStackStatus.md) |  | [optional] 

## Methods

### NewImageInfoSummary

`func NewImageInfoSummary(imageId string, region string, version string, imageBuildStatus ImageBuildStatus, ) *ImageInfoSummary`

NewImageInfoSummary instantiates a new ImageInfoSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageInfoSummaryWithDefaults

`func NewImageInfoSummaryWithDefaults() *ImageInfoSummary`

NewImageInfoSummaryWithDefaults instantiates a new ImageInfoSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *ImageInfoSummary) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ImageInfoSummary) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ImageInfoSummary) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetEc2AmiInfo

`func (o *ImageInfoSummary) GetEc2AmiInfo() Ec2AmiInfoSummary`

GetEc2AmiInfo returns the Ec2AmiInfo field if non-nil, zero value otherwise.

### GetEc2AmiInfoOk

`func (o *ImageInfoSummary) GetEc2AmiInfoOk() (*Ec2AmiInfoSummary, bool)`

GetEc2AmiInfoOk returns a tuple with the Ec2AmiInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEc2AmiInfo

`func (o *ImageInfoSummary) SetEc2AmiInfo(v Ec2AmiInfoSummary)`

SetEc2AmiInfo sets Ec2AmiInfo field to given value.

### HasEc2AmiInfo

`func (o *ImageInfoSummary) HasEc2AmiInfo() bool`

HasEc2AmiInfo returns a boolean if a field has been set.

### GetRegion

`func (o *ImageInfoSummary) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ImageInfoSummary) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ImageInfoSummary) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVersion

`func (o *ImageInfoSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ImageInfoSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ImageInfoSummary) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetCloudformationStackArn

`func (o *ImageInfoSummary) GetCloudformationStackArn() string`

GetCloudformationStackArn returns the CloudformationStackArn field if non-nil, zero value otherwise.

### GetCloudformationStackArnOk

`func (o *ImageInfoSummary) GetCloudformationStackArnOk() (*string, bool)`

GetCloudformationStackArnOk returns a tuple with the CloudformationStackArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudformationStackArn

`func (o *ImageInfoSummary) SetCloudformationStackArn(v string)`

SetCloudformationStackArn sets CloudformationStackArn field to given value.

### HasCloudformationStackArn

`func (o *ImageInfoSummary) HasCloudformationStackArn() bool`

HasCloudformationStackArn returns a boolean if a field has been set.

### GetImageBuildStatus

`func (o *ImageInfoSummary) GetImageBuildStatus() ImageBuildStatus`

GetImageBuildStatus returns the ImageBuildStatus field if non-nil, zero value otherwise.

### GetImageBuildStatusOk

`func (o *ImageInfoSummary) GetImageBuildStatusOk() (*ImageBuildStatus, bool)`

GetImageBuildStatusOk returns a tuple with the ImageBuildStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuildStatus

`func (o *ImageInfoSummary) SetImageBuildStatus(v ImageBuildStatus)`

SetImageBuildStatus sets ImageBuildStatus field to given value.


### GetCloudformationStackStatus

`func (o *ImageInfoSummary) GetCloudformationStackStatus() CloudFormationStackStatus`

GetCloudformationStackStatus returns the CloudformationStackStatus field if non-nil, zero value otherwise.

### GetCloudformationStackStatusOk

`func (o *ImageInfoSummary) GetCloudformationStackStatusOk() (*CloudFormationStackStatus, bool)`

GetCloudformationStackStatusOk returns a tuple with the CloudformationStackStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudformationStackStatus

`func (o *ImageInfoSummary) SetCloudformationStackStatus(v CloudFormationStackStatus)`

SetCloudformationStackStatus sets CloudformationStackStatus field to given value.

### HasCloudformationStackStatus

`func (o *ImageInfoSummary) HasCloudformationStackStatus() bool`

HasCloudformationStackStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


