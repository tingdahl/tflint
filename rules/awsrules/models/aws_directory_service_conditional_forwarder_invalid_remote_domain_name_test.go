// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsDirectoryServiceConditionalForwarderInvalidRemoteDomainNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_directory_service_conditional_forwarder" "foo" {
	remote_domain_name = "example^com"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsDirectoryServiceConditionalForwarderInvalidRemoteDomainNameRule(),
					Message: `"example^com" does not match valid pattern ^([a-zA-Z0-9]+[\\.-])+([a-zA-Z0-9])+[.]?$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_directory_service_conditional_forwarder" "foo" {
	remote_domain_name = "example.com"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsDirectoryServiceConditionalForwarderInvalidRemoteDomainNameRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
