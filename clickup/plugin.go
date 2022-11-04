package clickup

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-clickup",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo(),
		TableMap: map[string]*plugin.Table{
			"clickup_folder":        tableClickupFolder(),
			"clickup_goal":          tableClickupGoal(),
			"clickup_space":         tableClickupSpace(),
			"clickup_task":          tableClickupTask(),
			"clickup_task_assignee": tableClickupTaskAssignee(),
			"clickup_task_watcher":  tableClickupTaskWatcher(),
			"clickup_team":          tableClickupTeam(),
			"clickup_team_member":   tableClickupTeamMember(),
		},
	}

	return p
}
