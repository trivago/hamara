package grafana_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trivago/grafana-datasource-to-yaml/pkg/grafana"
	"gopkg.in/h2non/gock.v1"
)

func TestGetAllDatasources(t *testing.T) {
	assert := assert.New(t)
	defer gock.Off()
	datasourcesMock, err := ioutil.ReadFile("testdata/datasources.golden.json")
	if err != nil {
		t.Fatal(err)
	}

	datasourceMock, err := ioutil.ReadFile("testdata/datasources-1.golden.json")
	if err != nil {
		t.Fatal(err)
	}

	var datasourcesMap []map[string]interface{}
	var datasourceMap map[string]interface{}
	json.Unmarshal(datasourcesMock, &datasourcesMap)
	json.Unmarshal(datasourceMock, &datasourceMap)

	gock.New("http://localhost:3000").
		Get("/api/datasources").
		Reply(200).
		JSON(datasourcesMap)

	gock.New("http://localhost:3000").
		Get("/api/datasources/1").
		Reply(200).
		JSON(datasourceMap)

	client, err := grafana.NewRestClientFn(grafana.ClientConfig{Host: "http://localhost:3000", Key: "123"})
	if err != nil {
		t.Fatal(err)
	}

	datasources, err := client.GetAllDatasources()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(1, len(datasources))
	actualDatasource := datasources[0]
	assert.Equal(1, len(actualDatasource.SecureJsonData))
	assert.Equal("$GDEV-INFLUXDB-TELEGRAF_PASSWORD", actualDatasource.SecureJsonData["password"])
}

func TestIncompleteHost(t *testing.T) {
	assert := assert.New(t)
	defer gock.Off()
	gock.New("http://localhost:3000").
		Get("/api/datasources").
		Reply(200).
		JSON("[]")

	client, err := grafana.NewRestClientFn(grafana.ClientConfig{Host: "localhost:3000", Key: "123"})
	if err != nil {
		t.Fatal(err)
	}

	datasources, err := client.GetAllDatasources()
	assert.Equal(nil, err)
	assert.Equal([]*grafana.DataSource{}, datasources)
}

func TestInvalidJSON(t *testing.T) {
	assert := assert.New(t)
	defer gock.Off()
	gock.New("http://localhost:3000").
		Get("/api/datasources").
		Reply(200).
		JSON("<invalid json format>")

	client, err := grafana.NewRestClientFn(grafana.ClientConfig{Host: "http://localhost:3000", Key: "123"})
	_, err = client.GetAllDatasources()
	assert.EqualError(err, "Failed to parse JSON data from http://localhost:3000/api/datasources")
}
