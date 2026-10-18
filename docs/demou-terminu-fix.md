# demou 尚未符合 tdp 的地方

照 tdp v0.7.4 的條目清單（`checklist/items-zh_TW.md`）對過，下面是還沒符合的條目。都出在 terminu-go v0.1.13 的行為，demou 這邊
改不了；已經用 DM 編號回報給 terminu 的 session，lib 修好、`go.mod` 跟上之後逐條刪掉，全部修完就刪掉這個檔（tdp E7）。

## E5-10　popup 開著時把終端機縮到 5 列以下，app 當掉（DM4）

- **違反**：tdp E5「疊 popup 不能當掉……不可以 panic」。
- **在哪裡**：terminu-go 的 `menuView.draw`（`menu.go`）、`noteView.draw`（`note.go`）、`formView.draw`（`form.go`）、finder 的 draw
  （`finder.go`）用 `h-4` 當 slice 的容量，h 小於 4 時是負的。
- **現況**：開著 Space menu、`?`、NEW BOOK 表單（含上面的 text、select、confirm），把終端機縮到 80×5、80×1、1×1，demou 印出
  `Caught panic: runtime error: makeslice: cap out of range` 後結束。80×5 啟動再按 `Space` 也一樣。6 列以上正常。
- **該怎麼改**：lib 把容量改成 `max(h-4, 0)`。修好之後，demou 的 `TestE5PopupBiggerThanTheScreen` 加上 80×5、1×1
  （`TODO(lib DM4)`）。

## F7-6　終端機 10 列以下時，NEW BOOK 表單比畫面高度 − 2 還高，蓋住 footer（DM5）

- **違反**：tdp F7「最高是畫面高度 − 2：畫面 chip 列或 statusbar 與 footer 永遠看得到」。
- **在哪裡**：terminu-go `form.go`：`measure` 把高度限在畫面高度 − 2，`draw` 卻一定畫一列以上的欄位、空列、錯誤列、分隔線與按鈕。
- **現況**：80×10 開 NEW BOOK，框畫了 9 列，下框壓在最後一列、footer 不見；80×8 連下框都看不到。80×11 起正常。
- **該怎麼改**：lib 的 draw 照 measure 給的高度畫，放不下就捲。修好之後，demou 的 L4 測試在 24×8 也量 popup 自己的框
  （`TODO(lib DM5)`）。

## popup-50　終端機 8 列時，Title 的 text popup 看不到打字列（DM6）

- **違反**：components/layout/popup「內容放不下時：跟著 focus 捲……讓 focus 所在的那一個整個看得到」。
- **在哪裡**：terminu-go `input.go` 的 `lineInput.draw` 用 `Pad2(body, inner, h-4)` 從下面截掉，沒有捲到 focus。
- **現況**：80×8 打開 NEW BOOK 的 Title，只看得到說明列，打字列與錯誤列被截掉，打的字看不到。80×9 起正常。
- **該怎麼改**：lib 讓 text popup 放不下時捲到打字列。

## form-84、form-82　Title 打了字又刪掉，表單就算改過（DM7）

- **違反**：components/dialog/form「只有 input popup 確定了跟打開時不同的值，才算改過」；沒改過的表單 `Esc` 直接關。
- **在哪裡**：terminu-go `input.go` 的 `lineInput.confirm`：打過字（`touched`）就呼叫 `Done("")`，表單因此設成改過。
- **現況**：`P` `n` `Enter`、打 `a`、`Backspace`、`Enter`，Title 還是 `not set`，按 `Esc` 卻跳出 `Discard changes to New book?`。
- **該怎麼改**：lib 拿確定的值跟打開時的值比，一樣就只關 popup、不呼叫 `Done`。

## note-42　viewer 的 key reference 把書名轉成大寫（DM8）

- **違反**：components 的「大寫」：標題裡的名稱照原樣（components-21）；note 的 key reference 標題是 `<surface> keys`。
- **在哪裡**：terminu-go `note.go` 的 `keyReference` 把整個 surface 轉大寫，viewer 的 surface 含 Subject。
- **現況**：打開 Pride and Prejudice 的第一頁再按 `?`，標題是 `FIRST PAGE — PRIDE AND PREJUDICE KEYS`，應該是
  `FIRST PAGE — Pride and Prejudice KEYS`。作者的 viewer 也一樣。
- **該怎麼改**：lib 只把 Title 轉大寫，Subject 照原樣。

## M6-6　NEW BOOK 表單的 key reference 列了不存在的鍵（DM9）

- **違反**：tdp M6「`?` 的 key reference……對象不存在就不列」。
- **在哪裡**：terminu-go `form.go` 的 `keyRef` 是寫死的清單。
- **現況**：NEW BOOK 沒有 checkbox，`?` 卻列 `Space  toggle the checkbox option`；剛打開、沒有錯誤時也列 `e`。
- **該怎麼改**：lib 只在表單有 checkbox 時列 `Space`，有錯誤時才列 `e`。

## T1-2　從 Space menu 執行 Remove，confirm 按 `Enter` 之後 Space menu 也收掉（DM11）

- **違反**：tdp T1 的表「短的 confirm／訊息 → 保留」（清單：取消或完成後回到原 menu）。
- **在哪裡**：terminu-go `app.go` 的 `runAction`：Action 帶 `Ask` 時，接受之後關掉整疊（`closeAll`）。
- **現況**：`P` `Space` `r` 疊出 REMOVE，`Esc` 回到 Space menu（對）；`Enter` 拿掉書之後，Space menu 也不見了。
- **該怎麼改**：規範先定（tdp M4「執行了就關掉 menu」跟 T1 的表互相拉扯，見 DM11），lib 再照定案改。

## panel-37　powerline 符號佔兩格時，換 tab 字會位移（DM12）

- **違反**：components/layout/panel 的膠囊串「斜角與接縫一律各佔一個 powerline 符號的寬度……換選中的段時字不位移」。
- **在哪裡**：terminu-go `panel.go` 的 `capsules`：在第一個 tab 時，`[N]` 與 tab 之間固定一個空格，沒有跟著 powerline 的寬度。
- **現況**：`DEMOU__POWERLINE_WIDTH=2`，MOVE 的 `[2]` 從 BOOKS 換到 AUTHORS，`BOOKS`、`AUTHORS` 都往右移一格。
- **該怎麼改**：規範先寫明那個空格佔一個 powerline 符號的寬度（跟 panel-42「中間是一個空格」對不上，見 DM12），lib 再改。

## panel-24　panel 窄的時候，空狀態的提示貼著框（DM13）

- **違反**：components/layout/panel「內容左邊離框一格」。
- **在哪裡**：terminu-go `panel.go` 的 `emptyState` 把句子截到整個內寬，沒有左右各留一格。
- **現況**：40×12，MOVE 的 Reading 架清空後 `Tab` 到 `[2]`：`┃Press [Space] to see what you can do …┃`，左右都貼著框。
- **該怎麼改**：規範定空狀態的句子放不下時怎麼辦（見 DM13），lib 再改；最少截到內寬 − 2。

## input-62　powerline 符號佔兩格時，表單上 focus 的值多一個 `…`（DM14）

- **違反**：components/input「在表單與 panel 上怎麼顯示」：值太長才從尾端截、加 `…`。
- **在哪裡**：terminu-go `form.go` 先把值補滿 `room` 格，`block` 再照箭頭的寬度截一次，箭頭兩格時把補的空白截掉、加上 `…`。
- **現況**：`DEMOU__POWERLINE_WIDTH=2`，`P` `n` 之後 focus 那一欄是 `not set…`；Title 填 `Dune` 是 `Dune…`。
- **該怎麼改**：lib 補空白時扣掉箭頭的寬度。
