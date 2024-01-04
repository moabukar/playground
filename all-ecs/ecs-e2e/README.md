# Notes

## TF apply output

```bash

$ terraform apply -auto-approve
...
Apply complete! Resources: 4 added, 0 changed, 0 destroyed.

Outputs:

app_url = "tf-alb-1725491528.eu-west-3.elb.amazonaws.com"
aws_region = "eu-west-3"
container_name = "application"
ecr_repository_name = "app-registry"
ecr_url = "111111111111.dkr.ecr.eu-west-3.amazonaws.com/app-registry"
ecs_cluster = "ecs-cluster"
ecs_service = "ecs-service"
publisher_access_key = "AAAAAAAAAAAAAAAAAAAA"
publisher_secret_key = <sensitive>

```

## Secrets

```bash

$ gh secret set AWS_ACCESS_KEY_ID -b $(terraform output -raw publisher_access_key)
✓ Set secret AWS_ACCESS_KEY_ID for tbobm/tf-ecr-ecs-gh-deploy
$ gh secret set AWS_SECRET_ACCESS_KEY -b $(terraform output -raw publisher_secret_key)
✓ Set secret AWS_SECRET_ACCESS_KEY for tbobm/tf-ecr-ecs-gh-deploy
$ gh secret set AWS_REGION -b $(terraform output -raw aws_region)
✓ Set secret AWS_REGION for tbobm/tf-ecr-ecs-gh-deploy
$ gh secret set ECR_REPOSITORY_NAME -b $(terraform output -raw ecr_repository_name)
✓ Set secret ECR_REPOSITORY_NAME for tbobm/tf-ecr-ecs-gh-deploy
$ gh secret set ECS_CLUSTER -b $(terraform output -raw ecs_cluster)
✓ Set secret ECS_CLUSTER for tbobm/tf-ecr-ecs-gh-deploy
$ gh secret set ECS_SERVICE -b $(terraform output -raw ecs_service)
✓ Set secret ECS_SERVICE for tbobm/tf-ecr-ecs-gh-deploy
$ gh secret list
AWS_ACCESS_KEY_ID      Updated 2024-01-04
AWS_REGION             Updated 2024-01-04
AWS_SECRET_ACCESS_KEY  Updated 2024-01-04
ECS_CLUSTER            Updated 2024-01-04
ECR_REPOSITORY_NAME    Updated 2024-01-04
ECS_SERVICE            Updated 2024-01-04


```

## Test workflow

```bash

curl tf-alb-1725491528.eu-west-3.elb.amazonaws.com


curl tf-alb-1725491528.eu-west-3.elb.amazonaws.com
{
  "message": "Hello from ip-172-31-13-81.eu-west-3.compute.internal"
}

```

