package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexrios/workshop-cli/internal/diff"
	"github.com/fatih/color"
)

func Compare(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("missing arguments\nUsage: workshop compare <module> <lesson> [file]\n\nExample: workshop compare 01 1.2\nExample: workshop compare 04 4.1 cache.go")
	}

	module := normalizeModule(args[0])
	lessonNum := args[1]
	specificFile := ""
	if len(args) > 2 {
		specificFile = args[2]
	}

	// Construct paths
	moduleDir := fmt.Sprintf("module%s", module)
	lessonDir := filepath.Join(moduleDir, fmt.Sprintf("lesson%s", lessonNum))
	workDir := filepath.Join(lessonDir, "work")
	solutionDir := filepath.Join(lessonDir, "solution")

	// Check if directories exist
	if !dirExists(workDir) {
		return fmt.Errorf("work directory not found: %s\nRun 'workshop setup %s %s' first", workDir, module, lessonNum)
	}

	if !dirExists(solutionDir) {
		return fmt.Errorf("solution directory not found: %s", solutionDir)
	}

	// Compare specific file or all files
	if specificFile != "" {
		return compareFile(workDir, solutionDir, specificFile)
	}

	return compareAllFiles(workDir, solutionDir, lessonNum)
}

func compareFile(workDir, solutionDir, filename string) error {
	workFile := filepath.Join(workDir, filename)
	solutionFile := filepath.Join(solutionDir, filename)

	if !fileExists(workFile) {
		return fmt.Errorf("file not found in work directory: %s", workFile)
	}

	if !fileExists(solutionFile) {
		return fmt.Errorf("file not found in solution directory: %s", solutionFile)
	}

	color.Blue("Comparing %s...\n", filename)
	fmt.Println()

	return diff.ShowDiff(workFile, solutionFile)
}

func compareAllFiles(workDir, solutionDir, lessonNum string) error {
	color.Blue("Comparing all files in lesson %s...\n", lessonNum)
	fmt.Println()

	// Get all .go files from work directory
	workFiles, err := filepath.Glob(filepath.Join(workDir, "*.go"))
	if err != nil {
		return fmt.Errorf("failed to list work files: %w", err)
	}

	if len(workFiles) == 0 {
		return fmt.Errorf("no Go files found in work directory")
	}

	hasChanges := false
	for _, workFile := range workFiles {
		filename := filepath.Base(workFile)
		solutionFile := filepath.Join(solutionDir, filename)

		if !fileExists(solutionFile) {
			continue
		}

		color.Yellow("=== %s ===\n", filename)

		// Check if files are identical
		if filesIdentical(workFile, solutionFile) {
			color.Green("✓ Files are identical\n")
		} else {
			hasChanges = true
			if err := diff.ShowDiff(workFile, solutionFile); err != nil {
				return err
			}
		}
		fmt.Println()
	}

	if !hasChanges {
		color.Green("All files match the solution!\n")
	}

	return nil
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func filesIdentical(file1, file2 string) bool {
	f1, err := os.Open(file1)
	if err != nil {
		return false
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		return false
	}
	defer f2.Close()

	scanner1 := bufio.NewScanner(f1)
	scanner2 := bufio.NewScanner(f2)

	for scanner1.Scan() {
		if !scanner2.Scan() {
			return false
		}
		if scanner1.Text() != scanner2.Text() {
			return false
		}
	}

	// Check if file2 has more lines
	if scanner2.Scan() {
		return false
	}

	return scanner1.Err() == nil && scanner2.Err() == nil
}
