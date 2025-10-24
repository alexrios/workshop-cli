package commands

import (
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
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
		{
			name:    "work directory not found",
			args:    []string{"01", "1.2"},
			wantErr: true,
			errMsg:  "work directory not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Compare(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("Compare() error = %v, want error containing %v", err, tt.errMsg)
			}
		})
	}
}

func TestFileExists(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "non-existent file",
			path: "/non/existent/file.go",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileExists(tt.path); got != tt.want {
				t.Errorf("fileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilesIdentical(t *testing.T) {
	tests := []struct {
		name  string
		file1 string
		file2 string
		want  bool
	}{
		{
			name:  "non-existent files",
			file1: "/non/existent/file1.go",
			file2: "/non/existent/file2.go",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filesIdentical(tt.file1, tt.file2); got != tt.want {
				t.Errorf("filesIdentical() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareFile(t *testing.T) {
	tests := []struct {
		name        string
		workDir     string
		solutionDir string
		filename    string
		wantErr     bool
		errMsg      string
	}{
		{
			name:        "file not found in work directory",
			workDir:     "/non/existent/work",
			solutionDir: "/non/existent/solution",
			filename:    "test.go",
			wantErr:     true,
			errMsg:      "file not found in work directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := compareFile(tt.workDir, tt.solutionDir, tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("compareFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("compareFile() error = %v, want error containing %v", err, tt.errMsg)
			}
		})
	}
}

func TestCompareAllFiles(t *testing.T) {
	tests := []struct {
		name        string
		workDir     string
		solutionDir string
		lessonNum   string
		wantErr     bool
		errMsg      string
	}{
		{
			name:        "no go files found",
			workDir:     "/non/existent/work",
			solutionDir: "/non/existent/solution",
			lessonNum:   "1.2",
			wantErr:     true,
			errMsg:      "no Go files found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := compareAllFiles(tt.workDir, tt.solutionDir, tt.lessonNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("compareAllFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err != nil && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("compareAllFiles() error = %v, want error containing %v", err, tt.errMsg)
			}
		})
	}
}
