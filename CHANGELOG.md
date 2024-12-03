# CHANGELOG

## 1.1.0

**BUG FIX**

* Fix an issue that was causing terraform-apply failure when ParallelCluster API 3.11.x is used to deploy clusters with login nodes.
  https://github.com/aws/aws-parallelcluster/issues/6489

## 1.0.0

**CHANGES**

First official release of the AWS ParallelCluster Provider for Terraform, with support for ParallelCluster 3.8.0+.
With this release the user can deploy ParallelCluster clusters and build custom AMIs through an existing ParallelCluster API.
