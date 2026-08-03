variable "resource_group_name" {
  type = string
}

variable "location" {
  type = string
}

variable "environment" {
  type = string
}

variable "aks_subnet_name" {
  type = string
}

variable "database_subnet_name" {
  type = string
}

variable "address_space" {
  type = list(string)
}

variable "aks_subnet_prefix" {
  type = list(string)
}

variable "database_subnet_prefix" {
  type = list(string)
}