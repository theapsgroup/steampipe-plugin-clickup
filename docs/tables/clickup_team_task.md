# Table: clickup_team_task

Obtain information about tasks assigned to a specific team within your ClickUp environment.

However you **MUST** specify either an `id` (single) or `team_id` (for multiple tasks) in the WHERE or JOIN clause.

## Examples

### Get a task by id

```sql
select
  id,
  name,
  description,
  creator,
  status,
  priority
from
  clickup_team_task
where
  id = '69xca6m';
```

### List all tasks for a specific team

```sql
select
  id,
  status,
  date_created,
  date_closed,
  due_date,
  team_id,
  project_id
from
  clickup_team_task
where
  team_id = '2506756';
```

### Obtain tasks for my team that're of a specific status

```sql
select
  id,
  status,
  date_created,
  date_closed,
  due_date,
  team_id,
  space_id,
  list_id,
  folder_id,
from
  clickup_team_task
where
  team_id = '2506756'
and
  status = 'planned';
```
