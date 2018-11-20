package cmd

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/spf13/viper"

	"github.com/pkg/errors"
)

type Client struct {
	EndpointURL *url.URL
	HTTPClient  *http.Client
	Username    string
	Password    string
	authToken   string
	Logger      *log.Logger
}

func NewClient(endpointURL, username, password string, options ...func(*Client)) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(endpointURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse URL: %s", endpointURL)
	}
	if len(username) == 0 {
		return nil, errors.New("missing username")
	}
	if len(password) == 0 {
		return nil, errors.New("missing user password")
	}

	client := &Client{
		EndpointURL: parsedURL,
		Username:    username,
		Password:    password,
	}

	// part of functional option patterns
	for _, option := range options {
		option(client)
	}

	if client.HTTPClient == nil {
		client.HTTPClient = http.DefaultClient
	}

	discardLogger := log.New(ioutil.Discard, "", log.LstdFlags)
	if client.Logger == nil {
		client.Logger = discardLogger
	}

	return client, nil
}

func NewDefaultClient(options ...func(*Client)) (*Client, error) {
	endpointURL := viper.GetString("url")
	username := viper.GetString("username")
	password := viper.GetString("password")
	return NewClient(endpointURL, username, password, options...)
}

func (client *Client) NewRequest(ctx context.Context, method string, subPath string, body io.Reader) (*http.Request, error) {
	endpointURL := *client.EndpointURL
	endpointURL.Path = path.Join(client.EndpointURL.Path, subPath)

	req, err := http.NewRequest(method, endpointURL.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func decodeBody(resp *http.Response, out interface{}, f *os.File) error {
	defer resp.Body.Close()

	if f != nil {
		resp.Body = ioutil.NopCloser(io.TeeReader(resp.Body, f))
		defer f.Close()
	}

	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
