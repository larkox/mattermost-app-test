package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-test/routers/mattermost/bindings"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func fBindings(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	out := bindings.Get(c.Context)
	utils.WriteBindings(w, out)
}
