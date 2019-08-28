package grafana

type ClientConfig struct {
	Host string
	Key  string
}

type Client interface {
	GetAllDatasources() ([]*DataSource, error)
}

type NewClientFn = func(ClientConfig) (Client, error)
