resource "azurerm_resource_group" "shopops" {
  name     = var.name
  location = var.location

  tags = {
    project    = "shopops"
    managed_by = "terraform"
  }
}