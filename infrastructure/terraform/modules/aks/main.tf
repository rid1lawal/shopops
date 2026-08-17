resource "azurerm_kubernetes_cluster" "aks" {
  name                = var.name
  location            = var.location
  resource_group_name = var.resource_group_name

  dns_prefix = var.name

  sku_tier = "Free"

  oidc_issuer_enabled       = true
  workload_identity_enabled = true

  key_vault_secrets_provider {
  secret_rotation_enabled = true
}

  default_node_pool {
    name = "sys"

    vm_size        = var.vm_size
    node_count     = var.node_count
    vnet_subnet_id = var.subnet_id

    upgrade_settings {
      max_surge = "10%"
    }
  }

  identity {
    type = "SystemAssigned"
  }

  network_profile {
    network_plugin      = "azure"
    network_plugin_mode = "overlay"

    network_policy     = "cilium"
    network_data_plane = "cilium"

    load_balancer_sku = "standard"

    service_cidr   = "10.240.0.0/16"
    dns_service_ip = "10.240.0.10"
  }

  tags = var.tags
}
