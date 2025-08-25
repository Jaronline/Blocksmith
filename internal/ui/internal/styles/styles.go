package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/jaronline/blocksmith/ui/button"
	zb "github.com/jaronline/blocksmith/ui/zone/button"
)

var (
	NoStyle = lipgloss.NewStyle()

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#4CAF50")).
			SetString("Blocksmith").
			MarginBottom(1)

	Description = lipgloss.NewStyle().
			Foreground(lipgloss.Color("255")).
			MarginBottom(1)

	defaultPrimaryButton = lipgloss.NewStyle().
				Foreground(lipgloss.Color("255")).
				Background(lipgloss.Color("#FF6F00")).
				Padding(0, 2)

	defaultSuccessButton = defaultPrimaryButton.
				Background(lipgloss.Color("#4CAF50"))

	PrimaryButton = zb.Styles{
		Styles: button.Styles{
			DefaultStyle: defaultPrimaryButton,
			DisabledStyle: defaultPrimaryButton.
				Foreground(lipgloss.Color("#9E9E9E")).
				Background(lipgloss.Color("#4B4B4B")),
			FocusStyle: defaultPrimaryButton.Underline(true),
		},
		ActiveStyle: defaultPrimaryButton.
			Background(lipgloss.Color("#B35F00")),
	}

	SuccessButton = zb.Styles{
		Styles: button.Styles{
			DefaultStyle: defaultSuccessButton,
			DisabledStyle: defaultSuccessButton.
				Foreground(lipgloss.Color("#4CAF50")).
				Background(lipgloss.Color("#9E9E9E")),
			FocusStyle: defaultSuccessButton.Underline(true),
		},
		ActiveStyle: defaultSuccessButton.
			Background(lipgloss.Color("#387B38")),
	}

	FocusedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4CAF50"))

	CursorStyle = FocusedStyle

	DocStyle = lipgloss.NewStyle().
			Padding(1, 2, 1, 2)
)
