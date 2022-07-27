package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elek/abced/abcfile"
	"sort"
	"strconv"
)

type Front struct {
	list     list.Model
	abc      *abcfile.AbcFile
	filename string
	editor   *Editor
}

var _ tea.Model = &Front{}

func NewFront(filename string, in *abcfile.AbcFile) *Front {
	f := &Front{
		abc:      in,
		filename: filename,
	}

	f.recreateList()
	return f
}

func (f *Front) recreateList() {
	k := make([]list.Item, 0)
	for _, s := range f.abc.Tunes {
		k = append(k, abcFile{
			ID:   s.ID(),
			Name: s.Title(),
		})
	}
	sort.Slice(k, func(i, j int) bool {
		id1, _ := strconv.Atoi(k[i].(abcFile).ID)
		id2, _ := strconv.Atoi(k[j].(abcFile).ID)
		return id1 < id2
	})
	f.list = list.New(k, list.NewDefaultDelegate(), 30, 20)
}

type Loaded abcfile.AbcFile

type LoadError error

func (f *Front) Init() tea.Cmd {
	return nil
}

type abcFile struct {
	ID   string
	Name string
}

func (a abcFile) FilterValue() string {
	return a.Name
}

func (a abcFile) Title() string {
	return a.ID
}

func (a abcFile) Description() string {
	return a.Name
}

func (f *Front) Update(msg tea.Msg) (model tea.Model, cmd tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if f.editor == nil {
				item := f.list.SelectedItem().(abcFile)
				s := f.abc.GetScore(item.ID)
				f.editor = NewEditorFromTune(s)
				return f, nil
			}
		case "ctrl+n":
			if f.editor == nil {
				maxID := 0
				for _, s := range f.abc.Tunes {
					i, err := strconv.Atoi(s.ID())
					if err == nil {
						if i > maxID {
							maxID = i
						}
					}
				}
				f.editor = NewEditor(strconv.Itoa(maxID + 1))
			}
		case "ctrl+x":
			return f, tea.Quit
		case "ctrl+i":
			if f.editor != nil {
				f.editor = nil
				return f, nil
			}
		case "q", "ctrl+q":
			if f.editor != nil {
				content := f.editor.GetABC()
				score := abcfile.NewTune(content)
				id := score.ID()
				existing := f.abc.GetScore(id)
				if existing != nil {
					existing.SetValue(content)
				} else {
					f.abc.AddScore(f.editor.GetABC())
				}
				f.editor = nil
				f.recreateList()
				err := abcfile.WriteFile(f.filename, f.abc)
				if err != nil {
					fmt.Println(err)
				}
				return f, nil
			}
			return f, tea.Quit
		}
	}
	model = f
	if f.editor != nil {
		var updated tea.Model
		updated, cmd = f.editor.Update(msg)
		f.editor = updated.(*Editor)
	} else {
		f.list, cmd = f.list.Update(msg)
	}
	return

}

func (f *Front) View() string {
	if f.editor != nil {
		return f.editor.View()
	}
	return f.list.View()
}
