terraform {
  required_version = ">= 1.0"

  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.9.0"
    }
    azuread = {
      source  = "hashicorp/azuread"
      version = "~> 2.22.0"
    }
  }
}

provider "azurerm" {
  features {}
}
