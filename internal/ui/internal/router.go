package internal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jaronline/blocksmith/internal/ui/internal/home"
	initScreen "github.com/jaronline/blocksmith/internal/ui/internal/init"
	"github.com/jaronline/blocksmith/internal/ui/internal/keymap"
	"github.com/jaronline/blocksmith/ui/router"
)

func getRouter(keyMap keymap.KeyMap) *router.Router {
	return router.NewRouter(
		"home",
		[]router.Route{
			router.NewRoute(
				"home",
				func(ctx router.Context) tea.Model {
					return home.New(keyMap, func() (tea.Model, tea.Cmd) {
						return ctx.GoTo("init")
					})
				},
			),
			router.NewRoute(
				"init",
				func(ctx router.Context) tea.Model {
					return initScreen.New(keyMap, func() (tea.Model, tea.Cmd) {
						return ctx.GoTo("home")
					})
				},
			),
		},
	)
}
