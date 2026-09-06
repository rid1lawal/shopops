output "catalog_secret_arn" {
  value = aws_secretsmanager_secret.catalog.arn
}

output "order_secret_arn" {
  value = aws_secretsmanager_secret.order.arn
}