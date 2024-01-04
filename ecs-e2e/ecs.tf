resource "aws_ecs_cluster" "cluster" {
  name               = local.ecs["cluster_name"] # "ecs-cluster"
  capacity_providers = ["FARGATE"]

  default_capacity_provider_strategy {
    capacity_provider = "FARGATE"
    weight            = "100"
  }
}

## Task def

resource "aws_ecs_task_definition" "task" {
  family = "service"
  requires_compatibilities = [
    "FARGATE",
  ]
  execution_role_arn = aws_iam_role.fargate.arn
  network_mode       = "awsvpc"
  cpu                = 256
  memory             = 512
  container_definitions = jsonencode([
    {
      name      = local.container.name   # "application"
      image     = local.container.image  # "particule/helloworld"
      essential = true
      portMappings = [
        {
          containerPort = 80
          hostPort      = 80
        }
      ]
    }
  ])
}

## Service
resource "aws_ecs_service" "service" {
  name            = local.ecs.service_name
  cluster         = aws_ecs_cluster.cluster.id
  task_definition = aws_ecs_task_definition.task.arn
  desired_count   = 1

  network_configuration {
    subnets          = [for s in data.aws_subnet.subnets : s.id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.group.arn  # our target group
    container_name   = local.container.name           # "application"
    container_port   = 80
  }
  capacity_provider_strategy {
    base              = 0
    capacity_provider = "FARGATE"
    weight            = 100
  }
}
