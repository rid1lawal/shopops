module "resource_group" {
  source = "../../modules/resource-group"

  name     = var.resource_group_name
  location = var.location
}

resource "azurerm_container_registry" "shopops" {
  name                = var.acr_name
  resource_group_name = module.resource_group.name
  location            = module.resource_group.location

  sku           = "Basic"
  admin_enabled = false

  tags = {
    project    = "shopops"
    managed_by = "terraform"
  }
}