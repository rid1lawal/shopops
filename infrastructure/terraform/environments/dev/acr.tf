resource "azurerm_role_assignment" "aks_acr_pull" {
  scope                = data.terraform_remote_state.container_registry.outputs.acr_id
  role_definition_name = "AcrPull"

  principal_id = module.aks.kubelet_object_id
}