package cmd

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"
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

			actual := new(bytes.Buffer)
			rootCmd.SetOutput(actual)
			rootCmd.SetArgs(tt.args)
			Execute()

			expected, err := ioutil.ReadFile("testdata/" + tt.fixture)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(actual.String(), string(expected)) {
				t.Logf("actual = %s, expected = %s", actual, expected)
			}
		})
	}

}
