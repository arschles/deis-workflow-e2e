package controller

import (
	"net/url"
	"sync"

	cclient "github.com/deis/workflow-cli/controller/client"
	"github.com/goware/urlx"
)

// Client is a concurrency safe wrapper around the client library at
// github.com/deis/workflow-cli/controller/client
type Client struct {
	rwm        *sync.RWMutex
	deisClient *cclient.Client
}

// NewClient creates a new Client instance. It will not have a token or be logged in
func NewClient(urlStr string, responseLimit int, sslVerify bool) (*Client, error) {
	u, err := urlx.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	httpCl := cclient.CreateHTTPClient(sslVerify)
	return &Client{
		rwm: new(sync.RWMutex),
		deisClient: &cclient.Client{
			HTTPClient:    httpCl,
			SSLVerify:     sslVerify,
			ControllerURL: *u,
			Token:         "",
			Username:      "",
			ResponseLimit: responseLimit,
		},
	}, nil
}

// IsLoggedIn returns true if the client has a valid token, false otherwise
func (c *Client) IsLoggedIn() bool {
	c.rwm.RLock()
	defer c.rwm.RUnlock()
	return c.deisClient.Token != ""
}

// ControllerURL returns the URL that c is configured to talk to
func (c *Client) ControllerURL() url.URL {
	c.rwm.RLock()
	c.rwm.RUnlock()
	return c.deisClient.ControllerURL
}
