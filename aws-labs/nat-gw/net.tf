locals {
  availability_zones = ["us-east-1a", "us-east-1b", "us-east-1c"]
}

resource "aws_vpc" "main" {
  cidr_block = "10.16.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
}

resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.main.id
}

resource "aws_subnet" "db_subnet" {
  count             = length(local.availability_zones)
  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(aws_vpc.main.cidr_block, 4, count.index)
  availability_zone = local.availability_zones[count.index]
}

resource "aws_subnet" "app_subnet" {
  count             = length(local.availability_zones)
  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(aws_vpc.main.cidr_block, 4, 3 + count.index)
  availability_zone = local.availability_zones[count.index]
}

resource "aws_subnet" "web_subnet" {
  count             = length(local.availability_zones)
  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(aws_vpc.main.cidr_block, 4, 6 + count.index)
  availability_zone = local.availability_zones[count.index]
}



resource "aws_nat_gateway" "nat_gw_web" {
  count         = length(local.availability_zones)
  allocation_id = aws_eip.nat_eip_web.id
  subnet_id     = aws_subnet.web_subnet.id
}

resource "aws_nat_gateway" "nat_gw_db" {
  count         = length(local.availability_zones)
  allocation_id = aws_eip.nat_eip_db.id
  subnet_id     = aws_subnet.db_subnet.id
}

resource "aws_nat_gateway" "nat_gw_app" {
  count         = length(local.availability_zones)
  allocation_id = aws_eip.nat_eip_app.id
  subnet_id     = aws_subnet.app_subnet.id
}

resource "aws_eip" "nat_eip_web" {
  vpc = true
}

resource "aws_eip" "nat_eip_db" {
  vpc = true
}

resource "aws_eip" "nat_eip_app" {
  vpc = true
}

## Web association
resource "aws_route_table" "web_route_table" {
  count = length(local.availability_zones)
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw.id
  }
}

resource "aws_route_table_association" "web_subnet_route_table_association" {
  count          = length(local.availability_zones)
  subnet_id      = aws_subnet.web_subnet.id
  route_table_id = aws_route_table.web_route_table.id
}

###

# Route Tables and Associations for APP Subnets (with NAT Gateways)
resource "aws_route_table" "app_route_table" {
  count  = length(local.availability_zones)
  vpc_id = aws_vpc.main.id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.nat_gw_app.id
  }
}

resource "aws_route_table_association" "app_route_table_association" {
  count          = length(local.availability_zones)
  subnet_id      = aws_subnet.app_subnet.id
  route_table_id = aws_route_table.app_route_table.id
}

# Route Tables for DB Subnets (no NAT Gateway as these are private)
resource "aws_route_table" "db_route_table" {
  count  = length(local.availability_zones)
  vpc_id = aws_vpc.main.id
}

resource "aws_route_table_association" "db_route_table_association" {
  count          = length(local.availability_zones)
  subnet_id      = aws_subnet.db_subnet.id
  route_table_id = aws_route_table.db_route_table.id
}
