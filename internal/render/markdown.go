package render

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func Markdown(w *os.File, header []string,rows [][]string) {
	table := tablewriter.NewWriter(w)

	table.SetHeader(header)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(rows) // Add Bulk Data
	table.Render()
}