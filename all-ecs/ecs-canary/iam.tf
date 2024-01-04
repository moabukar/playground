module "ecr_ecs_ci_user" {
  source         = "github.com/moabukar/tf-mods//iam/ecr?ref=v1.0.15"
  env            = var.env
  project_id     = var.project_id
  create_ci_user = true
  # This is the ECR ARN - Feel free to add other repository as required (if you want to re-use role for CI/CD in other projects)
  ecr_resource_arns = [
    "arn:aws:ecr:${var.aws_region}:${data.aws_caller_identity.current.account_id}:repository/web/${var.project_id}",
    "arn:aws:ecr:${var.aws_region}:${data.aws_caller_identity.current.account_id}:repository/web/${var.project_id}/*"
  ]

  other_iam_statements = {
    codedeploy = {
      actions = [
        "codedeploy:GetDeploymentGroup",
        "codedeploy:CreateDeployment",
        "codedeploy:GetDeployment",
        "codedeploy:GetDeploymentConfig",
        "codedeploy:RegisterApplicationRevision"
      ]
      effect = "Allow"
      resources = [
        "*"
      ]
    }
  }
}


## Execution role and task roles
module "ecs_roles" {
  source                    = "github.com/moabukar/tf-mods//iam/ecs?ref=v1.0.1"
  create_ecs_execution_role = true
  create_ecs_task_role      = true

  # Extend baseline policy statements (ignore for now)
  ecs_execution_policies_extension = {}

}

data "aws_iam_policy_document" "codedeploy_assume_role" {
  version = "2012-10-17"
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type = "Service"
      identifiers = [
        "codedeploy.amazonaws.com"
      ]
    }
  }
}

resource "aws_iam_role" "codedeploy_role" {
  name               = "CodeDeployRole${var.project_id}"
  description        = "CodeDeployRole for ${var.project_id} in ${var.env}"
  assume_role_policy = data.aws_iam_policy_document.codedeploy_assume_role.json
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_iam_role_policy_attachment" "codedeploy_policy_attachment" {
  policy_arn = "arn:aws:iam::aws:policy/AWSCodeDeployRoleForECS"
  role       = aws_iam_role.codedeploy_role.name
}
