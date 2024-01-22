# github-runner-check

Command-line tool to check Github self-hosted runners.

## Usage

The command `github-runner-check` can be called with or without parameters. It scans default directory recursively to find `svc.sh` scripts and executes them with `status` command.

### Parameters

| Name      | Values | Description |
| ----------- | ----------- | --- |
| `-p`, `--path`  | Example: `/path/to/runners` | Path to the directory which contains runner folder(s) |
