resource "pcluster_compute_fleet_status" "example_fleet" {
  cluster_name   = var.cluster_name
  status_request = "START_REQUESTED"
}
