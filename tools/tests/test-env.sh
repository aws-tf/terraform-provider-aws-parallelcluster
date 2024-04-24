#!/bin/bash

# Variables you can set in your local test environment
#
# TEST_REGION:                    AWS region;
#                                 This is the region where the tests will run.
#                                 Default is AWS_DEFAULT_REGION or us-east-1 if AWS_DEFAULT_REGION is not set.
#
# TEST_AVAILABILITY_ZONE:         AWS availability zone;
#                                 This is the availability zone where the tests will run.
#                                 Default is us-east-1a.
#
# TEST_PCAPI_STACK_NAME:          CloudFormation stack name;
#                                 This is the name of the existing ParallelCluster API stack that tests will use.
#                                 Default is ParallelCluster.
#
# TEST_USE_USER_ROLE:             [true|false];
#                                 If true, the tests will use the user role ParallelClusterApiUserRole returned by the ParallelCluster API.
#                                 Default is false.
#
# TEST_CLUSTER_NAME:              [Cluster name];
#                                 This is the name of the cluster that tests will create.
#                                 Default is test-cluster.

export TEST_REGION="us-east-1"
export TEST_AVAILABILITY_ZONE="us-east-1a"
export TEST_PCAPI_STACK_NAME="ParallelCluster"
export TEST_USE_USER_ROLE="true"
export TEST_CLUSTER_NAME="test-cluster-$(whoami)"
