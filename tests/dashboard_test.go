package tests

import (
	"testing"

	"lazypm/internal/app"

	tea "github.com/charmbracelet/bubbletea"
)

func TestDashboard_Init(t *testing.T) {
	dash := app.NewDashboard()
	cmd := dash.Init()
	if cmd != nil {
		t.Errorf("Init() = %v, want nil", cmd)
	}
}

func TestDashboard_Update(t *testing.T) {
	dash := app.NewDashboard()

	_, cmd := dash.Update(tea.KeyMsg{Type: tea.KeyCtrlC})

	if cmd == nil {
		t.Errorf("Update(tea.KeyCtrlC) cmd = nil, want not nil")
	}

	_, cmd = dash.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	if cmd != nil {
		t.Errorf("Update(tea.WindowSizeMsg) cmd = %v, want nil", cmd)
	}
}

func TestDashboard_View(t *testing.T) {
	dash := app.NewDashboard()
	view := dash.View()
	if view == "" {
		t.Errorf("View() = %v, want not empty", view)
	}
}
