package commands

import (
	"testing"
	"testing/fstest"
)

func TestList(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "no modules",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := List()
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExtractTitle(t *testing.T) {
	tests := []struct {
		name     string
		fs       fstest.MapFS
		filename string
		want     string
	}{
		{
			name:     "file not found",
			filename: "/non/existent/file.md",
			want:     "No title",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractTitle(tt.filename)
			if got != tt.want {
				t.Errorf("extractTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
