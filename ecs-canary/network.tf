module "networking" {
  source     = "github.com/moabukar/tf-mods//networking?ref=v1.0.1"
  env        = var.env
  project_id = var.project_id
  subnet_public_cidrblock = [
    "10.0.1.0/24",
    "10.0.2.0/24"
  ]
  subnet_private_cidrblock = [
    "10.0.11.0/24",
    "10.0.22.0/24"
  ]
  azs = ["us-east-1a", "us-east-1b"]
}

#### Security groups
resource "aws_security_group" "alb_ecs_sg" {
  vpc_id = module.networking.vpc_id

  ## Allow inbound on port 80 from internet (all traffic)
  ingress {
    protocol    = "tcp"
    from_port   = 80
    to_port     = 80
    cidr_blocks = ["0.0.0.0/0"]
  }

  ## Allow outbound to ecs instances in private subnet
  egress {
    protocol    = "tcp"
    from_port   = local.target_port
    to_port     = local.target_port
    cidr_blocks = module.networking.private_subnets[*].cidr_block
  }
}

resource "aws_security_group" "ecs_sg" {
  vpc_id = module.networking.vpc_id
  ingress {
    protocol        = "tcp"
    from_port       = local.target_port
    to_port         = local.target_port
    security_groups = [aws_security_group.alb_ecs_sg.id]
  }

  ## Allow ECS service to reach out to internet (download packages, pull images etc)
  egress {
    protocol    = -1
    from_port   = 0
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }
}
