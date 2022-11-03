# Table: clickup_task_watcher

Obtain information about watchers of a specific task.

You **MUST** specify a `task_id` in the WHERE or JOIN clause.

### Get watchers for a specific task

```sql
select
    id,
    username,
    email,
    color,
    profile_picture,
    initials
from
    clickup_task_watcher
where
    task_id = '4g9milk'
```
