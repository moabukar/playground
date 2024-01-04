

resource "aws_ecs_cluster" "web_cluster" {
  name = "web-cluster-${var.project_id}-${var.env}"
  setting {
    name  = "containerInsights"
    value = "enabled"
  }
}

resource "local_file" "output_task_def" {
  content         = data.template_file.task_def_generated.rendered
  file_permission = "644"
  filename        = "./task-definitions/service.latest.json"
}

resource "aws_ecs_task_definition" "nextjs" {
  family             = "task-definition-node"
  execution_role_arn = module.ecs_roles.ecs_execution_role_arn
  task_role_arn      = module.ecs_roles.ecs_task_role_arn

  requires_compatibilities = [local.ecs_launch_type]
  network_mode             = local.ecs_network_mode
  cpu                      = local.ecs_cpu
  memory                   = local.ecs_memory
  container_definitions = jsonencode(
    jsondecode(data.template_file.task_def_generated.rendered).containerDefinitions
  )
}

resource "aws_ecs_service" "web_ecs_service" {
  name            = "web-service-${var.project_id}-${var.env}"
  cluster         = aws_ecs_cluster.web_cluster.id
  task_definition = aws_ecs_task_definition.nextjs.arn
  desired_count   = local.ecs_desired_count
  launch_type     = local.ecs_launch_type

  load_balancer {
    target_group_arn = module.ecs_tg_blue.tg.arn
    container_name   = local.ecs_container_name
    container_port   = local.target_port
  }

  network_configuration {
    subnets         = module.networking.private_subnets[*].id
    security_groups = [aws_security_group.ecs_sg.id]
  }

  deployment_controller {
    type = "CODE_DEPLOY"
  }

  tags = {
    Name = "web-service-${var.project_id}-${var.env}"
  }

  depends_on = [
    module.alb.lb,
    module.ecs_tg_blue.tg
  ]
}
