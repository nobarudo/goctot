package cmd

import (
	"encoding/csv"
	"fmt"
	"goctot/internal/render"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var isMarkdown bool

var rootCmd = &cobra.Command{
	Use:   "goctot [file]",
	Short: "This command displays a CSV file in a table format.",
	Long: `This command displays a CSV file in a table format.`,
	RunE: func(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
			return cmd.Help()
		}

		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		defer f.Close()

	return run(f)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&isMarkdown, "markdown", "m", false, "Change format to Markdown.")
}

func run(f io.Reader) error {
	table := csv.NewReader(f)
	records, err := table.ReadAll()
	if err != nil {
		return err
	}

	if len(records) == 0 {
		return fmt.Errorf("missing file argument")
	}

	header := records[0]
	rows := records[1:]
	if isMarkdown {
		render.Markdown(os.Stdout, header, rows)
	}else {
		render.Table(os.Stdout, header, rows)
	}
	
	return nil
}