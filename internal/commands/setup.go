package commands

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func Setup(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("missing arguments\nUsage: workshop setup <module> <lesson>\nExample: workshop setup 01 1.2")
	}

	module := normalizeModule(args[0])
	lessonNum := args[1]

	// Construct paths
	moduleDir := fmt.Sprintf("module%s", module)
	lessonDir := filepath.Join(moduleDir, fmt.Sprintf("lesson%s", lessonNum))
	starterDir := filepath.Join(lessonDir, "starter")
	solutionDir := filepath.Join(lessonDir, "solution")
	workDir := filepath.Join(lessonDir, "work")

	// Check if lesson exists
	if !dirExists(lessonDir) {
		return fmt.Errorf("lesson directory not found: %s", lessonDir)
	}

	// Check if starter directory exists
	if !dirExists(starterDir) {
		return fmt.Errorf("starter directory not found: %s\nThis lesson may not have starter files available.", starterDir)
	}

	// Check if work directory already exists
	if dirExists(workDir) {
		color.Yellow("Warning: Work directory already exists: %s\n", workDir)
		fmt.Print("Do you want to overwrite it? (y/N): ")

		var response string
		fmt.Scanln(&response)

		if !strings.EqualFold(strings.TrimSpace(response), "y") {
			color.Blue("Keeping existing work directory.\n")
			return nil
		}

		if err := os.RemoveAll(workDir); err != nil {
			return fmt.Errorf("failed to remove existing work directory: %w", err)
		}
	}

	// Create work directory and copy starter files
	color.Blue("Setting up lesson %s.%s...\n", module, lessonNum)

	if err := copyDir(starterDir, workDir); err != nil {
		return fmt.Errorf("failed to copy starter files: %w", err)
	}

	// Count files copied
	fileCount, err := countFiles(workDir)
	if err != nil {
		return fmt.Errorf("failed to count files: %w", err)
	}

	color.Green("✓ Copied %d starter files to: %s\n", fileCount, workDir)
	fmt.Println()

	fmt.Println("Files ready for implementation:")
	files, _ := findGoFiles(workDir)
	for _, file := range files {
		fmt.Printf("  - %s\n", file)
	}

	fmt.Println()
	color.Blue("Next steps:\n")
	fmt.Printf("  1. cd %s\n", workDir)
	fmt.Printf("  2. Read the lesson: cat ../LESSON.md\n")
	fmt.Printf("  3. Read the exercise: cat ../EXERCISE.md\n")
	fmt.Println("  4. Implement the TODOs in the starter files")
	fmt.Println("  5. Run tests: go test -v")
	fmt.Println()
	color.Blue("Reference:\n")
	fmt.Printf("  Solution files available at: %s/\n", solutionDir)
	fmt.Println()

	return nil
}

func normalizeModule(module string) string {
	if len(module) == 1 {
		return "0" + module
	}
	return module
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func copyDir(src, dst string) error {
	// Create destination directory
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}

	// Read source directory
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// Copy each entry
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			// Recursively copy subdirectory
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			// Copy file
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	// Copy permissions
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	return os.Chmod(dst, srcInfo.Mode())
}

func countFiles(dir string) (int, error) {
	count := 0
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			count++
		}
		return nil
	})
	return count, err
}

func findGoFiles(dir string) ([]string, error) {
	var files []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			name := entry.Name()
			if strings.HasSuffix(name, ".go") || name == "go.mod" {
				files = append(files, filepath.Join(dir, name))
			}
		}
	}
	return files, nil
}
