data "terraform_remote_state" "container_registry" {
  backend = "azurerm"

  config = {
    resource_group_name  = "shopops-tfstate-rg"
    storage_account_name = "shopopstfstate"
    container_name       = "tfstate"
    key                  = "shopops-acr.tfstate"
  }
}