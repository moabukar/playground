provider "aws" {
  region = "eu-west-2"
}

resource "aws_ecr_repository" "ecs_devops_sandbox_repository" {
  name = "ecs-devops-sandbox-repository"
}

resource "aws_vpc" "ecs_devops_sandbox_vpc" {
  cidr_block = "10.0.0.0/16" # Define your CIDR block
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name = "ecs-devops-sandbox-vpc"
  }
}

resource "aws_ecs_cluster" "ecs_devops_sandbox_cluster" {
  name = "ecs-devops-sandbox-cluster"
}

resource "aws_iam_role" "ecs_execution_role" {
  name = "ecs-devops-sandbox-execution-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        },
      },
    ]
  })
}

resource "aws_iam_role_policy" "ecs_execution_role_policy" {
  name = "ecs_execution_role_policy"
  role = aws_iam_role.ecs_execution_role.id

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "ecr:GetAuthorizationToken",
          "ecr:BatchCheckLayerAvailability",
          "ecr:GetDownloadUrlForLayer",
          "ecr:BatchGetImage",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        Resource = "*"
      },
    ]
  })
}

resource "aws_ecs_task_definition" "ecs_devops_sandbox_task_definition" {
  family                   = "ecs-devops-sandbox-task-definition"
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.ecs_execution_role.arn
  cpu                      = "256" # Adjust as needed
  memory                   = "512" # Adjust as needed
  requires_compatibilities = ["FARGATE"]

  container_definitions = jsonencode([{
    name  = "ecs-devops-sandbox",
    image = "amazon/amazon-ecs-sample",
    essential = true,
    // Add other container properties as needed
  }])
}

resource "aws_ecs_service" "ecs_devops_sandbox_service" {
  name            = "ecs-devops-sandbox-service"
  cluster         = aws_ecs_cluster.ecs_devops_sandbox_cluster.id
  task_definition = aws_ecs_task_definition.ecs_devops_sandbox_task_definition.arn
  launch_type     = "FARGATE"

  network_configuration {
    subnets = [aws_vpc.ecs_devops_sandbox_vpc.subnet_ids] # Replace with your subnet IDs
    // Include security groups if necessary
  }

  desired_count = 1
  // Add other service properties as needed
}
