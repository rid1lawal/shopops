resource "azurerm_key_vault" "shopops" {
  name                = "shopops-${local.environment}-kv"
  location            = module.resource_group.location
  resource_group_name = module.resource_group.name

  tenant_id = data.azurerm_client_config.current.tenant_id

  sku_name = "standard"

  rbac_authorization_enabled = true

  soft_delete_retention_days = 7
  purge_protection_enabled   = false

  tags = local.tags
}

resource "azurerm_role_assignment" "catalog_key_vault_secrets_user" {
  scope                = azurerm_key_vault.shopops.id
  role_definition_name = "Key Vault Secrets User"
  principal_id         = azurerm_user_assigned_identity.catalog_secrets.principal_id
}