# CHANGELOG

## 1.2.0

**CHANGES**

* Upgrade Go minimum version to 1.23.0 (from 1.21).
* Upgrade aws-sdk-go-v2 to v1.36.3 (from v1.32.7).
* Upgrade aws-sdk-go-v2/config to v1.29.13 (from v1.28.7).
* Upgrade aws-sdk-go-v2/service/apigateway to v1.30.1 (from v1.28.2).
* Upgrade aws-sdk-go-v2/service/cloudformation to v1.59.1 (from v1.56.2).
* Upgrade aws-sdk-go-v2/service/sts to v1.33.18 (from v1.33.3).
* Upgrade terraform-plugin-framework to v1.14.1 (from v1.9.0).
* Upgrade terraform-plugin-go to v0.26.0 (from v0.23.0).
* Upgrade terraform-plugin-sdk/v2 to v2.36.1 (from v2.34.0).

## 1.1.0

**BUG FIX**

* Fix an issue that was causing terraform-apply failure when ParallelCluster API 3.11.x is used to deploy clusters with login nodes.

## 1.0.0

**CHANGES**

First official release of the AWS ParallelCluster Provider for Terraform, with support for ParallelCluster 3.8.0+.
With this release the user can deploy ParallelCluster clusters and build custom AMIs through an existing ParallelCluster API.
