package grafana

// ClientConfig is the configuration for a supported Grafana client
type ClientConfig struct {
	Host string
	Key  string
}

// Client contains the supported operations over a Grafana instance
type Client interface {
	GetAllDatasources() ([]*DataSource, error)
}

// NewClientFn is a convenient type offered for instantiating new clients
type NewClientFn = func(ClientConfig) (Client, error)
