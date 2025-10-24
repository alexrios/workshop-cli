package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fatih/color"
)

func List() error {
	color.Blue("Available Lessons:\n")
	fmt.Println()

	// Iterate through modules 01-13
	for i := 1; i <= 13; i++ {
		moduleDir := fmt.Sprintf("module%02d", i)
		if !dirExists(moduleDir) {
			continue
		}

		fmt.Printf("=== %s ===\n", moduleDir)

		// Find all lesson directories
		lessons, err := filepath.Glob(filepath.Join(moduleDir, "lesson*"))
		if err != nil {
			continue
		}

		// Sort lessons
		sort.Strings(lessons)

		// List each lesson with its title
		for _, lessonPath := range lessons {
			if !dirExists(lessonPath) {
				continue
			}

			lessonName := filepath.Base(lessonPath)
			lessonMd := filepath.Join(lessonPath, "LESSON.md")

			title := "No title"
			if fileExists(lessonMd) {
				title = extractTitle(lessonMd)
			}

			fmt.Printf("  %s: %s\n", lessonName, title)
		}

		fmt.Println()
	}

	return nil
}

func extractTitle(lessonFile string) string {
	file, err := os.Open(lessonFile)
	if err != nil {
		return "No title"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}

	return "No title"
}
