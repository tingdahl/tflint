// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSesReceiptFilterInvalidPolicyRule checks the pattern is valid
type AwsSesReceiptFilterInvalidPolicyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSesReceiptFilterInvalidPolicyRule returns new rule with default attributes
func NewAwsSesReceiptFilterInvalidPolicyRule() *AwsSesReceiptFilterInvalidPolicyRule {
	return &AwsSesReceiptFilterInvalidPolicyRule{
		resourceType:  "aws_ses_receipt_filter",
		attributeName: "policy",
		enum: []string{
			"Block",
			"Allow",
		},
	}
}

// Name returns the rule name
func (r *AwsSesReceiptFilterInvalidPolicyRule) Name() string {
	return "aws_ses_receipt_filter_invalid_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSesReceiptFilterInvalidPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSesReceiptFilterInvalidPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSesReceiptFilterInvalidPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSesReceiptFilterInvalidPolicyRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as policy`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
