package home

import (
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jaronline/blocksmith/internal/lib"
	"github.com/jaronline/blocksmith/internal/ui/internal/keymap"
	"github.com/jaronline/blocksmith/internal/ui/internal/styles"
	"github.com/jaronline/blocksmith/ui/zone/button"
	"golang.org/x/term"
)

type Model struct {
	keys     keymap.KeyMap
	help     help.Model
	buttons  []*button.Model
	packName string
}

func New(keyMap keymap.KeyMap, onInit func() (tea.Model, tea.Cmd)) *Model {
	m := &Model{
		keys:     keyMap,
		help:     help.New(),
		buttons:  make([]*button.Model, 1),
		packName: "",
	}

	if pkg, err := lib.GetCurrentPackage(); err == nil && pkg != nil && pkg.Name != nil {
		m.packName = *pkg.Name
	}

	var b *button.Model
	for i := range m.buttons {
		b = button.New("init")
		b.Styles = styles.PrimaryButton
		b.KeyMap = button.KeyMap{
			Click: m.keys.Click,
		}

		switch i {
		case 0:
			b.Text = "Initialize Modpack"
			if m.packName != "" {
				b.Text = "Edit Modpack"
			}
			b.OnClick = func() tea.Msg {
				return button.NewMsg(onInit())
			}
			b.Focus()
		}

		m.buttons[i] = b
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	var cmds []tea.Cmd

	for _, btn := range m.buttons {
		switch btnMsg := btn.Update(msg).(type) {
		case button.Msg:
			cmds = append(cmds, btnMsg.Cmd())
			return btnMsg.Model(), tea.Batch(cmds...)
		}
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}
	docStyle := styles.DocStyle

	// Title
	{
		title := styles.TitleStyle.Render()
		doc.WriteString(title + "\n")
	}

	// Content
	{
		desc := "No modpack found"
		if m.packName != "" {
			desc = "Modpack: " + m.packName
		}
		doc.WriteString(styles.Description.Render(desc) + "\n")
	}

	// Dialog
	{
		for _, btn := range m.buttons {
			doc.WriteString(btn.View() + "\n\n")
		}
	}

	// Help
	{
		helpView := m.help.View(m.keys)
		doc.WriteString(helpView + "\n")
	}

	if width > 0 {
		docStyle = docStyle.MaxWidth(width)
	}

	return docStyle.Render(doc.String())
}
