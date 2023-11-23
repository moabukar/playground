module "alb" {
  source             = "github.com/moabukar/tf-mods//alb?ref=v1.0.2"
  create_alb         = true
  enable_https       = false
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.alb_ecs_sg.id]
  subnets            = module.networking.public_subnets[*].id
  target_group       = module.ecs_tg_blue.tg.arn
}

module "ecs_tg_blue" {
  source              = "github.com/moabukar/tf-mods//alb?ref=v1.0.2"
  create_target_group = true
  port                = local.target_port
  protocol            = "HTTP"
  target_type         = "ip"
  vpc_id              = module.networking.vpc_id
}

# Target group for new infrastructure
module "ecs_tg_green" {
  project_id          = "${var.project_id}-green"
  source              = "github.com/moabukar/tf-mods//alb?ref=v1.0.2"
  create_target_group = true
  port                = local.target_port
  protocol            = "HTTP"
  target_type         = "ip"
  vpc_id              = module.networking.vpc_id
}
