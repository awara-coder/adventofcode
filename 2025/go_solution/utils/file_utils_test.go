package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadFileContents(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		setupFile   bool
		want        []string
		wantErr     bool
	}{
		{
			name:        "success with multiple lines",
			fileContent: "line1\nline2\nline3",
			setupFile:   true,
			want:        []string{"line1", "line2", "line3"},
			wantErr:     false,
		},
		{
			name:        "empty file",
			fileContent: "",
			setupFile:   true,
			want:        []string{},
			wantErr:     false,
		},
		{
			name:        "single line",
			fileContent: "single line",
			setupFile:   true,
			want:        []string{"single line"},
			wantErr:     false,
		},
		{
			name:        "trailing newline",
			fileContent: "line1\nline2\n",
			setupFile:   true,
			want:        []string{"line1", "line2"},
			wantErr:     false,
		},
		{
			name:        "whitespace lines preserved",
			fileContent: "line1\n  \nline3",
			setupFile:   true,
			want:        []string{"line1", "  ", "line3"},
			wantErr:     false,
		},
		{
			name:      "file not found",
			setupFile: false,
			want:      nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var testFile string

			if tt.setupFile {
				tmpDir := t.TempDir()
				testFile = filepath.Join(tmpDir, "test.txt")
				err := os.WriteFile(testFile, []byte(tt.fileContent), 0644)
				if err != nil {
					t.Fatalf("Failed to create test file: %v", err)
				}
			} else {
				testFile = "/path/to/nonexistent/file.txt"
			}

			got, err := ReadFileContents(testFile)

			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileContents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if got != nil {
					t.Errorf("ReadFileContents() = %v, want nil for error case", got)
				}
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("ReadFileContents() returned %d lines, want %d", len(got), len(tt.want))
				return
			}

			for i, line := range got {
				if line != tt.want[i] {
					t.Errorf("ReadFileContents() line %d = %q, want %q", i, line, tt.want[i])
				}
			}
		})
	}
}
