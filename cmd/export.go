package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var host string
var token string

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export datasources",
	Long:  `Retrieve datasources from existing Grafana and export it into a YAML provisioning file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(host, token)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.PersistentFlags().StringVarP(&host, "host", "H", "", "Grafana host")
	exportCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "API key with Admin rights from Grafana")
	exportCmd.MarkPersistentFlagRequired("host")
	exportCmd.MarkPersistentFlagRequired("token")
}
