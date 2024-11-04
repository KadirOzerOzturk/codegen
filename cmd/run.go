package cmd

import (
	"fmt"

	"github.com/KadirOzerOzturk/deneme/helpers"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "create",
	Long:  `This func use for create new files`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		helpers.GenerateFiles(name)
	},
}

// runCmd represents the run command
var deleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "delete",
	Long:  `This func use for delete created files`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		fmt.Println("delete template : " + name)
	},
}
func init() {

	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(deleteCmd)
}
