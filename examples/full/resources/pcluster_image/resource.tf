/** Copyright 2024 Amazon.com, Inc. or its affiliates. All Rights Reserved.
*
* Licensed under the Apache License, Version 2.0 (the "License"). You may not
* use this file except in compliance with the License. A copy of the License is
* located at
*
* http://aws.amazon.com/apache2.0/
*
* or in the "LICENSE.txt" file accompanying this file. This file is distributed
* on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, express or
* implied. See the License for the specific language governing permissions and
* limitations under the License.
*/

resource "aws_vpc" "build" {
  enable_dns_support   = true
  enable_dns_hostnames = true
  cidr_block           = "10.0.0.0/16"

  tags = {
    Name = "ImageBuildTestVpc"
  }
}

resource "aws_subnet" "private_subnet" {
  vpc_id                  = aws_vpc.build.id
  cidr_block              = "10.0.2.0/24"
  availability_zone       = var.default_az
  map_public_ip_on_launch = true
  tags = {
    Name = "ImageBuildTest"
  }
}

resource "aws_security_group" "build" {
  vpc_id = aws_vpc.build.id

  egress {
    from_port        = 443
    to_port          = 443
    protocol         = "tcp"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }
}

resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.build.id

  tags = {
    Name = "ImageBuildTest"
  }
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.build.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw.id
  }

  tags = {
    Name = "image-build-test-public-route-table"
  }
}

resource "aws_route_table_association" "public" {
  subnet_id      = aws_subnet.private_subnet.id
  route_table_id = aws_route_table.public.id
}

data "pcluster_list_official_images" "parent_image" {
  region       = var.region
  os           = "alinux2"
  architecture = "x86_64"
}

resource "pcluster_image" "build-demo" {
  image_id            = "imageBuilderDemo2"
  rollback_on_failure = false
  image_configuration = yamlencode({
    "Build" : {
      "InstanceType" : "c5.2xlarge",
      "ParentImage" : data.pcluster_list_official_images.parent_image.official_images[0].amiId,
      "SubnetId" : aws_subnet.private_subnet.id,
      "SecurityGroupIds" : [aws_security_group.build.id],
      "UpdateOsPackages" : { "Enabled" : false }
    }
  })

}
