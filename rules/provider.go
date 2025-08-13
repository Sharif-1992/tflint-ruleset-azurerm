package rules

import (
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "github.com/terraform-linters/tflint-ruleset-azurerm/rules/apispec"
    "github.com/terraform-linters/tflint-ruleset-azurerm/rules/storage_account" // import your package
)

// Rules is a list of all rules
var Rules = append([]tflint.Rule{
    NewAzurermKubernetesClusterDefaultNodePoolInvalidVMSizeRule(),
    NewAzurermKubernetesClusterNodePoolInvalidVMSizeRule(),
    NewAzurermLinuxVirtualMachineInvalidSizeRule(),
    NewAzurermLinuxVirtualMachineScaleSetInvalidSkuRule(),
    NewAzurermVirtualMachineInvalidVMSizeRule(),
    NewAzurermWindowsVirtualMachineInvalidNameRule(),
    NewAzurermWindowsVirtualMachineInvalidSizeRule(),
    NewAzurermWindowsVirtualMachineScaleSetInvalidSkuRule(),
    NewAzurermResourceMissingTagsRule(),

    // ✅ Add your custom rule here
    storage_account.NewRequiresTagsRule(),
}, apispec.Rules...)
