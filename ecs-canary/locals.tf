locals {
  # Target port to expose
  target_port = 3000

  ## ECS Service config
  ecs_launch_type    = "FARGATE"
  ecs_desired_count  = 2
  ecs_network_mode   = "awsvpc"
  ecs_cpu            = 512
  ecs_memory         = 1024
  ecs_container_name = "nextjs-image"
  ecs_log_group      = "/aws/ecs/${var.project_id}-${var.env}"
  # Retention in days
  ecs_log_retention = 1

  # Deployment Configuration
  ecs_deployment_type = "TimeBasedCanary"
  ## In minutes
  ecs_deployment_config_interval = 5
  ## In percentage
  ecs_deployment_config_pct = 25
}
