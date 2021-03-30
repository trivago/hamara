<p align="center">
  <img alt="Logo" src="https://drive.google.com/uc?export=view&id=1ezQnuq5VN1pjwx1mdTFAI6RX3ooXqsWY" height="150">
</p>

<p align="center">
  [![Build Status](https://github.com/trivago/hamara/workflows/Go/badge.svg)](https://github.com/trivago/hamara/actions)
  [![Codecov branch](https://img.shields.io/codecov/c/github/trivago/hamara/master?color=codecov&label=codecov&logo=codecov&logoColor=codecov)](https://codecov.io/gh/trivago/hamara)
  [![Report Card](https://goreportcard.com/badge/github.com/trivago/hamara?style=flat)(https://goreportcard.com/report/github.com/trivago/hamara)
</p>
<p align="center">
  [![Build Status](https://img.shields.io/docker/cloud/build/mrlioncub/hamara)](https://hub.docker.com/r/mrlioncub/hamara)
  [![Docker Automated build](https://img.shields.io/docker/cloud/automated/mrlioncub/hamara)](https://hub.docker.com/r/mrlioncub/hamara)
  [![Docker Image Size](https://img.shields.io/docker/image-size/mrlioncub/wdmrc/hamara)](https://hub.docker.com/r/mrlioncub/hamara)
</p>

`hamara` is a tool to export datasources from the existing Grafana DB into a YAML provisioning file by utilizing the Grafana REST API.

**Usage**
---

```
Retrieve datasources from existing Grafana and export it into a YAML provisioning file

Usage:
  hamara export [flags]

Flags:
  -h, --help          help for export
  -H, --host string   Grafana host
  -k, --key string    API key with Admin rights from Grafana
```

or docker run:

```bash
docker run --rm mrlioncub/hamara
```

**Example**
---

```bash
export GRAFANA_API_KEY=<your API key here>
hamara export --host=localhost:3000 --key=$GRAFANA_API_KEY > datasources.yaml
cat datasources.yaml
```

or docker run:

```bash
export GRAFANA_API_KEY=<your API key here>
docker run --rm mrlioncub/hamara export --host=localhost:3000 --key=$GRAFANA_API_KEY > datasources.yaml
cat datasources.yaml
```

**Installation Options**
---

1. Download the `hamara` binary from Releases tab.

2. Install with `go get` (Installed Go required)
    + `$ go get -u github.com/trivago/hamara`
    + `$ hamara`

**How to Contribute**
---

1. Clone repo and create a new branch: `$ git checkout https://github.com/trivago/hamara -b name_for_new_branch`.
2. Make changes and test
3. Submit Pull Request with comprehensive description of changes
