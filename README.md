# Pyro

CLI and Go package to check if any of the routes in a list of critical routes of an app are failing.

Pyro was designed to be invoked in a test runner or uptime checker, so that rather than simply checking that a server pings back or that a single route works, the test runner / checker can verify that all critical routes of an application render or resolve to expected HTTP status codes. Pyro simplifies that check to a simple CLI invocation with a human-friedly test spec file.

## CLI

- `pyro help`
- `pyro check <url> <status>`
- `pyro test <suite file>`

### Flags

* flags: https://github.com/spf13/cobra

- `--silent, -S`: nothing to stdout, the return code will report if any routes failed
- `--no-color, -N`: output will not contain ANSI color text
- `--timeout, -T`: set HTTP timeout, in seconds. Defaults to 5.

## Todo items
- [ ] Cookie authentication

