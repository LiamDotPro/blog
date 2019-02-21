package users

import (
	"testing"
)

func Test_validNewUser(t *testing.T) {
	tests := []struct {
		model      newUserModel
		errorCount int
	}{
		{newUserModel{Admin: false, Password: "", Name: ""}, 1},
		{newUserModel{Admin: false, Password: "Password1234", Name: ""}, 0},
		{newUserModel{Admin: false, Password: "TooShor", Name: ""}, 1},
	}

	for i, tt := range tests {
		if errors := validNewUser(tt.model); len(errors) != tt.errorCount {
			t.Errorf("validNewUser() case %v = %v, want %v", i, tt.errorCount, len(errors))
		}
	}
}
