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

### Getting open-source software metrics

Most of the metrics here were based from [CHAOSS's open-source software
metrics](https://github.com/chaoss/metrics). To obtain them for a particular
repository, simply run the following command:

```s
$ blossom report <owner>/<repository>
```

For example,

```s
$ blossom report ljvmiranda921/blossom
```

So far, I am still updating the metrics that we have. If you would like to help
out, then I'd definitely appreciate it!

## Contributing

Simply fork this repository and [make a Pull
Request](https://help.github.com/en/articles/creating-a-pull-request)! I'm
open to any kind of contribution, but I'd definitely appreciate:

- Implementation of some open-source software metrics
- Proposal for new metrics
- Documentation
- Testing

Also, we have a
[CONTRIBUTING.md](https://github.com/ljvmiranda921/blossom/blob/master/CONTRIBUTING.md)
and a [Code of
Conduct](https://github.com/ljvmiranda921/blossom/blob/master/CODE_OF_CONDUCT.md),
so please check that one out!

## FAQ

- **Why make this if there's already CHAOSS?** The CHAOSS toolset is really
    good and comprehensive, but I want something that's more lightweight and
    "atomic." I want to immediately get software metrics without setting up my
    own webserver etc., so this fits into my use-case.
- **Can I already use this in my organization?** This is still experimental and
     not yet production-ready. I also don't know (yet) how this can be
     deployed. If you've successfully used this in any of your projects then
     feel free to ping me!
- **Your Golang skills needs improvement** I know, I'm still learning Go so
    please bear with me. If there are any code smells, please don't hesitate to
     make a PR!
