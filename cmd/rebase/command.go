package rebase

import "github.com/spf13/cobra"

func Execute() *cobra.Command {
	return &cobra.Command{
		Use:   "rebase [arguments]",
		Short: "makes the rebase with specified branch. If there are any conficts it will open up the merge tool",
		Run:   rebase,
	}
}

func rebase(cmd *cobra.Command, args []string) {
}
