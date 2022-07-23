package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/elek/abced/abcfile"
	"github.com/elek/abced/research"
	"github.com/elek/abced/ui"
	"github.com/spf13/cobra"
	"github.com/zeebo/errs/v2"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	c := cobra.Command{
		Use: "abced <file>",
	}

	{
		c.AddCommand(&cobra.Command{
			Use: "list",
			RunE: func(cmd *cobra.Command, args []string) error {
				files, err := fileList(args)
				if err != nil {
					return errs.Wrap(err)
				}
				research.List(files)
				return nil
			},
		})
	}

	{
		c.AddCommand(&cobra.Command{
			Use: "edit",
			RunE: func(cmd *cobra.Command, args []string) error {
				f, err := abcfile.ReadFile(args[0])
				if err != nil {
					panic(err)
				}

				p := tea.NewProgram(ui.NewFront(args[0], &f))
				return p.Start()
			},
		})
	}

	err := c.Execute()
	if err != nil {
		log.Fatalf("%++v", err)
	}
}

func fileList(files []string) ([]string, error) {
	if len(files) == 0 {
		files, err := os.ReadDir(".")
		if err != nil {
			return []string{}, err
		}
		abcFiles := make([]string, 0)
		for _, f := range files {
			if !f.IsDir() && strings.HasSuffix(f.Name(), "abc") {
				abcFiles = append(abcFiles, f.Name())
			}
		}
		return abcFiles, nil
	}
	return files, nil
}
