package main

import (
	"fmt"
	"os"

	"github.com/alexrios/workshop-cli/internal/commands"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		printUsage()
		return nil
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "setup":
		return commands.Setup(args)
	case "compare":
		return commands.Compare(args)
	case "reset":
		if len(args) > 0 && args[0] == "all" {
			return commands.ResetAll()
		}
		return commands.Reset(args)
	case "list":
		return commands.List()
	case "status":
		return commands.Status()
	case "validate":
		return commands.Validate()
	case "help", "-h", "--help":
		printUsage()
		return nil
	default:
		return fmt.Errorf("unknown command: %s", cmd)
	}
}

func printUsage() {
	usage := `Workshop CLI - Alex Rios Workshop Management

Usage:
  workshop <command> [arguments]

Commands:
  setup <module> <lesson>       Setup a lesson workspace
  compare <module> <lesson> [file]  Compare work with solution
  reset <module> <lesson>       Reset lesson to starter state
  reset all                     Reset ALL lessons to starter state
  list                          List all available lessons
  status                        Show lessons in progress
  validate                      Validate all starter files compile
  help                          Show this help message

Examples:
  workshop setup 01 1.2         # Setup lesson 1.2 from module 01
  workshop compare 04 4.1       # Compare all files
  workshop compare 04 4.1 cache.go  # Compare specific file
  workshop reset 01 1.2         # Reset lesson
  workshop reset all            # Reset ALL lessons
  workshop list                 # List all lessons
  workshop status               # Show progress

`
	fmt.Print(usage)
}
