package internal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jaronline/blocksmith/internal/ui/internal/keymap"
	"github.com/jaronline/blocksmith/ui/router"
	zone "github.com/lrstanley/bubblezone"
)

type RootScreenModel struct {
	router *router.Router
	model  tea.Model
}

func NewScreen(keyMap keymap.KeyMap) *RootScreenModel {
	screenRouter := getRouter(keyMap)

	model := &RootScreenModel{
		router: screenRouter,
	}
	model.model = screenRouter.BuildCurrentRoute(model)

	return model
}

func (m RootScreenModel) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Blocksmith"),
		m.model.Init(),
	)
}

func (m RootScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.model, cmd = m.model.Update(msg)
	return m, cmd
}

func (m RootScreenModel) View() string {
	return zone.Scan(m.model.View())
}

func (m RootScreenModel) GoTo(name string) (tea.Model, tea.Cmd) {
	model := m.router.GoTo(m, name)
	return model, model.Init()
}
