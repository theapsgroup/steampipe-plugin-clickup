package clickup

import (
	"context"
	"fmt"
	"github.com/raksul/go-clickup/clickup"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableClickupTaskWatcher() *plugin.Table {
	return &plugin.Table{
		Name:        "clickup_task_watcher",
		Description: "Obtain watchers for a specific task by specifying a task_id.",
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
	plugin.Logger(ctx).Debug("listTaskWatchers", "taskId", taskId)

	opts := &clickup.GetTaskOptions{}

	task, _, err := client.Tasks.GetTask(ctx, taskId, opts)
	if err != nil {
		plugin.Logger(ctx).Error(fmt.Sprintf("unable to obtain task with id '%s': %v", taskId, err))
		return nil, fmt.Errorf("unable to obtain task with id '%s': %v", taskId, err)
	}

	plugin.Logger(ctx).Debug("listTaskWatchers", "taskId", taskId, "results", len(task.Watchers))
	for _, watcher := range task.Watchers {
		d.StreamListItem(ctx, watcher)
	}

	return nil, nil
}
