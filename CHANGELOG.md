# Changelog

Each release names the version of terminu design (tdp) it follows.

## Unreleased

- Built on terminu-go v1.1.0-beta.3, following tdp v1.1.0-beta.3
- The release templates from tdp v1.1.0-beta.2: a beta tag (`vX.Y.Z-beta.N`) makes a GitHub pre-release that the
  Homebrew tap does not get, and the release notes are written outside the checkout
- goreleaser from the tdp v1.1.0-beta.3 template: the ldflags that stamp the version take the module path from
  `go.mod`
- Claude Code's files stay out of git: `.gitignore` from the tdp v1.1.0-beta.1 template, and git no longer tracks
  `CLAUDE.md`; harness's `.claude/rules/terminu.md` stays in
- The README's terminu family section links tdp, terminu-go and terminu-harness instead of listing the other members

## v1.0.0 — tdp v1.0.0

- The repo, from the tdp v1.0.0 project templates: `family.mk`, the release workflow, goreleaser, `install.sh` and
  `uninstall.sh`, `internal/version`; `demou version` and `demou help`
- Three screens over one bookshelf of twenty classics: `[M]ove` (shelves and their books, with an Authors tab), `[A]ctions`
  (open, begin, finish, sort) and `[P]opups` (remove, after asking; the NEW BOOK form with a select over it); `q` asks
  first
- On a first page or an author, `v` selects and `y` copies to the terminal's clipboard
- `!` opens the App Log: the screens switched to, the pages and authors opened, the books moved, removed, added and
  sorted, newest last
- Needs a terminal of at least 60×24; a smaller one shows demou's name, the size it needs and how to quit
- demou's mark, `docs/icon.svg`
