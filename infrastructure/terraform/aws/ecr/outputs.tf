output "repository_url" {
  value = aws_ecr_repository.shopops_catalog.repository_url
}

output "repository_arn" {
  description = "ARN of the ShopOps ECR repository"
  value       = aws_ecr_repository.shopops_catalog.arn
}