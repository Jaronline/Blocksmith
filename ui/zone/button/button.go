package button

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jaronline/blocksmith/ui/button"
	zone "github.com/lrstanley/bubblezone"
)

type KeyMap = button.KeyMap
type Msg = button.Msg

var DefaultKeyMap = button.DefaultKeyMap
var NewMsg = button.NewMsg

type Styles struct {
	button.Styles
	ActiveStyle lipgloss.Style
}

var DefaultStyles = Styles{
	Styles: button.DefaultStyles,
	ActiveStyle: button.DefaultStyles.DefaultStyle.
		Background(lipgloss.Color("238")),
}

type Model struct {
	button.Model

	// Styles These will be applied as style depending on the button state.
	Styles Styles

	// OnMousePress listener for when a mouse is pressed while on this button.
	OnMousePress func(msg tea.MouseMsg) tea.Msg
	// OnMouseRelease listener for when a mouse is released while on this button.
	OnMouseRelease func(msg tea.MouseMsg) tea.Msg
	// OnMouseMotion listener for when a mouse is moved while on this button.
	OnMouseMotion func(msg tea.MouseMsg) tea.Msg

	// wasPressed indicates whether the button was pressed earlier
	wasPressed bool
}

func New(id string) *Model {
	m := WithButton(button.New())
	m.Id = id
	return m
}

func WithButton(b button.Model) *Model {
	return &Model{
		Model:  b,
		Styles: DefaultStyles,
		OnMousePress: func(msg tea.MouseMsg) tea.Msg {
			return nil
		},
		OnMouseRelease: func(msg tea.MouseMsg) tea.Msg {
			return nil
		},
		OnMouseMotion: func(msg tea.MouseMsg) tea.Msg {
			return nil
		},
		wasPressed: false,
	}
}

func (m *Model) Update(msg tea.Msg) tea.Msg {
	if m.Disabled {
		return nil
	}

	switch msg := msg.(type) {
	case tea.MouseMsg:
		if zone.Get(m.Id).InBounds(msg) {
			switch msg.Action {
			case tea.MouseActionPress:
				m.wasPressed = true
				if m.OnMousePress != nil {
					return m.OnMousePress(msg)
				}
			case tea.MouseActionRelease:
				if m.wasPressed {
					m.wasPressed = false
					if m.OnClick != nil {
						return m.OnClick()
					}
				}
				if m.OnMouseRelease != nil {
					return m.OnMouseRelease(msg)
				}
			case tea.MouseActionMotion:
				if m.OnMouseMotion != nil {
					return m.OnMouseMotion(msg)
				}
			}
		}
		if msg.Action == tea.MouseActionRelease && m.wasPressed {
			m.wasPressed = false
		}
	}

	return m.Model.Update(msg)
}

func (m Model) View() string {
	style := m.Styles.DefaultStyle
	if m.Disabled {
		style = m.Styles.DisabledStyle
	} else if m.Focused() && m.wasPressed {
		style = m.Styles.ActiveStyle.Inherit(m.Styles.FocusStyle)
	} else if m.Focused() {
		style = m.Styles.FocusStyle
	} else if m.wasPressed {
		style = m.Styles.ActiveStyle
	}
	return zone.Mark(m.Id, style.Render(m.Text))
}
