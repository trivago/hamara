package services

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/trivago/grafana-datasource-to-yaml/pkg/grafana"
	"gopkg.in/yaml.v2"
)

type ConverterService interface {
	Convert(io.Reader, io.Writer) error
}

type GrafanaConverter struct{}

func (grafanaConverter *GrafanaConverter) Convert(reader io.Reader, writer io.Writer) error {
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

func NewGrafanaConverter() *GrafanaConverter {
	return &GrafanaConverter{}
}
