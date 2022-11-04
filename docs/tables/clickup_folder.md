# Table: clickup_folder

Obtain information about folders within your ClickUp environment.

However you **MUST** specify either an `id` (single) or `space_id` (for multiple tasks) in the WHERE or JOIN clause.

## Examples

### Get a folder by id

```sql
select
    id,
    name,
    order_index,
    hidden,
    space_id,
    task_count,
    archived,
from
    clickup_folder
where
    id = '7fsd72'
```

### List all folders for a specific space

```sql
select
    id,
    name,
    order_index,
    hidden,
    space_id,
    task_count,
    archived,
from
    clickup_folder
where
    space_id = '54649813'
```
