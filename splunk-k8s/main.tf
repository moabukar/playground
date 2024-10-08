terraform {
  required_version = ">= 0.13"

  required_providers {
    # https://github.com/terraform-providers/terraform-provider-azurerm/releases
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 2.64.0"
    }

    # https://github.com/terraform-providers/terraform-provider-azuread/releases
    azuread = {
      source  = "hashicorp/azuread"
      version = "~> 1.5.1"
    }

    # https://github.com/hashicorp/terraform-provider-kubernetes/releases
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.3.2"
    }

    # https://github.com/hashicorp/terraform-provider-helm/releases
    helm = {
      source  = "hashicorp/helm"
      version = "~> 2.2.0"
    }
  }
}

provider "azurerm" {
  features {}
}

resource "random_string" "aks" {
  length  = 4
  special = false
  upper   = false
}

locals {
  # az aks get-versions --location uksouth --output table
  kubernetes_version  = "1.20.7"
  location            = "uksouth"
  prefix              = "ar${random_string.aks.result}" # aks dns_prefix must start with a letter
  resource_group_name = "${local.prefix}-rg-azurerm-kubernetes-cluster"
  name                = "${local.prefix}-aks-cluster"

  tags = {
    App    = "splunk"
    Env    = "Dev"
    Owner  = "Mo"
    Source = "terraform"
    # Must add following tags to avoid automatic changes by Azure Policy
    application      = "splunk-testing"
    project          = "splunk-testing"
    service-offering = "splunk-testing"
    uin              = "splunk-testing"
  }
}

resource "azurerm_resource_group" "aks" {
  name     = local.resource_group_name
  location = local.location
  tags     = local.tags
}

module "aks" {
  source = "moabukar/aks/azurerm"

  kubernetes_version  = local.kubernetes_version
  location            = azurerm_resource_group.aks.location
  resource_group_name = azurerm_resource_group.aks.name
  name                = local.name
  tags                = local.tags

  # Add existing group to the new AKS cluster admin group
  aks_admin_group_member_name = "splunk_admins"

  # override defaults
  default_node_pool = {
    vm_size  = "Standard_DS5_v2" # 16 vCPU,	56GiB Mem
    count    = 3
    max_pods = 99
  }
}

output "aks_credentials_command" {
  value = "az aks get-credentials --resource-group ${azurerm_resource_group.aks.name} --name ${module.aks.name} --overwrite-existing --admin"
}

output "full_object" {
  value     = module.aks.full_object
  sensitive = true
}
