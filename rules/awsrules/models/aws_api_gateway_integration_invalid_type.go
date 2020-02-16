// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsAPIGatewayIntegrationInvalidTypeRule checks the pattern is valid
type AwsAPIGatewayIntegrationInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAPIGatewayIntegrationInvalidTypeRule returns new rule with default attributes
func NewAwsAPIGatewayIntegrationInvalidTypeRule() *AwsAPIGatewayIntegrationInvalidTypeRule {
	return &AwsAPIGatewayIntegrationInvalidTypeRule{
		resourceType:  "aws_api_gateway_integration",
		attributeName: "type",
		enum: []string{
			"HTTP",
			"AWS",
			"MOCK",
			"HTTP_PROXY",
			"AWS_PROXY",
		},
	}
}

// Name returns the rule name
func (r *AwsAPIGatewayIntegrationInvalidTypeRule) Name() string {
	return "aws_api_gateway_integration_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAPIGatewayIntegrationInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAPIGatewayIntegrationInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAPIGatewayIntegrationInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAPIGatewayIntegrationInvalidTypeRule) Check(runner *tflint.Runner) error {
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
