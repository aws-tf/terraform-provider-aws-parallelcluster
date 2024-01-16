---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pcluster_compute_fleet_status Data Source - pcluster"
subcategory: ""
description: |-
  Describe the status of the compute fleet.
---

# pcluster_compute_fleet_status (Data Source)

Describe the status of the compute fleet.

## Example Usage

```terraform
data "pcluster_compute_fleet_status" "example" {
  cluster_name = var.cluster_name
  region       = var.region
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_name` (String) Filter by architecture. The default is no filtering.

### Optional

- `region` (String) The AWS Region that official images are listed in.

### Read-Only

- `last_status_update_time` (String) The timestamp representing the last status update time.
- `status` (String) The status of the compute fleet.