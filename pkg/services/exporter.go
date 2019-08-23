package services

import "fmt"

// ExporterService interface to export datasources
type ExporterService interface {
	Export(string, string) error
}

// GrafanaExporter implementation to export datasources
type GrafanaExporter struct{}

// Export will retrieve the datasources from Grafana and convert it to YAML provisioning file
func (grafana GrafanaExporter) Export(host string, token string) error {
	fmt.Printf("Exporting Grafana %s %s", host, token)
	return nil
}

// NewGrafanaExporter initialize the GrafanaExporter struct
func NewGrafanaExporter() *GrafanaExporter {
	return &GrafanaExporter{}
}
