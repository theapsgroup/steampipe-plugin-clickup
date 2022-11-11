package clickup

import (
	"context"
	"fmt"
	"github.com/raksul/go-clickup/clickup"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableClickupTeamTask() *plugin.Table {
	return &plugin.Table{
		Name:        "clickup_team_task",
		Description: "Obtain tasks by specifying an id or a team_id, as well as other filters.",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "team_id",
					Require:   plugin.Required,
					Operators: []string{"="},
				},
				{
					Name:      "status",
					Require:   plugin.Optional,
					Operators: []string{"="},
				},
			},
			Hydrate: listTeamTasks,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTask,
		},
		Columns: taskColumns(),
	}
}

func listTeamTasks(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	opts := &clickup.GetTasksOptions{
		Page:          0,
		Archived:      true,
		IncludeClosed: true,
		Subtasks:      true,
	}

	q := d.KeyColumnQuals

	teamId := q["team_id"].GetStringValue()

	// TODO: Enhance with folderId, listId, spaceId, etc once support is captured in the SDK.
	if q["status"] != nil {
		opts.Statuses = []string{q["status"].GetStringValue()}
	}

	for {
		tasks, _, err := client.Tasks.GetFilteredTeamTasks(ctx, teamId, opts)
		if err != nil {
			return nil, fmt.Errorf("unable to obtain tasks for team id '%s': %v", teamId, err)
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
