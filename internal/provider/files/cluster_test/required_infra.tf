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

resource "aws_default_vpc" "default" {
  tags = {
    Name = "Default VPC"
  }
}

resource "aws_default_subnet" "public_az1" {
  availability_zone = "${var.region}a"
}
