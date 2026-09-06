output "github_actions_ecr_role_arn" {
  description = "IAM role assumed by GitHub Actions to push images to ECR"
  value       = aws_iam_role.github_actions_ecr.arn
}
