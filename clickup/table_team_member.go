package clickup

import (
	"context"
	"fmt"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableClickupTeamMember() *plugin.Table {
	return &plugin.Table{
		Name:        "clickup_team_member",
		Description: "Obtain members for a specific team by specifying a task_id.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("team_id"),
			Hydrate:    listTeamMembers,
		},
		Columns: teamMemberColumns(),
	}
}

func listTeamMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	teamId := d.KeyColumnQuals["team_id"].GetStringValue()

	// Note: No SDK method to obtain a single team
	teams, _, err := client.Teams.GetTeams(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to obtain teams: %v", err)
	}

	for _, team := range teams {
		if team.ID == teamId {
			for _, member := range team.Members {
				d.StreamListItem(ctx, member.User)
			}
			break
		}
	}

	return nil, nil
}

func teamMemberColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Type:        proto.ColumnType_INT,
			Description: "Identifier for the user.",
		},
		{
			Name:        "username",
			Type:        proto.ColumnType_STRING,
			Description: "Username for the user.",
		},
		{
			Name:        "email",
			Type:        proto.ColumnType_STRING,
			Description: "Email address for the user.",
		},
		{
			Name:        "color",
			Type:        proto.ColumnType_STRING,
			Description: "Color associated with the user.",
		},
		{
			Name:        "profile_picture",
			Type:        proto.ColumnType_STRING,
			Description: "URL for the profile picture of the user.",
		},
		{
			Name:        "initials",
			Type:        proto.ColumnType_STRING,
			Description: "The initials of the user.",
		},
		{
			Name:        "last_active",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when the user was last active.",
			Transform:   transform.FromField("LastActive").Transform(unixTimeTransform),
		},
		{
			Name:        "date_joined",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when the user actually joined the team.",
			Transform:   transform.FromField("DateJoined").Transform(unixTimeTransform),
		},
		{
			Name:        "date_invited",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when the user was invited to join the team.",
			Transform:   transform.FromField("DateInvited").Transform(unixTimeTransform),
		},
		{
			Name:        "team_id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier for the team.",
			Transform:   transform.FromQual("team_id"),
		},
	}
}
