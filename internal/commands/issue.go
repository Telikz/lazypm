package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "issue",
	Short: "Manage issues",
	Long:  `The 'hello' command prints a simple greeting message to the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from LazyPM CLI!")
	},
}

func init() {}
