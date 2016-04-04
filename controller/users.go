package controller

import (
	"github.com/deis/workflow-cli/controller/api"
	"github.com/deis/workflow-cli/controller/models/users"
)

// ListUsers lists the users on the controller using the given client
func ListUsers(cl *Client, num int) ([]api.User, int, error) {
	cl.rwm.RLock()
	defer cl.rwm.RUnlock()
	return users.List(cl.deisClient, num)
}
