// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsEc2TransitGatewayInvalidDNSSupportRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_ec2_transit_gateway" "foo" {
	dns_support = "enabled"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsEc2TransitGatewayInvalidDNSSupportRule(),
					Message: `"enabled" is an invalid value as dns_support`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_ec2_transit_gateway" "foo" {
	dns_support = "enable"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsEc2TransitGatewayInvalidDNSSupportRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
