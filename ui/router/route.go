package router

import tea "github.com/charmbracelet/bubbletea"

type Route struct {
	Name  string
	Build func(context Context) tea.Model
}

func NewRoute(name string, build func(context Context) tea.Model) Route {
	return Route{
		Name:  name,
		Build: build,
	}
}
