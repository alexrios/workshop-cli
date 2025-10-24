package diff

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// ShowDiff displays a unified diff between two files
func ShowDiff(file1, file2 string) error {
	lines1, err := readFileLines(file1)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", file1, err)
	}

	lines2, err := readFileLines(file2)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", file2, err)
	}

	// Simple line-by-line diff
	diffs := computeDiff(lines1, lines2)

	// Print diff header
	fmt.Printf("--- %s\n", file1)
	fmt.Printf("+++ %s\n", file2)

	// Print changes
	for _, d := range diffs {
		switch d.Type {
		case DiffAdd:
			color.Green("+%s\n", d.Line)
		case DiffRemove:
			color.Red("-%s\n", d.Line)
		case DiffContext:
			fmt.Printf(" %s\n", d.Line)
		}
	}

	return nil
}

type DiffType int

const (
	DiffContext DiffType = iota
	DiffAdd
	DiffRemove
)

type DiffLine struct {
	Type DiffType
	Line string
}

// computeDiff creates a simple diff between two slices of lines
func computeDiff(lines1, lines2 []string) []DiffLine {
	var result []DiffLine

	// Simple longest common subsequence approach
	maxLen := len(lines1)
	if len(lines2) > maxLen {
		maxLen = len(lines2)
	}

	i, j := 0, 0
	for i < len(lines1) || j < len(lines2) {
		if i >= len(lines1) {
			// Remaining lines in file2 are additions
			result = append(result, DiffLine{Type: DiffAdd, Line: lines2[j]})
			j++
		} else if j >= len(lines2) {
			// Remaining lines in file1 are removals
			result = append(result, DiffLine{Type: DiffRemove, Line: lines1[i]})
			i++
		} else if lines1[i] == lines2[j] {
			// Lines match - context
			result = append(result, DiffLine{Type: DiffContext, Line: lines1[i]})
			i++
			j++
		} else {
			// Lines differ - check if it's a replacement or separate add/remove
			// Simple heuristic: if next line matches, it's a removal
			if i+1 < len(lines1) && lines1[i+1] == lines2[j] {
				result = append(result, DiffLine{Type: DiffRemove, Line: lines1[i]})
				i++
			} else if j+1 < len(lines2) && lines1[i] == lines2[j+1] {
				result = append(result, DiffLine{Type: DiffAdd, Line: lines2[j]})
				j++
			} else {
				// Both are different - mark as remove and add
				result = append(result, DiffLine{Type: DiffRemove, Line: lines1[i]})
				result = append(result, DiffLine{Type: DiffAdd, Line: lines2[j]})
				i++
				j++
			}
		}
	}

	return result
}

func readFileLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimRight(scanner.Text(), "\r\n"))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
