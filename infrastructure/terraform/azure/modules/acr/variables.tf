variable "location" {
  description = "Azure region for the ACR"
  type        = string
}

variable "resource_group_name" {
  description = "Resource group for the ShopOps ACR"
  type        = string
}

variable "acr_name" {
  description = "Globally unique Azure Container Registry name"
  type        = string
}