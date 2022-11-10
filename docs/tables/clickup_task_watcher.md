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

### List watchers for a selection of tasks

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
  clickup_task_watcher w
left join
  some_tasks
on
  w.task_id = some_tasks.id
```
