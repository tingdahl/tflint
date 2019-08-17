// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudformationStackSetInvalidAdministrationRoleArnRule checks the pattern is valid
type AwsCloudformationStackSetInvalidAdministrationRoleArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudformationStackSetInvalidAdministrationRoleArnRule returns new rule with default attributes
func NewAwsCloudformationStackSetInvalidAdministrationRoleArnRule() *AwsCloudformationStackSetInvalidAdministrationRoleArnRule {
	return &AwsCloudformationStackSetInvalidAdministrationRoleArnRule{
		resourceType:  "aws_cloudformation_stack_set",
		attributeName: "administration_role_arn",
		max:           2048,
		min:           20,
	}
}

// Name returns the rule name
func (r *AwsCloudformationStackSetInvalidAdministrationRoleArnRule) Name() string {
	return "aws_cloudformation_stack_set_invalid_administration_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudformationStackSetInvalidAdministrationRoleArnRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsCloudformationStackSetInvalidAdministrationRoleArnRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudformationStackSetInvalidAdministrationRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudformationStackSetInvalidAdministrationRoleArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"administration_role_arn must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"administration_role_arn must be 20 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}