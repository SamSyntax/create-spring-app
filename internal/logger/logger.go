package logger

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	BoldStyle = lipgloss.NewStyle().Bold(true)
	DoneStyle = lipgloss.NewStyle().MarginLeft(1).Foreground(lipgloss.Color("42"))
	InfoStyle = lipgloss.NewStyle().MarginLeft(1).Foreground(lipgloss.Color("33"))
)

func LogSuccess(msg string) {
	fmt.Printf("%s %s\n", DoneStyle.Render("✔ "), BoldStyle.Render(msg))
}

func LogInfo(msg string) {
	fmt.Printf("%s %s\n", InfoStyle.Render("🛈 "), msg)
}
