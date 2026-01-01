package misc

import (
	"testing"
)

func TestValidateNoSpaces(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		wantErr bool
	}{
		{"No spaces", "validname", false},
		{"With spaces", "invalid name", true},
		{"Empty", "", false},
		{"With tab", "invalid\tname", false}, // Code checks strings.Contains(str, " ")
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateNoSpaces(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("ValidateNoSpaces() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
