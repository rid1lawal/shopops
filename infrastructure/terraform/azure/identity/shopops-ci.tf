data "azurerm_client_config" "current" {}

data "azurerm_container_registry" "shopops" {
  name                = "shopopsacr"
  resource_group_name = "shopops-rg"
}

resource "azuread_application" "shopops_ci" {
  display_name = "shopops-ci"
}

resource "azuread_service_principal" "shopops_ci" {
  client_id = azuread_application.shopops_ci.client_id
}

resource "azuread_application_federated_identity_credential" "github_main" {
  application_id = azuread_application.shopops_ci.id

  display_name = "github-main"
  description  = "Allow GitHub Actions to authenticate from the main branch"

  audiences = [
    "api://AzureADTokenExchange"
  ]

  issuer  = "https://token.actions.githubusercontent.com"
  subject = "repo:${var.github_owner}@${var.github_owner_id}/${var.github_repository}@${var.github_repository_id}:ref:refs/heads/main"
}

resource "azurerm_role_assignment" "acr_push" {
  scope                = data.azurerm_container_registry.shopops.id
  role_definition_name = "AcrPush"
  principal_id         = azuread_service_principal.shopops_ci.object_id
}