# Table: clickup_folderless_list

Obtain information about lists that aren't associated to folders within your ClickUp environment.

However you **MUST** specify a `space_id` in the WHERE or JOIN clause.

## Examples

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
    clickup_folderless_list
where
    space_id = '7423465';
```
