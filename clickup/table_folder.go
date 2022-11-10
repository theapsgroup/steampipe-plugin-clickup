package clickup

import (
	"context"
	"fmt"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableClickupFolder() *plugin.Table {
	return &plugin.Table{
		Name:        "clickup_folder",
		Description: "Obtain folders by specifying either an id or a space_id.",
		List: &plugin.ListConfig{
			Hydrate:    listFolders,
			KeyColumns: plugin.SingleColumn("space_id"),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getFolder,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: folderColumns(),
	}
}

func listFolders(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	spaceId := d.KeyColumnQuals["space_id"].GetStringValue()

	folders, _, err := client.Folders.GetFolders(ctx, spaceId, true)
	if err != nil {
		return nil, fmt.Errorf("unable to obtain folders for space id '%s': %v", spaceId, err)
	}

	for _, folder := range folders {
		d.StreamListItem(ctx, folder)
	}

	return nil, nil
}

func getFolder(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	folderId := d.KeyColumnQuals["id"].GetStringValue()

	folder, _, err := client.Folders.GetFolder(ctx, folderId)
	if err != nil {
		return nil, fmt.Errorf("unable to obtain folder with id '%s': %v", folderId, err)
	}

	return folder, nil
}

func folderColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier for the folder.",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Name of the given folder.",
		},
		{
			Name:        "order_index",
			Type:        proto.ColumnType_INT,
			Description: "The order index of the folder.",
			Transform:   transform.FromField("Orderindex"),
		},
		{
			Name:        "override_statuses",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if statuses can be overriden within the folder.",
		},
		{
			Name:        "hidden",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the folder is hidden.",
		},
		{
			Name:        "space_id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier of the space to which the folder belongs.",
			Transform:   transform.FromField("Space.ID"),
		},
		{
			Name:        "task_count",
			Type:        proto.ColumnType_INT,
			Description: "Count of tasks within the folder.",
		},
		{
			Name:        "archived",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the folder is archived.",
		},
		{
			Name:        "statuses",
			Type:        proto.ColumnType_JSON,
			Description: "An array of status objects for the valid statuses within the folder.",
		},
	}
}
