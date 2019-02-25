# blossom

**blossom** is a **b**undle of **l**ightweight **o**pen-**s**ource **so**ftware **m**etrics. 
It is a tool that enables you to fetch [software metrics](https://opensource.guide/metrics/) 
from any of your organization's repository in Github.

## Setup

This application requires [go1.11](https://golang.org/doc/go1.11) in order to
work. First, clone the repository and enter it: 

```s
$ git clone git@github.com:ljvmiranda921/blossom.git
$ cd blossom
```

Then get the dependencies and build the project:

```s
$ go get
$ go build .
```

Optionally, you can also install `blossom` in your system:

```s
$ go install
```

## Usage

### Authentication

Because we rely on the [Github v4 GraphQL
API](https://developer.github.com/v4/), an [API
token](https://github.blog/2013-05-16-personal-api-tokens/) is required. We
recommend creating a `.env` file with the `GITHUB_TOKEN` variable:

```s
# .env
GITHUB_TOKEN=<your API token here>
```

All authentication is handled by `pkg/auth`.


