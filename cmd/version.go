package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display version",
	Run: func(cmd *cobra.Command, args []string) {
		// 最新のバージョンを表示する
		fmt.Fprintln(cmd.OutOrStdout(),"goctot version 0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
