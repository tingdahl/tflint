// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsNetworkACLRuleInvalidRuleActionRule checks the pattern is valid
type AwsNetworkACLRuleInvalidRuleActionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsNetworkACLRuleInvalidRuleActionRule returns new rule with default attributes
func NewAwsNetworkACLRuleInvalidRuleActionRule() *AwsNetworkACLRuleInvalidRuleActionRule {
	return &AwsNetworkACLRuleInvalidRuleActionRule{
		resourceType:  "aws_network_acl_rule",
		attributeName: "rule_action",
		enum: []string{
			"allow",
			"deny",
		},
	}
}

// Name returns the rule name
func (r *AwsNetworkACLRuleInvalidRuleActionRule) Name() string {
	return "aws_network_acl_rule_invalid_rule_action"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsNetworkACLRuleInvalidRuleActionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsNetworkACLRuleInvalidRuleActionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsNetworkACLRuleInvalidRuleActionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsNetworkACLRuleInvalidRuleActionRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as rule_action`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
