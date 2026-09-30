variable "catalog_db_url" {
  description = "Database connection URL for the Catalog service"
  type        = string
  sensitive   = true
}

variable "order_db_url" {
  description = "Database connection URL for the Order service"
  type        = string
  sensitive   = true
}