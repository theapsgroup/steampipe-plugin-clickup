# Table: clickup_space

Obtain information about the spaces within your ClickUp environment.

However you **MUST** specify either an `id` (single) or `team_id` (multiple) in the WHERE or JOIN clause.

> Note: In the event of using the space `id` to obtain a single entity, the `team_id` field will be null as no value is returned by the API.

## Examples

### Get a space by id

```sql
    select
        id,
        name,
        private,
        sprints,
        tags,
        milestones
from
    clickup_space
where
    id = '7423465';
```

### List all spaces for a specific team

```sql
    select
        id,
        name,
        private,
        sprints,
        tags,
        milestones,
        statuses,
        multiple_assignees,
        due_dates,
        points,
        time_estimates,
        time_tracking,
        archived
from
    clickup_space
where
    team_id = '969532885';
```
