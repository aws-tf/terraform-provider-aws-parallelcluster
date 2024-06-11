locals {
  example_config = {
    Region = var.region
    Image = {
      Os = "alinux2"
    }
    HeadNode = {
      InstanceType = var.head_node_type
      Networking = {
        ElasticIp = true
        SubnetId  = var.subnet
      }
      Ssh = {
        KeyName = var.key_pair
      }
    }
    Scheduling = {
      Scheduler = "slurm"
      SlurmQueues = [
        {
          Name         = "queue1"
          CapacityType = "ONDEMAND"
          Networking = {
            SubnetIds      = [var.subnet]
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

resource "aws-parallelcluster_cluster" "example" {
  cluster_name             = var.cluster_name
  region                   = var.region
  rollback_on_failure      = true
  validation_failure_level = "WARNING"
  suppress_validators      = []

  cluster_configuration = yamlencode(local.example_config)
}
