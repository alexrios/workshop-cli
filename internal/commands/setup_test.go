package commands

import (
	"strings"
	"testing"
	"testing/fstest"
)

func TestSetup(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		fs      fstest.MapFS
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
		{
			name:    "lesson directory not found",
			args:    []string{"01", "1.2"},
			fs:      fstest.MapFS{},
			wantErr: true,
			errMsg:  "lesson directory not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Setup(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Setup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("Setup() error = %v, want error containing %v", err, tt.errMsg)
			}
		})
	}
}

func TestNormalizeModule(t *testing.T) {
	tests := []struct {
		name   string
		module string
		want   string
	}{
		{
			name:   "single digit",
			module: "1",
			want:   "01",
		},
		{
			name:   "double digit",
			module: "10",
			want:   "10",
		},
		{
			name:   "already normalized",
			module: "05",
			want:   "05",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeModule(tt.module); got != tt.want {
				t.Errorf("normalizeModule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountFiles(t *testing.T) {
	tests := []struct {
		name    string
		dir     string
		want    int
		wantErr bool
	}{
		{
			name:    "non-existent directory",
			dir:     "/non/existent/path",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := countFiles(tt.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("countFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("countFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirExists(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "non-existent path",
			path: "/non/existent/path",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dirExists(tt.path); got != tt.want {
				t.Errorf("dirExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
