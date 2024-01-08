# DescribeImageResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | **string** | Id of the Image to retrieve detailed information for. | 
**Region** | **string** | AWS region where the image is created. | 
**Version** | **string** | ParallelCluster version used to build the image. | 
**ImageBuildStatus** | [**ImageBuildStatus**](ImageBuildStatus.md) |  | 
**ImageBuildLogsArn** | Pointer to **string** | ARN of the logs for the image build process. | [optional] 
**CloudformationStackStatus** | Pointer to [**CloudFormationStackStatus**](CloudFormationStackStatus.md) |  | [optional] 
**CloudformationStackStatusReason** | Pointer to **string** | Reason for the CloudFormation stack status. | [optional] 
**CloudformationStackArn** | Pointer to **string** | ARN of the main CloudFormation stack. | [optional] 
**CreationTime** | Pointer to **time.Time** | Timestamp representing the image creation time. | [optional] 
**CloudformationStackCreationTime** | Pointer to **time.Time** | Timestamp representing the CloudFormation stack creation time. | [optional] 
**CloudformationStackTags** | Pointer to [**[]Tag**](Tag.md) | Tags for the CloudFormation stack. | [optional] 
**ImageConfiguration** | [**ImageConfigurationStructure**](ImageConfigurationStructure.md) |  | 
**ImagebuilderImageStatus** | Pointer to [**ImageBuilderImageStatus**](ImageBuilderImageStatus.md) |  | [optional] 
**ImagebuilderImageStatusReason** | Pointer to **string** | Reason for the ImageBuilder Image status. | [optional] 
**Ec2AmiInfo** | Pointer to [**Ec2AmiInfo**](Ec2AmiInfo.md) |  | [optional] 

## Methods

### NewDescribeImageResponseContent

`func NewDescribeImageResponseContent(imageId string, region string, version string, imageBuildStatus ImageBuildStatus, imageConfiguration ImageConfigurationStructure, ) *DescribeImageResponseContent`

NewDescribeImageResponseContent instantiates a new DescribeImageResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeImageResponseContentWithDefaults

`func NewDescribeImageResponseContentWithDefaults() *DescribeImageResponseContent`

NewDescribeImageResponseContentWithDefaults instantiates a new DescribeImageResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *DescribeImageResponseContent) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *DescribeImageResponseContent) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *DescribeImageResponseContent) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetRegion

`func (o *DescribeImageResponseContent) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DescribeImageResponseContent) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DescribeImageResponseContent) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVersion

`func (o *DescribeImageResponseContent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DescribeImageResponseContent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DescribeImageResponseContent) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetImageBuildStatus

`func (o *DescribeImageResponseContent) GetImageBuildStatus() ImageBuildStatus`

GetImageBuildStatus returns the ImageBuildStatus field if non-nil, zero value otherwise.

### GetImageBuildStatusOk

`func (o *DescribeImageResponseContent) GetImageBuildStatusOk() (*ImageBuildStatus, bool)`

GetImageBuildStatusOk returns a tuple with the ImageBuildStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuildStatus

`func (o *DescribeImageResponseContent) SetImageBuildStatus(v ImageBuildStatus)`

SetImageBuildStatus sets ImageBuildStatus field to given value.


### GetImageBuildLogsArn

`func (o *DescribeImageResponseContent) GetImageBuildLogsArn() string`

GetImageBuildLogsArn returns the ImageBuildLogsArn field if non-nil, zero value otherwise.

### GetImageBuildLogsArnOk

`func (o *DescribeImageResponseContent) GetImageBuildLogsArnOk() (*string, bool)`

GetImageBuildLogsArnOk returns a tuple with the ImageBuildLogsArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuildLogsArn

`func (o *DescribeImageResponseContent) SetImageBuildLogsArn(v string)`

SetImageBuildLogsArn sets ImageBuildLogsArn field to given value.

### HasImageBuildLogsArn

`func (o *DescribeImageResponseContent) HasImageBuildLogsArn() bool`

HasImageBuildLogsArn returns a boolean if a field has been set.

### GetCloudformationStackStatus

`func (o *DescribeImageResponseContent) GetCloudformationStackStatus() CloudFormationStackStatus`

GetCloudformationStackStatus returns the CloudformationStackStatus field if non-nil, zero value otherwise.

### GetCloudformationStackStatusOk

`func (o *DescribeImageResponseContent) GetCloudformationStackStatusOk() (*CloudFormationStackStatus, bool)`

GetCloudformationStackStatusOk returns a tuple with the CloudformationStackStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudformationStackStatus

`func (o *DescribeImageResponseContent) SetCloudformationStackStatus(v CloudFormationStackStatus)`

SetCloudformationStackStatus sets CloudformationStackStatus field to given value.

### HasCloudformationStackStatus

`func (o *DescribeImageResponseContent) HasCloudformationStackStatus() bool`

HasCloudformationStackStatus returns a boolean if a field has been set.

### GetCloudformationStackStatusReason

`func (o *DescribeImageResponseContent) GetCloudformationStackStatusReason() string`

GetCloudformationStackStatusReason returns the CloudformationStackStatusReason field if non-nil, zero value otherwise.

### GetCloudformationStackStatusReasonOk

`func (o *DescribeImageResponseContent) GetCloudformationStackStatusReasonOk() (*string, bool)`

GetCloudformationStackStatusReasonOk returns a tuple with the CloudformationStackStatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudformationStackStatusReason

`func (o *DescribeImageResponseContent) SetCloudformationStackStatusReason(v string)`

SetCloudformationStackStatusReason sets CloudformationStackStatusReason field to given value.

### HasCloudformationStackStatusReason

`func (o *DescribeImageResponseContent) HasCloudformationStackStatusReason() bool`

HasCloudformationStackStatusReason returns a boolean if a field has been set.

### GetCloudformationStackArn

`func (o *DescribeImageResponseContent) GetCloudformationStackArn() string`

GetCloudformationStackArn returns the CloudformationStackArn field if non-nil, zero value otherwise.

### GetCloudformationStackArnOk

`func (o *DescribeImageResponseContent) GetCloudformationStackArnOk() (*string, bool)`

GetCloudformationStackArnOk returns a tuple with the CloudformationStackArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudformationStackArn

`func (o *DescribeImageResponseContent) SetCloudformationStackArn(v string)`

SetCloudformationStackArn sets CloudformationStackArn field to given value.

### HasCloudformationStackArn

`func (o *DescribeImageResponseContent) HasCloudformationStackArn() bool`

HasCloudformationStackArn returns a boolean if a field has been set.

### GetCreationTime

`func (o *DescribeImageResponseContent) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *DescribeImageResponseContent) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *DescribeImageResponseContent) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *DescribeImageResponseContent) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCloudformationStackCreationTime

`func (o *DescribeImageResponseContent) GetCloudformationStackCreationTime() time.Time`

GetCloudformationStackCreationTime returns the CloudformationStackCreationTime field if non-nil, zero value otherwise.

### GetCloudformationStackCreationTimeOk

`func (o *DescribeImageResponseContent) GetCloudformationStackCreationTimeOk() (*time.Time, bool)`

GetCloudformationStackCreationTimeOk returns a tuple with the CloudformationStackCreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudformationStackCreationTime

`func (o *DescribeImageResponseContent) SetCloudformationStackCreationTime(v time.Time)`

SetCloudformationStackCreationTime sets CloudformationStackCreationTime field to given value.

### HasCloudformationStackCreationTime

`func (o *DescribeImageResponseContent) HasCloudformationStackCreationTime() bool`

HasCloudformationStackCreationTime returns a boolean if a field has been set.

### GetCloudformationStackTags

`func (o *DescribeImageResponseContent) GetCloudformationStackTags() []Tag`

GetCloudformationStackTags returns the CloudformationStackTags field if non-nil, zero value otherwise.

### GetCloudformationStackTagsOk

`func (o *DescribeImageResponseContent) GetCloudformationStackTagsOk() (*[]Tag, bool)`

GetCloudformationStackTagsOk returns a tuple with the CloudformationStackTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudformationStackTags

`func (o *DescribeImageResponseContent) SetCloudformationStackTags(v []Tag)`

SetCloudformationStackTags sets CloudformationStackTags field to given value.

### HasCloudformationStackTags

`func (o *DescribeImageResponseContent) HasCloudformationStackTags() bool`

HasCloudformationStackTags returns a boolean if a field has been set.

### GetImageConfiguration

`func (o *DescribeImageResponseContent) GetImageConfiguration() ImageConfigurationStructure`

GetImageConfiguration returns the ImageConfiguration field if non-nil, zero value otherwise.

### GetImageConfigurationOk

`func (o *DescribeImageResponseContent) GetImageConfigurationOk() (*ImageConfigurationStructure, bool)`

GetImageConfigurationOk returns a tuple with the ImageConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageConfiguration

`func (o *DescribeImageResponseContent) SetImageConfiguration(v ImageConfigurationStructure)`

SetImageConfiguration sets ImageConfiguration field to given value.


### GetImagebuilderImageStatus

`func (o *DescribeImageResponseContent) GetImagebuilderImageStatus() ImageBuilderImageStatus`

GetImagebuilderImageStatus returns the ImagebuilderImageStatus field if non-nil, zero value otherwise.

### GetImagebuilderImageStatusOk

`func (o *DescribeImageResponseContent) GetImagebuilderImageStatusOk() (*ImageBuilderImageStatus, bool)`

GetImagebuilderImageStatusOk returns a tuple with the ImagebuilderImageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagebuilderImageStatus

`func (o *DescribeImageResponseContent) SetImagebuilderImageStatus(v ImageBuilderImageStatus)`

SetImagebuilderImageStatus sets ImagebuilderImageStatus field to given value.

### HasImagebuilderImageStatus

`func (o *DescribeImageResponseContent) HasImagebuilderImageStatus() bool`

HasImagebuilderImageStatus returns a boolean if a field has been set.

### GetImagebuilderImageStatusReason

`func (o *DescribeImageResponseContent) GetImagebuilderImageStatusReason() string`

GetImagebuilderImageStatusReason returns the ImagebuilderImageStatusReason field if non-nil, zero value otherwise.

### GetImagebuilderImageStatusReasonOk

`func (o *DescribeImageResponseContent) GetImagebuilderImageStatusReasonOk() (*string, bool)`

GetImagebuilderImageStatusReasonOk returns a tuple with the ImagebuilderImageStatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagebuilderImageStatusReason

`func (o *DescribeImageResponseContent) SetImagebuilderImageStatusReason(v string)`

SetImagebuilderImageStatusReason sets ImagebuilderImageStatusReason field to given value.

### HasImagebuilderImageStatusReason

`func (o *DescribeImageResponseContent) HasImagebuilderImageStatusReason() bool`

HasImagebuilderImageStatusReason returns a boolean if a field has been set.

### GetEc2AmiInfo

`func (o *DescribeImageResponseContent) GetEc2AmiInfo() Ec2AmiInfo`

GetEc2AmiInfo returns the Ec2AmiInfo field if non-nil, zero value otherwise.

### GetEc2AmiInfoOk

`func (o *DescribeImageResponseContent) GetEc2AmiInfoOk() (*Ec2AmiInfo, bool)`

GetEc2AmiInfoOk returns a tuple with the Ec2AmiInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEc2AmiInfo

`func (o *DescribeImageResponseContent) SetEc2AmiInfo(v Ec2AmiInfo)`

SetEc2AmiInfo sets Ec2AmiInfo field to given value.

### HasEc2AmiInfo

`func (o *DescribeImageResponseContent) HasEc2AmiInfo() bool`

HasEc2AmiInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


