resource "aws_codedeploy_deployment_config" "custom_canary" {
  deployment_config_name = "EcsCanary25Percent20Minutes"
  compute_platform       = "ECS"
  traffic_routing_config {
    type = local.ecs_deployment_type
    time_based_canary {
      interval   = local.ecs_deployment_config_interval
      percentage = local.ecs_deployment_config_pct
    }
  }
}

resource "aws_codedeploy_app" "node_app" {
  compute_platform = "ECS"
  name             = "deployment-app-${var.project_id}-${var.env}"
}

resource "aws_codedeploy_deployment_group" "node_deploy_group" {
  app_name               = aws_codedeploy_app.node_app.name
  deployment_config_name = aws_codedeploy_deployment_config.custom_canary.id
  deployment_group_name  = "deployment-group-${var.project_id}-${var.env}"
  service_role_arn       = aws_iam_role.codedeploy_role.arn

  auto_rollback_configuration {
    enabled = true
    events  = ["DEPLOYMENT_FAILURE", "DEPLOYMENT_STOP_ON_ALARM"]
  }

  alarm_configuration {
    alarms = [
      module.http_error_alarm.name,
      module.application_error_alarm.name
    ]
    enabled = true
  }

  blue_green_deployment_config {
    deployment_ready_option {
      action_on_timeout = "CONTINUE_DEPLOYMENT"
    }

    terminate_blue_instances_on_deployment_success {
      action                           = "TERMINATE"
      termination_wait_time_in_minutes = 0
    }
  }

  deployment_style {
    deployment_option = "WITH_TRAFFIC_CONTROL"
    deployment_type   = "BLUE_GREEN"
  }

  ecs_service {
    cluster_name = aws_ecs_cluster.web_cluster.name
    service_name = aws_ecs_service.web_ecs_service.name
  }

  load_balancer_info {
    target_group_pair_info {
      prod_traffic_route {
        listener_arns = [module.alb.http_listener.arn]
      }

      target_group {
        name = module.ecs_tg_blue.tg.name
      }

      target_group {
        name = module.ecs_tg_green.tg.name
      }
    }
  }
}
