# Packages

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [Packages](#packages)
    - [Configuration](#configuration)
        - [Configuration format](#configuration-format)
        - [Directory architecture](#directory-architecture)

<!-- markdown-toc end -->

## Configuration

`config` package manage configurations for ghctl.
The configuration directory locates `$HOME/.ghctl/config` in default.
If you set `--ghconfig`, ghctl use the path.

### Configuration format

The configuration file write in the YAML style.

```yaml
current-user: [name]
users:
  - name: <string>
    token: <string>
```

### Directory architecture

```bash
❯❯❯ tree .ghctl/
.ghctl/
├── cache/
└── config

1 directory, 1 file
```

`.ghctl` contains configuration file and `cache` directory that stores
ecache of repositories, issues and pull-requests.
