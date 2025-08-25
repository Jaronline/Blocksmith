package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	root "github.com/jaronline/blocksmith/internal/ui/internal"
	"github.com/jaronline/blocksmith/internal/ui/internal/keymap"
	zone "github.com/lrstanley/bubblezone"
)

func StartTUI() {
	zone.NewGlobal()
	if _, err := tea.NewProgram(root.NewScreen(keymap.DefaultKeyMap), tea.WithAltScreen(), tea.WithMouseCellMotion()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
