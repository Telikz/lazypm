package commands

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"lazypm/internal/app"
)

var rootCmd = &cobra.Command{
	Use:   "lazypm",
	Short: "LazyPM is a distributed project management Terminal UI solution.",
	Long:  "LazyPM is a local-first, Git-powered project management tool with a terminal user interface.",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(app.NewDashboard(), tea.WithAltScreen())
		if err, _ := p.Run(); err != nil {
			log.Fatalf("Error running LazyPM: %v", err)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func GetRootCmd() *cobra.Command {
	return rootCmd
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
