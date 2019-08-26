package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

// ExporterService interface to export datasources
type ExporterService interface {
	Export(string, string) error
}

// GrafanaExporter implementation to export datasources
type GrafanaExporter struct{}

type DataSource struct {
	OrgId             int64                  `json:"orgId" yaml:"orgId,omitempty"`
	Version           int                    `json:"version" yaml:"version,omitempty"`
	Name              string                 `json:"name" yaml:"name"`
	Type              string                 `json:"type" yaml:"type"`
	Access            string                 `json:"access" yaml:"access"`
	Url               string                 `json:"url" yaml:"url,omitempty"`
	Password          string                 `json:"password" yaml:"password,omitempty"`
	User              string                 `json:"user" yaml:"user,omitempty"`
	Database          string                 `json:"database" yaml:"database,omitempty"`
	BasicAuth         bool                   `json:"basicAuth" yaml:"basicAuth,omitempty"`
	BasicAuthUser     string                 `json:"basicAuthUser" yaml:"basicAuthUser,omitempty"`
	BasicAuthPassword string                 `json:"basicAuthPassword" yaml:"basicAuthPassword,omitempty"`
	WithCredentials   bool                   `json:"withCredentials" yaml:"withCredentials,omitempty"`
	IsDefault         bool                   `json:"isDefault" yaml:"isDefault,omitempty"`
	JsonData          map[string]interface{} `json:"jsonData" yaml:"jsonData,omitempty"`
	SecureJsonData    map[string]string      `json:"secureJsonData" yaml:"secureJsonData,omitempty"`
	Editable          bool                   `json:"editable" yaml:"editable,omitempty"`
}

type DataSourceProvisioning struct {
	ApiVersion int64

	Datasources []*DataSource `yaml:"datasources"`
}

// Export will retrieve the datasources from Grafana and convert it to YAML provisioning file
func (grafana *GrafanaExporter) Export(host string, token string) error {

	apiPartURL := "/api/datasources"
	url1 := fmt.Sprintf("%s%s", host, apiPartURL)

	req, err := http.NewRequest("GET", url1, nil)
	key := fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Failed with status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: proceed with converion only if authorization was OK

	var yamlBytes []byte
	var ds []*DataSource

	// deserialize the content of JSON file (the bytes) into a struct
	if err = json.Unmarshal(body, &ds); err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprintf("%s\n", err))
	}

	dsProv := DataSourceProvisioning{ApiVersion: 1, Datasources: ds}

	// serialize the struct
	if yamlBytes, err = yaml.Marshal(dsProv); err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprintf("%s\n", err))
	}

	fmt.Println(string(yamlBytes))

	// // write to file
	// if _, err = file.Write(yamlBytes); err != nil {
	// 	fmt.Fprint(os.Stderr, fmt.Sprintf("%s\n", err))
	// }

	return nil
}

// NewGrafanaExporter initialize the GrafanaExporter struct
func NewGrafanaExporter() *GrafanaExporter {
	return &GrafanaExporter{}
}
