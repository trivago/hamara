package services

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/trivago/grafana-datasource-to-yaml/pkg/grafana"
	"gopkg.in/yaml.v2"
)

// ExporterService interface to export datasources
type ExporterService interface {
	Export(io.Reader, io.Writer) error
}

// GrafanaExporter implementation to export datasources
type GrafanaExporter struct{}

// Export will retrieve the datasources from Grafana and convert it to YAML provisioning file
func (ge *GrafanaExporter) Export(reader io.Reader, writer io.Writer) error {
	var (
		yamlBytes []byte
		content   []byte
		ds        []*grafana.DataSource
		err       error
	)

	if content, err = ioutil.ReadAll(reader); err != nil {
		return err
	}

	if err = json.Unmarshal(content, &ds); err != nil {
		return err
	}

	dsProv := grafana.DataSourceProvisioning{ApiVersion: 1, Datasources: ds}
	if yamlBytes, err = yaml.Marshal(dsProv); err != nil {
		return err
	}

	writer.Write(yamlBytes)

	return nil
}

// NewGrafanaExporter initialize the GrafanaExporter struct
func NewGrafanaExporter() *GrafanaExporter {
	return &GrafanaExporter{}
}
