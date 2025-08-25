package router

import tea "github.com/charmbracelet/bubbletea"

type Navigation interface {
	GoTo(name string) (tea.Model, tea.Cmd)
}
