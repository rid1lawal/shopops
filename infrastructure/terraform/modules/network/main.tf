resource "azurerm_virtual_network" "shopops-vnet" {
  name                = "shopops-${var.environment}-vnet"
  location            = var.location
  resource_group_name = var.resource_group_name

  address_space = var.address_space

  tags = {
    environment = var.environment
    managed_by  = "terraform"
    project     = "shopops"
  }
}

resource "azurerm_subnet" "aks" {
  name                 = var.aks_subnet_name
  resource_group_name  = var.resource_group_name
  virtual_network_name = azurerm_virtual_network.shopops-vnet.name

  address_prefixes = var.aks_subnet_prefix
}

resource "azurerm_subnet" "database" {
  name                 = var.database_subnet_name
  resource_group_name  = var.resource_group_name
  virtual_network_name = azurerm_virtual_network.shopops-vnet.name

  address_prefixes = var.database_subnet_prefix

  delegation {
    name = "postgresql-flexible-server"

    service_delegation {
      name = "Microsoft.DBforPostgreSQL/flexibleServers"

      actions = [
        "Microsoft.Network/virtualNetworks/subnets/join/action",
      ]
    }
  }
}
