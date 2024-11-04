package cmd

import (
	"github.com/KadirOzerOzturk/deneme/helpers"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [name]",
	Short: "Print a greeting message",
	Long:  `This command prints a greeting message with the provided name.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		helpers.GenerateFiles(name)
	},
}

func init() {

	rootCmd.AddCommand(runCmd)
}
