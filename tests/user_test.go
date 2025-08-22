package tests

import (
	"testing"
	"grpc-app/utils"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		email string
		valid bool
	}{
		{"test@example.com", true},
		{"invalid-email", false},
		{"", false},
	}

	for _, test := range tests {
		result := utils.ValidateEmail(test.email)
		if result != test.valid {
			t.Errorf("ValidateEmail(%s) = %v; want %v", test.email, result, test.valid)
		}
	}
}

func TestValidateName(t *testing.T) {
	tests := []struct {
		name  string
		valid bool
	}{
		{"João Silva", true},
		{"A", false},
		{"", false},
	}

	for _, test := range tests {
		result := utils.ValidateName(test.name)
		if result != test.valid {
			t.Errorf("ValidateName(%s) = %v; want %v", test.name, result, test.valid)
		}
	}
}
