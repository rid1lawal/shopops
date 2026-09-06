resource "azurerm_user_assigned_identity" "catalog_secrets" {
  name                = "shopops-${local.environment}-catalog-secrets"
  resource_group_name = module.resource_group.name
  location            = module.resource_group.location

  tags = local.tags
}

resource "azurerm_federated_identity_credential" "catalog" {
  name = "catalog"

  user_assigned_identity_id = azurerm_user_assigned_identity.catalog_secrets.id

  audience = [
    "api://AzureADTokenExchange"
  ]

  issuer = module.aks.oidc_issuer_url

  subject = "system:serviceaccount:shopops-dev:catalog"
}