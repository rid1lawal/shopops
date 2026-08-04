locals {
  project     = "shopops"
  environment = "dev"
  location    = "East US 2"

  tags = {
    project     = local.project
    environment = local.environment
    managed_by  = "terraform"
  }
}

module "resource_group" {
  source = "../../modules/resource-group"

  name     = "${local.project}-${local.environment}-rg"
  location = local.location
}

module "network" {
  source = "../../modules/network"

  resource_group_name = module.resource_group.name
  location            = module.resource_group.location

  environment            = local.environment
  address_space          = ["10.10.0.0/16"]
  aks_subnet_prefix      = ["10.10.1.0/24"]
  database_subnet_prefix = ["10.10.2.0/24"]

  aks_subnet_name      = "aks-subnet"
  database_subnet_name = "db-subnet"
}

module "aks" {
  source = "../../modules/aks"

  name                = "${local.project}-${local.environment}-aks"
  location            = module.resource_group.location
  resource_group_name = module.resource_group.name

  subnet_id = module.network.aks_subnet_id

  vm_size    = "Standard_D2ads_v6"
  node_count = 1

  tags = local.tags
}