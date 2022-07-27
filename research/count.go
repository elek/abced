package research

import (
	"github.com/elek/abced/abcfile"
	"github.com/olekukonko/tablewriter"
	"github.com/zeebo/errs/v2"
	"os"
	"path/filepath"
	"strconv"
)

func Count(files []string) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"File", "#"})
	table.SetAutoWrapText(false)
	s := 0
	for _, path := range files {
		abc, err := abcfile.ReadFile(path)
		if err != nil {
			return errs.Errorf("Couldn't read %s: %++v", path, err)
		}

		table.Append([]string{
			filepath.Base(path),
			strconv.Itoa(len(abc.Tunes)),
		})
		s += len(abc.Tunes)

	}
	table.Append([]string{
		"SUM",
		strconv.Itoa(s),
	})

	table.Render()
	return nil
}
