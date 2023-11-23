resource "aws_cloudwatch_log_group" "ecs" {
  name              = local.ecs_log_group
  retention_in_days = local.ecs_log_retention
}

## Cloudwatch log errors
module "application_error_alarm" {
  source             = "github.com/Jareechang/tf-modules//cloudwatch/alarms/application-log-errors?ref=v1.0.12"
  evaluation_periods = "2"
  threshold          = "10"
  arn_suffix         = module.alb.lb.arn_suffix
  project_id         = var.project_id
  env                = var.env
  # Keyword to match for this can be changed
  pattern          = "Error"
  log_group_name   = aws_cloudwatch_log_group.ecs.name
  metric_name      = "ApplicationErrorCount"
  metric_namespace = "ECS/${var.project_id}-${var.env}"
}

## ALB errors (5xx)
module "http_error_alarm" {
  source             = "github.com/Jareechang/tf-modules//cloudwatch/alarms/alb-http-errors?ref=v1.0.8"
  evaluation_periods = "2"
  threshold          = "10"
  arn_suffix         = module.alb.lb.arn_suffix
  project_id         = var.project_id
}
