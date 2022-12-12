package clickup

import (
	"context"
	"fmt"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableClickupTeam() *plugin.Table {
	return &plugin.Table{
		Name:        "clickup_team",
		Description: "Obtain teams associated with your user token.",
		List: &plugin.ListConfig{
			Hydrate: listTeams,
		},
		Columns: teamColumns(),
	}
}

func listTeams(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	plugin.Logger(ctx).Debug("listTeams")

	teams, _, err := client.Teams.GetTeams(ctx)
	if err != nil {
		plugin.Logger(ctx).Error(fmt.Sprintf("unable to obtain teams: %v", err))
		return nil, fmt.Errorf("unable to obtain teams: %v", err)
	}

	plugin.Logger(ctx).Debug("listTeams", "results", len(teams))
	for _, team := range teams {
		d.StreamListItem(ctx, team)
	}

	return nil, nil
}

func teamColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier for the team.",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Name given to the team.",
		},
		{
			Name:        "color",
			Type:        proto.ColumnType_STRING,
			Description: "Color associated with the team.",
		},
	}
}
