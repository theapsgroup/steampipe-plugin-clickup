package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"steampipe-plugin-clickup/clickup"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: clickup.Plugin})
}
