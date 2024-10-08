# Test AKS module
locals {
  # version used for both main AKS API service, and default node pool
  # https://github.com/Azure/AKS/releases
  # az aks get-versions --location uksouth --output table
  kubernetes_version  = "1.21.9"
  prefix              = "rush"
  location            = "uksouth"
  resource_group_name = "${local.prefix}-rg-azurerm-kubernetes-cluster"
  name                = "${local.prefix}-aks-cluster"

  tags = {
    Env    = "Dev"
    Owner  = "Mo"
    Source = "terraform"
  }
}

resource "azurerm_resource_group" "aks" {
  name     = local.resource_group_name
  location = local.location
  tags     = local.tags
}

module "aks" {
  source = "../"

  kubernetes_version   = local.kubernetes_version
  location             = azurerm_resource_group.aks.location
  resource_group_name  = azurerm_resource_group.aks.name
  name                 = local.name
  aad_auth_enabled     = true
  azure_policy_enabled = false
  tags                 = local.tags

  # override defaults
  default_node_pool = {
    count                = 1
    max_pods             = 99
    orchestrator_version = local.kubernetes_version
  }

  # Add existing "AKS-Admins" group to the new AKS cluster admin group
  aks_admin_group_member_name = "AKS-Admins"
}
