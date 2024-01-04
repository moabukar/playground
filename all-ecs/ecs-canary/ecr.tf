resource "aws_ecr_repository" "main" {
  name                 = "web/${var.project_id}/nextjs"
  image_tag_mutability = "IMMUTABLE"
}
