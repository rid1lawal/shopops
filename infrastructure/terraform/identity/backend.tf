terraform {
  backend "azurerm" {
    resource_group_name  = "shopops-tfstate-rg"
    storage_account_name = "shopopstfstate"
    container_name       = "tfstate"
    key                  = "identity.tfstate"
  }
}