package clickup

import (
	"context"
	"fmt"
	"github.com/raksul/go-clickup/clickup"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableClickupTaskAssignee() *plugin.Table {
	return &plugin.Table{
		Name:        "clickup_task_assignee",
		Description: "Obtain assignees for a specific task by specifying a task_id.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("task_id"),
			Hydrate:    listTaskAssignees,
		},
		Columns: taskUserColumns(),
	}
}

func listTaskAssignees(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	taskId := d.EqualsQuals["task_id"].GetStringValue()
	plugin.Logger(ctx).Debug("listTaskAssignees", "taskId", taskId)

	opts := &clickup.GetTaskOptions{}

	task, _, err := client.Tasks.GetTask(ctx, taskId, opts)
	if err != nil {
		plugin.Logger(ctx).Error(fmt.Sprintf("unable to obtain task with id '%s': %v", taskId, err))
		return nil, fmt.Errorf("unable to obtain task with id '%s': %v", taskId, err)
	}

	plugin.Logger(ctx).Debug("listTaskAssignees", "taskId", taskId, "results", len(task.Assignees))
	for _, assignee := range task.Assignees {
		d.StreamListItem(ctx, assignee)
	}

	return nil, nil
}

func taskUserColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Type:        proto.ColumnType_INT,
			Description: "Identifier for the task assignee.",
		},
		{
			Name:        "username",
			Type:        proto.ColumnType_STRING,
			Description: "Username for the task assignee.",
		},
		{
			Name:        "email",
			Type:        proto.ColumnType_STRING,
			Description: "Email address for the task assignee.",
		},
		{
			Name:        "color",
			Type:        proto.ColumnType_STRING,
			Description: "Color associated with the task assignee.",
		},
		{
			Name:        "profile_picture",
			Type:        proto.ColumnType_STRING,
			Description: "URL for the profile picture of the task assignee.",
		},
		{
			Name:        "initials",
			Type:        proto.ColumnType_STRING,
			Description: "The initials of the task assignee.",
		},
		{
			Name:        "task_id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier for the task.",
			Transform:   transform.FromQual("task_id"),
		},
	}
}
