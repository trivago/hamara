package cmd_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/trivago/grafana-datasource-to-yaml/cmd"
	"github.com/trivago/grafana-datasource-to-yaml/pkg/grafana"
)

type grafanaRestClientMock struct {
	mock.Mock
}

func (m *grafanaRestClientMock) GetAllDatasources() ([]*grafana.DataSource, error) {
	args := m.Called()
	return args.Get(0).([]*grafana.DataSource), args.Error(1)
}

func newGrafanaRestClientMockFn(grafanaRestClient *grafanaRestClientMock) func(grafana.ClientConfig) (grafana.Client, error) {
	return func(grafana.ClientConfig) (grafana.Client, error) {
		return grafanaRestClient, nil
	}
}

func TestExportCmd(t *testing.T) {
	tests := []struct {
		name string
		args []string
		host string
		key  string
	}{
		{"no arguments", []string{}, "", ""},
		{"missing host", []string{"--key=123"}, "", "123"},
		{"missing key", []string{"--host=http://localhost:3000"}, "http://localhost:3000", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			exportCmd := cmd.NewExportCmd(ioutil.Discard, grafana.NewRestClientFn)
			exportCmd.SetArgs(tt.args)
			exportCmd.Execute()

			assert.Equal(tt.host, exportCmd.Flag("host").Value.String())
			assert.Equal(tt.key, exportCmd.Flag("key").Value.String())
		})
	}
}

func TestGetAllDatasources(t *testing.T) {
	grafanaRestClient := new(grafanaRestClientMock)
	grafanaRestClient.On("GetAllDatasources").Return([]*grafana.DataSource{}, nil)

	exportCmd := cmd.NewExportCmd(ioutil.Discard, newGrafanaRestClientMockFn(grafanaRestClient))
	exportCmd.SetArgs([]string{"--host=http://localhost:3000", "--key=123"})
	exportCmd.Execute()

	grafanaRestClient.AssertExpectations(t)
}
