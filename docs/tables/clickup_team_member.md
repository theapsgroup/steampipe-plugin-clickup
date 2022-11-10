# Table: clickup_team_member

Obtain information about members of a specific team.

You **MUST** specify a `team_id` in the WHERE or JOIN clause.

## Examples

### Get members for a specific team

```sql
select
  id,
  username,
  email,
  color,
  profile_picture,
  initials,
  last_active,
  date_joined,
  date_invited
from
  clickup_team_member
where
  team_id = 21596865;
```

### List all members for all teams you have access to

```sql
select
  t.id as team_id,
  t.name as team_name,
  t.color as team_color,
  m.id as member_id,
  m.username as member,
  m.email as member_email
from
  clickup_team t
left join
  clickup_team_member m
on
  t.id = m.team_id
```
