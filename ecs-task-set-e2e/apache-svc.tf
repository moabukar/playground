locals {
  apache_name = "${local.name}-apache"
}

resource "aws_ecs_task_set" "apache" {
  service         = aws_ecs_service.both.id
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.apache.arn

  # load_balancer {
  #   target_group_arn = aws_lb_target_group.apache.arn
  #   container_name   = "apache"
  #   container_port   = 80
  # }

  external_id      = "alpha"
  force_delete     = true
  launch_type      = "FARGATE"
  platform_version = "1.4.0"

  network_configuration {
    subnets          = local.subnet_ids
    security_groups  = [aws_security_group.apache.id]
    assign_public_ip = false
  }

  scale {
    unit  = "PERCENT"
    value = 50
  }

  service_registries {
    registry_arn = aws_service_discovery_service.apache.arn
  }

  tags = local.tags
}

# resource "aws_ecs_service" "apache" {
#   name    = local.apache_name
#   cluster = aws_ecs_cluster.main.id

#   deployment_controller {
#     type = "EXTERNAL"
#   }

#   desired_count           = 2
#   enable_ecs_managed_tags = true
#   enable_execute_command  = true
#   force_new_deployment    = false
#   propagate_tags          = "SERVICE"

#   service_connect_configuration {
#     enabled = false
#   }

#   tags = local.tags
# }

resource "aws_service_discovery_service" "apache" {
  name        = "apache"
  description = "Service discovery for apache service."

  dns_config {
    namespace_id = aws_service_discovery_public_dns_namespace.main.id

    dns_records {
      ttl  = 10
      type = "A"
    }

    routing_policy = "MULTIVALUE"
  }

  force_destroy = true

  health_check_custom_config {
    failure_threshold = 1
  }

  tags = local.tags
}

resource "aws_ecs_task_definition" "apache" {
  family = local.apache_name
  # container_definitions = jsonencode(local.apache_container_definition)
  container_definitions = jsonencode([
    {
      name      = "apache"
      image     = "AWS_ACCOUNT_ID.dkr.ecr.eu-west-1.amazonaws.com/mo-apache:v2"
      cpu       = 10
      memory    = 512
      essential = true
      portMappings = [
        {
          containerPort = 80
          hostPort      = 80
        }
      ]
    }
  ])
  cpu                = 256
  execution_role_arn = aws_iam_role.apache_exec.arn
  memory             = 512
  network_mode       = "awsvpc"

  runtime_platform {
    operating_system_family = "LINUX"
    cpu_architecture        = "ARM64"
  }

  requires_compatibilities = ["FARGATE"]
  tags                     = local.tags
  task_role_arn            = aws_iam_role.apache_task.arn
}

## LB

# resource "aws_lb" "apache" {
#   name                             = local.apache_name
#   internal                         = true
#   load_balancer_type               = "network"
#   subnets                          = local.subnet_ids
#   ip_address_type                  = "ipv4"
#   enable_cross_zone_load_balancing = true

#   tags = local.tags
# }

# resource "aws_lb_target_group" "apache" {
#   connection_termination = true
#   deregistration_delay   = 30
#   name                   = local.apache_name
#   port                   = 80
#   protocol               = "HTTP"
#   tags                   = local.tags
#   target_type            = "ip"
#   vpc_id                 = local.vpc_id

#   health_check {
#     enabled             = true
#     healthy_threshold   = 3
#     interval            = 5
#     protocol            = "HTTP"
#     timeout             = 5
#     unhealthy_threshold = 3
#   }
# }

# resource "aws_lb_listener" "apache" {
#   load_balancer_arn = aws_lb.both.arn
#   port              = "80"
#   protocol          = "TCP"

#   default_action {
#     type             = "forward"
#     target_group_arn = aws_lb_target_group.apache.arn
#   }

#   tags = local.tags
# }

# resource "aws_route53_record" "apache_lb" {
#   zone_id = "Z048359310Y6AD9XSL491"
#   name    = "mo-apache-lb.sandbox.<hosted_zone>.services"
#   type    = "A"

#   alias {
#     name                   = aws_lb.both.dns_name
#     zone_id                = aws_lb.both.zone_id
#     evaluate_target_health = true
#   }
# }
