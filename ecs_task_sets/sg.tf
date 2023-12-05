# for ECS Cluster allow port 80
resource "aws_security_group" "ecs_sg" {
  vpc_id = module.network.vpc_id

  ingress {
    from_port = 80
    to_port   = 80
    protocol  = "tcp"
    # it should be modified in production
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port = 0
    to_port   = 65535
    protocol  = "tcp"
    # it should be modified in production
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# for RDS Cluster default port for MYSQL:3306
resource "aws_security_group" "rds_sg" {
  vpc_id = module.network.vpc_id

  ingress {
    protocol  = "tcp"
    from_port = 3306
    to_port   = 3306
    # it should be modified in production
    cidr_blocks     = ["0.0.0.0/0"]
    security_groups = [aws_security_group.ecs_sg.id]
  }

  egress {
    from_port = 0
    to_port   = 65535
    protocol  = "tcp"
    # it should be modified in production
    cidr_blocks = ["0.0.0.0/0"]
  }
}
