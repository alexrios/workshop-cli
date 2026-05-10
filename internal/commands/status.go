package commands

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/fatih/color"
)

func Status() error {
	color.Blue("Lessons in Progress:\n")
	fmt.Println()

	found := false

	// Iterate through modules 01-15
	for i := 1; i <= 15; i++ {
		moduleDir := fmt.Sprintf("module%02d", i)
		if !dirExists(moduleDir) {
			continue
		}

		hasWork := false
		var workLessons []string

		// Find all lesson directories with work subdirectories
		lessons, err := filepath.Glob(filepath.Join(moduleDir, "lesson*"))
		if err != nil {
			continue
		}

		sort.Strings(lessons)

		for _, lessonPath := range lessons {
			workDir := filepath.Join(lessonPath, "work")
			if dirExists(workDir) {
				hasWork = true
				found = true

				lessonName := filepath.Base(lessonPath)

				// Count Go files
				goFiles, _ := filepath.Glob(filepath.Join(workDir, "*.go"))
				fileCount := len(goFiles)

				workLessons = append(workLessons, fmt.Sprintf("  ✓ %s (%d Go files)", lessonName, fileCount))
			}
		}

		// Print module header and lessons if any have work
		if hasWork {
			fmt.Printf("=== %s ===\n", moduleDir)
			for _, lesson := range workLessons {
				fmt.Println(lesson)
			}
			fmt.Println()
		}
	}

	if !found {
		color.Yellow("No lessons in progress yet.\n")
		fmt.Println("Run 'workshop setup <module> <lesson>' to get started!")
	}

	return nil
}
