# ttrk

`ttrk` is a keyboard-first time tracker that runs entirely in your terminal. It provides a Bubble Tea powered TUI for organising projects, capturing tasks, and logging focused work sessions without leaving the command line.

## Features

- Create projects and attach as many tasks as you need.
- Start and stop timers to log work; every session is stored automatically.
- Review per-task and per-project totals without leaving the keyboard.
- Data is kept in a simple JSON file under your config directory, making backups straightforward.
- Ships with Makefile helpers for building, linting, testing, and releasing.

## Installation

### Using `go install`

```bash
go install github.com/MBH999/ttrk/cmd/ttrk@latest
```

The binary will be placed in `$GOBIN` (or `$GOPATH/bin` if `GOBIN` is unset).

### From source

```bash
git clone https://github.com/MBH999/ttrk.git
cd ttrk
make build
```

The compiled binary is written to `bin/ttrk`.

## Quick start

1. Run `./bin/ttrk` (or `go run ./cmd/ttrk`) to launch the interface.
2. Use the arrow keys or `j/k` to move around the menus.
3. Pick **New Project** to add a project, then choose it and add tasks.
4. Press `enter` on a task to start tracking time and `space` or `enter` again to log the session.

A `debug.log` file is created alongside the binary with Bubble Tea's debug output. This can be helpful when reporting issues.

## Keyboard reference

| Screen | Keys | Action |
| --- | --- | --- |
| Any | `ctrl+c` | Quit immediately |
| Home | `↑/↓` or `j/k` | Move the selection |
| Home | `enter` | Activate the highlighted option |
| Project list | `enter` | Open the highlighted project |
| Project list | `a` or `n` | Create a project |
| Project/Task lists | `esc` or `backspace` | Return to the previous screen |
| Task list | `enter` | Start the timer for the selected task |
| Task list | `a` or `n` | Add a task to the active project |
| Timer | `space` or `enter` | Stop and log the running session |
| Timer | `esc` | Cancel the running session without logging |

## Data and configuration

`ttrk` looks for its configuration in `~/.config/ttrk`. The directory is created automatically on first run.

- `config.ini` (optional): customise settings, currently only `data_dir` is supported. Example:
  ```ini
  # ~/.config/ttrk/config.ini
  data_dir = ~/Dropbox/ttrk-data
  ```
- `data.json`: generated automatically and stores all projects, tasks, and timer entries.

If you set `data_dir`, `ttrk` will keep `data.json` inside that directory.

## Development

- Go 1.22 or newer
- `make build` – build the binary
- `make test` – run the test suite
- `make lint` – run `golangci-lint` (must be installed locally)
- `make run` – rebuild and start the TUI

Before opening a pull request, run `make all` to execute the common checks.

---

Happy tracking!
