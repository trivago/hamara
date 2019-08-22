package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestExportCmd(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		fixture string
	}{
		{"no arguments", []string{}, "export-no-args.golden"},
		{"valid arguments", []string{"--host=http://localhost:3000", "--token=123"}, "datasources.golden.yaml"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			defer gock.Off() // Flush pending mocks after test execution

			mockResponse, err := ioutil.ReadFile("testdata/datasources.golden.json")
			if err != nil {
				t.Fatal(err)
			}

			var datasourcesMap []map[string]interface{}
			json.Unmarshal(mockResponse, &datasourcesMap)
			gock.New("http://localhost:3000").
				Get("/api/datasources").
				Reply(200).
				JSON(datasourcesMap)

			args := []string{"export"}
			args = append(args, tt.args...)

			r, w, _ := os.Pipe()
			os.Stdout = w
			os.Stderr = w

			rootCmd.SetArgs(args)
			rootCmd.Execute()

			w.Close()
			actual, _ := ioutil.ReadAll(r)

			expected, err := ioutil.ReadFile("testdata/" + tt.fixture)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(actual, expected)
		})
	}

}
