package clickup

import (
	"context"
	"fmt"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableClickupGoal() *plugin.Table {
	return &plugin.Table{
		Name:        "clickup_goal",
		Description: "Obtain goals by specifying either an id or a team_id.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("team_id"),
			Hydrate:    listGoals,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getGoal,
		},
		Columns: goalColumns(),
	}
}

func listGoals(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	teamId := d.KeyColumnQuals["team_id"].GetStringValue()

	for {
		goals, _, _, err := client.Goals.GetGoals(ctx, teamId, true)
		if err != nil {
			return nil, fmt.Errorf("unable to obtain goals for team id '%s': %v", teamId, err)
		}

		for _, goal := range goals {
			d.StreamListItem(ctx, goal)
		}
	}

	return nil, nil
}

func getGoal(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	goalId := d.KeyColumnQuals["id"].GetStringValue()

	goal, _, err := client.Goals.GetGoal(ctx, goalId)
	if err != nil {
		return nil, fmt.Errorf("unable to obtain goal with id '%s': %v", goalId, err)
	}

	return goal, nil
}

func goalColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier for the goal.",
		},
		{
			Name:        "pretty_id",
			Type:        proto.ColumnType_STRING,
			Description: "Pretty identifier for the goal.",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "The name of the given goal.",
		},
		{
			Name:        "team_id",
			Type:        proto.ColumnType_STRING,
			Description: "Identifier for the team to which the goal is assigned.",
		},
		{
			Name:        "creator_id",
			Type:        proto.ColumnType_INT,
			Description: "Identifier for the user whom created the goal.",
			Transform:   transform.FromField("Creator"),
		},
		{
			Name:        "owner_id",
			Type:        proto.ColumnType_INT,
			Description: "Identifier for the user designated as the owner of the goal.",
			Transform:   transform.FromField("Owner.ID"),
		},
		{
			Name:        "owner_email",
			Type:        proto.ColumnType_STRING,
			Description: "Email address for the user designated as the owner of the goal.",
			Transform:   transform.FromField("Owner.Email"),
		},
		{
			Name:        "owner_username",
			Type:        proto.ColumnType_STRING,
			Description: "Username of the user designated as the owner of the goal.",
			Transform:   transform.FromField("Owner.Username"),
		},
		{
			Name:        "color",
			Type:        proto.ColumnType_STRING,
			Description: "The color associated with the goal.",
		},
		{
			Name:        "date_created",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when the goal was created.",
			Transform:   transform.From(unixTimeTransform),
		},
		{
			Name:        "start_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when work on the goal was started.",
			Transform:   transform.From(unixTimeTransform),
		},
		{
			Name:        "due_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when the goal is due.",
			Transform:   transform.From(unixTimeTransform),
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Textual description of the goal.",
		},
		{
			Name:        "private",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the goal is private.",
		},
		{
			Name:        "archived",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the goal is archived.",
		},
		{
			Name:        "multiple_owners",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the goal has/supports multiple owners.",
		},
		{
			Name:        "date_updated",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when the goal was last updated.",
			Transform:   transform.From(unixTimeTransform),
		},
		{
			Name:        "folder_id",
			Type:        proto.ColumnType_STRING,
			Description: "identifier for the folder of the goal.",
		},
		{
			Name:        "pinned",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the goal is pinned.",
		},
		{
			Name:        "owners",
			Type:        proto.ColumnType_JSON,
			Description: "An array of owners if multiple owners are assigned to the goal.",
		},
		{
			Name:        "percent_completed",
			Type:        proto.ColumnType_INT,
			Description: "Numeric representation of how complete the goal is.",
		},
	}
}
