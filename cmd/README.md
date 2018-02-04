# Commands

This package provides implementation of ghctl commands. ghctl is using
[Cobra][0] for CLI framework.

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [Commands](#commands)
    - [Global](#global)
        - [Commands](#commands-1)
        - [Options](#options)
    - [Context](#context)
        - [Create Context](#create-context)
        - [Get Context](#get-context)
        - [Current Context](#current-context)
    - [Repository](#repository)
    - [Issue](#issue)
    - [Pull-request](#pull-request)
    - [Notification](#notification)

<!-- markdown-toc end -->

## Global

### Commands

| COMMANDS         | SHORT NAME   | DESCRIPTION                           |
| :--------------- | ------------ | :------------------------------------ |
| context          | ctx          | see [Context](#context)               |
| repository       | repo         | see [Repository](#repository)         |
| issue            | -            | see [Issue](#issue)                   |
| pull-request     | pr           | see [Pull-Request](#pull-request)     |
| notification     | notif        | see [Notification](#notification)     |

### Options

| OPTIONS    | ARGS                | DESCRIPTION                                        |
| :--------- | ------------------- | :------------------------------------------------- |
| ghconfig   | filePath <string>   | Path to the configuration file to use for ghctl.   |

## Context

The context command manage context data for accessing to GitHub API. 

```
ghctl context [command]
```

| COMMANDS | description                     |
| :------- | :-----------                    |
| create   | see [Create](#create-context)   |
| get      | see [Get](#get-context)         |
| current  | see [Current](#current-context) |

### Create Context

```
ghctl context create [options] <access_token>
```

`create` command save new context data to the configuration file. This command
don't handle the context authentication. You have to get the access token manually.

| OPTIONS    | ARGS   | DEFAULT                        | DESCRIPTION                                                        |
| :--------- | ------ | ------------------------------ | :----------------------------------------------------------------- |
| endpoint   | string | https://api.github.com/graphql | GitHub API endpoint. (Only GraphQL API is supported)               |
| name       | string | `GitHub contextname`           | Context identification.                                            |

### Get Context

```
ghctl context get [options] [context_name]
```

`get` command outputs context data. If no arguments, all contexts are printed.
Otherwise, specified context is printed. Also, you can chose output format.

| OPTIONS    | ARGS      | DEFAULT                        | DESCRIPTION                                                        |
| :--------- | ------    | ------------------------------ | :----------------------------------------------------------------- |
| output     | json/yaml | yaml                           | Output format specification.                                       |

### Current Context

```
ghctl context current [options]
```

`current` command displays the current context for calling GitHub API.
By setting the switch option, you can switch the current context.

| OPTIONS    | ARGS   | DEFAULT                        | DESCRIPTION                                                        |
| :--------- | ------ | ------------------------------ | :----------------------------------------------------------------- |
| switch     | string | `context name`                 | Switch current context.                                            |
| simple     | bool   | false                          | Displays only current context name without description.            |

## Repository

```
// TBD
```

## Issue

```
// TBD
```

## Pull-request

```
// TBD
```

## Notification

```
// TBD
```

[0]: https://github.com/spf13/cobra
