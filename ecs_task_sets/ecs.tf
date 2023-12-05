
resource "aws_ecs_cluster" "ecs_cluster" {
  name = "my-cluster"
}


resource "aws_ecs_task_definition" "task_definition" {
  # family if required
  family                = "worker"
  container_definitions = file("./container_specs.json")
}


resource "aws_ecs_service" "worker" {
  name            = "worker"
  cluster         = aws_ecs_cluster.ecs_cluster.id
  task_definition = aws_ecs_task_definition.task_definition.arn
  desired_count   = 1
}

resource "aws_ecs_task_set" "example" {
  service         = aws_ecs_service.worker.id
  cluster         = aws_ecs_cluster.ecs_cluster.id
  task_definition = aws_ecs_task_definition.task_definition.arn

  load_balancer {
    target_group_arn = aws_lb_target_group.alb-example.arn
    container_name   = "nodejs"
    container_port   = 8080
  }
}
