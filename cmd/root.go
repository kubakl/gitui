package cmd

import (
	"fmt"

	git "github.com/libgit2/git2go/v33"
	"github.com/spf13/cobra"
)

func Execute() {
	root := &cobra.Command{
		Use:   "gitui",
		Short: "A command line user interface for managing git repositories",
		Run:   mergetool,
	}

	if err := root.Execute(); err != nil {
		panic(err)
	}
}

func mergetool(cmd *cobra.Command, args []string) {
	repo, err := git.OpenRepository("./.git")
	if err != nil {
		cmd.PrintErrln("There's no repository in the current repository")
	}

	index, err := repo.Index()
	fmt.Println(index.HasConflicts())
}
