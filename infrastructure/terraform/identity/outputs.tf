output "client_id" {
  description = "Client ID used by GitHub Actions"
  value       = azuread_application.shopops_ci.client_id
}

output "tenant_id" {
  description = "Azure tenant ID"
  value       = data.azurerm_client_config.current.tenant_id
}

output "subscription_id" {
  description = "Azure subscription ID"
  value       = data.azurerm_client_config.current.subscription_id
}