package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	abc "github.com/elek/abced/types"
	"io/ioutil"
	"os/exec"
	"strings"
)

type Editor struct {
	headers        []*Header
	lines          []*Line
	currentLine    int
	selectedHeader string
}

func NewEditorFromData(header map[string]string, lines []string) *Editor {
	e := &Editor{
		headers: []*Header{},
	}

	for k, v := range header {
		e.headers = append(e.headers, NewHeader(k, v))
	}

	for _, l := range lines {
		e.lines = append(e.lines, NewLine(l, abc.NewBeat(1, 4)))
	}
	if len(e.lines) > 0 {
		e.lines[0].Focus()
	}
	return e
}
func NewEditor(id string) *Editor {
	e := &Editor{
		headers: []*Header{
			NewHeader("X", id),
			NewHeader("T", ""),
			NewHeader("M", "4/4"),
			NewHeader("L", "1/4"),
			NewHeader("K", "C"),
		},
	}

	e.lines = append(e.lines, NewLine("", abc.NewBeat(1, 4)))
	e.lines[e.currentLine].Focus()
	return e
}

func (e *Editor) Init() tea.Cmd {
	return nil
}

func (e *Editor) UpdateHeaderMode(msg tea.Msg) (*Editor, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			for _, h := range e.headers {
				if h.Key == e.selectedHeader {
					h.Text.Blur()
				}
			}
			e.selectedHeader = ""
			return e, nil
		}
	}

	for _, h := range e.headers {
		if h.Key == e.selectedHeader {
			h.Update(msg)
		}
	}
	return e, nil
}

func (e *Editor) UpdateEditorMode(msg tea.Msg) (*Editor, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {
		case "p":
			go func() {
				err := e.Play()
				if err != nil {
					fmt.Println(err)
				}
			}()
			return e, nil
		case "ctrl+o":
			e.lines = append(e.lines, NewLine(e.lines[e.currentLine].validContent, abc.NewBeat(1, 4)))
			e.currentLine = len(e.lines) - 1
			e.setFocus()
			return e, nil
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

		case "enter":

			e.lines = append(e.lines, NewLine("", abc.NewBeat(1, 4)))
			e.currentLine++
			e.setFocus()
			return e, nil

		}

		// edit one of the headers
		for _, h := range e.headers {
			if h.Key == strings.ToUpper(msg.String()) {
				e.selectedHeader = strings.ToUpper(msg.String())
				h.Text.Focus()
			} else {
				h.Text.Blur()
			}
		}

	}

	// forward update to the right line editor
	var cmds []tea.Cmd
	for i, _ := range e.lines {
		if i == e.currentLine {
			l, c := e.lines[i].Update(msg)
			e.lines[i] = l.(*Line)
			cmds = append(cmds, c)
		}
	}
	return e, tea.Batch(cmds...)
}

func (e *Editor) Update(msg tea.Msg) (*Editor, tea.Cmd) {
	if e.selectedHeader == "" {
		return e.UpdateEditorMode(msg)
	} else {
		return e.UpdateHeaderMode(msg)
	}
}

func (e *Editor) setFocus() {
	for ix, l := range e.lines {
		if ix == e.currentLine {
			l.Focus()
		} else {
			l.Blur()
		}
	}
}

func (e *Editor) Play() error {
	c := e.GetABC()
	firstLine := strings.Split(c, "\n")[0]
	id := strings.TrimSpace(strings.Split(firstLine, ":")[1])
	err := ioutil.WriteFile("/tmp/a.abc", []byte(c), 0644)
	if err != nil {
		return err
	}

	command := exec.Command("abc2midi", "/tmp/a.abc", id, "-o", "/tmp/out.midi")
	_, err = command.CombinedOutput()
	if err != nil {
		return err
	}

	command = exec.Command("timidity", "/tmp/out.midi")
	_, err = command.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func (e *Editor) GetABC() string {
	s := ""
	for _, h := range e.headers {
		s += h.Key + ":" + h.Text.Value() + "\n"
	}
	for _, l := range e.lines {
		s += l.content.Value() + "\n"
	}
	return s
}
func (e *Editor) View() string {
	s := ""
	for _, h := range e.headers {
		s += h.View()
	}
	for i, _ := range e.lines {
		s += e.lines[i].View() + "\n"
	}
	return s
}
