// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsAthenaWorkgroupInvalidDescriptionRule checks the pattern is valid
type AwsAthenaWorkgroupInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsAthenaWorkgroupInvalidDescriptionRule returns new rule with default attributes
func NewAwsAthenaWorkgroupInvalidDescriptionRule() *AwsAthenaWorkgroupInvalidDescriptionRule {
	return &AwsAthenaWorkgroupInvalidDescriptionRule{
		resourceType:  "aws_athena_workgroup",
		attributeName: "description",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsAthenaWorkgroupInvalidDescriptionRule) Name() string {
	return "aws_athena_workgroup_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAthenaWorkgroupInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAthenaWorkgroupInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAthenaWorkgroupInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAthenaWorkgroupInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
