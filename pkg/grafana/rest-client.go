package grafana

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

type RestClient struct {
	baseURL string
	key     string
	client  *http.Client
}

func NewRestClient(host string, key string) *RestClient {
	baseURL, _ := url.Parse(host)
	key = fmt.Sprintf("Bearer %s", key)
	return &RestClient{baseURL: baseURL.String(), key: key, client: http.DefaultClient}
}

func (r *RestClient) Get(query string, params url.Values) (io.Reader, int, error) {
	return r.doRequest("GET", query, params, nil)
}

func (r *RestClient) doRequest(method, query string, params url.Values, buf io.Reader) (io.Reader, int, error) {
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
	return resp.Body, resp.StatusCode, err
}

func (r *RestClient) GetAllDatasources() (io.Reader, error) {
	var (
		raw  io.Reader
		code int
		err  error
	)

	if raw, code, err = r.Get("api/datasources", nil); err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("HTTP error %d", code)
	}

	return raw, err
}
