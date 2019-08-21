package cmd

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestExportCmd(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		fixture string
	}{
		{"no arguments", []string{}, "export-no-args.golden"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual := new(bytes.Buffer)
			args := []string{"export"}
			args = append(args, tt.args...)

			rootCmd.SetOutput(actual)
			rootCmd.SetArgs(args)
			rootCmd.Execute()

			expected, err := ioutil.ReadFile("testdata/" + tt.fixture)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(actual.String(), string(expected)) {
				t.Fatalf("actual = %s, expected = %s", actual, expected)
			}
		})
	}

}
