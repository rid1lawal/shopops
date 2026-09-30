resource "aws_secretsmanager_secret" "catalog" {
  name                    = "shopops/catalog"
  description             = "Secrets for the ShopOps Catalog service"
  recovery_window_in_days = 0

  tags = {
    Project = "shopops"
  }
}

resource "aws_secretsmanager_secret_version" "catalog" {
  secret_id = aws_secretsmanager_secret.catalog.id

  secret_string = jsonencode({
    DB_URL = var.catalog_db_url
  })
}

resource "aws_secretsmanager_secret" "order" {
  name                    = "shopops/order"
  description             = "Secrets for the ShopOps Order service"
  recovery_window_in_days = 0

  tags = {
    Project = "shopops"
  }
}

resource "aws_secretsmanager_secret_version" "order" {
  secret_id = aws_secretsmanager_secret.order.id

  secret_string = jsonencode({
    DB_URL = var.order_db_url
  })
}