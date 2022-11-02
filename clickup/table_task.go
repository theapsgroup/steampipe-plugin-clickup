package clickup

import (
	"context"
	"fmt"
	"github.com/raksul/go-clickup/clickup"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableClickupTask() *plugin.Table {
	return &plugin.Table{
		Name: "clickup_task",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("list_id"),
			Hydrate:    listTasks,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTask,
		},
		Columns: taskColumns(),
	}
}

func listTasks(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	listId := d.KeyColumnQuals["list_id"].GetStringValue()

	opts := &clickup.GetTasksOptions{
		Page:          0,
		Archived:      true,
		IncludeClosed: true,
		Subtasks:      true,
	}

	for {
		tasks, _, err := client.Tasks.GetTasks(ctx, listId, opts)
		if err != nil {
			return nil, fmt.Errorf("unable to obtain tasks for list id '%s': %v", listId, err)
		}

		for _, task := range tasks {
			d.StreamListItem(ctx, task)
		}

		if len(tasks) < 100 {
			break
		}

		opts.Page++
	}

	return nil, nil
}

func getTask(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	taskId := d.KeyColumnQuals["id"].GetStringValue()

	opts := &clickup.GetTaskOptions{}

	task, _, err := client.Tasks.GetTask(ctx, taskId, opts)
	if err != nil {
		return nil, fmt.Errorf("unable to obtain task with id '%s': %v", taskId, err)
	}

	return task, nil
}

func taskColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique string identifier for the task.",
		},
		{
			Name:        "custom_id",
			Type:        proto.ColumnType_STRING,
			Description: "Custom identifier for the task.",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Name given to the task.",
		},
		{
			Name:        "text_content",
			Type:        proto.ColumnType_STRING,
			Description: "The textual content of the task.",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "The description on the task.",
		},
		{
			Name:        "status",
			Type:        proto.ColumnType_STRING,
			Description: "Current status of the task.",
			Transform:   transform.FromField("Status.Status"),
		},
		{
			Name:        "order_index",
			Type:        proto.ColumnType_STRING,
			Description: "Order index of the task.",
			Transform:   transform.FromField("Orderindex"),
		},
		{
			Name:        "date_created",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when the task was created.",
			Transform:   transform.FromField("DateCreated").Transform(unixTimeTransform),
		},
		{
			Name:        "date_updated",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when the task was last updated.",
			Transform:   transform.FromField("DateUpdated").Transform(unixTimeTransform),
		},
		{
			Name:        "date_closed",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when the task was closed.",
			Transform:   transform.FromField("DateClosed").Transform(unixTimeTransform),
		},
		{
			Name:        "archived",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the task is archived.",
		},
		// creator
		// assignees
		// watchers
		// checklists
		// tags
		{
			Name:        "parent",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier for the parent task.",
		},
		{
			Name:        "priority",
			Type:        proto.ColumnType_STRING,
			Description: "The priority of the task.",
			Transform:   transform.FromField("Priority.Priority"),
		},
		{
			Name:        "due_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when the task is due.",
			Transform:   transform.FromField("DueDate").Transform(unixTimeTransform),
		},
		{
			Name:        "start_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when work on the task was started.",
			Transform:   transform.FromField("StartDate").Transform(unixTimeTransform),
		},
		{
			Name:        "points",
			Type:        proto.ColumnType_INT,
			Description: "Points attributed to the task.",
		},
		{
			Name:        "time_estimate",
			Type:        proto.ColumnType_INT,
			Description: "Estimate (in ms) of how long is required to complete the task.",
		},
		// custom fields
		// dependencies
		// linked tasks
		{
			Name:        "team_id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier for the team associated with the task.",
		},
		{
			Name:        "url",
			Type:        proto.ColumnType_STRING,
			Description: "Direct URL to the task.",
			Transform:   transform.FromField("URL"),
		},
		{
			Name:        "permission_level",
			Type:        proto.ColumnType_STRING,
			Description: "Permission level associated to the task.",
		},
		{
			Name:        "list_id",
			Type:        proto.ColumnType_STRING,
			Description: "Identifier of the list the task belongs to.",
			Transform:   transform.FromField("List.ID"),
		},
		{
			Name:        "list",
			Type:        proto.ColumnType_STRING,
			Description: "Name of the list the task belongs to.",
			Transform:   transform.FromField("List.Name"),
		},
		{
			Name:        "project_id",
			Type:        proto.ColumnType_STRING,
			Description: "Identifier of the project the task belongs to.",
			Transform:   transform.FromField("Project.ID"),
		},
		{
			Name:        "project",
			Type:        proto.ColumnType_STRING,
			Description: "Name of the project the task belongs to.",
			Transform:   transform.FromField("Project.Name"),
		},
		{
			Name:        "folder_id",
			Type:        proto.ColumnType_STRING,
			Description: "Identifier of the folder the task belongs to.",
			Transform:   transform.FromField("Folder.ID"),
		},
		{
			Name:        "folder",
			Type:        proto.ColumnType_STRING,
			Description: "Name of the folder the task belongs to.",
			Transform:   transform.FromField("Folder.Name"),
		},
		{
			Name:        "space_id",
			Type:        proto.ColumnType_STRING,
			Description: "Identifier of the space the task belongs to.",
			Transform:   transform.FromField("Space.ID"),
		},
	}
}
