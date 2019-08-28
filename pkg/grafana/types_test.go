package grafana_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trivago/grafana-datasource-to-yaml/pkg/grafana"
)

func TestWriteTo(t *testing.T) {
	assert := assert.New(t)
	datasources := []*grafana.DataSource{
		&grafana.DataSource{
			Name:           "test-db",
			Type:           "influxdb",
			Access:         "proxy",
			JsonData:       map[string]interface{}{"key": "value"},
			SecureJsonData: map[string]string{"password": "test"},
		},
	}
	dsp := &grafana.DataSourceProvisioning{ApiVersion: 1, Datasources: datasources}
	output := new(bytes.Buffer)
	dsp.WriteTo(output)
	expected, _ := ioutil.ReadFile("testdata/raw-datasources.yaml")
	assert.Equal(output.String(), string(expected))
}
