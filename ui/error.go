package ui

import tea "github.com/charmbracelet/bubbletea"

type ErrorMessage struct {
	err error
}

type ErrorMsg error

func (e *ErrorMessage) Init() tea.Cmd {
	return nil
}

func (e *ErrorMessage) Update(msg tea.Msg) (*ErrorMessage, tea.Cmd) {
	if err, ok := msg.(ErrorMsg); ok {
		e.err = err
	}
	return e, nil
}

func (e *ErrorMessage) View() string {
	if e.err != nil {
		return e.err.Error()
	}
	return ""
}
