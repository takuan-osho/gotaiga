package cmd

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

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

	for _, option := range options {
		option(client)
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
