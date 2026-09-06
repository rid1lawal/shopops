resource "aws_iam_policy" "shopops_secrets" {
  name        = "shopops-secrets"
  description = "Allow the shopops services to read its Secrets Manager secret"

  policy = jsonencode({
    Version = "2012-10-17"

    Statement = [
      {
        Effect = "Allow"

        Action = [
          "secretsmanager:GetSecretValue",
          "secretsmanager:DescribeSecret"
        ]

        Resource = [data.terraform_remote_state.secrets_manager.outputs.catalog_secret_arn,
          data.terraform_remote_state.secrets_manager.outputs.catalog_secret_arn,
        ]
      }
    ]
  })
}

resource "aws_iam_role" "shopops_secrets" {
  name = "shopops-secrets"

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

resource "aws_iam_role_policy_attachment" "shopops_secrets" {
  role       = aws_iam_role.shopops_secrets.name
  policy_arn = aws_iam_policy.shopops_secrets.arn
}

resource "aws_eks_pod_identity_association" "shopops" {
  cluster_name    = "shopops"
  namespace       = "external-secrets"
  service_account = "external-secrets"
  role_arn        = aws_iam_role.shopops_secrets.arn
}
