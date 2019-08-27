package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/trivago/grafana-datasource-to-yaml/pkg/grafana"
	"github.com/trivago/grafana-datasource-to-yaml/pkg/services"
)

type exportCmd struct {
	host     string
	key      string
	out      io.Writer
	exporter services.ExporterService
}

func newExportCmd(out io.Writer, exporter services.ExporterService) *cobra.Command {
	ec := &exportCmd{out: out, exporter: exporter}

	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export datasources",
		Long:  `Retrieve datasources from existing Grafana and export it into a YAML provisioning file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return ec.run()
		},
	}

	cmd.Flags().StringVarP(&ec.host, "host", "H", "", "Grafana host")
	cmd.Flags().StringVarP(&ec.key, "key", "k", "", "API key with Admin rights from Grafana")
	cmd.MarkFlagRequired("host")
	cmd.MarkFlagRequired("key")
	return cmd
}

func (c *exportCmd) run() error {
	client := grafana.NewRestClient(c.host, c.key)
	raw, err := client.GetAllDatasources()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return c.exporter.Export(raw, c.out)
}
