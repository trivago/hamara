package cmd

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		{"valid arguments", []string{"--host=http://localhost:3000", "--key=123"}, "http://localhost:3000", "123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			exportCmd := newExportCmd(ioutil.Discard)
			exportCmd.SetArgs(tt.args)
			exportCmd.Execute()

			assert.Equal(tt.host, exportCmd.Flag("host").Value.String())
			assert.Equal(tt.key, exportCmd.Flag("key").Value.String())
		})
	}
}
