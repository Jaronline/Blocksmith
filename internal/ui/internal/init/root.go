package init

import (
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jaronline/blocksmith/internal/lib"
	"github.com/jaronline/blocksmith/internal/ui/internal/keymap"
	"github.com/jaronline/blocksmith/internal/ui/internal/styles"
	"github.com/jaronline/blocksmith/ui/zone/button"
	"golang.org/x/term"
)

type Model struct {
	keys          keymap.KeyMap
	help          help.Model
	inputs        []textinput.Model
	confirmButton *button.Model
	focusIndex    int
	pkg           *lib.Package
}

func New(keyMap keymap.KeyMap, onConfirm func() (tea.Model, tea.Cmd)) *Model {
	m := &Model{
		keys:          keyMap,
		help:          help.New(),
		inputs:        make([]textinput.Model, 2),
		confirmButton: button.New("confirm"),
		focusIndex:    0,
	}

	m.pkg, _ = lib.GetDefaultPackage()

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = styles.CursorStyle
		t.CharLimit = 32
		t.Width = 32

		switch i {
		case 0:
			t.Placeholder = "Name"
			t.Focus()
			t.PromptStyle = styles.FocusedStyle
			t.TextStyle = styles.FocusedStyle
			t.SetValue(m.pkg.Name)
		case 1:
			t.Placeholder = "Version"
			t.SetValue(m.pkg.Version)
		}

		m.inputs[i] = t
	}

	m.confirmButton.Text = "Initialize"
	m.confirmButton.Styles = styles.SuccessButton
	m.confirmButton.KeyMap = button.KeyMap{
		Click: m.keys.Click,
	}
	m.confirmButton.OnClick = func() tea.Msg {
		m.pkg.Write()
		return button.NewMsg(onConfirm())
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.UpDown):
			s := msg.String()

			if m.confirmButton.IsClick(msg) {
				break
			}

			if s == "up" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range cmds {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = styles.FocusedStyle
					m.inputs[i].TextStyle = styles.FocusedStyle
					continue
				}
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = styles.NoStyle
				m.inputs[i].TextStyle = styles.NoStyle
			}

			if m.focusIndex == len(m.inputs) {
				m.confirmButton.Focus()
			} else {
				m.confirmButton.Blur()
			}

			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)
	m.updatePackageData()

	switch btnMsg := m.confirmButton.Update(msg).(type) {
	case button.Msg:
		return btnMsg.Model(), tea.Batch(cmd, btnMsg.Cmd())
	}

	return m, cmd
}

func (m *Model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *Model) updatePackageData() {
	m.pkg.Name = m.inputs[0].Value()
	m.pkg.Version = m.inputs[1].Value()
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

	// Dialog
	{
		for i := range m.inputs {
			doc.WriteString(m.inputs[i].View())
			if i < len(m.inputs)-1 {
				doc.WriteRune('\n')
			}
		}

		doc.WriteString("\n\n" + m.confirmButton.View() + "\n\n")
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
