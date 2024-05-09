/**
 *  Copyright 2024 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"). You may not
 *  use this file except in compliance with the License. A copy of the License is
 *  located at
 *
 *  http://aws.amazon.com/apache2.0/
 *
 *  or in the "LICENSE.txt" file accompanying this file. This file is distributed
 *  on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, express or
 *  implied. See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  # Required Infrastructure Submodule
 *  The required infra submodule deploys a vpc, subnets, routes, gateways, and creates a
 *  key pair. These are necessary resources for the API to deploy and manage clusters.
 */

locals {
  test_config = {
    Region = var.region
    Image = {
      Os = "alinux2"
    }
    HeadNode = {
      InstanceType = var.head_node_type
      Networking = {
        SubnetId  = var.subnet != null ? var.subnet : aws_default_subnet.public_az1.id
      }
      Iam = {
        AdditionalIamPolicies = [
          {
            Policy: "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
          }
        ]
      }
    }
    Scheduling = {
      Scheduler = "slurm"
      SlurmQueues = [
        {
          Name         = "queue1"
          CapacityType = "ONDEMAND"
          Networking = {
            SubnetIds      = [var.subnet != null ? var.subnet : aws_default_subnet.public_az1.id]
          }
          Iam = {
            AdditionalIamPolicies = [
              {
                Policy: "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
              }
            ]
          }
          ComputeResources = [
            {
              Name                              = "compute"
              InstanceType                      = var.compute_node_type
              MinCount                          = var.min_nodes
              MaxCount                          = var.max_nodes
            }
          ]
        }
      ]
      SlurmSettings = {
        QueueUpdateStrategy = "TERMINATE"
      }
    }
  }
}

resource "pcluster_cluster" "test" {
  cluster_name        = var.cluster_name
  region              = var.region
  suppress_validators = []

  cluster_configuration = yamlencode(local.test_config)

  depends_on = [
    aws_default_vpc.default,
    aws_default_subnet.public_az1,
  ]
}

resource "pcluster_compute_fleet_status" "test" {
  cluster_name   = var.cluster_name
  status_request = "START_REQUESTED"
  depends_on     = [pcluster_cluster.test]
}
