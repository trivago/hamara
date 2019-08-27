package grafana

import (
	"io"

	"gopkg.in/yaml.v2"
)

type DataSource struct {
	Id                int64                  `json:"id" yaml:"-"`
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
	SecureJsonFields  map[string]bool        `json:"secureJsonFields" yaml:"-"`
	SecureJsonData    map[string]string      `json:"secureJsonData" yaml:"secureJsonData,omitempty"`
	Editable          bool                   `json:"editable" yaml:"editable,omitempty"`
	ReadOnly          bool                   `json:"readOnly" yaml:"-"`
}

type DataSourceProvisioning struct {
	ApiVersion  int64         `yaml:"apiVersion"`
	Datasources []*DataSource `yaml:"datasources"`
}

func (dsp *DataSourceProvisioning) WriteTo(out io.Writer) error {
	var (
		yamlBytes []byte
		err       error
	)

	if yamlBytes, err = yaml.Marshal(dsp); err != nil {
		return err
	}

	if _, err = out.Write(yamlBytes); err != nil {
		return err
	}

	return nil
}
