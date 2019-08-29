<p align="center">
  <img alt="Logo" src="https://drive.google.com/uc?export=view&id=1ezQnuq5VN1pjwx1mdTFAI6RX3ooXqsWY" height="150">
</p>

<p align="center">
  <a href="https://github.com/trivago/grafana-datasource-to-yaml/actions">
    <img alt="Build Status" src="https://github.com/trivago/grafana-datasource-to-yaml/workflows/Go/badge.svg" />
  </a>
  <a href="https://codecov.io/gh/trivago/grafana-datasource-to-yaml">
    <img alt="Coverage Status" src="https://codecov.io/gh/trivago/grafana-datasource-to-yaml/branch/master/graph/badge.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/trivago/grafana-datasource-to-yaml">
    <img alt="Report Card" src="https://goreportcard.com/badge/github.com/trivago/grafana-datasource-to-yaml" />
  </a>
</p>

`hamara` is a tool to export datasources from the existing Grafana DB into a YAML provisioning file by utilizing the Grafana REST API.

**Usage**
---

```
Retrieve datasources from existing Grafana and export it into a YAML provisioning file

Usage:
  grafana-datasource-to-yaml export [flags]

Flags:
  -h, --help          help for export
  -H, --host string   Grafana host
  -k, --key string    API key with Admin rights from Grafana
```

**Example**
---

```
export GRAFANA_API_KEY=<your API key here>
grafana-datasource-to-yaml export --host=localhost:3000 --key=$GRAFANA_API_KEY > datasources.yaml
cat datasources.yaml
```

**Installation Options**
---

1. Download the `grafana-datasource-to-yaml` binary from Releases tab.

2. Install with `go get` (Installed Go required)
    + `$ go get -u github.com/trivago/grafana-datasource-to-yaml`
    + `$ grafana-datasource-to-yaml`

**How to Contribute**
---

1. Clone repo and create a new branch: `$ git checkout https://github.com/trivago/grafana-datasource-to-yaml -b name_for_new_branch`.
2. Make changes and test
3. Submit Pull Request with comprehensive description of changes
