package render

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

func Table(w io.Writer, records [][]string) {
	table := tablewriter.NewWriter(w)

	header := records[0]
	rows := records[1:]
	
	table.SetHeader(header)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	// --- ASCII 固定（ズレ防止） ---
	table.SetBorder(true)
	table.SetCenterSeparator("|")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")

	table.AppendBulk(rows)
	table.Render()
}

func NoHeaderTable(w io.Writer, records [][]string) {
	table := tablewriter.NewWriter(w)

	rows := records[0:]
	
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	// --- ASCII 固定（ズレ防止） ---
	table.SetBorder(true)
	table.SetCenterSeparator("|")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")

	table.AppendBulk(rows)
	table.Render()
}