package cmd

import (
	"encoding/csv"
	"fmt"
	"goctot/internal/finder"
	"goctot/internal/render"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var isMarkdown bool
var isNoHeader bool
var isDir bool
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
	Args: func(cmd *cobra.Command, args []string) error {
		if isDir && len(args) != 1 {
			return fmt.Errorf("-d requires exactly one argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		switch {
			// ファイル引数あり
			case len(args) == 1:
				if isDir {
					return runDir(args)
				}
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
	rootCmd.Flags().BoolVarP(&isDir, "directory", "d", false, "display all CSV files in a directory")
	rootCmd.Flags().StringVarP(&outputfile, "output", "o", "","Specify the file to output from standard output.")
}

func run(f io.Reader) error {

	var out io.Writer = os.Stdout

	if outputfile != "" {
		f, err := os.Create(outputfile)
		if err != nil {
			return err
		}
		defer f.Close()
		out = f
	}

	table := csv.NewReader(f)
	records, err := table.ReadAll()
	if err != nil {
		return err
	}

	if len(records) == 0 {
		return fmt.Errorf("CSV file is empty")
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

func runDir(args []string) error {
	var out io.Writer = os.Stdout

	if outputfile != "" {
		f, err := os.Create(outputfile)
		if err != nil {
			return err
		}
		defer f.Close()
		out = f
	}

	csvFiles, err := finder.FindCSVFiles(args[0])
	if err != nil {
		return err
	}

	if len(csvFiles) == 0 {
		return fmt.Errorf("no csv files found in %s", args[0])
	}
	for _, path := range csvFiles {
		f, err := os.Open(path)
		if err != nil {
			return err
		}

		r := csv.NewReader(f)
		records, err := r.ReadAll()
		f.Close()

		if err != nil {
			return err
		}

		fmt.Fprintln(out, path)

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
		fmt.Fprintln(out)
	}
	return nil
}