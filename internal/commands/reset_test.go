package commands

import (
	"strings"
	"testing"
)

func TestReset(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "missing arguments",
			args:    []string{},
			wantErr: true,
			errMsg:  "missing arguments",
		},
		{
			name:    "missing lesson argument",
			args:    []string{"01"},
			wantErr: true,
			errMsg:  "missing arguments",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Reset(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("Reset() error = %v, want error containing %v", err, tt.errMsg)
			}
		})
	}
}

func TestResetAll(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "no work directories",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ResetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
