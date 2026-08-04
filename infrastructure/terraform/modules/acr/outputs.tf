output "acr_id" {
  value = azurerm_container_registry.shopops.id
}

output "resource_group_name" {
  value = module.resource_group.name
}

output "acr_name" {
  value = azurerm_container_registry.shopops.name
}

output "acr_login_server" {
  value = azurerm_container_registry.shopops.login_server
}