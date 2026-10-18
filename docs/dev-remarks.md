# demou 開發者備忘

demou 是家族的操作教學 app，遵循 [terminu design](https://github.com/vulcanshen/terminu) principle（tdp v0.7.0），
畫面用 [terminu-go](https://github.com/vulcanshen/terminu-go)。

## 運作方式

- `cmd/demou/main.go` 先看命令列：`version`、`help`（Rules E1）；沒有參數時開 terminu-go 的 `App`。
- 版本在 `internal/version`（tdp project 的範本）：本機 build 是 `dev`，goreleaser 發版時注入 tag 去掉 v 的版號。

## 設計決定

- **repo 照 tdp v0.7.0 的 project 從零建**：demou 是 project 的參考實作。範本檔（`family.mk`、`release.yml`、
  `.goreleaser.yaml`、`install.sh`、`uninstall.sh`、`internal/version/version.go`、`.gitignore`）照原樣複製，只改 project
  「範本檔」表最後一欄的地方；用起來的問題回報成 tdp 的缺漏，不在這裡改範本。
- **License 是 GPL-3.0**：terminu-go 是 GPL-3.0。

## 已否決，不要重提

## 已知的牆與未做

- 還沒符合 tdp 的地方逐條寫在 [`demou-terminu-fix.md`](demou-terminu-fix.md)。

## 偏離 tdp

沒有。

## 設計文件導讀

- [`demou-terminu-fix.md`](demou-terminu-fix.md)：還沒符合 tdp 的地方。

## 建置與開發

- `make`（或 `make help`）列出所有指令；家族共用的 target 在 `family.mk`，app 自己的放在 `Makefile` 的最後一段。
- 本機用 `go.work` 指著 `../terminu-go`（git 不收）。`make check` 有 `go.work` 時會再關掉它、照 `go.mod` 的版本跑一次
  （`check-lib`）：過不了，就是用到了 lib 還沒發版的東西。

## 發布

照 tdp project 的「發布」：`make check` → CHANGELOG 的 `## Unreleased` 改成 `## vX.Y.Z — tdp vA.B.C` → commit、打 tag →
push commit 與 tag，`release.yml` 接著測試、goreleaser 發 release 與推 Homebrew tap。
