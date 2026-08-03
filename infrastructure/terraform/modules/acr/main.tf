resource "azurerm_resource_group" "shopops" {
  name     = var.resource_group_name
  location = var.location

  tags = {
    project    = "shopops"
    managed_by = "terraform"
  }
}

resource "azurerm_container_registry" "shopops" {
  name                = var.acr_name
  resource_group_name = azurerm_resource_group.shopops.name
  location            = azurerm_resource_group.shopops.location

  sku           = "Basic"
  admin_enabled = false

  tags = {
    project    = "shopops"
    managed_by = "terraform"
  }
}