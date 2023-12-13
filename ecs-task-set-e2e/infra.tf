locals {
  name = "mo-sandbox"

  vpc_id = "vpc-<VPC_ID>"

  subnet_ids = [
    "subnet-<ID>",
    "subnet-<ID>",
    "subnet-<ID>",
  ]

  tags = {
    Name        = local.name
    Role        = local.name
    Environment = "sandbox"
    OwningTeam  = "TEAM_NAME"
  }
}

## Base infrastructure

resource "aws_ecs_cluster" "main" {
  name = local.name

  configuration {
    execute_command_configuration {
      logging = "OVERRIDE"
      log_configuration {
        cloud_watch_log_group_name = aws_cloudwatch_log_group.main.name
      }
    }
  }

  setting {
    name  = "containerInsights"
    value = "enabled"
  }

  tags = local.tags
}

resource "aws_ecs_cluster_capacity_providers" "main" {
  cluster_name = aws_ecs_cluster.main.name

  capacity_providers = ["FARGATE"]

  default_capacity_provider_strategy {
    base              = 1
    weight            = 100
    capacity_provider = "FARGATE"
  }
}

resource "aws_cloudwatch_log_group" "main" {
  name              = local.name
  retention_in_days = 7
  tags              = local.tags
}

resource "aws_service_discovery_public_dns_namespace" "main" {
  name        = "${local.name}.sandbox.<hosted_zone>.services"
  description = "PoC for task sets and deployments."
  tags        = local.tags
}

data "aws_route53_zone" "sandbox_<hosted_zone>_services" {
  zone_id = aws_service_discovery_public_dns_namespace.main.hosted_zone
}

# Need to create zone ID
resource "aws_route53_record" "discovery_namespace" {
  zone_id         = "<ZONE_ID>"
  allow_overwrite = true
  name            = "${local.name}.sandbox.<hosted_zone>.services"
  type            = "NS"
  ttl             = 300
  records         = data.aws_route53_zone.sandbox_<hosted_zone>_services.name_servers
}


