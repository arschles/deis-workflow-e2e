package controller

import (
	"github.com/deis/workflow-cli/controller/api"
	"github.com/deis/workflow-cli/controller/models/apps"
)

// NewApp creates a new app with the given ID
func NewApp(cl *Client, appID string) (api.App, error) {
	return apps.New(cl.deisClient, appID)
}

// DeleteApp deletes the app with the given appID
func DeleteApp(cl *Client, appID string) error {
	return apps.Delete(cl.deisClient, appID)
}
