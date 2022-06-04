package constants

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

var DocStyle = lipgloss.NewStyle().Margin(1, 2)

var HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

var ErrStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd534b")).Render

type keymap struct {
	Create key.Binding
	Select key.Binding
	Rename key.Binding
	Delete key.Binding
	Back   key.Binding
	Up     key.Binding
	Down   key.Binding
	Quit   key.Binding
}

// Keymap reusable key mappings shared across models
var Keymap = keymap{
	Create: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "create"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("enter", "select"),
	),
	Rename: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "rename"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Up: key.NewBinding(
		key.WithKeys("up", "k", "w"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j", "s"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
