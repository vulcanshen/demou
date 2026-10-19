# demou 開發者備忘

demou 是家族的操作教學 app，遵循 [terminu design](https://github.com/vulcanshen/terminu) principle（tdp v0.8.0），
畫面用 [terminu-go](https://github.com/vulcanshen/terminu-go)。

## 運作方式

- `cmd/demou/main.go` 先看命令列：`version`、`help`（tdp E1）；沒有參數時開 `internal/ui` 做好的 terminu-go `App`。
- `internal/ui`：三個畫面（`move.go`、`actions.go`、`popups.go`）看的是同一個書架（`books.go`）。書架只在記憶體裡：書搬了、
  拿掉了、加了，`refresh` 把三個畫面的清單一起重畫。什麼都不存，每次啟動都是同一份。
- splash 畫 `internal/ui/icon.svg`，它是 `docs/icon.svg`（正本）的副本：`go:embed` 讀不到 `docs/`，測試守住兩份一樣
  （tdp E6、project 的「目錄」）。換 mark 時兩份一起換。
- 版本在 `internal/version`（tdp project 的範本）：本機 build 是 `dev`，goreleaser 發版時注入 tag 去掉 v 的版號。

## 設計決定

2026-10-18 跟 user 一題一題定的：

- **為什麼有 demou**：user 的 YouTube 影片每個成員分三段：成員自己的功能 → 家族共用基礎操作 → 成員功能細節。中間那段用 demou
  錄一次，每支影片剪同一段進去；看完影片的人也可以裝來自己按一遍。
- **假資料是一個書架**：20 本公版經典（作者都過世很久，沒有版權問題）、三個書架（To read 7、Reading 3、Done 10）、18 位作者。
  要中性，放進每個成員的影片都合適，也不能像任何一個成員的領域（檔案像 filu、host 像 sshu、pod 像 kbu）。內容用英文，跟家族的
  畫面一致。沒有隨機的東西，每次啟動都一樣，錄出來才一樣。
- **每章一個畫面**，用畫面 chip 列換：`[M]ove`、`[A]ctions`、`[P]opups`。影片的一章就是一個畫面，觀眾跳到哪一章，畫面上就只有
  那一章要講的東西；換章時按 chip 列的鍵，也順便示範了家族多畫面 app 的 chip 列。
  1. **MOVE（看與移動）**：`[1] SHELVES` 是三個書架；`[2] BOOKS` 是 cursor 那個書架上的書，跟著 cursor 即時換，另有 `AUTHORS` tab。
     兩個 panel 都有 cursor，都示範得到繞回。`[1]` 的 `Enter` 把 focus 送到 `[2]`（tdp K3 的明細）。
  2. **ACTIONS（找動作）**：一個 panel 列出全部的書。item 區 `[Enter] Open`、`[b]egin`（移到 Reading，只有 To read 的書能用）、
     `[f]inish`（移到 Done，只有 Reading 的書能用），其他書上是停用色，cursor 一移就看得到變化；panel 區 `[s]ort`（按年份，再按一次
     回到按書名）。
  3. **POPUPS（popup）**：`[r]emove` 先問（confirm 寫出書名）；`[n]ew book` 開 NEW BOOK 表單，Title 是文字、Author 是 select（18 位
     作者：dialog/form 規定超過五個選項用 select），兩欄都必填，Add 之後書放進 To read、cursor 停在新書上。「疊兩層」就是表單上的
     select。
- **三個畫面共用一個書架**：在 ACTIONS 搬的書，回到 MOVE 書架上的數字跟著變。一份假資料講完三章，觀眾不用每章重新認識內容；
  在一章做的事，換到另一章還看得到結果，像一個真的 app。
- **書的 `Enter` 在每個畫面都是 Open**，打開第一頁（tdp K3：同一種項目在同一個 app 裡永遠是同一個動作）。AUTHORS tab 上的是作者，
  不是書：`Enter` 打開作者的生卒年與書。在 demou 裡加的書沒有第一頁，Open 是停用色（tdp M6）。
- **Remove 的問句帶書名**：`Action.Ask` 是登記時給的，問句卻要跟著 cursor 換。做法是 Ask 用一個指標，cursor 換書時（`List.OnMove`）
  改它的 `Question`。不在 `Do` 裡自己開 confirm：那樣從 Space menu 執行時 menu 會先關掉，`Esc` 回不到 menu（tdp F4）；而 demou
  教的正是「`Esc` 一次關一層」。接受之後 menu 一起收掉：刪除讓那一項不在了（tdp T1 的表）。
- **從 Space menu 打開的書與作者，`Esc` 回到 menu**：Open 是「打開東西來看」，照 tdp T1 的表保留 source（`Action.Stay`）。
- **從 Space menu 開 NEW BOOK，menu 跟著收掉（暫定）**：新增不在 tdp T1 的表上，由 app 推演。加好之後 cursor 停在新書上（tdp K3），
  menu 的 item 區講的卻是原本 cursor 上那本書，留著只會對不上；所以表單打開時 menu 就收掉（v0.1.13 的 Action 沒有 `Stay` 就是這樣）。
  代價是在表單上 `Esc` 取消時回到 panel，不回到 menu。已回報（DM38）。
- **Title 是必填，在 text popup 就擋**：Title 的 text popup 有自己的檢查，空著按 `Enter` 寫出 `Title is required`、popup 不關（tdp K3）；
  表單的 Add 也檢查一次，給沒打開過 Title 就按 Add 的人。
- **最低尺寸 80×24（暫定）**：tdp L1 讓 app 把最低尺寸從預設的 80×40 調低。demou 選 80×24：新開的終端機視窗常常就是 80×24（macOS 的
  Terminal 預設），看完影片想跟著按的人一打開就能用，不必先拉大視窗；demou 的畫面最多兩個 panel、popup 都不大，畫面測試在 80×24
  跑過每一種 popup。比它小（例：40×12）時，等 lib 有最低尺寸，就只顯示名稱與 `needs a terminal of at least 80×24`。已回報（DM39）。
- **`q` 先問**：`Quit demou?`，每個畫面都一樣（tdp K9：離開要不要先問由 app 決定）。popup 那一章要示範「`q` 離開前先問」
  （user 定的章節內容）；看完影片自己按的人，也不會一按 `q` 就離開、還沒按完的書架從頭來過。
- **家族 mark**：user 用 kamonu 畫的，U 框裡上面 D、E，下面 M、O；`docs/icon.svg` 是 kamonu 匯出的檔，分字註解照 tdp S5
  （U D E M O）。tdp S5 要每個成員有自己的 mark，kamonu 是家族畫 mark 的工具，所以由 user 用它畫。
- **錄法**：YouTube 的影片用 vulaude 的 tape-video，一個畫面一章（三章），字幕寫在 tape 裡、錄的時候就畫在下方；尺寸與配色照 ccdrawu 的
  `_common.tape`（1920×960，字 18），剪進成員的影片時才一致。README 的 gif 照 tdp E6 用 VHS 錄（`make gif`）。user 原本選「VHS 錄、
  剪輯時加字幕」，看到 ccdrawu 已經用 tape-video 錄好 YouTube 之後改用它。
- **兩個 viewer 都有選取模式**：tdp v0.8.0 定了每個 viewer 都有（tdp F1，DM25 的定案）；lib 下一版讓每個 viewer 預設就有，複製也由
  lib 做（剪貼簿加 Success toast），demou 不另外接。v0.1.13 還沒有，寫在 `docs/demou-terminu-fix.md`。
- **CHANGELOG 不寫 splash**：tdp S2：splash 不寫進 README、CHANGELOG 與 release notes（DM35 的定案）。
- **License 是 GPL-3.0**：terminu-go 是 GPL-3.0。

## 已否決，不要重提

- 假資料用太陽系（只有 8 列，示範不到捲動；「新增行星」的表單不自然）或購物清單（內容單薄）。
- MOVE 用「左邊是書、右邊是第一頁」：右邊沒有 cursor，示範不到在第二個 panel 裡繞回。
- ACTIONS 用兩個 panel：看得到 Space menu 跟著 focus 換，但書架那個 panel 也要有自己的動作，畫面上的東西變多。
- NEW BOOK 表單多一欄 Shelf（radio）：超出已定的「一個文字、一個 select」，影片要多講一種操作。
- 影片只有一章「terminu basics」：觀眾不能直接跳到中間某一段。
- 手動錄影：tdp 改了就要整段重按，每次都不一樣。
- demou 自己在畫面上顯示剛按的鍵：家族 app 沒有這個東西，而且 terminu-go 的按鍵不會傳給 app。

## 已知的牆與未做

- 尚未符合 tdp 的地方逐條寫在 `docs/demou-terminu-fix.md`（tdp E7）：照 tdp v0.7.4 的清單對過，再照 v0.8.0 改了的地方對過，
  剩下的都出在 terminu-go v0.1.13 的行為，等 lib 下一版。程式裡的 TODO：
  - `TODO(lib DM4)`：最低尺寸以下到 1×1 的測試（tdp E5、L1）。
  - `TODO(lib DM7)`：空的 Title 按 `Enter` 的測試。v0.1.13 會擋；本機還沒發版的 lib 把它當成「值沒改」直接關，等 lib 照 tdp v0.8.0 的
    K3 改回來，兩邊都擋了才加測試（`make check` 兩邊都要過）。
  - `TODO(lib DM25)`：兩個 viewer 的選取模式。
- 規範的盲點、lib 要補的、清單的問題回報給 terminu 的 session（DM4 起），由它記進紀錄（`terminu 的筆記`）。
  DM4–DM37 都在 tdp v0.8.0 定了。
- 之前的 DM1–DM3，tdp v0.7.1 都定了：
  - DM1：`docs/icon.svg` 怎麼 embed。規範照 demou 的做法寫進 tdp E6 與 project 的「目錄」（見上面的「運作方式」）。
  - DM2：viewer 在字的中間折行。規範定了（components 的「折行」：在空白處斷），terminu-go v0.1.13 跟上了；demou 原本先用
    `terminu.Wrap` 把第一頁折成最寬 60 格的幾行，現在拿掉，第一頁是一行一段，由 viewer 照 popup 的寬度折。
  - DM3：select 要捲動時，寬度沒算進捲動位置，`Esc:cancel` 被捨掉（80×24 打開 Author 就會）。terminu-go v0.1.13 修好了，
    demou 不用改程式。

## 偏離 tdp

沒有。

## 設計文件導讀

沒有另外的設計文件：畫面與決定都寫在這一份。

## 建置與開發

- `make`（或 `make help`）列出所有指令；家族共用的 target 在 `family.mk`，app 自己的放在 `Makefile` 的最後一段。
- project 目錄圖以外的頂層檔案只有 `CLAUDE.md`：給 AI 協作的說明（demou 是什麼、規範在哪、盲點怎麼記、跟 user 合作的方式），
  由 terminu 的 session 寫。
- 本機用 `go.work` 指著 `../terminu-go`（git 不收）。`make check` 有 `go.work` 時會再關掉它、照 `go.mod` 的版本跑一次
  （`check-lib`）：過不了，就是用到了 lib 還沒發版的東西。
- **測試**（`internal/ui`）：`harness_test.go` 不開終端機跑 demou，送按鍵、跑完動畫、把畫面當文字讀出來（照 kamonu 的 harness）。
  - tdp L4 的畫面測試：5 種尺寸 × 每個畫面與每一種 popup（含疊在別的 popup 上的），再加「打開之後把終端機縮小」，讓 popup 比畫面寬、
    比畫面高、兩邊都大（tdp E5）；icon 與 powerline 符號一格、icon 兩格、兩種都兩格各跑一輪。
  - 只測整個畫面：單獨量每一種 popup 的框是 lib 的事，popup 是 lib 畫的（tdp E8）。以前 harness 從畫面找出 popup 的框來量
    （`popupBox`），tdp v0.8.0 定了分工之後拿掉。
  - 動畫在測試裡也是真的時間（一個 popup 288ms，lib 沒有關掉的開關），所以畫面測試平行跑；一輪 `go test` 約 40 秒。icon 寬度是
    lib 的全域設定，只在兩輪之間、沒有測試在跑時改。
- **錄影**：
  - README 的 gif：`make gif`（`.local/demos/demo.tape` → `docs/demo.gif`）。
  - YouTube 的影片：`node tape-video.mjs .local/demos/youtube [章號…]`，
    在 repo 根目錄跑。三個 tape（`01-move`、`02-actions`、`03-popups`）各自先 `make build`，在 tape-video 的 tmux 裡開 demou；給了章號就只
    重錄那幾章。成品在 `.local/demos/youtube/out/`（`video.mp4`、`chapters.txt`、`sheet.png`）。
  - 字幕寫成 `# caption: …`，接著 `Ctrl+B`、`Type "]"` 切換；`[鍵]` 畫成鍵帽，所以字幕裡不是鍵的字不能加方括號。

## 發布

照 tdp project 的「發布」：`make check` → CHANGELOG 的 `## Unreleased` 改成 `## vX.Y.Z — tdp vA.B.C` → commit、打 tag →
push commit 與 tag，`release.yml` 接著測試、goreleaser 發 release 與推 Homebrew tap。
