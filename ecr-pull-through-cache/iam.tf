resource "aws_iam_role" "eks_node_group" {
  name = "node-group"

  assume_role_policy = data.aws_iam_policy_document.eks_node_group_assume_role_policy.json

  managed_policy_arns = [
    "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
    "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
    "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
    "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore",
    "arn:aws:iam::aws:policy/CloudWatchAgentServerPolicy"
  ]

  inline_policy {
    name   = "ecr-cache-policy"
    policy = data.aws_iam_policy_document.eks_node_custom_inline_policy.json
  }
}

data "aws_iam_policy_document" "eks_node_custom_inline_policy" {
  statement {
    actions = [
      "ecr:CreateRepository",
      "ecr:ReplicateImage",
      "ecr:BatchImportUpstreamImage"
    ]

    resources = ["*"]
  }
}
