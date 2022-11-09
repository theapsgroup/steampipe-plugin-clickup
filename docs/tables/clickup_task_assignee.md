# Table: clickup_task_assignee

Obtain information about assignees of a specific task.

You **MUST** specify a `task_id` in the WHERE or JOIN clause.

## Examples

### Get assignees for a specific task

```sql
select
  id,
  username,
  email,
  color,
  profile_picture,
  initials
from
  clickup_task_assignee
where
  task_id = '4g9milk'
```
