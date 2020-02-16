// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCodepipelineInvalidRoleArnRule checks the pattern is valid
type AwsCodepipelineInvalidRoleArnRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsCodepipelineInvalidRoleArnRule returns new rule with default attributes
func NewAwsCodepipelineInvalidRoleArnRule() *AwsCodepipelineInvalidRoleArnRule {
	return &AwsCodepipelineInvalidRoleArnRule{
		resourceType:  "aws_codepipeline",
		attributeName: "role_arn",
		max:           1024,
		pattern:       regexp.MustCompile(`^arn:aws(-[\w]+)*:iam::[0-9]{12}:role/.*$`),
	}
}

// Name returns the rule name
func (r *AwsCodepipelineInvalidRoleArnRule) Name() string {
	return "aws_codepipeline_invalid_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodepipelineInvalidRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodepipelineInvalidRoleArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodepipelineInvalidRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodepipelineInvalidRoleArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"role_arn must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:aws(-[\w]+)*:iam::[0-9]{12}:role/.*$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
