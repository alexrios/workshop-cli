package diff

import (
	"reflect"
	"testing"
)

func TestComputeDiff(t *testing.T) {
	tests := []struct {
		name   string
		lines1 []string
		lines2 []string
		want   []DiffLine
	}{
		{
			name:   "identical files",
			lines1: []string{"line1", "line2", "line3"},
			lines2: []string{"line1", "line2", "line3"},
			want: []DiffLine{
				{Type: DiffContext, Line: "line1"},
				{Type: DiffContext, Line: "line2"},
				{Type: DiffContext, Line: "line3"},
			},
		},
		{
			name:   "addition at end",
			lines1: []string{"line1", "line2"},
			lines2: []string{"line1", "line2", "line3"},
			want: []DiffLine{
				{Type: DiffContext, Line: "line1"},
				{Type: DiffContext, Line: "line2"},
				{Type: DiffAdd, Line: "line3"},
			},
		},
		{
			name:   "removal at end",
			lines1: []string{"line1", "line2", "line3"},
			lines2: []string{"line1", "line2"},
			want: []DiffLine{
				{Type: DiffContext, Line: "line1"},
				{Type: DiffContext, Line: "line2"},
				{Type: DiffRemove, Line: "line3"},
			},
		},
		{
			name:   "modification in middle",
			lines1: []string{"line1", "old", "line3"},
			lines2: []string{"line1", "new", "line3"},
			want: []DiffLine{
				{Type: DiffContext, Line: "line1"},
				{Type: DiffRemove, Line: "old"},
				{Type: DiffAdd, Line: "new"},
				{Type: DiffContext, Line: "line3"},
			},
		},
		{
			name:   "insertion in middle",
			lines1: []string{"line1", "line3"},
			lines2: []string{"line1", "line2", "line3"},
			want: []DiffLine{
				{Type: DiffContext, Line: "line1"},
				{Type: DiffAdd, Line: "line2"},
				{Type: DiffContext, Line: "line3"},
			},
		},
		{
			name:   "removal in middle",
			lines1: []string{"line1", "line2", "line3"},
			lines2: []string{"line1", "line3"},
			want: []DiffLine{
				{Type: DiffContext, Line: "line1"},
				{Type: DiffRemove, Line: "line2"},
				{Type: DiffContext, Line: "line3"},
			},
		},
		{
			name:   "empty files",
			lines1: []string{},
			lines2: []string{},
			want:   nil,
		},
		{
			name:   "empty to non-empty",
			lines1: []string{},
			lines2: []string{"line1", "line2"},
			want: []DiffLine{
				{Type: DiffAdd, Line: "line1"},
				{Type: DiffAdd, Line: "line2"},
			},
		},
		{
			name:   "non-empty to empty",
			lines1: []string{"line1", "line2"},
			lines2: []string{},
			want: []DiffLine{
				{Type: DiffRemove, Line: "line1"},
				{Type: DiffRemove, Line: "line2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := computeDiff(tt.lines1, tt.lines2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShowDiff(t *testing.T) {
	tests := []struct {
		name    string
		file1   string
		file2   string
		wantErr bool
	}{
		{
			name:    "non-existent file1",
			file1:   "/non/existent/file1.go",
			file2:   "/non/existent/file2.go",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ShowDiff(tt.file1, tt.file2)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShowDiff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadFileLines(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			name:     "non-existent file",
			filename: "/non/existent/file.go",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := readFileLines(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFileLines() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
