package render

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

func Table(w io.Writer, header []string, rows [][]string) {
	table := tablewriter.NewWriter(w)

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