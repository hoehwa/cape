/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package principle

import (
	"github.com/jipilmuk/cape/utils"
	"github.com/spf13/cobra"
)

// managementCmd represents the management command.
var managementCmd = &cobra.Command{
	Use:   "management",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintContent("/principle/management")
	},
}

func init() {
	principleCmd.AddCommand(managementCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// managementCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// managementCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
