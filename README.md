# ghctl

The ghctl CLI tool for GitHub notifications, issue and pull-requests.

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [ghctl](#ghctl)
    - [Install](#install)
    - [Quick Starts](#quick-starts)
        - [Add a new context](#add-a-new-context)
        - [Get Pull-Requests](#get-pull-requests)
    - [Complete documentation](#complete-documentation)

<!-- markdown-toc end -->

## Install

Clone this repository, and run the following make command.

```
❯❯❯ make install
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

## Complete documentation

Plunge into the [commands](./cmd/README.md) or [packages](./pkg/README.md).
