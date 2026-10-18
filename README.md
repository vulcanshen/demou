# demou

[![Release](https://img.shields.io/github/v/release/vulcanshen/demou)](https://github.com/vulcanshen/demou/releases)
[![License](https://img.shields.io/github/license/vulcanshen/demou)](LICENSE)

**Language**: English · [繁體中文](README-zh_TW.md)

**The terminu family's app for learning the operations, hands-on.**

<!-- What it does, in one paragraph. -->

<!-- One demo gif, recorded with VHS from .local/demos/demo.tape. -->

## Features

<!-- What the user gets, not how it is done. -->

## Install

### Requirements

- macOS or Linux (WSL too).
- A [Nerd Font](https://www.nerdfonts.com) terminal profile: the icons and the frames are Nerd Font glyphs.
- A truecolor (24-bit) terminal: the faded colours and the layer colours cannot be told apart in 256 colours.

demou measures how many cells an icon and a powerline symbol take when it starts. When it gets them wrong, set them:

| Variable | What it sets |
|---|---|
| `DEMOU__ICON_WIDTH` | The cells an icon takes, for demou |
| `DEMOU__POWERLINE_WIDTH` | The cells a powerline symbol takes, for demou |
| `TERMINU__ICON_WIDTH` | The cells an icon takes, for every terminu family app: set it once and they all read it |
| `TERMINU__POWERLINE_WIDTH` | The cells a powerline symbol takes, for every terminu family app |

Run inside a family app's terminal, demou gets the `TERMINU__` ones from the app around it.

### Homebrew

```sh
brew install vulcanshen/tap/demou
```

### Script

```sh
curl -fsSL https://raw.githubusercontent.com/vulcanshen/demou/main/install.sh | sh
```

No sudo: it installs to `~/.local/bin`.

### Go

```sh
go install github.com/vulcanshen/demou/cmd/demou@latest
```

<!-- What happens the first time demou starts. -->

### Uninstall

```sh
brew uninstall demou
```

Installed with the script:

```sh
curl -fsSL https://raw.githubusercontent.com/vulcanshen/demou/main/uninstall.sh | sh
```

## Quick start

<!-- The first few steps. -->

## Usage

```
demou            start
demou version    print the version (-v, --version)
demou help       print the command-line usage (-h, --help)
```

<!-- The keys. -->

## Settings and data

<!-- Where demou keeps its settings and data, and the family setting keys it uses. -->

## Limits

<!-- In the user's words. -->

## Links

- [CHANGELOG](CHANGELOG.md)
- [Developer notes](docs/dev-remarks.md) (in Traditional Chinese)

## The terminu family

demou follows [terminu design](https://github.com/vulcanshen/terminu) and is built with
[terminu-go](https://github.com/vulcanshen/terminu-go). The other apps in the family:

| App | What it is |
|---|---|
| [kbu](https://github.com/vulcanshen/kbu) | A Kubernetes TUI dashboard |
| [filu](https://github.com/vulcanshen/filu) | A terminal file manager you don't have to learn |
| [sshu](https://github.com/vulcanshen/sshu) | A terminal front end for ssh and sftp |
| [webu](https://github.com/vulcanshen/webu) | A terminal browser that reads a web page as a document |
| [locku](https://github.com/vulcanshen/locku) | A screensaver with a PIN lock, for the terminal |
| [kamonu](https://github.com/vulcanshen/kamonu) | A terminal editor for the family's marks (kamon), drawn letter by letter |
| [ccdrawu](https://github.com/vulcanshen/ccdrawu) | A sidebar for Claude Code: usage, context, cost and git beside the transcript |

## License

[GPL-3.0](LICENSE)
