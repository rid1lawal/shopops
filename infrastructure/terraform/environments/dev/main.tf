module "resource_group" {
  source = "../../modules/resource-group"

  name     = "shopops-dev-rg"
  location = "East US"
}

module "network" {
  source = "../../modules/network"

  resource_group_name = module.resource_group.name
  location            = module.resource_group.location

  environment            = "dev"
  address_space          = ["10.0.0.0/16"]
  aks_subnet_prefix      = ["10.0.1.0/24"]
  database_subnet_prefix = ["10.0.2.0/24"]

  aks_subnet_name      = "aks-subnet"
  database_subnet_name = "db-subnet"
}