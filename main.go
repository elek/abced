package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elek/abced/abcfile"
	"github.com/elek/abced/ui"
	"os"
)

func main() {
	filename := os.Args[1]

	f, err := abcfile.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(ui.NewFront(filename, &f))
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
