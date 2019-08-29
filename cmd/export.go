package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/trivago/hamara/pkg/grafana"
)

type exportCmd struct {
	host     string
	key      string
	out      io.Writer
	clientFn grafana.NewClientFn
}

func NewExportCmd(out io.Writer, fn grafana.NewClientFn) *cobra.Command {
	ec := &exportCmd{out: out, clientFn: fn}

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
	client, err := c.clientFn(grafana.ClientConfig{Host: c.host, Key: c.key})
	datasources, err := client.GetAllDatasources()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dsProv := grafana.DataSourceProvisioning{ApiVersion: 1, Datasources: datasources}
	return dsProv.WriteTo(c.out)
}
