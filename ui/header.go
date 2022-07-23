package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Header struct {
	Key      string
	Text     textinput.Model
	KeyStyle lipgloss.Style
}

func NewHeader(key string, value string) *Header {
	ti := textinput.New()
	ti.SetValue(value)
	ti.Blur()
	ti.Prompt = ":"
	return &Header{
		Key:      key,
		Text:     ti,
		KeyStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")),
	}
}
func (h *Header) Init() tea.Cmd {
	return nil
}

func (h *Header) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	h.Text, cmd = h.Text.Update(msg)
	return h, cmd
}

func (h *Header) View() string {
	return h.KeyStyle.Render(h.Key) + h.Text.View()
}

var _ tea.Model = &Header{}
