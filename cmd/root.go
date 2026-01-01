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
var isNoHeader bool
var outputfile string

func hasStdin() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (info.Mode() & os.ModeCharDevice) == 0
}

var rootCmd = &cobra.Command{
	Use:   "goctot [file]",
	Short: "This command displays a CSV file in a table format.",
	Long: `This command displays a CSV file in a table format.`,
	Args: cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		switch {
			// ファイル引数あり
			case len(args) == 1:
				f, err := os.Open(args[0])
				if err != nil {
					return err
				}
				defer f.Close()
				return run(f)

			// stdin が pipe
			case len(args) == 0 && hasStdin():
				return run(os.Stdin)

			// それ以外（引数なし・stdinなし）
			default:
				return cmd.Help()
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&isMarkdown, "markdown", "m", false, "format to Markdown.")
	rootCmd.Flags().BoolVarP(&isNoHeader, "no-header", "n", false, "no header option.")
	rootCmd.Flags().StringVarP(&outputfile, "output", "o", "","Specify the file to output from standard output.")
}

func run(f io.Reader) error {
	table := csv.NewReader(f)
	records, err := table.ReadAll()
	if err != nil {
		return err
	}

	if len(records) == 0 {
		return fmt.Errorf("CSV file is empty")
	}

	var out io.Writer = os.Stdout

	if outputfile != "" {
		f, err := os.Create(outputfile)
		if err != nil {
			return err
		}
		defer f.Close()
		out = f
	}

	if isMarkdown {
		if isNoHeader {
			render.NoHeaderMarkdown(out, records)
		} else {
			render.Markdown(out, records)
		}
	}else {
		if isNoHeader {
			render.NoHeaderTable(out, records)
		} else {
			render.Table(out, records)
		}
	}
	
	return nil
}