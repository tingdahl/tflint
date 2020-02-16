// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsCognitoUserPoolInvalidSmsVerificationMessageRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cognito_user_pool" "foo" {
	sms_verification_message = "Verification code"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCognitoUserPoolInvalidSmsVerificationMessageRule(),
					Message: `"Verification code" does not match valid pattern ^.*\{####\}.*$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cognito_user_pool" "foo" {
	sms_verification_message = "Verification code is {####}"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCognitoUserPoolInvalidSmsVerificationMessageRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
