package research

import (
	"github.com/elek/abced/abcfile"
	"github.com/elek/abced/parser"
	"github.com/olekukonko/tablewriter"
	"os"
)

func List(files []string) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Title", "V"})
	table.SetAutoWrapText(false)
	for _, path := range files {
		abc, err := abcfile.ReadFile(path)
		if err != nil {
			return err
		}

		for _, tune := range abc.Tunes {

			e := ""
			_, err := parser.ParseTune(*tune)
			if err != nil {
				e = err.Error()
			}
			table.Append([]string{
				tune.ID(),
				tune.Title(),
				e,
			})
		}
	}
	table.Render()
	return nil
}
