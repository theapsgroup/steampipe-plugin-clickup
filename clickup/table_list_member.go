package clickup

import (
	"context"
	"fmt"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableClickupListMember() *plugin.Table {
	return &plugin.Table{
		Name:        "clickup_list_member",
		Description: "Obtain members of a specific list by specifying a list_id.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("list_id"),
			Hydrate:    listListMembers,
		},
		Columns: memberColumns(),
	}
}

func listListMembers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	listId := d.KeyColumnQuals["list_id"].GetStringValue()
	plugin.Logger(ctx).Debug("listListMembers", "listId", listId)

	members, _, err := client.Members.GetListMembers(ctx, listId)
	if err != nil {
		plugin.Logger(ctx).Error(fmt.Sprintf("unable to obtain members for list id '%s': %v", listId, err))
		return nil, fmt.Errorf("unable to obtain members for list id '%s': %v", listId, err)
	}

	plugin.Logger(ctx).Debug("listListMembers", "listId", listId, "results", len(members))
	for _, member := range members {
		d.StreamListItem(ctx, member)
	}

	return nil, nil
}

func memberColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Type:        proto.ColumnType_INT,
			Description: "Unique Identifier of the user.",
		},
		{
			Name:        "username",
			Type:        proto.ColumnType_STRING,
			Description: "Username of the user.",
		},
		{
			Name:        "email",
			Type:        proto.ColumnType_STRING,
			Description: "Email address of the user.",
		},
		{
			Name:        "color",
			Type:        proto.ColumnType_STRING,
			Description: "The color associated with the user.",
		},
		{
			Name:        "profile_picture",
			Type:        proto.ColumnType_STRING,
			Description: "URL of the profile picture for the user.",
			Transform:   transform.FromField("ProfilePicture"),
		},
		{
			Name:        "initials",
			Type:        proto.ColumnType_STRING,
			Description: "Initials of the user.",
		},
		{
			Name:        "list_id",
			Type:        proto.ColumnType_STRING,
			Description: "Identifier for the list the user is a member of.",
			Transform:   transform.FromQual("list_id"),
		},
	}
}
