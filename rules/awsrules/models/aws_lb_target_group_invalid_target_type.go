// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsLbTargetGroupInvalidTargetTypeRule checks the pattern is valid
type AwsLbTargetGroupInvalidTargetTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsLbTargetGroupInvalidTargetTypeRule returns new rule with default attributes
func NewAwsLbTargetGroupInvalidTargetTypeRule() *AwsLbTargetGroupInvalidTargetTypeRule {
	return &AwsLbTargetGroupInvalidTargetTypeRule{
		resourceType:  "aws_lb_target_group",
		attributeName: "target_type",
		enum: []string{
			"instance",
			"ip",
			"lambda",
		},
	}
}

// Name returns the rule name
func (r *AwsLbTargetGroupInvalidTargetTypeRule) Name() string {
	return "aws_lb_target_group_invalid_target_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLbTargetGroupInvalidTargetTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLbTargetGroupInvalidTargetTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLbTargetGroupInvalidTargetTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLbTargetGroupInvalidTargetTypeRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as target_type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
