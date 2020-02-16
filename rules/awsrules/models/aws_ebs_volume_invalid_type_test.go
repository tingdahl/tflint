// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsEbsVolumeInvalidTypeRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_ebs_volume" "foo" {
	type = "gp3"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsEbsVolumeInvalidTypeRule(),
					Message: `"gp3" is an invalid value as type`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_ebs_volume" "foo" {
	type = "gp2"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsEbsVolumeInvalidTypeRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
