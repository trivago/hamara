package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

type exportCmd struct {
	host  string
	token string
	out   io.Writer
}

func newExportCmd(out io.Writer) *cobra.Command {
	ec := &exportCmd{out: out}

	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export datasources",
		Long:  `Retrieve datasources from existing Grafana and export it into a YAML provisioning file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return ec.run()
		},
	}

	cmd.Flags().StringVarP(&ec.host, "host", "H", "", "Grafana host")
	cmd.Flags().StringVarP(&ec.token, "token", "t", "", "API key with Admin rights from Grafana")
	cmd.MarkFlagRequired("host")
	cmd.MarkFlagRequired("token")
	return cmd
}

func (c *exportCmd) run() error {
	fmt.Fprintf(c.out, "Exporting from %s\n", c.host)
	return nil
}
