# DescribeClusterResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** | Name of the cluster. | 
**Region** | **string** | AWS region where the cluster is created. | 
**Version** | **string** | ParallelCluster version used to create the cluster. | 
**CloudFormationStackStatus** | [**CloudFormationStackStatus**](CloudFormationStackStatus.md) |  | 
**ClusterStatus** | [**ClusterStatus**](ClusterStatus.md) |  | 
**Scheduler** | Pointer to [**Scheduler**](Scheduler.md) |  | [optional] 
**CloudformationStackArn** | **string** | ARN of the main CloudFormation stack. | 
**CreationTime** | **time.Time** | Timestamp representing the cluster creation time. | 
**LastUpdatedTime** | **time.Time** | Timestamp representing the last cluster update time. | 
**ClusterConfiguration** | [**ClusterConfigurationStructure**](ClusterConfigurationStructure.md) |  | 
**ComputeFleetStatus** | [**ComputeFleetStatus**](ComputeFleetStatus.md) |  | 
**Tags** | [**[]Tag**](Tag.md) | Tags associated with the cluster. | 
**HeadNode** | Pointer to [**EC2Instance**](EC2Instance.md) |  | [optional] 
**LoginNodes** | Pointer to [**[]LoginNodesPool**](LoginNodesPool.md) |  | [optional] 
**Failures** | Pointer to [**[]Failure**](Failure.md) | Failures array containing failures reason and code when the stack is in CREATE_FAILED status. | [optional] 

## Methods

### NewDescribeClusterResponseContent

`func NewDescribeClusterResponseContent(clusterName string, region string, version string, cloudFormationStackStatus CloudFormationStackStatus, clusterStatus ClusterStatus, cloudformationStackArn string, creationTime time.Time, lastUpdatedTime time.Time, clusterConfiguration ClusterConfigurationStructure, computeFleetStatus ComputeFleetStatus, tags []Tag, ) *DescribeClusterResponseContent`

NewDescribeClusterResponseContent instantiates a new DescribeClusterResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeClusterResponseContentWithDefaults

`func NewDescribeClusterResponseContentWithDefaults() *DescribeClusterResponseContent`

NewDescribeClusterResponseContentWithDefaults instantiates a new DescribeClusterResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *DescribeClusterResponseContent) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *DescribeClusterResponseContent) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *DescribeClusterResponseContent) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetRegion

`func (o *DescribeClusterResponseContent) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DescribeClusterResponseContent) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DescribeClusterResponseContent) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVersion

`func (o *DescribeClusterResponseContent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DescribeClusterResponseContent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DescribeClusterResponseContent) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetCloudFormationStackStatus

`func (o *DescribeClusterResponseContent) GetCloudFormationStackStatus() CloudFormationStackStatus`

GetCloudFormationStackStatus returns the CloudFormationStackStatus field if non-nil, zero value otherwise.

### GetCloudFormationStackStatusOk

`func (o *DescribeClusterResponseContent) GetCloudFormationStackStatusOk() (*CloudFormationStackStatus, bool)`

GetCloudFormationStackStatusOk returns a tuple with the CloudFormationStackStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudFormationStackStatus

`func (o *DescribeClusterResponseContent) SetCloudFormationStackStatus(v CloudFormationStackStatus)`

SetCloudFormationStackStatus sets CloudFormationStackStatus field to given value.


### GetClusterStatus

`func (o *DescribeClusterResponseContent) GetClusterStatus() ClusterStatus`

GetClusterStatus returns the ClusterStatus field if non-nil, zero value otherwise.

### GetClusterStatusOk

`func (o *DescribeClusterResponseContent) GetClusterStatusOk() (*ClusterStatus, bool)`

GetClusterStatusOk returns a tuple with the ClusterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStatus

`func (o *DescribeClusterResponseContent) SetClusterStatus(v ClusterStatus)`

SetClusterStatus sets ClusterStatus field to given value.


### GetScheduler

`func (o *DescribeClusterResponseContent) GetScheduler() Scheduler`

GetScheduler returns the Scheduler field if non-nil, zero value otherwise.

### GetSchedulerOk

`func (o *DescribeClusterResponseContent) GetSchedulerOk() (*Scheduler, bool)`

GetSchedulerOk returns a tuple with the Scheduler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduler

`func (o *DescribeClusterResponseContent) SetScheduler(v Scheduler)`

SetScheduler sets Scheduler field to given value.

### HasScheduler

`func (o *DescribeClusterResponseContent) HasScheduler() bool`

HasScheduler returns a boolean if a field has been set.

### GetCloudformationStackArn

`func (o *DescribeClusterResponseContent) GetCloudformationStackArn() string`

GetCloudformationStackArn returns the CloudformationStackArn field if non-nil, zero value otherwise.

### GetCloudformationStackArnOk

`func (o *DescribeClusterResponseContent) GetCloudformationStackArnOk() (*string, bool)`

GetCloudformationStackArnOk returns a tuple with the CloudformationStackArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudformationStackArn

`func (o *DescribeClusterResponseContent) SetCloudformationStackArn(v string)`

SetCloudformationStackArn sets CloudformationStackArn field to given value.


### GetCreationTime

`func (o *DescribeClusterResponseContent) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *DescribeClusterResponseContent) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *DescribeClusterResponseContent) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetLastUpdatedTime

`func (o *DescribeClusterResponseContent) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *DescribeClusterResponseContent) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *DescribeClusterResponseContent) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetClusterConfiguration

`func (o *DescribeClusterResponseContent) GetClusterConfiguration() ClusterConfigurationStructure`

GetClusterConfiguration returns the ClusterConfiguration field if non-nil, zero value otherwise.

### GetClusterConfigurationOk

`func (o *DescribeClusterResponseContent) GetClusterConfigurationOk() (*ClusterConfigurationStructure, bool)`

GetClusterConfigurationOk returns a tuple with the ClusterConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterConfiguration

`func (o *DescribeClusterResponseContent) SetClusterConfiguration(v ClusterConfigurationStructure)`

SetClusterConfiguration sets ClusterConfiguration field to given value.


### GetComputeFleetStatus

`func (o *DescribeClusterResponseContent) GetComputeFleetStatus() ComputeFleetStatus`

GetComputeFleetStatus returns the ComputeFleetStatus field if non-nil, zero value otherwise.

### GetComputeFleetStatusOk

`func (o *DescribeClusterResponseContent) GetComputeFleetStatusOk() (*ComputeFleetStatus, bool)`

GetComputeFleetStatusOk returns a tuple with the ComputeFleetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeFleetStatus

`func (o *DescribeClusterResponseContent) SetComputeFleetStatus(v ComputeFleetStatus)`

SetComputeFleetStatus sets ComputeFleetStatus field to given value.


### GetTags

`func (o *DescribeClusterResponseContent) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DescribeClusterResponseContent) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DescribeClusterResponseContent) SetTags(v []Tag)`

SetTags sets Tags field to given value.


### GetHeadNode

`func (o *DescribeClusterResponseContent) GetHeadNode() EC2Instance`

GetHeadNode returns the HeadNode field if non-nil, zero value otherwise.

### GetHeadNodeOk

`func (o *DescribeClusterResponseContent) GetHeadNodeOk() (*EC2Instance, bool)`

GetHeadNodeOk returns a tuple with the HeadNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadNode

`func (o *DescribeClusterResponseContent) SetHeadNode(v EC2Instance)`

SetHeadNode sets HeadNode field to given value.

### HasHeadNode

`func (o *DescribeClusterResponseContent) HasHeadNode() bool`

HasHeadNode returns a boolean if a field has been set.

### GetLoginNodes

`func (o *DescribeClusterResponseContent) GetLoginNodes() []LoginNodesPool`

GetLoginNodes returns the LoginNodes field if non-nil, zero value otherwise.

### GetLoginNodesOk

`func (o *DescribeClusterResponseContent) GetLoginNodesOk() (*[]LoginNodesPool, bool)`

GetLoginNodesOk returns a tuple with the LoginNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginNodes

`func (o *DescribeClusterResponseContent) SetLoginNodes(v []LoginNodesPool)`

SetLoginNodes sets LoginNodes field to given value.

### HasLoginNodes

`func (o *DescribeClusterResponseContent) HasLoginNodes() bool`

HasLoginNodes returns a boolean if a field has been set.

### GetFailures

`func (o *DescribeClusterResponseContent) GetFailures() []Failure`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *DescribeClusterResponseContent) GetFailuresOk() (*[]Failure, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *DescribeClusterResponseContent) SetFailures(v []Failure)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *DescribeClusterResponseContent) HasFailures() bool`

HasFailures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


