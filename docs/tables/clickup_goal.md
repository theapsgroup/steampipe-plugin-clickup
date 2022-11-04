# Table: clickup_goal

Obtain information about folders within your ClickUp environment.

However you **MUST** specify either an `id` (single) or `team_id` (for multiple tasks) in the WHERE or JOIN clause.

## Examples

### Get a specific goal by id

```sql
select
    id,
    pretty_id,
    name,
    team_id,
    folder_id,
    date_created,
    start_date,
    due_date,
    date_updated,
    private,
    archived,
    pinned,
    percent_complete
from
    clickup_goal
where 
    id = 'v47b5vb'
```

### List goals for a specific team

```sql
select
    id,
    pretty_id,
    name,
    team_id,
    folder_id,
    date_created,
    start_date,
    due_date,
    date_updated,
    private,
    archived,
    pinned,
    percent_complete
from
    clickup_goal
where
    team_id = '46446546'
```
