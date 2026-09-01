output "vnet_id" {
  value = azurerm_virtual_network.shopops-vnet.id
}

output "vnet_name" {
  value = azurerm_virtual_network.shopops-vnet.name
}

output "aks_subnet_id" {
  value = azurerm_subnet.aks.id
}

output "database_subnet_id" {
  value = azurerm_subnet.database.id
}