output "repository_url" {
  value = aws_ecr_repository.shopops_ecr.repository_url
}

output "repository_arn" {
  description = "ARN of the ShopOps ECR repository"
  value       = aws_ecr_repository.shopops_ecr.arn
}