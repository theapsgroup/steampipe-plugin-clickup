package clickup

import (
    "context"
    "errors"
    "github.com/raksul/go-clickup/clickup"
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin"
    "os"
)

func connect(_ context.Context, d *plugin.QueryData) (*clickup.Client, error) {
    token := os.Getenv("CLICKUP_TOKEN")

    clickupConfig := GetConfig(d.Connection)
    if clickupConfig.Token != nil {
        token = *clickupConfig.Token
    }

    if token == "" {
        return nil, errors.New("the 'token' must be set in connection configuration file or 'CLICKUP_TOKEN' environment variable must be set. Please set and then restart Steampipe")
    }

    client := clickup.NewClient(nil, token)
    return client, nil
}
