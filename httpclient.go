package openweathermap

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	baseURL            = "https://api.openweathermap.org/data/2.5"
	userAgent          = "go-openweathermap/1.0"
	defaultAppIDEnvKey = "OPENWEATHERMAP_APPID"
)

var (
	ErrorNoAppID = errors.New("error no app id set")
)

var _ Client = (*HTTPClient)(nil)

type HTTPClient struct {
	httpClient *http.Client

	appID   string
	baseURL string
}

func NewHTTPClient(httpClient *http.Client, appID string) (*HTTPClient, error) {
	if appID == "" {
		envAppIDValue := os.Getenv(defaultAppIDEnvKey)
		if envAppIDValue == "" {
			return &HTTPClient{}, ErrorNoAppID
		} else {
			appID = envAppIDValue
		}
	}

	h := &HTTPClient{
		httpClient: httpClient,

		appID:   appID,
		baseURL: baseURL,
	}

	return h, nil
}

func (h *HTTPClient) request(ctx context.Context, method, path string, query url.Values, body io.Reader, out interface{}) (*http.Response, error) {
	query.Add("appid", h.appID)
	if len(query) > 0 {
		path = fmt.Sprintf("%s?%s", path, query.Encode())
	}

	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("/%s", path)
	}

	reqURL := fmt.Sprintf("%s%s", h.baseURL, path)

	req, err := http.NewRequestWithContext(ctx, method, reqURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", userAgent)

	res, err := h.httpClient.Do(req)
	if err != nil {
		return res, err
	}

	responseError := res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res, err
	}
	res.Body.Close()
	res.Body = ioutil.NopCloser(bytes.NewBuffer(resBody))

	if responseError {
		err := &apiError{}
		e := json.Unmarshal(resBody, err)
		if e != nil {
			return res, e
		}

		err.HTTPResponse = res

		return res, err
	}

	if out != nil {
		err = json.Unmarshal(resBody, out)
		if err != nil {
			return res, fmt.Errorf("error unmarshaling response body: %s", err.Error())
		}
	}

	return res, nil
}

func (h *HTTPClient) get(ctx context.Context, path string, query url.Values, out interface{}) (*http.Response, error) {
	return h.request(ctx, http.MethodGet, path, query, nil, out)
}
