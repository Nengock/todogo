// Suggested code may be subject to a license. Learn more: ~LicenseLog:3665122573.
package main

import (
	"flag"
	"fmt"
	"os"

	"todo-app/commands"
)

func main() {
	// Define CLI commands using flag.NewFlagSet
	// Each command has its own FlagSet for parsing arguments specific to that command.
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	// Check if at least one command-line argument is provided.
	// If not, print an error message and exit.
	if len(os.Args) < 2 {
		fmt.Println("expected 'add', 'list', 'complete' or 'delete' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		// Parse arguments specific to the "add" command.
		addCmd.Parse(os.Args[2:])
		if addCmd.Parsed() {
			if len(addCmd.Args()) < 1 {
				fmt.Println("Description is required for 'add' command")
				os.Exit(1)
			}
			commands.AddTask(addCmd.Arg(0))
		}
	case "list":
		// Parse arguments specific to the "list" command.
		listCmd.Parse(os.Args[2:])
		if listCmd.Parsed() {
			commands.ListTasks()
		}
	case "complete":
		// Parse arguments specific to the "complete" command.
		completeCmd.Parse(os.Args[2:])
		if completeCmd.Parsed() {
			if len(completeCmd.Args()) < 1 {
				fmt.Println("Task ID is required for 'complete' command")
				os.Exit(1)
			}
			commands.CompleteTask(completeCmd.Arg(0))
		}
	case "delete":
		// Parse arguments specific to the "delete" command.
		deleteCmd.Parse(os.Args[2:])
		if deleteCmd.Parsed() {
			if len(deleteCmd.Args()) < 1 {
				fmt.Println("Task ID is required for 'delete' command")
				os.Exit(1)
			}
			commands.DeleteTask(deleteCmd.Arg(0))
		}
	default:
		// Handle unknown subcommands.
		fmt.Println("expected 'add', 'list', 'complete' or 'delete' subcommands")
		os.Exit(1)
	}
}
