package clickup

import (
	"context"
	"fmt"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableClickupFolderlessList() *plugin.Table {
	return &plugin.Table{
		Name:        "clickup_folderless_list",
		Description: "Obtain lists not associated to a folder by providing a space_id.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("space_id"),
			Hydrate:    listFolderlessLists,
		},
		Columns: listColumns(),
	}
}

func listFolderlessLists(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	spaceId := d.KeyColumnQuals["space_id"].GetStringValue()
	plugin.Logger(ctx).Debug("listFolderlessLists", "spaceId", spaceId)

	lists, _, err := client.Lists.GetFolderlessLists(ctx, spaceId, true)
	if err != nil {
		plugin.Logger(ctx).Error(fmt.Sprintf("unable to obtain folderless lists for space id '%s': %v", spaceId, err))
		return nil, fmt.Errorf("unable to obtain folderless lists for space id '%s': %v", spaceId, err)
	}

	plugin.Logger(ctx).Debug("listFolderlessLists", "spaceId", spaceId, "results", len(lists))
	for _, list := range lists {
		d.StreamListItem(ctx, list)
	}

	return nil, nil
}
