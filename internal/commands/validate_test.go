package commands

import (
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "no starter directories",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuildInDir(t *testing.T) {
	tests := []struct {
		name    string
		dir     string
		wantErr bool
	}{
		{
			name:    "non-existent directory",
			dir:     "/non/existent/path",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := buildInDir(tt.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildInDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
