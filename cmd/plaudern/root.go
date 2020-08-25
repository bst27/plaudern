package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "plaudern",
		Short: "An app to manage comments for static websites",
		Long:  "Plaudern is an app to manage comments for static websites.",
	}
)

func execute() error {
	return rootCmd.Execute()
}
