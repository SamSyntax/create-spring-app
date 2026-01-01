package core

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func GetCustomTheme() *huh.Theme {
	t := huh.ThemeBase()

	primaryOrange := lipgloss.Color("208")
	subtleOrange := lipgloss.Color("166")
	cream := lipgloss.Color("255")
	darkGray := lipgloss.Color("240")

	t.Focused.Base = t.Focused.Base.
		BorderForeground(primaryOrange).
		PaddingLeft(1)

	t.Focused.Title = t.Focused.Title.
		Foreground(primaryOrange).
		Bold(true)

	t.Focused.Description = t.Focused.Description.
		Foreground(darkGray)

	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.
		Foreground(primaryOrange)

	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.
		Foreground(primaryOrange)

	t.FieldSeparator = t.FieldSeparator.
		Foreground(subtleOrange)

	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.
		Foreground(primaryOrange)

	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.
		Foreground(primaryOrange).
		Bold(true)

	t.Focused.SelectedOption = t.Focused.SelectedOption.
		Foreground(primaryOrange).
		Bold(true)

	t.Focused.SelectedPrefix = t.Focused.SelectedPrefix.
		Foreground(primaryOrange).
		SetString("🗸 ")

	t.Focused.UnselectedPrefix = t.Focused.UnselectedPrefix.
		Foreground(darkGray).
		SetString("○ ")

	t.Focused.FocusedButton = t.Focused.FocusedButton.
		Background(primaryOrange).
		Foreground(cream).
		Bold(true)

	t.Focused.ErrorMessage = t.Focused.ErrorMessage.
		Foreground(lipgloss.Color("196")).
		Bold(true)

	return t
}
