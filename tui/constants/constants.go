package constants

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

var CustomBorder = table.Border{
	Top:    "─",
	Left:   "│",
	Right:  "│",
	Bottom: "─",

	TopRight:    "╮",
	TopLeft:     "╭",
	BottomRight: "╯",
	BottomLeft:  "╰",

	TopJunction:    "╥",
	LeftJunction:   "├",
	RightJunction:  "┤",
	BottomJunction: "╨",
	InnerJunction:  "╫",

	InnerDivider: "║",
}

var DocStyle = lipgloss.NewStyle().Margin(1, 2)
var HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
var ErrStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd534b"))

type keymap struct {
	Create key.Binding
	Select key.Binding
	Edit   key.Binding
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
	Edit: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit"),
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
		key.WithKeys("k", "up", "w"),
		key.WithHelp("k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down", "s"),
		key.WithHelp("j", "down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
