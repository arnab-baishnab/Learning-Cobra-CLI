package main

import (
	"fmt"
	"os"

	// adjust import path as needed
	"github.com/arnab-baishnab/Learning-Cobra-CLI"
	"github.com/spf13/cobra"
)

const (
	todoFile = ".todos.json"
)

var todos todo.Todos

func main() {
	var rootCmd = &cobra.Command{Use: "todo"}

	var addCmd = &cobra.Command{
		Use:   "add [task]",
		Short: "Add a new todo",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			task := args[0]
			todos.Add(task)
			/*
				if err := todos.Store(todoFile); err != nil {
					fmt.Println("Error saving todos:", err)
					os.Exit(1)
				}
				/**/
			fmt.Println("Added task:", task)
		},
	}

	var completeCmd = &cobra.Command{
		Use:   "complete [index]",
		Short: "Mark a todo as completed",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			index, err := parseIndex(args[0])
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			if err := todos.Complete(index); err != nil {
				fmt.Println("Error completing task:", err)
				os.Exit(1)
			}
			if err := todos.Store(todoFile); err != nil {
				fmt.Println("Error saving todos:", err)
				os.Exit(1)
			}
			fmt.Println("Completed task:", index)
		},
	}

	var deleteCmd = &cobra.Command{
		Use:   "delete [index]",
		Short: "Delete a todo",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			index, err := parseIndex(args[0])
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			if err := todos.Delete(index); err != nil {
				fmt.Println("Error deleting task:", err)
				os.Exit(1)
			}
			if err := todos.Store(todoFile); err != nil {
				fmt.Println("Error saving todos:", err)
				os.Exit(1)
			}
			fmt.Println("Deleted task:", index)
		},
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all todos",
		Run: func(cmd *cobra.Command, args []string) {
			if err := todos.Load(todoFile); err != nil {
				fmt.Println("Error loading todos:", err)
				os.Exit(1)
			}
			todos.Print()
		},
	}

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(listCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseIndex(arg string) (int, error) {
	var index int
	_, err := fmt.Sscanf(arg, "%d", &index)
	if err != nil {
		return 0, fmt.Errorf("invalid index")
	}
	return index, nil
}
