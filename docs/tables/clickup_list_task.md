# Table: clickup_list_task

Obtain information about tasks assigned to a specific list within your ClickUp environment.

However you **MUST** specify either an `id` (single) or `list_id` (for multiple tasks) in the WHERE or JOIN clause.

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
  clickup_list_task
where
  id = '69xca6m';
```

### List all tasks for a specific list

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
  clickup_list_task
where
  list_id = '194506756';
```
