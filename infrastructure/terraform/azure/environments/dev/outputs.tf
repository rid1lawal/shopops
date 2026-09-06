output "resource_group_name" {
  description = "Resource group for the dev environment"
  value       = module.resource_group.name
}

output "vnet_name" {
  description = "Virtual network name"
  value       = module.network.vnet_name
}

output "aks_cluster_name" {
  description = "AKS cluster name"
  value       = module.aks.name
}

output "aks_cluster_id" {
  description = "AKS cluster ID"
  value       = module.aks.id
}

output "aks_node_resource_group" {
  description = "AKS managed node resource group"
  value       = module.aks.node_resource_group
}

output "aks_oidc_issuer_url" {
  description = "OIDC issuer URL for Workload Identity"
  value       = module.aks.oidc_issuer_url
}

output "vnet_id" {
  description = "Virtual network ID"
  value       = module.network.vnet_id
}

output "database_subnet_id" {
  description = "Subnet ID for the PostgreSQL database"
  value       = module.network.database_subnet_id
}