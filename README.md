# ClickUp plugin for Steampipe

Use SQL to instantly query ClickUp tasks, goals and more. Open source CLI. No DB required.

- **[Get started ->](https://hub.steampipe.io/plugins/theapsgroup/clickup)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/theapsgroup/clickup/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/theapsgroup/steampipe-plugin-clickup/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install theapsgroup/clickup
```

Setup the configuration:

```shell
vi ~/.steampipe/config/clickup.spc
```

or set the following Environment Variables

- `CLICKUP_TOKEN` : The API Key / Token to use.

Run a query:

Interactive Mode:
```sql
select
  *
from
  clickup_team;
```

or from CLI:
```shell
steampipe query "select * from clickup_team"
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)
- [ClickUp API Token](https://clickup.com/api/developer-portal/authentication#personal-token)

Clone:

```sh
git clone https://github.com/theapsgroup/steampipe-plugin-clickup.git
cd steampipe-plugin-clickup
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```shell
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/clickup.spc
```

Try it!

```
steampipe query
> .inspect clickup
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)
