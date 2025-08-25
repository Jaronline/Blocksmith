package button

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	Click key.Binding
}

var DefaultKeyMap = KeyMap{
	Click: key.NewBinding(key.WithKeys("enter")),
}

// Styles is a set of available style definitions for the Button bubble.
type Styles struct {
	DefaultStyle  lipgloss.Style
	DisabledStyle lipgloss.Style
	FocusStyle    lipgloss.Style
}

var defaultStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("255")).
	Background(lipgloss.Color("240")).
	Padding(1, 2).
	MarginTop(1)

var DefaultStyles = Styles{
	DefaultStyle: defaultStyle,
	DisabledStyle: defaultStyle.
		Background(lipgloss.Color("247")),
	FocusStyle: defaultStyle.
		Background(lipgloss.Color("237")).
		MarginRight(2).
		Underline(true),
}

type Msg struct {
	model tea.Model
	cmd   tea.Cmd
}

func NewMsg(model tea.Model, cmd tea.Cmd) Msg {
	return Msg{
		model: model,
		cmd:   cmd,
	}
}

func (m Msg) Model() tea.Model {
	return m.model
}

func (m Msg) Cmd() tea.Cmd {
	return m.cmd
}

type Model struct {
	Err error

	// General settings.
	Text     string
	Disabled bool
	Id       string

	// Styles These will be applied as style depending on the button state.
	Styles Styles

	// KeyMap encodes the keybindings recognized by the widget.
	KeyMap KeyMap

	// OnClick listener for when this button is clicked.
	OnClick func() tea.Msg

	// focus indicates whether user input focus should be on this button
	// component. When false, ignore keyboard input.
	focus bool
}

// New creates a new model with default settings.
func New() Model {

	return Model{
		Text:     "",
		Disabled: false,
		Id:       "",
		Styles:   DefaultStyles,
		KeyMap:   DefaultKeyMap,
		OnClick:  func() tea.Msg { return nil },
		focus:    false,
	}
}

// Focused returns the focus state on the model.
func (m *Model) Focused() bool {
	return m.focus
}

// Focus sets the focus state on the model. When the model is in focus it can
// receive keyboard input.
func (m *Model) Focus() {
	m.focus = true
}

// Blur removes the focus state on the model.  When the model is blurred it can
// not receive keyboard input.
func (m *Model) Blur() {
	m.focus = false
}

// Update is the Bubble Tea update loop.
func (m Model) Update(msg tea.Msg) tea.Msg {
	if !m.focus || m.Disabled {
		return nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Click):
			if m.OnClick != nil {
				return m.OnClick()
			}
		}
	}

	return nil
}

func (m Model) IsClick(msg tea.Msg) bool {
	if !m.focus || m.Disabled {
		return false
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return key.Matches(msg, m.KeyMap.Click) && m.OnClick != nil
	}
	return false
}

func (m Model) View() string {
	style := m.Styles.DefaultStyle
	if m.Disabled {
		style = m.Styles.DisabledStyle
	} else if m.focus {
		style = m.Styles.FocusStyle
	}
	return style.Render(m.Text)
}
