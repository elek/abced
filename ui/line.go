package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elek/abced/ast"
	"github.com/elek/abced/music"
	"github.com/elek/abced/parser"
	"strings"
)

type Line struct {
	content       textinput.Model
	validContent  string
	defaultLength music.Beat
	lastPitch     *music.Pitch
}

func NewLine(value string, defaultLength music.Beat) *Line {
	ti := textinput.New()
	ti.Prompt = ""
	ti.SetValue(value)
	return &Line{
		defaultLength: defaultLength,
		content:       ti,
	}
}

func (l *Line) Init() tea.Cmd {
	return textinput.Blink
}

func (l *Line) GuessOctave(orig string) string {
	if l.lastPitch != nil {

		//up means one octave up --> lower case
		up := music.NewPitchFromString(strings.ToLower(orig))
		low := music.NewPitchFromString(strings.ToUpper(orig))

		if l.lastPitch.Distance(up) < l.lastPitch.Distance(low) {
			orig = strings.ToLower(orig)
		} else {
			orig = strings.ToUpper(orig)
		}
	}
	return orig
}

func (l *Line) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		switch msg.String() {

		case "a", "b", "c", "d", "e", "f", "g":
			s := l.GuessOctave(msg.String())
			l.content, cmd = l.content.Update(tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune{rune(s[0])},
			})
			l.ReIndex()
			l.addBar()
			return l, cmd
		case "A", "B", "C", "D", "E", "F", "G":
			l.content, cmd = l.content.Update(msg)
			l.ReIndex()
			l.addBar()
			return l, cmd
		case "ctrl+a", "ctrl+b", "ctrl+c", "ctrl+d", "ctrl+e", "ctrl+f", "ctrl+g":
			s := strings.Split(msg.String(), "+")[1]
			l.content, cmd = l.content.Update(tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune{rune(s[0])},
			})
			l.ReIndex()
			l.addBar()
			return l, cmd
		case "left", "right":
			l.content, cmd = l.content.Update(msg)
			return l, cmd
		case "backspace", "delete":
			l.content, cmd = l.content.Update(msg)
			l.ReIndex()
			return l, cmd
		case "/", "1", "2", "3", "4", "5", "6", "7", "8", "9", "<", ">", ".", ",", " ", "z", "|", "_", "=", "^", "(", ")", ":":
			if msg.String() == " " && l.content.Cursor() > 0 && l.content.Value()[l.content.Cursor()-1] == ' ' {
				return l, nil
			}
			l.content, cmd = l.content.Update(msg)
			l.ReIndex()
			l.addBar()
			return l, cmd
		}
	}
	return l, nil
}

func (l *Line) View() string {
	return l.content.View()
}

func (l *Line) ReIndex() {
	for i := len(l.content.Value()); i >= 1; i-- {
		l.validContent = l.content.Value()[0:i]
		parsed, err := ast.BuildAST(l.validContent)
		if err != nil {
			continue
		}
		lastNote := LastNote{}
		ast.Walk(&lastNote, parsed)
		l.lastPitch = lastNote.LastPitch

		break
	}
}

func (l *Line) addBar() {
	if l.content.Value() != l.validContent {
		return
	}

	if l.content.Value() == "" {
		return
	}

	if l.content.Value()[len(l.content.Value())-1] != ' ' {
		return
	}

	if l.content.Cursor() != len(l.content.Value()) {
		return
	}

	parsed, _ := ast.BuildAST(l.validContent)

	// this contains the postfix of line since last bar
	nl := ast.Line{}
	for i := len(parsed.Nodes) - 1; i >= 0; i-- {
		if _, ok := parsed.Nodes[i].(ast.Bar); ok {
			break
		}
		nl.Nodes = append([]ast.Node{parsed.Nodes[i]}, nl.Nodes...)
	}

	lb, err := parser.ParseAST(&nl, music.NewBeat(1, 4), "C")
	if err != nil {
		fmt.Println(err)
	}

	if lb.Length() == music.NewBeat(4, 4).Simplify() {
		l.content.Focus()
		l.content, _ = l.content.Update(tea.KeyMsg{
			Type:  tea.KeyRunes,
			Runes: []rune{'|', ' '},
		})
	}
}

func (l *Line) Focus() {
	l.content.Focus()
}

func (l *Line) Blur() {
	l.content.Blur()
}

var _ tea.Model = &Line{}

type LastNote struct {
	LastPitch *music.Pitch
}

func (l *LastNote) Visit(node ast.Node) (w ast.Visitor) {
	switch n := node.(type) {
	case ast.Note:
		p := music.NewPitchFromString(n.Letter.Pitch)
		l.LastPitch = &p
	}
	return l
}
