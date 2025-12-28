package render

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

func Markdown(w io.Writer, records [][]string) {
	table := tablewriter.NewWriter(w)

	header := records[0]
	rows := records[1:]
	
	table.SetHeader(header)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(rows) // Add Bulk Data
	table.Render()
}

func NoHeaderMarkdown(w io.Writer, records [][]string) {
	table := tablewriter.NewWriter(w)

	rows := records[0:]
	
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(rows) // Add Bulk Data
	table.Render()
}