#tfsec:ignore:aws-ssm-secret-use-customer-key
resource "aws_secretsmanager_secret" "ecr_pullthroughcache_docker_hub" {
  name = "ecr-pullthroughcache/docker-hub"

  recovery_window_in_days = 7
}

#tfsec:ignore:aws-ssm-secret-use-customer-key
resource "aws_secretsmanager_secret" "ecr_pullthroughcache_github" {
  name = "ecr-pullthroughcache/github"

  recovery_window_in_days = 7
}
