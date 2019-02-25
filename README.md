# blossom

**blossom** is a tool that enables you to get [open-source software
metrics](https://opensource.guide/metrics/) from any of your organization's
repository.

## Authentication

This utility relies on the [Github v4 GraphQL API](https://developer.github.com/v4/),
and requires an [API
token](https://github.blog/2013-05-16-personal-api-tokens/) to access. We
recommend creating a `.env` file with the `GITHUB_TOKEN` variable:

```
# .env
GITHUB_TOKEN=<your API token here>
```

All authentication is handled by `pkg/auth`.
