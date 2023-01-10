package clickup

import (
	"context"
	"fmt"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableClickupList() *plugin.Table {
	return &plugin.Table{
		Name:        "clickup_list",
		Description: "Obtain lists that are associated to a folder by specifying either an id or a folder_id.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("folder_id"),
			Hydrate:    listLists,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getList,
		},
		Columns: listColumns(),
	}
}

func listLists(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	folderId := d.EqualsQuals["folder_id"].GetStringValue()
	plugin.Logger(ctx).Debug("listLists", "folderId", folderId)

	lists, _, err := client.Lists.GetLists(ctx, folderId, true)
	if err != nil {
		plugin.Logger(ctx).Error(fmt.Sprintf("unable to obtain lists for folder id '%s': %v", folderId, err))
		return nil, fmt.Errorf("unable to obtain lists for folder id '%s': %v", folderId, err)
	}

	plugin.Logger(ctx).Debug("listLists", "folderId", folderId, "results", len(lists))
	for _, list := range lists {
		d.StreamListItem(ctx, list)
	}

	return nil, nil
}

func getList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("unable to establish a connection: %v", err)
	}

	listId := d.EqualsQuals["id"].GetStringValue()
	plugin.Logger(ctx).Debug("getList", "id", listId)

	list, _, err := client.Lists.GetList(ctx, listId)
	if err != nil {
		plugin.Logger(ctx).Error(fmt.Sprintf("unable to obtain list with id '%s': %v", listId, err))
		return nil, fmt.Errorf("unable to obtain list with id '%s': %v", listId, err)
	}

	return list, nil
}

func listColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier for the list.",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Name of the list.",
		},
		{
			Name:        "order_index",
			Type:        proto.ColumnType_INT,
			Description: "Order index of the list.",
			Transform:   transform.FromField("Orderindex"),
		},
		{
			Name:        "content",
			Type:        proto.ColumnType_STRING,
			Description: "Content description of the list.",
		},
		{
			Name:        "status",
			Type:        proto.ColumnType_STRING,
			Description: "Status of the list.",
			Transform:   transform.FromField("Status.Status"),
		},
		{
			Name:        "priority",
			Type:        proto.ColumnType_STRING,
			Description: "Priority of the list.",
			Transform:   transform.FromField("Priority.Priority"),
		},
		{
			Name:        "assignee_id",
			Type:        proto.ColumnType_INT,
			Description: "Unique identifier of the user to whom the list is assigned.",
			Transform:   transform.FromField("Assignee.ID"),
		},
		{
			Name:        "assignee",
			Type:        proto.ColumnType_STRING,
			Description: "Username of the user to whom the list is assigned.",
			Transform:   transform.FromField("Assignee.Username"),
		},
		{
			Name:        "assignee_email",
			Type:        proto.ColumnType_STRING,
			Description: "Email address of the user to whom the list is assigned.",
			Transform:   transform.FromField("Assignee.Email"),
		},
		{
			Name:        "task_count",
			Type:        proto.ColumnType_STRING,
			Description: "Count of tasks on the list.",
		},
		{
			Name:        "due_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when the list is due.",
			Transform:   transform.From(unixTimeTransform),
		},
		{
			Name:        "start_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "Timestamp when work on the list was started.",
			Transform:   transform.From(unixTimeTransform),
		},
		{
			Name:        "folder_id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier of the folder the list belongs to.",
			Transform:   transform.FromField("Folder.ID"),
		},
		{
			Name:        "folder",
			Type:        proto.ColumnType_STRING,
			Description: "Name of the folder the list belongs to.",
			Transform:   transform.FromField("Folder.Name"),
		},
		{
			Name:        "space_id",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier of the space the list belongs to.",
			Transform:   transform.FromField("Space.ID"),
		},
		{
			Name:        "space",
			Type:        proto.ColumnType_STRING,
			Description: "Name of the space the list belongs to.",
			Transform:   transform.FromField("Space.Name"),
		},
		{
			Name:        "statuses",
			Type:        proto.ColumnType_JSON,
			Description: "An array of status objects for available statuses within the list.",
		},
		{
			Name:        "archived",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the list is archived.",
		},
	}
}
