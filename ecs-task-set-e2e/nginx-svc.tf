locals {
  nginx_name = "${local.name}-nginx"
  # ecs_service = "both"
}

resource "aws_ecs_task_set" "nginx" {
  service         = aws_ecs_service.both.id
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.nginx.arn

  # load_balancer {
  #   target_group_arn = aws_lb_target_group.nginx.arn
  #   container_name   = "nginx"
  #   container_port   = 80
  # }

  external_id      = "alpha"
  force_delete     = true
  launch_type      = "FARGATE"
  platform_version = "1.4.0"

  network_configuration {
    subnets          = local.subnet_ids
    security_groups  = [aws_security_group.nginx.id]
    assign_public_ip = false
  }

  scale {
    unit  = "PERCENT"
    value = 50
  }

  service_registries {
    registry_arn = aws_service_discovery_service.nginx.arn
  }

  tags = local.tags
}

resource "aws_ecs_service" "both" {
  # name    = local.nginx_name
  name    = "both-svc"
  cluster = aws_ecs_cluster.main.id

  deployment_controller {
    type = "EXTERNAL"
  }

  desired_count           = 2
  enable_ecs_managed_tags = true
  enable_execute_command  = true
  force_new_deployment    = false
  propagate_tags          = "SERVICE"

  service_connect_configuration {
    enabled = false
  }

  tags = local.tags
}

resource "aws_service_discovery_service" "both" {
  name        = "both"
  description = "Service discovery for both service."

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

resource "aws_service_discovery_service" "nginx" {
  name        = "nginx"
  description = "Service discovery for nginx service."

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

resource "aws_ecs_task_definition" "nginx" {
  family = local.nginx_name
  # container_definitions = jsonencode(local.nginx_container_definition)
  container_definitions = jsonencode([
    {
      name      = "nginx"
      image     = "<AWS_ACCOUNT_ID>.dkr.ecr.eu-west-1.amazonaws.com/mo-nginx:v2"
      cpu       = 10
      memory    = 512
      essential = true
      portMappings = [
        {
          containerPort = 80
          hostPort      = 80
        }
      ]
      healthCheck = {
        command     = ["CMD-SHELL", "/usr/bin/curl -f http://localhost || exit 1 "]
        interval    = 5
        retries     = 3
        startPeriod = 15
        timeout     = 2
      }
    }
  ])
  cpu                = 256
  execution_role_arn = aws_iam_role.nginx_exec.arn
  memory             = 512
  network_mode       = "awsvpc"

  runtime_platform {
    operating_system_family = "LINUX"
    cpu_architecture        = "ARM64"
  }

  requires_compatibilities = ["FARGATE"]
  tags                     = local.tags
  task_role_arn            = aws_iam_role.nginx_task.arn
}

## LB

# resource "aws_lb" "both" {
#   name                             = local.nginx_name
#   internal                         = true
#   load_balancer_type               = "network"
#   subnets                          = local.subnet_ids
#   ip_address_type                  = "ipv4"
#   enable_cross_zone_load_balancing = true

#   tags = local.tags
# }

# resource "aws_lb_target_group" "nginx" {
#   connection_termination = true
#   deregistration_delay   = 30
#   name                   = local.nginx_name
#   port                   = 80
#   protocol               = "TCP"
#   tags                   = local.tags
#   target_type            = "ip"
#   vpc_id                 = local.vpc_id

# health_check {
#   # path                = "/ping"
#   enabled             = true
#   healthy_threshold   = 3
#   interval            = 5
#   protocol            = "TCP"
#   timeout             = 5
#   unhealthy_threshold = 3
# }
#   health_check {
#   path     = "/ping"
#   interval = 30
#   timeout  = 6
#   matcher  = "200"
# }
#}

# resource "aws_lb_listener" "nginx" {
#   load_balancer_arn = aws_lb.both.arn
#   port              = "80"
#   protocol          = "TCP"

#   default_action {
#     type             = "forward"
#     target_group_arn = aws_lb_target_group.nginx.arn
#   }

#   tags = local.tags
# }

# resource "aws_route53_record" "nginx_lb" {
#   zone_id = "Z048359310Y6AD9XSL491"
#   name    = "mo-nginx-lb.sandbox.<hosted_zone>.services"
#   type    = "A"

#   alias {
#     name                   = aws_lb.both.dns_name
#     zone_id                = aws_lb.both.zone_id
#     evaluate_target_health = true
#   }
# }

## LB listener rule

# resource "aws_lb_listener_rule" "host_based_routing" {
#   listener_arn = aws_lb_listener.nginx.arn
#   priority     = 100

#   action {
#     type = "forward"
#     forward {
#       target_group {
#         arn    = aws_lb_target_group.nginx.arn
#         weight = 80
#       }

#       target_group {
#         arn    = aws_lb_target_group.apache.arn
#         weight = 20
#       }

#       stickiness {
#         enabled  = true
#         duration = 600
#       }
#     }
#   }

#   condition {
#     host_header {
#       values = ["*mo-sandbox.sandbox.<hosted_zone>.services"]
#     }
#   }
# }


# resource "aws_lb_listener_rule" "nginx_routing" {
#   listener_arn = aws_lb_listener.nginx.arn
#   priority     = 100

#   action {
#     type             = "forward"
#     target_group_arn = aws_lb_target_group.nginx.arn

#     forward {
#       target_group {
#         arn    = aws_lb_target_group.nginx.arn
#         weight = 50
#       }
#       target_group {
#         arn    = aws_lb_target_group.apache.arn
#         weight = 50
#       }
#     }
#   }

#   condition {
#     host_header {
#       values = ["mo-nginx-lb.sandbox.<hosted_zone>.services"]
#     }
#   }
# }

# resource "aws_lb_listener_rule" "apache_routing" {
#   listener_arn = aws_lb_listener.apache.arn
#   priority     = 200

#   action {
#     type             = "forward"
#     target_group_arn = aws_lb_target_group.apache.arn

#     forward {
#       target_group {
#         arn    = aws_lb_target_group.nginx.arn
#         weight = 50
#       }
#       target_group {
#         arn    = aws_lb_target_group.apache.arn
#         weight = 50
#       }
#     }
#   }

#   condition {
#     host_header {
#       values = ["mo-apache-lb.sandbox.<hosted_zone>.services"]
#     }
#   }
# }
