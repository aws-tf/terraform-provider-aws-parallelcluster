locals {
  test_config = {
    Region = var.region
    Image = {
      Os = "alinux2"
    }
    HeadNode = {
      InstanceType = var.head_node_type
      Networking = {
        ElasticIp = true
        SubnetId  = var.subnet != null ? var.subnet : aws_subnet.pcluster-test-subnet.id
      }
      Ssh = {
        KeyName = var.key_pair != null ? var.key_pair : aws_key_pair.pcluster-test-key-pair.id
      }
    }
    Scheduling = {
      Scheduler = "slurm"
      SlurmQueues = [
        {
          Name         = "queue1"
          CapacityType = "ONDEMAND"
          Networking = {
            SubnetIds      = [var.subnet != null ? var.subnet : aws_subnet.pcluster-test-subnet.id]
            AssignPublicIp = true
          }
          ComputeResources = [
            {
              Name                              = "compute"
              InstanceType                      = var.compute_node_type
              MinCount                          = var.min_nodes
              MaxCount                          = var.max_nodes
              DisableSimultaneousMultithreading = true
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
  cluster_name             = var.cluster_name
  region                   = var.region
  suppress_validators      = []

  cluster_configuration = yamlencode(local.test_config)

  depends_on = [
    aws_default_route_table.pcluster-test-route-table,
    aws_default_vpc.default,
    aws_internet_gateway.pcluster-test-internet-gateway,
    aws_internet_gateway_attachment.pcluster-test-inetenet-gateway-attachment,
    aws_key_pair.pcluster-test-key-pair,
    aws_subnet.pcluster-test-subnet,
    aws_vpc.pcluster-test-vpc
  ]
}

resource "pcluster_compute_fleet_status" "test" {
  cluster_name   = var.cluster_name
  status_request = "START_REQUESTED"
  depends_on     = [pcluster_cluster.test]
}
