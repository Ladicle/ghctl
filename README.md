# ghctl

The ghctl CLI tool for GitHub notifications, issue and pull-requests.

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [ghctl](#ghctl)
    - [Install](#install)
    - [Usage](#usage)
    - [Quick Starts](#quick-starts)
    - [Complete documentation](#complete-documentation)

<!-- markdown-toc end -->

## Install

```
// TBD
```

## Quick Starts

<!-- TODO pasete a nice screen-animation :) -->

### Add a new context

First of all, please generate a private access-token from [here](https://github.com/settings/tokens).
OAuth scopes require only `repo` and `read:org`. After that run the following command using the token and create a new context.

```
❯❯❯ ghctl ctx create <ACCESS TOKEN>
Register <NAME> context is successfully.
```

### Get Pull-Requests

```
❯❯❯ ghctl pr get --repo-match "status = open & count(LGTM*) = 5" --pr-match "status = open" Ladicle/*

REPOSITORY   TITLE     LABELS                   URL
dotfiles     piyoyoy   LGTM-ladicle,waitReview  http://github.com/Ladicle
emacs.d      hogegho                            http://github.com/Ladicle
```

## Complete documentation

Plunge into the [commands](./cmd/README.md) or [packages](./pkg/README.md).
