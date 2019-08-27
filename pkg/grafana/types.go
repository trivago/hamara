package grafana

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