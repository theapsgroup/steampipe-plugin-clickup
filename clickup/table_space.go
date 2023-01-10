package clickup

import (
	"context"
	"fmt"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableClickupSpace() *plugin.Table {
	return &plugin.Table{
		Name:        "clickup_space",
		Description: "Obtain spaces by specifying either an id or a team_id.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("team_id"),
			Hydrate:    listSpace,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSpace,
		},
		Columns: spaceColumns(),
	}
}

func listSpace(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	teamId := d.EqualsQuals["team_id"].GetStringValue()
	plugin.Logger(ctx).Debug("listSpace", "teamId", teamId)

	spaces, _, err := client.Spaces.GetSpaces(ctx, teamId)
	if err != nil {
		plugin.Logger(ctx).Error(fmt.Sprintf("unable to obtain spacess for team id '%s': %v", teamId, err))
		return nil, fmt.Errorf("unable to obtain spacess for team id '%s': %v", teamId, err)
	}

	plugin.Logger(ctx).Debug("listSpace", "teamId", teamId, "results", len(spaces))
	for _, space := range spaces {
		d.StreamListItem(ctx, space)
	}

	return nil, nil
}

func getSpace(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	spaceId := d.EqualsQuals["id"].GetStringValue()
	plugin.Logger(ctx).Debug("getSpace", "id", spaceId)

	space, _, err := client.Spaces.GetSpace(ctx, spaceId)
	if err != nil {
		plugin.Logger(ctx).Error(fmt.Sprintf("unable to obtain space with id '%s': %v", spaceId, err))
		return nil, fmt.Errorf("unable to obtain space with id '%s': %v", spaceId, err)
	}

	return space, nil
}

func spaceColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier for the space.",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "The name of the space.",
		},
		{
			Name:        "private",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the space is designated as a private space.",
		},
		{
			Name:        "multiple_assignees",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the space supports multiple assignees to tasks, etc.",
		},
		{
			Name:        "due_dates",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the space has due dates enabled.",
			Transform:   transform.FromField("Features.DueDates.Enabled"),
		},
		{
			Name:        "sprints",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the space has sprints enabled.",
			Transform:   transform.FromField("Features.Sprints.Enabled"),
		},
		{
			Name:        "time_tracking",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the space has time tracking enabled.",
			Transform:   transform.FromField("Features.TimeTracking.Enabled"),
		},
		{
			Name:        "time_estimates",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the space has time estimates enabled.",
			Transform:   transform.FromField("Features.TimeEstimates.Enabled"),
		},
		{
			Name:        "points",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the space has points enabled.",
			Transform:   transform.FromField("Features.Points.Enabled"),
		},
		{
			Name:        "custom_items",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the space has custom items enabled.",
			Transform:   transform.FromField("Features.CustomItems.Enabled"),
		},
		{
			Name:        "tags",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the space has tags enabled.",
			Transform:   transform.FromField("Features.Tags.Enabled"),
		},
		{
			Name:        "milestones",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the space has milestones enabled.",
			Transform:   transform.FromField("Features.Milestones.Enabled"),
		},
		{
			Name:        "archived",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the space is archived.",
		},
		{
			Name:        "statuses",
			Type:        proto.ColumnType_JSON,
			Description: "An array os statuses available on the space.",
		},
		{
			Name:      "team_id",
			Type:      proto.ColumnType_STRING,
			Transform: transform.FromQual("team_id"),
		},
	}
}
