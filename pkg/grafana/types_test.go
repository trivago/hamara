package grafana_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trivago/hamara/pkg/grafana"
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
	actual := new(bytes.Buffer)
	dsp.WriteTo(actual)
	expected, _ := ioutil.ReadFile("testdata/datasources.golden.yaml")
	assert.Equal(string(expected), actual.String())
}
