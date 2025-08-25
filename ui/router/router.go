package router

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Router struct {
	routes       []Route
	currentRoute string
}

func NewRouter(initialRoute string, routes []Route) *Router {
	return &Router{
		currentRoute: initialRoute,
		routes:       routes,
	}
}

func (r Router) GoTo(ctx Context, name string) tea.Model {
	var result *Route
	var err error

	if result, err = r.findRoute(name); err != nil {
		panic(err)
	}

	r.currentRoute = result.Name
	return r.BuildCurrentRoute(ctx)
}

func (r Router) BuildCurrentRoute(ctx Context) tea.Model {
	var route *Route
	var err error

	if route, err = r.findRoute(r.currentRoute); err != nil {
		panic(err)
	}

	return route.Build(ctx)
}

func (r Router) findRoute(name string) (*Route, error) {
	var result *Route = nil
	for _, route := range r.routes {
		if route.Name == name {
			result = &route
			break
		}
	}

	if result == nil {
		return nil, fmt.Errorf("no route found with name \"%s\"", name)
	}

	return result, nil
}
