package grafana

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/gosimple/slug"
	"github.com/iancoleman/strcase"
)

type RestClient struct {
	baseURL string
	key     string
	client  *http.Client
}

func NewRestClientFn(config ClientConfig) (Client, error) {
	baseURL, _ := url.Parse(config.Host)
	if baseURL.Host == "" {
		baseURL, _ = url.Parse("http://" + config.Host)
	}
	key := fmt.Sprintf("Bearer %s", config.Key)
	return &RestClient{baseURL: baseURL.String(), key: key, client: http.DefaultClient}, nil
}

func (r *RestClient) Get(query string, params url.Values) ([]byte, int, error) {
	return r.doRequest("GET", query, params, nil)
}

func (r *RestClient) doRequest(method, query string, params url.Values, buf io.Reader) ([]byte, int, error) {
	u, _ := url.Parse(r.baseURL)
	u.Path = path.Join(u.Path, query)
	if params != nil {
		u.RawQuery = params.Encode()
	}
	req, err := http.NewRequest(method, u.String(), buf)
	req.Header.Set("Authorization", r.key)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return data, resp.StatusCode, err
}

func (r *RestClient) GetAllDatasources() ([]*DataSource, error) {
	var (
		raw         []byte
		datasources []*DataSource
		code        int
		err         error
	)

	datasourcesPath := "/api/datasources"
	if raw, code, err = r.Get(datasourcesPath, nil); err != nil {
		return nil, err
	}

	if code != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d", code)
	}

	if err = json.Unmarshal(raw, &datasources); err != nil {
		return nil, fmt.Errorf("Failed to parse JSON data from %s%s", r.baseURL, datasourcesPath)
	}

	existedEnv := make(map[string]bool)
	for idx, ds := range datasources {
		var newDs DataSource
		if newDs, err = r.GetDatasource(ds.Id); err != nil {
			return nil, err
		}

		// assign the current datasource with the new datasource with json data field
		ds = &newDs
		datasources[idx] = ds

		ds.SecureJsonData = make(map[string]string)
		for key, value := range ds.SecureJsonFields {
			if value {
				sanitized := slug.Make(fmt.Sprintf("%s_%s", ds.Name, key))
				placeholder := strcase.ToScreamingSnake("$" + sanitized)
				if existedEnv[placeholder] {
					// duplicated env existed
					return nil, fmt.Errorf("Duplicated ENV variable: `%s`. Rename `%s` datasource", placeholder, ds.Name)
				}
				existedEnv[placeholder] = true
				ds.SecureJsonData[key] = placeholder
			}
		}

		if ds.Access == "" {
			ds.Access = "proxy"
		}
	}

	return datasources, err
}

func (r *RestClient) GetDatasource(id int64) (DataSource, error) {
	var (
		raw  []byte
		ds   DataSource
		code int
		err  error
	)

	datasourcePath := fmt.Sprintf("/api/datasources/%d", id)
	if raw, code, err = r.Get(datasourcePath, nil); err != nil {
		return ds, err
	}
	if code != http.StatusOK {
		return ds, fmt.Errorf("HTTP error %d", code)
	}
	if err = json.Unmarshal(raw, &ds); err != nil {
		return ds, fmt.Errorf("Failed to parse JSON data from %s%s", r.baseURL, datasourcePath)
	}

	return ds, err
}
