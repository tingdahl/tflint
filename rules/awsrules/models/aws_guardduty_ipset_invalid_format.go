// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsGuarddutyIpsetInvalidFormatRule checks the pattern is valid
type AwsGuarddutyIpsetInvalidFormatRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	enum          []string
}

// NewAwsGuarddutyIpsetInvalidFormatRule returns new rule with default attributes
func NewAwsGuarddutyIpsetInvalidFormatRule() *AwsGuarddutyIpsetInvalidFormatRule {
	return &AwsGuarddutyIpsetInvalidFormatRule{
		resourceType:  "aws_guardduty_ipset",
		attributeName: "format",
		max:           300,
		min:           1,
		enum: []string{
			"TXT",
			"STIX",
			"OTX_CSV",
			"ALIEN_VAULT",
			"PROOF_POINT",
			"FIRE_EYE",
		},
	}
}

// Name returns the rule name
func (r *AwsGuarddutyIpsetInvalidFormatRule) Name() string {
	return "aws_guardduty_ipset_invalid_format"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGuarddutyIpsetInvalidFormatRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGuarddutyIpsetInvalidFormatRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGuarddutyIpsetInvalidFormatRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGuarddutyIpsetInvalidFormatRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"format must be 300 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"format must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as format`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
