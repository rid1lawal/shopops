resource "aws_iam_policy" "catalog_secrets" {
  name        = "shopops-catalog-secrets"
  description = "Allow the Catalog service to read its Secrets Manager secret"

  policy = jsonencode({
    Version = "2012-10-17"

    Statement = [
      {
        Effect = "Allow"

        Action = [
          "secretsmanager:GetSecretValue",
          "secretsmanager:DescribeSecret"
        ]

        Resource = data.terraform_remote_state.secrets_manager.outputs.catalog_secret_arn
      }
    ]
  })
}

resource "aws_iam_role" "catalog_secrets" {
  name = "shopops-catalog-secrets"

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

resource "aws_iam_role_policy_attachment" "catalog_secrets" {
  role       = aws_iam_role.catalog_secrets.name
  policy_arn = aws_iam_policy.catalog_secrets.arn
}

resource "aws_eks_pod_identity_association" "catalog" {
  cluster_name    = "shopops"
  namespace       = "external-secrets"
  service_account = "external-secrets"
  role_arn        = aws_iam_role.catalog_secrets.arn
}