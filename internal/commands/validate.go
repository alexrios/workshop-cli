package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"

	"github.com/fatih/color"
)

func Validate() error {
	color.Blue("Validating starter files...\n")
	fmt.Println()

	failed := 0
	total := 0

	// Find all starter directories
	starterDirs, err := filepath.Glob("module*/lesson*/starter")
	if err != nil {
		return fmt.Errorf("failed to find starter directories: %w", err)
	}

	sort.Strings(starterDirs)

	for _, starterDir := range starterDirs {
		if !dirExists(starterDir) {
			continue
		}

		total++

		// Extract module and lesson from path
		lessonPath := filepath.Dir(starterDir)
		lessonName := filepath.Base(lessonPath)
		moduleName := filepath.Base(filepath.Dir(lessonPath))

		fmt.Printf("Checking %s/%s... ", moduleName, lessonName)

		// Try to build in starter directory
		if err := buildInDir(starterDir); err != nil {
			color.Red("✗ FAILED\n")
			fmt.Printf("  Error: %v\n", err)
			failed++
		} else {
			color.Green("✓\n")
		}
	}
	
	if failed > 0 {
		color.Red("Results: %d/%d passed\n", total-failed, total)
		return fmt.Errorf("%d starter files failed to compile", failed)
	}

	color.Green("Results: %d/%d passed\n", total, total)
	return nil
}

func buildInDir(dir string) error {
	// Save current directory
	origDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(origDir)

	// Change to target directory
	if err := os.Chdir(dir); err != nil {
		return err
	}

	// Run go build
	cmd := exec.Command("go", "build", "./...")
	cmd.Stdout = nil
	cmd.Stderr = nil

	return cmd.Run()
}
