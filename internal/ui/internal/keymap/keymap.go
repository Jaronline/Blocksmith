package keymap

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Help   key.Binding
	Quit   key.Binding
	Click  key.Binding
	UpDown key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.UpDown, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Click},
		{k.Help, k.UpDown, k.Quit},
	}
}

var DefaultKeyMap = KeyMap{
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "esc"),
		key.WithHelp("ctrl+c", "quit"),
	),
	Click: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "click button"),
	),
	UpDown: key.NewBinding(
		key.WithKeys("up", "down", "enter"),
		key.WithHelp("↑/↓", "up/down"),
	),
}
