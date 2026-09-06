variable "location" {
  description = "Azure region where the Terraform state resources will be created"
  type        = string
  default     = "East US"
}

variable "resource_group_name" {
  description = "Resource group for Terraform state resources"
  type        = string
  default     = "shopops-tfstate-rg"
}

variable "storage_account_name" {
  description = "Globally unique Azure Storage Account name"
  type        = string
}