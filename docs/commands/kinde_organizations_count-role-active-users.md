## kinde organizations count-role-active-users

Get organization role active users count

### Synopsis

Get the number of active users that hold a given role within a specific
organization.

A user is counted as active if they were issued at least one access token during
the requested period, regardless of organization context on the token. Only
users who currently hold the role in the organization are included. The count is
scoped to the current environment.

Both `date_time_from` and `date_time_to` are required, inclusive, and must be
ISO 8601 datetimes in UTC. Provide them at second precision (no fractional
seconds). If a value includes fractional seconds, it is normalized before any
other processing: `date_time_from` is rounded down and `date_time_to` is rounded
up to the nearest second. Window validation, the active-user query, and the
echoed response bounds all use those normalized values.

The requested window must not exceed 3 days. Because both bounds are inclusive,
this means `date_time_to` must be earlier than `date_time_from` plus 3 days. For
example, `2026-07-01T00:00:00Z` to `2026-07-03T23:59:59Z` is a full 3 day window
covering the whole of 1, 2 and 3 July. A longer window is rejected with
`DATE_TIME_RANGE_TOO_LARGE`.

Required scope: read:organization_user_roles

API: GET /api/v1/organizations/{org_code}/roles/{role_id}/active_users/count

```
kinde organizations count-role-active-users [flags]
```

### Options

```
      --date-time-from string   Start of the active period (inclusive), as an ISO 8601 datetime in UTC at second precision (RFC 3339 timestamp, e.g. 2026-01-31T09:00:00Z; required)
      --date-time-to string     End of the active period (inclusive), as an ISO 8601 datetime in UTC at second precision (RFC 3339 timestamp, e.g. 2026-01-31T09:00:00Z; required)
  -h, --help                    help for count-role-active-users
      --org-code string         The organization's code (required)
      --role-id string          The role's public id (required)
```

### Options inherited from parent commands

```
      --domain string   Kinde domain to use for this command (default: the domain you last logged in to; or set KINDE_DOMAIN)
```

### SEE ALSO

* [kinde organizations](kinde_organizations.md)	 - Manage organizations, their users and settings

