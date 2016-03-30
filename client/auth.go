package client

import (
	"github.com/deis/workflow-cli/controller/models/auth"
)

// Login logs the client into the the controller at cl.ControllerURL()
func Login(cl *Client, user, pass string) (string, error) {
	cl.rwm.RLock()
	tkn, err := auth.Login(cl.cl, user, pass)
	if err != nil {
		cl.rwm.RUnlock()
		return "", err
	}
	cl.rwm.RUnlock()
	cl.rwm.Lock()
	defer cl.rwm.Unlock()
	cl.cl.Token = tkn
	cl.cl.Username = user
	return tkn, nil
}

func Register(cl *Client, user, pass, email string) error {
	return auth.Register(cl.cl, user, pass, email)
}
