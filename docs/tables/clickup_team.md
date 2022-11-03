# Table: clickup_team

Obtain information about teams in your ClickUp environment.

## Examples

### List all teams

```sql
select
    id,
    name,
    color
from
    clickup_team;
```

### List all users for all teams

```sql
select
    t.name as team,
    u.username,
    u.email,
    u.last_active
from
    clickup_team t
left join
    clickup_team_member u
on t.id = u.team_id
```
