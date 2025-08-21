package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "issue",
	Short: "Manage issues",
	Long:  `The 'issue' command lets you manage issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from LazyPM CLI!")
	},
}

func init() {}
