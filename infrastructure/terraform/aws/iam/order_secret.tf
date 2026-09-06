resource "aws_iam_policy" "order_secrets" {
  name        = "shopops-order_secrets"
  description = "Allow the Order service to read its Secrets Manager secret"

  policy = jsonencode({
    Version = "2012-10-17"

    Statement = [
      {
        Effect = "Allow"

        Action = [
          "secretsmanager:GetSecretValue",
          "secretsmanager:DescribeSecret"
        ]

        Resource = data.terraform_remote_state.secrets_manager.outputs.order_secret_arn
      }
    ]
  })
}

resource "aws_iam_role" "order_secrets" {
  name = "shopops-order_secrets"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"

    Statement = [
      {
        Effect = "Allow"

        Principal = {
          Service = "pods.eks.amazonaws.com"
        }

        Action = [
          "sts:AssumeRole",
          "sts:TagSession"
        ]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "order_secrets" {
  role       = aws_iam_role.order_secrets.name
  policy_arn = aws_iam_policy.order_secrets.arn
}

resource "aws_eks_pod_identity_association" "order" {
  cluster_name    = "shopops"
  namespace       = "external-secrets"
  service_account = "order-secrets"
  role_arn        = aws_iam_role.order_secrets.arn
}