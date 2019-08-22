package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		fixture string
	}{
		{"no arguments", []string{}, "no-args.golden"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := new(bytes.Buffer)
			rootCmd.SetOutput(actual)
			rootCmd.SetArgs(tt.args)
			rootCmd.Execute()

			expected, err := ioutil.ReadFile("testdata/" + tt.fixture)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(actual.String(), string(expected))
		})
	}

}
