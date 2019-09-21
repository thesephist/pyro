# Pyro

CLI and Go package to check if any of the routes in a list of critical routes of an app are failing.

## CLI

- `pyro help`
- `pyro ok <url>`
- `pyro check <url> <status>`
- `pyro test <suite file>`

### Flags
- `--silent, -S`: nothing to stdout, the return code will report if any routes failed
- `--no-color, -N`: output will not contain ANSI color text
- `--timeout, -T`: set HTTP timeout, in seconds. Defaults to 5.

## Configuration

## Go API

