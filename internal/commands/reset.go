package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func Reset(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("missing arguments\nUsage: workshop reset <module> <lesson>\nExample: workshop reset 01 1.2")
	}

	module := normalizeModule(args[0])
	lessonNum := args[1]

	// Construct paths
	moduleDir := fmt.Sprintf("module%s", module)
	lessonDir := filepath.Join(moduleDir, fmt.Sprintf("lesson%s", lessonNum))
	workDir := filepath.Join(lessonDir, "work")

	// Check if work directory exists
	if !dirExists(workDir) {
		color.Yellow("No work directory found for lesson %s.%s\n", module, lessonNum)
		return nil
	}

	// Confirm deletion
	color.Yellow("Warning: This will delete all your work for lesson %s.%s\n", module, lessonNum)
	fmt.Printf("Work directory: %s\n", workDir)
	fmt.Print("Are you sure? (y/N): ")

	var response string
	fmt.Scanln(&response)

	if !strings.EqualFold(strings.TrimSpace(response), "y") {
		color.Blue("Cancelled. Your work is safe.\n")
		return nil
	}

	// Remove work directory
	if err := os.RemoveAll(workDir); err != nil {
		return fmt.Errorf("failed to remove work directory: %w", err)
	}

	color.Green("✓ Lesson %s.%s has been reset\n", module, lessonNum)
	fmt.Println()
	fmt.Printf("Run 'workshop setup %s %s' to start fresh\n", module, lessonNum)

	return nil
}

func ResetAll() error {
	// Find all work directories
	workDirs, err := filepath.Glob("module*/lesson*/work")
	if err != nil {
		return fmt.Errorf("failed to find work directories: %w", err)
	}

	if len(workDirs) == 0 {
		color.Yellow("No work directories found.\n")
		return nil
	}

	// Confirm deletion
	color.Yellow("Warning: This will delete ALL your work for ALL lessons\n")
	fmt.Printf("Found %d work directories:\n", len(workDirs))
	for _, dir := range workDirs {
		fmt.Printf("  - %s\n", dir)
	}
	fmt.Println()
	fmt.Print("Are you sure you want to reset everything? (y/N): ")

	var response string
	fmt.Scanln(&response)

	if !strings.EqualFold(strings.TrimSpace(response), "y") {
		color.Blue("Cancelled. Your work is safe.\n")
		return nil
	}

	// Remove all work directories
	removed := 0
	failed := 0
	for _, workDir := range workDirs {
		if err := os.RemoveAll(workDir); err != nil {
			color.Red("✗ Failed to remove %s: %v\n", workDir, err)
			failed++
		} else {
			removed++
		}
	}

	fmt.Println()
	if failed > 0 {
		color.Yellow("Reset complete: %d removed, %d failed\n", removed, failed)
	} else {
		color.Green("✓ Successfully reset all %d lessons\n", removed)
	}

	return nil
}
