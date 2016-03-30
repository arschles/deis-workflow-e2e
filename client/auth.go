package client

import (
	"github.com/deis/workflow-cli/controller/models/auth"
)

// Login logs the client into the the controller at cl.ControllerURL()
func Login(cl *Client, user, pass string) (string, error) {
	cl.rwm.RLock()
	tkn, err := auth.Login(cl.deisClient, user, pass)
	if err != nil {
		cl.rwm.RUnlock()
		return "", err
	}
	cl.rwm.RUnlock()
	cl.rwm.Lock()
	defer cl.rwm.Unlock()
	cl.deisClient.Token = tkn
	cl.deisClient.Username = user
	return tkn, nil
}

// Register registers a new user with the controller at cl.ControllerURL()
func Register(cl *Client, user, pass, email string) error {
	return auth.Register(cl.deisClient, user, pass, email)
}
