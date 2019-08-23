package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
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

			actual := new(bytes.Buffer)
			exportCmd := newExportCmd(actual)
			exportCmd.SetArgs(tt.args)
			exportCmd.Execute()

			assert.Equal(tt.host, exportCmd.Flag("host").Value.String())
			assert.Equal(tt.token, exportCmd.Flag("token").Value.String())
		})
	}
}
