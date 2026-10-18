# demou

[![Release](https://img.shields.io/github/v/release/vulcanshen/demou)](https://github.com/vulcanshen/demou/releases)
[![License](https://img.shields.io/github/license/vulcanshen/demou)](LICENSE)

**Language**: [English](README.md) · 繁體中文

**家族的操作教學 app：動手按，學會 terminu 家族的操作。**

<!-- 一段「它能做什麼」。 -->

<!-- 一張 demo gif，用 VHS 照 .local/demos/demo.tape 錄。 -->

## 特色

<!-- 使用者拿到什麼，不寫怎麼做到的。 -->

## 安裝

### 前置需求

- macOS 或 Linux（WSL 也可以）。
- [Nerd Font](https://www.nerdfonts.com) 的終端機設定檔：icon 與框線都是 Nerd Font 的字。
- truecolor（24-bit 色）的終端機：淡化的顏色與層色在 256 色下分不出來。

demou 啟動時會量 icon 與 powerline 符號佔幾格。量錯時自己設：

| 變數 | 設的是 |
|---|---|
| `DEMOU__ICON_WIDTH` | icon 佔幾格，只給 demou |
| `DEMOU__POWERLINE_WIDTH` | powerline 符號佔幾格，只給 demou |
| `TERMINU__ICON_WIDTH` | icon 佔幾格，給 terminu 家族每一個 app：設一次，全家族都讀到 |
| `TERMINU__POWERLINE_WIDTH` | powerline 符號佔幾格，給 terminu 家族每一個 app |

在家族 app 的終端機裡跑時，`TERMINU__` 的兩個由外層的 app 替它設好。

### Homebrew

```sh
brew install vulcanshen/tap/demou
```

### 安裝腳本

```sh
curl -fsSL https://raw.githubusercontent.com/vulcanshen/demou/main/install.sh | sh
```

不用 sudo：裝到 `~/.local/bin`。

### Go

```sh
go install github.com/vulcanshen/demou/cmd/demou@latest
```

<!-- 第一次啟動會發生什麼。 -->

### 移除

```sh
brew uninstall demou
```

用安裝腳本裝的：

```sh
curl -fsSL https://raw.githubusercontent.com/vulcanshen/demou/main/uninstall.sh | sh
```

## 快速開始

<!-- 頭幾步。 -->

## 使用方式

```
demou            啟動
demou version    印出版號（-v、--version）
demou help       印出命令列的用法（-h、--help）
```

<!-- 按鍵。 -->

## 設定與資料

<!-- demou 把設定與資料放在哪裡，用到哪些家族共用的設定 key。 -->

## 限制

<!-- 用使用者的語言寫。 -->

## 相關連結

- [CHANGELOG](CHANGELOG.md)
- [開發者備忘](docs/dev-remarks.md)

## terminu family

demou 遵循 [terminu design](https://github.com/vulcanshen/terminu)，用 [terminu-go](https://github.com/vulcanshen/terminu-go)
做成。家族的其他成員：

| App | 是什麼 |
|---|---|
| [kbu](https://github.com/vulcanshen/kbu) | Kubernetes TUI dashboard |
| [filu](https://github.com/vulcanshen/filu) | 不用學的 terminal 檔案管理員 |
| [sshu](https://github.com/vulcanshen/sshu) | ssh 與 sftp 的 terminal 前端 |
| [webu](https://github.com/vulcanshen/webu) | 把網頁當文件讀的 terminal 瀏覽器 |
| [locku](https://github.com/vulcanshen/locku) | 帶 PIN 鎖的 terminal 螢幕保護程式 |
| [kamonu](https://github.com/vulcanshen/kamonu) | 一個字一個字畫家族 mark（家紋）的 terminal 編輯器 |
| [ccdrawu](https://github.com/vulcanshen/ccdrawu) | Claude Code 的側邊欄：用量、context、花費與 git 放在對話旁邊 |

## License

[GPL-3.0](LICENSE)
