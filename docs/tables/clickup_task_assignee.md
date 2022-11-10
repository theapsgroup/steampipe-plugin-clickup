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

### List assignees for a selection of tasks

```sql
with some_tasks as
(
  select
    t.id
  from
    clickup_task t
  where
    t.list_id = 'xczx34'
  order by
    id desc
  limit
    10
)
select
  id,
  username,
  email,
  color,
  profile_picture,
  initials,
  task_id
from
    clickup_task_assignee a
left join
  some_tasks
on
  a.task_id = some_tasks.id
```
