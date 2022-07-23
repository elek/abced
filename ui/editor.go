package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elek/abced/abcfile"
	abc "github.com/elek/abced/types"
)

type Editor struct {
	lines       []tea.Model
	currentLine int
	errorMsg    *ErrorMessage
}

func NewEditorFromTune(tune *abcfile.Tune) *Editor {
	e := &Editor{
		lines:    []tea.Model{},
		errorMsg: &ErrorMessage{},
	}

	for _, v := range tune.Items() {
		switch i := v.(type) {
		case abcfile.Score:
			e.lines = append(e.lines, NewLine(string(i), abc.NewBeat(1, 4)))
		case abcfile.Header:
			e.lines = append(e.lines, NewHeader(i.Key, i.Value))
		}
	}
	e.FocusFirstScore()
	return e
}
func NewEditor(id string) *Editor {
	e := &Editor{
		errorMsg: &ErrorMessage{},
	}

	e.lines = append(e.lines, NewHeader("X", id))
	e.lines = append(e.lines, NewHeader("T", ""))
	e.lines = append(e.lines, NewHeader("M", "4/4"))
	e.lines = append(e.lines, NewHeader("L", "1/4"))
	e.lines = append(e.lines, NewHeader("K", "C"))
	e.lines = append(e.lines, NewLine("", abc.NewBeat(1, 4)))
	e.FocusFirstScore()
	return e
}

func (e *Editor) Init() tea.Cmd {
	return nil
}

func (e *Editor) UpdateHeaderMode(current tea.Model, msg tea.Msg) (*Editor, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			current.(*Header).Text.Blur()
			e.FocusFirstScore()
			return e, nil
		}
	}
	var cmd tea.Cmd
	e.lines[e.currentLine], cmd = current.Update(msg)
	return e, cmd
}

func (e *Editor) UpdateEditorMode(current *Line, msg tea.Msg) (*Editor, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {
		case "p":
			return e, func() tea.Msg {
				err := Play(e)
				if err != nil {
					return ErrorMsg(err)
				}
				return nil
			}
		case "v":
			return e, func() tea.Msg {
				err := Show(e)
				if err != nil {
					return ErrorMsg(err)
				}
				return nil
			}

		case "ctrl+o":
			e.lines = append(e.lines, NewLine(current.validContent, abc.NewBeat(1, 4)))
			e.currentLine = len(e.lines) - 1
			e.setFocus()
			return e, nil

		case "enter":
			e.lines = append(e.lines, NewLine("", abc.NewBeat(1, 4)))
			e.currentLine++
			e.setFocus()
			return e, nil

		}
	}
	var cmd tea.Cmd
	e.lines[e.currentLine], cmd = current.Update(msg)
	return e, cmd

}

func (e *Editor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	e.errorMsg, _ = e.errorMsg.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if e.currentLine > 0 {
				e.currentLine--
			}
			e.setFocus()
			return e, nil

		case "down":
			e.currentLine = (e.currentLine + 1) % len(e.lines)
			e.setFocus()
			return e, nil
		}
	}
	model := e.lines[e.currentLine]
	switch i := model.(type) {
	case *Line:
		return e.UpdateEditorMode(i, msg)
	case *Header:
		return e.UpdateHeaderMode(i, msg)
	}
	return e, nil
}

func (e *Editor) setFocus() {
	for ix, i := range e.lines {
		switch l := i.(type) {
		case *Line:
			if ix == e.currentLine {
				l.Focus()
			} else {
				l.Blur()
			}
		case *Header:
			if ix == e.currentLine {
				l.Text.Focus()
			} else {
				l.Text.Blur()
			}
		}

	}
}

func (e *Editor) GetABC() string {
	s := ""
	for _, h := range e.lines {
		switch i := h.(type) {
		case *Line:
			s += i.content.Value() + "\n"
		case *Header:
			s += i.Key + ":" + i.Text.Value() + "\n"
		}
	}
	return s
}
func (e *Editor) View() string {
	s := ""
	for _, h := range e.lines {
		s += h.View() + "\n"
	}
	s += e.errorMsg.View()
	return s
}

func (e *Editor) FocusFirstScore() {
	for ix, l := range e.lines {
		if _, ok := l.(*Line); ok {
			e.currentLine = ix
			break
		}
	}
	e.setFocus()
}
