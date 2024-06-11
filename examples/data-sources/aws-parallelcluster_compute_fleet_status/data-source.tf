data "aws-parallelcluster_compute_fleet_status" "example" {
  cluster_name = var.cluster_name
  region       = var.region
}
