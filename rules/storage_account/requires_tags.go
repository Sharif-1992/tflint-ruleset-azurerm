// adding comment
package storage_account

import (
    "fmt"

    "github.com/terraform-linters/tflint-plugin-sdk/hclext"
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type RequiresTagsRule struct {
    tflint.DefaultRule
}

func NewRequiresTagsRule() *RequiresTagsRule {
    return &RequiresTagsRule{}
}

func (r *RequiresTagsRule) Name() string     { return "azurerm_storage_account_requires_tags" }
func (r *RequiresTagsRule) Enabled() bool    { return true }
func (r *RequiresTagsRule) Severity() tflint.Severity { return tflint.ERROR }
func (r *RequiresTagsRule) Link() string     { return "" }

func (r *RequiresTagsRule) Check(runner tflint.Runner) error {
    // Load all azurerm_storage_account resources
    resources, err := runner.GetResourceContent("azurerm_storage_account", &hclext.BodySchema{
        Attributes: []hclext.AttributeSchema{
            {Name: "tags"},
        },
    }, nil)
    if err != nil {
        return err
    }

    // Loop through each resource block
    for _, block := range resources.Blocks {
        if _, exists := block.Body.Attributes["tags"]; !exists {
            runner.EmitIssue(
                r,
                fmt.Sprintf("Storage account '%s' must have at least one tag", block.Labels[0]),
                block.DefRange,
            )
        }
    }
    return nil
}