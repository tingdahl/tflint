// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSsmParameterInvalidTypeRule checks the pattern is valid
type AwsSsmParameterInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSsmParameterInvalidTypeRule returns new rule with default attributes
func NewAwsSsmParameterInvalidTypeRule() *AwsSsmParameterInvalidTypeRule {
	return &AwsSsmParameterInvalidTypeRule{
		resourceType:  "aws_ssm_parameter",
		attributeName: "type",
		enum: []string{
			"String",
			"StringList",
			"SecureString",
		},
	}
}

// Name returns the rule name
func (r *AwsSsmParameterInvalidTypeRule) Name() string {
	return "aws_ssm_parameter_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmParameterInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmParameterInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmParameterInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmParameterInvalidTypeRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
