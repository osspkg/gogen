# Contributing

## Prerequisites

- Go 1.26 or newer
- `goppy` for the repository's Makefile tasks

## Checks

Run checks from the repository root:

```sh
make tests
make lint
make build
```

Before opening a pull request, run the relevant checks and include the user-visible behavior and validation results in the description. Keep changes focused and add regression tests for behavior changes.

`make ci` runs the same workflow used by GitHub Actions. It also installs `goppy@latest` and invokes `goppy setup-lib`, which may require network access and modify local setup.
