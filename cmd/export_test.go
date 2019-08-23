package cmd

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/trivago/grafana-datasource-to-yaml/pkg/services"
)

func TestExportCmd(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		host  string
		token string
	}{
		{"no arguments", []string{}, "", ""},
		{"missing host", []string{"--token=123"}, "", "123"},
		{"missing token", []string{"--host=http://localhost:3000"}, "http://localhost:3000", ""},
		{"valid arguments", []string{"--host=http://localhost:3000", "--token=123"}, "http://localhost:3000", "123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			exportCmd := newExportCmd(ioutil.Discard, services.NewGrafanaExporter())
			exportCmd.SetArgs(tt.args)
			exportCmd.Execute()

			assert.Equal(tt.host, exportCmd.Flag("host").Value.String())
			assert.Equal(tt.token, exportCmd.Flag("token").Value.String())
		})
	}
}

type grafanaExporterMock struct {
	mock.Mock
}

func (m *grafanaExporterMock) Export(host string, token string) error {
	args := m.Called(host, token)
	return args.Error(0)
}

func TestExporterFunction(t *testing.T) {
	exporterMock := new(grafanaExporterMock)
	exporterMock.On("Export", "http://localhost:3000", "123").Return(nil)

	exportCmd := newExportCmd(ioutil.Discard, exporterMock)
	exportCmd.SetArgs([]string{"--host=http://localhost:3000", "--token=123"})
	exportCmd.Execute()

	exporterMock.AssertExpectations(t)
}
