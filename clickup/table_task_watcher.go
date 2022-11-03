package clickup

import (
	"context"
	"fmt"
	"github.com/raksul/go-clickup/clickup"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableClickupTaskWatcher() *plugin.Table {
	return &plugin.Table{
		Name: "clickup_task_watcher",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("task_id"),
			Hydrate:    listTaskWatchers,
		},
		Columns: taskUserColumns(),
	}
}

func listTaskWatchers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	taskId := d.KeyColumnQuals["task_id"].GetStringValue()

	opts := &clickup.GetTaskOptions{}

	task, _, err := client.Tasks.GetTask(ctx, taskId, opts)
	if err != nil {
		return nil, fmt.Errorf("unable to obtain task with id '%s': %v", taskId, err)
	}

	for _, watcher := range task.Watchers {
		d.StreamListItem(ctx, watcher)
	}

	return nil, nil
}
