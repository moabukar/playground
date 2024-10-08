terraform {
  required_version = ">= 1.0"
  required_providers {

    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 3.0"
    }

    azuread = {
      source  = "hashicorp/azuread"
      version = ">= 2.0"
    }

    random = {
      source  = "hashicorp/random"
      version = ">= 2.0"
    }

    tls = {
      source  = "hashicorp/tls"
      version = ">= 2.0"
    }
  }
}
