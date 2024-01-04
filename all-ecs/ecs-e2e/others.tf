resource "aws_ecr_repository" "repository" {
  name                 = local.ecr["repository_name"]  # "repository"
  image_tag_mutability = "MUTABLE"
}


resource "aws_iam_user" "publisher" {
  name = "ecr-publisher"
  path = "/serviceaccounts/"
}

resource "aws_iam_access_key" "publisher" {
  user = aws_iam_user.publisher.name
}

## Outputs

output "publisher_access_key" {
  value       = aws_iam_access_key.publisher.id
  description = "AWS_ACCESS_KEY to publish to ECR"
}

output "publisher_secret_key" {
  value       = aws_iam_access_key.publisher.secret
  description = "AWS_SECRET_ACCESS_KEY to upload to the ECR"
  sensitive   = true
}

output "ecr_url" {
  value       = aws_ecr_repository.repository.repository_url
  description = "The ECR repository URL"
}
