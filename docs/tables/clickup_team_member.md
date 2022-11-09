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
