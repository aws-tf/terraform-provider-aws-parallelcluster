# ClusterInfoSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** | Name of the cluster. | 
**Region** | **string** | AWS region where the cluster is created. | 
**Version** | **string** | ParallelCluster version used to create the cluster. | 
**CloudformationStackArn** | **string** | ARN of the main CloudFormation stack. | 
**CloudformationStackStatus** | [**CloudFormationStackStatus**](CloudFormationStackStatus.md) |  | 
**ClusterStatus** | [**ClusterStatus**](ClusterStatus.md) |  | 
**Scheduler** | Pointer to [**Scheduler**](Scheduler.md) |  | [optional] 

## Methods

### NewClusterInfoSummary

`func NewClusterInfoSummary(clusterName string, region string, version string, cloudformationStackArn string, cloudformationStackStatus CloudFormationStackStatus, clusterStatus ClusterStatus, ) *ClusterInfoSummary`

NewClusterInfoSummary instantiates a new ClusterInfoSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInfoSummaryWithDefaults

`func NewClusterInfoSummaryWithDefaults() *ClusterInfoSummary`

NewClusterInfoSummaryWithDefaults instantiates a new ClusterInfoSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *ClusterInfoSummary) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ClusterInfoSummary) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ClusterInfoSummary) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetRegion

`func (o *ClusterInfoSummary) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterInfoSummary) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterInfoSummary) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVersion

`func (o *ClusterInfoSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterInfoSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterInfoSummary) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetCloudformationStackArn

`func (o *ClusterInfoSummary) GetCloudformationStackArn() string`

GetCloudformationStackArn returns the CloudformationStackArn field if non-nil, zero value otherwise.

### GetCloudformationStackArnOk

`func (o *ClusterInfoSummary) GetCloudformationStackArnOk() (*string, bool)`

GetCloudformationStackArnOk returns a tuple with the CloudformationStackArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudformationStackArn

`func (o *ClusterInfoSummary) SetCloudformationStackArn(v string)`

SetCloudformationStackArn sets CloudformationStackArn field to given value.


### GetCloudformationStackStatus

`func (o *ClusterInfoSummary) GetCloudformationStackStatus() CloudFormationStackStatus`

GetCloudformationStackStatus returns the CloudformationStackStatus field if non-nil, zero value otherwise.

### GetCloudformationStackStatusOk

`func (o *ClusterInfoSummary) GetCloudformationStackStatusOk() (*CloudFormationStackStatus, bool)`

GetCloudformationStackStatusOk returns a tuple with the CloudformationStackStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudformationStackStatus

`func (o *ClusterInfoSummary) SetCloudformationStackStatus(v CloudFormationStackStatus)`

SetCloudformationStackStatus sets CloudformationStackStatus field to given value.


### GetClusterStatus

`func (o *ClusterInfoSummary) GetClusterStatus() ClusterStatus`

GetClusterStatus returns the ClusterStatus field if non-nil, zero value otherwise.

### GetClusterStatusOk

`func (o *ClusterInfoSummary) GetClusterStatusOk() (*ClusterStatus, bool)`

GetClusterStatusOk returns a tuple with the ClusterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStatus

`func (o *ClusterInfoSummary) SetClusterStatus(v ClusterStatus)`

SetClusterStatus sets ClusterStatus field to given value.


### GetScheduler

`func (o *ClusterInfoSummary) GetScheduler() Scheduler`

GetScheduler returns the Scheduler field if non-nil, zero value otherwise.

### GetSchedulerOk

`func (o *ClusterInfoSummary) GetSchedulerOk() (*Scheduler, bool)`

GetSchedulerOk returns a tuple with the Scheduler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduler

`func (o *ClusterInfoSummary) SetScheduler(v Scheduler)`

SetScheduler sets Scheduler field to given value.

### HasScheduler

`func (o *ClusterInfoSummary) HasScheduler() bool`

HasScheduler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


