---
organization: The APS Group
category: ["saas"]
icon_url: "/images/plugins/theapsgroup/clickup.svg"
brand_color: "#7B68EE"
display_name: "ClickUp"
short_name: "clickup"
description: "Steampipe plugin for querying ClickUp Tasks, Lists and other resources."
og_description: Query ClickUp with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/theapsgroup/clickup-social-graphic.png"
---

# ClickUp + Turbot Steampipe

[ClickUp](https://clickup.com/) is a SaaS specialising in Project/Task Management similar to Jira or Asana.

[Steampipe](https://steampipe.io/) is an open source CLI for querying cloud APIs using SQL from [Turbot](https://turbot.com/)

## Documentation

- [Table definitions / examples](https://hub.steampipe.io/plugins/theapsgroup/clickup/tables)

## Getting Started

### Installation

```shell
steampipe plugin install theapsgroup/clickup
```

### Prerequisites

- ClickUp Account
- [ClickUp API Token](https://clickup.com/api/developer-portal/authentication#personal-token)

### Configuration

Configuration can take place in the config file (which takes precedence) `~/.steampipe/config/clickup.spc` or in Environment Variables.

Environment Variables:
- `CLICKUP_TOKEN` for the API token (ex: `pk_t348t9v3UYFG535ti`)

Configuration File:

```hcl
connection "clickup" {
  plugin  = "theapsgroup/clickup"
  token   = "pk_t348t9v3UYFG535ti"
}
```

### Testing

A quick test can be performed from your terminal with:

```shell
steampipe query "select * from clickup_task"
```
