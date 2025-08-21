package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Dashboard struct {
	width    int
	height   int
	quitting bool
}

func NewDashboard() Dashboard {
	return Dashboard{}
}

func (dash Dashboard) Init() tea.Cmd {
	return nil // No initial commands
}

func (dash Dashboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			dash.quitting = true
			return dash, tea.Quit
		}

	case tea.WindowSizeMsg:
		dash.width = msg.Width
		dash.height = msg.Height
		return dash, nil
	}
	return dash, nil
}

func (dash Dashboard) View() string {
	if dash.quitting {
		return "" // Clear screen on quit
	}

	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63")).PaddingBottom(1)

	content := lipgloss.JoinVertical(lipgloss.Left,
		titleStyle.Render("Welcome to LazyPM!"),
	)

	return lipgloss.Place(
		dash.width,
		dash.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}
