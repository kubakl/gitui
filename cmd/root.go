package cmd

import "github.com/spf13/cobra"

func Execute() {
	root := &cobra.Command{
		Use:   "gitui <command> [arguments]",
		Short: "A command line user interface for managing git repositories",
	}

	root.AddCommand()

	if err := root.Execute(); err != nil {
		panic(err)
	}
}
