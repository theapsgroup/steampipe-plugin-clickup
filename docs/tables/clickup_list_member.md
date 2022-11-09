# Table: clickup_list_member

Obtain information about members of a list within your ClickUp environment.

You **MUST** specify a `list_id` in the WHERE or JOIN clause.

## Examples

### List all members of a specific list

```sql
select
  id,
  username,
  email,
  color,
  profile_picture,
  initials
from
  clickup_list_member
where
  list_id = '9429v3v2390'
```
