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
		Name:        "clickup_task",
		Description: "Obtain tasks by a specific id or by providing either a team_id or list_id",
		List: &plugin.ListConfig{
			Hydrate: listTasks,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "team_id",
					Require: plugin.AnyOf,
				},
				{
					Name:    "list_id",
					Require: plugin.AnyOf,
				},
				{
					Name:    "status",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getTask,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: taskColumns(),
	}
}

func listTasks(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	q := d.KeyColumnQuals

	// Default options
	opts := &clickup.GetTasksOptions{
		Page:          0,
		Archived:      true,
		IncludeClosed: true,
		Subtasks:      true,
	}

	teamId := q["team_id"].GetStringValue()
	listId := q["list_id"].GetStringValue()

	// TODO: Enhance with folderId, listId, spaceId, etc once support is captured in the SDK.
	if q["status"] != nil {
		plugin.Logger(ctx).Debug("listTasks", "status", q["status"].GetStringValue())
		opts.Statuses = []string{q["status"].GetStringValue()}
	}

	var ts []clickup.Task
	for {
		if listId != "" {
			plugin.Logger(ctx).Debug("listTasks", "listId", listId, "page", opts.Page)
			tasks, _, err := client.Tasks.GetTasks(ctx, listId, opts)
			if err != nil {
				plugin.Logger(ctx).Error(fmt.Sprintf("unable to obtain tasks for list id '%s': %v", listId, err))
				return nil, fmt.Errorf("unable to obtain tasks for list id '%s': %v", listId, err)
			}
			plugin.Logger(ctx).Debug("listTasks", "listId", listId, "page", opts.Page, "results", len(tasks))
			ts = tasks
		} else {
			plugin.Logger(ctx).Debug("listTasks", "teamId", teamId, "page", opts.Page)
			tasks, _, err := client.Tasks.GetFilteredTeamTasks(ctx, teamId, opts)
			if err != nil {
				plugin.Logger(ctx).Error(fmt.Sprintf("unable to obtain tasks for team id '%s': %v", teamId, err))
				return nil, fmt.Errorf("unable to obtain tasks for team id '%s': %v", teamId, err)
			}
			plugin.Logger(ctx).Debug("listTasks", "teamId", teamId, "page", opts.Page, "results", len(tasks))
			ts = tasks
		}

		for _, t := range ts {
			d.StreamListItem(ctx, t)
		}

		if len(ts) < 100 {
			plugin.Logger(ctx).Debug("listTasks - exiting as page returned < 100 items", "page", opts.Page, "results", len(ts))
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
		{
			Name:        "creator_id",
			Type:        proto.ColumnType_INT,
			Description: "Identifier for the user whom created the task.",
			Transform:   transform.FromField("Creator.ID"),
		},
		{
			Name:        "creator",
			Type:        proto.ColumnType_STRING,
			Description: "Username of the user whom created the task.",
			Transform:   transform.FromField("Creator.Username"),
		},
		{
			Name:        "creator_email",
			Type:        proto.ColumnType_STRING,
			Description: "Email address of the user whom created the task.",
			Transform:   transform.FromField("Creator.Email"),
		},
		// checklists
		{
			Name:        "tags",
			Type:        proto.ColumnType_JSON,
			Description: "An array of tags associated with the task.",
		},
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
			Type:        proto.ColumnType_STRING,
			Description: "Timestamp when the task is due.",
			Transform:   transform.FromField("DueDate"),
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
			Transform:   transform.FromField("Points.IntVal"),
		},
		{
			Name:        "time_estimate",
			Type:        proto.ColumnType_INT,
			Description: "Estimate (in ms) of how long is required to complete the task.",
		},
		// custom fields
		{
			Name:        "dependencies",
			Type:        proto.ColumnType_JSON,
			Description: "An array of task dependencies.",
		},
		{
			Name:        "linked_tasks",
			Type:        proto.ColumnType_JSON,
			Description: "An array of json objects to identify linked tasks.",
		},
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
