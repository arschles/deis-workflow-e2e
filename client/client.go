package client

import (
	"net/url"
	"sync"

	cclient "github.com/deis/workflow-cli/controller/client"
	"github.com/goware/urlx"
)

type Client struct {
	rwm *sync.RWMutex
	cl  *cclient.Client
}

// New creates a new Client instance. It will not have a token or be logged in
func New(urlStr string, responseLimit int, sslVerify bool) (*Client, error) {
	u, err := urlx.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	httpCl := cclient.CreateHTTPClient(sslVerify)
	return &Client{
		rwm: new(sync.RWMutex),
		cl: &cclient.Client{
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
	return c.cl.Token != ""
}

// ControllerURL returns the URL that c is configured to talk to
func (c *Client) ControllerURL() url.URL {
	return c.cl.ControllerURL
}
