package clickup

import (
	"context"
	"errors"
	"github.com/raksul/go-clickup/clickup"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
	"os"
	"strconv"
	"time"
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

// Transform Functions
func unixTimeTransform(_ context.Context, input *transform.TransformData) (interface{}, error) {
	if input.Value == nil || input.Value == "" {
		return nil, nil
	}

	value, err := strconv.ParseInt(input.Value.(string), 0, 64)
	if err != nil {
		return nil, err
	}
	return time.Unix(value/1000, 0), nil
}
