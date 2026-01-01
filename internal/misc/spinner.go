package misc

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type SpinnerModel struct {
	Spinner  spinner.Model
	ErrChan  chan error
	Quitting bool
	Err      error
}

func (m SpinnerModel) Init() tea.Cmd {
	return m.Spinner.Tick
}

func (m SpinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	select {
	case err := <-m.ErrChan:
		m.Err = err
		return m, tea.Quit
	default:
	}

	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	default:
		return m, nil
	}
}

func (m SpinnerModel) View() string {
	if m.Err != nil {
		return ""
	}
	str := fmt.Sprintf("\n  %s Generating Spring Boot project...", m.Spinner.View())
	return str
}
