package cmd

import (
    "os"
    "github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
    Use:   "codegen",
    Short: "useful application for creating folders and classes",
    Long:  `This app creates service, entities, controller, route file while user creates a new folder`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

func init() {

    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.codegen.yaml)")
}
