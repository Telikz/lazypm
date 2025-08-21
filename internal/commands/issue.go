package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Manage issues",
	Long:  `The 'issue' command lets you manage issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from LazyPM CLI!")
	},
}

var createIssueCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new issue",
	Long:  `The 'create' command lets you create a new issue.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	issueCmd.AddCommand(createIssueCmd)

	createIssueCmd.Flags().StringP("title", "t", "", "Issue title")
}
