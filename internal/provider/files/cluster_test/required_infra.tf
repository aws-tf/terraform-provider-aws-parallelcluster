resource "aws_subnet" "pcluster-test-subnet" {
  vpc_id     = aws_vpc.pcluster-test-vpc.id
  cidr_block = "10.0.111.0/24"

  tags = {
    Name = "pcluster-test-subnet"
  }
}

resource "aws_default_vpc" "default" {
  tags = {
    Name = "Default VPC"
  }
}

resource "aws_vpc" "pcluster-test-vpc" {
  cidr_block       = "10.0.0.0/16"
  instance_tenancy = "default"

  tags = {
    Name = "pcluster-test-vpc"
  }
}

resource "aws_internet_gateway_attachment" "pcluster-test-inetenet-gateway-attachment" {
  internet_gateway_id = aws_internet_gateway.pcluster-test-internet-gateway.id
  vpc_id              = aws_vpc.pcluster-test-vpc.id
}

resource "aws_internet_gateway" "pcluster-test-internet-gateway" {}

resource "aws_default_route_table" "pcluster-test-route-table" {
  default_route_table_id = aws_vpc.pcluster-test-vpc.default_route_table_id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.pcluster-test-internet-gateway.id
  }

  tags = {
    Name = "pcluster-test-route-table"
  }
}

resource "aws_key_pair" "pcluster-test-key-pair" {

  key_name_prefix = "pcluster"
  public_key      = trimspace(tls_private_key.pcluster-test-private-key.public_key_openssh)
  tags = {
    Name = "pcluster-test-key-pair"
  }
}

resource "tls_private_key" "pcluster-test-private-key" {
  algorithm = "RSA"
  rsa_bits  = 4096
}
