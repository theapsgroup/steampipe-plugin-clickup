# Table: clickup_list

Obtain information about lists within your ClickUp environment.

However you **MUST** specify either an `id` (single) or `folder_id` (for multiple tasks) in the WHERE or JOIN clause.

> Note: by default `archived` items won't be returned, to return archived items only set `archived = true` in the where clause.

## Examples

### Get a list by id

```sql
select
  id,
  name,
  order_index,
  content,
  status,
  priority,
  assignee,
  task_count,
  due_date,
  start_date,
  folder,
  space,
  archived
from
  clickup_list
where
  id = '6fs7dfm';
```

### List lists for a specific folder

```sql
select
  id,
  name,
  order_index,
  content,
  status,
  priority,
  assignee,
  task_count,
  due_date,
  start_date,
  folder,
  space,
  archived
from
  clickup_list
where
  folder_id = 's6fsd8fds';
```
