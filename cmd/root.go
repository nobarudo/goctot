/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var inputFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goctot [file]",
	Short: "This command displays a CSV file in a table format.",
	Long: `This command displays a CSV file in a table format.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	RunE: func(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
			return cmd.Help()
		}

		// ここから通常処理
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		defer f.Close()

	return run(f)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goctot.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

func run(f io.Reader) error {
	table := csv.NewReader(f)
	records, err := table.ReadAll()
	if err != nil {
		return err
	}

	if len(records) == 0 {
		return fmt.Errorf("CSV が空です")
	}

	header := records[0]
	rows := records[1:]

	renderTable(os.Stdout, header, rows)

	return nil
}

func renderTable(w *os.File, header []string, rows [][]string) {
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