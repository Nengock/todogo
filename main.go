package main

import (
	"flag"
	"fmt"
	"os"

	"todo-app/commands"
)

func main() {
	// Define CLI commands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	// Parse the command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("expected 'add', 'list', 'complete' or 'delete' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if addCmd.Parsed() {
			if len(addCmd.Args()) < 1 {
				fmt.Println("Description is required for 'add' command")
				os.Exit(1)
			}
			commands.AddTask(addCmd.Arg(0))
		}
	case "list":
		listCmd.Parse(os.Args[2:])
		if listCmd.Parsed() {
			commands.ListTasks()
		}
	case "complete":
		completeCmd.Parse(os.Args[2:])
		if completeCmd.Parsed() {
			if len(completeCmd.Args()) < 1 {
				fmt.Println("Task ID is required for 'complete' command")
				os.Exit(1)
			}
			commands.CompleteTask(completeCmd.Arg(0))
		}
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if deleteCmd.Parsed() {
			if len(deleteCmd.Args()) < 1 {
				fmt.Println("Task ID is required for 'delete' command")
				os.Exit(1)
			}
			commands.DeleteTask(deleteCmd.Arg(0))
		}
	default:
		fmt.Println("expected 'add', 'list', 'complete' or 'delete' subcommands")
		os.Exit(1)
	}
}
