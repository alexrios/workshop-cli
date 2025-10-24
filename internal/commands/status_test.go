package commands

import (
	"testing"
)

func TestStatus(t *testing.T) {
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
			err := Status()
			if (err != nil) != tt.wantErr {
				t.Errorf("Status() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
