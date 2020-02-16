// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsCloudhsmV2ClusterInvalidHsmTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudhsm_v2_cluster" "foo" {
	hsm_type = "hsm1.micro"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCloudhsmV2ClusterInvalidHsmTypeRule(),
					Message: `"hsm1.micro" does not match valid pattern ^(hsm1\.medium)$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudhsm_v2_cluster" "foo" {
	hsm_type = "hsm1.medium"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCloudhsmV2ClusterInvalidHsmTypeRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
