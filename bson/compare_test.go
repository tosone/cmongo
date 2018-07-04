package bson

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		oid string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{"5b3d005279bc20391c35f0c1"}, true},
		{"unusual", args{"5b3d005279bc20391c35f0c"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.oid); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
