# demou 尚未符合 tdp 的地方

照 tdp v0.7.4 的清單對過一遍，再照 v0.8.0 改了的規則與條目對過，下面是還沒符合的地方。都出在 terminu-go v0.1.13 的行為，demou
這邊改不了；已經回報給 terminu 的 session（DM 編號），規範也都定了（tdp v0.8.0）。lib 發新版、`go.mod` 跟上之後逐條刪掉，全部修完就
刪掉這個檔（tdp E7）。條目編號是 tdp v0.8.0 清單（`checklist/items-zh_TW.md`）的。

## 比最低尺寸小時，畫的不是名稱與提示（DM34；DM4、DM5、DM6、DM13、DM17、DM37）

- **違反**：tdp L1「比最低尺寸小時……不畫 app 的畫面，只在正中央畫兩列」、「一直縮到 1×1 都不能當掉」；tdp E5 的測試邊界要含比最低
  尺寸小、一直到 1×1 的尺寸。清單 L1-2～L1-6、E5-10、E5-13。
- **在哪裡**：terminu-go v0.1.13 的 `App` 沒有最低尺寸。demou 的最低尺寸是 80×24（暫定，見 `docs/dev-remarks.md`）。
- **現況**：比 80×24 小時照樣畫三個畫面與 popup，在那些尺寸上還會：
  - 開著 popup 縮到 5 列以下，`makeslice: cap out of range` 後結束（DM4）；
  - 10 列以下 NEW BOOK 表單比畫面高度 − 2 還高，蓋住 footer（DM5）；8 列時 Title 的 text popup 看不到打字列（DM6）；
  - 40×12 時空狀態的句子貼著框（DM13）、popup 的 hint 捨掉 `Esc`（DM17）；15 欄以下膠囊與欄被直接切掉（DM37）。
- **該怎麼改**：lib 加最低尺寸（預設 80×40、app 可以調低）與「只顯示名稱與提示」的畫面；demou 設 80×24，`TestE5PopupBiggerThanTheScreen`
  加 80×5、1×1，驗名稱與提示（`TODO(lib DM4)`）。

## key reference 列了不存在的鍵、移動鍵的寫法不對（DM9、DM10）

- **違反**：tdp M6、K6「對象不存在就不列」「一維清單寫 `h|k` 往前、`j|l` 往後」。清單 M6-6、K6-7、K6-8。
- **在哪裡**：terminu-go v0.1.13 `form.go` 的 `keyRef` 是寫死的清單；menu、表單的移動鍵；`app.go` 的 panel key reference 一律放
  `p.Body.MoveKeys()`。
- **現況**：
  - NEW BOOK 沒有 checkbox，`?` 卻列 `Space  toggle the checkbox option`，`Enter` 也寫 choose or toggle the option；沒有錯誤時也列 `e`。
  - 表單與 Space menu 的 `move` 寫 `j|k`，沒有 `h|k`、`j|l`。
  - 清單空了（例：POPUPS 把書都拿掉）按 `?` 仍列 `j|k`、`u|d`、`gg|G`；viewer 內容全放得下時也列 `j|k scroll`。
- **該怎麼改**：lib 照表單實際有的欄位與狀態列鍵；一維清單的 move 寫 `h|k`、`j|l`；移動鍵、捲動鍵照當下有沒有對象給。

## viewer 的 key reference 把書名轉成大寫（DM8）

- **違反**：components 的「大寫」：標題裡的名稱照原樣；note 的 key reference 標題是 `<surface> keys`。清單 note-42。
- **在哪裡**：terminu-go v0.1.13 `note.go` 的 `keyReference` 把整個 surface 轉大寫，viewer 的 surface 含 Subject。
- **現況**：打開 Pride and Prejudice 的第一頁再按 `?`，標題是 `FIRST PAGE — PRIDE AND PREJUDICE KEYS`，應該是
  `FIRST PAGE — Pride and Prejudice KEYS`。作者的 viewer 也一樣。
- **該怎麼改**：lib 只把 Title 轉大寫，Subject 照原樣。

## viewer 沒有選取模式（DM25）

- **違反**：tdp F1「每個 viewer 都有選取模式」、components/dialog/note 的「選取模式」。清單 note-70，以及選取模式的其他條目。
- **在哪裡**：`internal/ui/ui.go` 的 `openBook`、`internal/ui/move.go` 的 `openAuthor` 開的兩個 viewer（`TODO(lib DM25)`）。
- **現況**：在第一頁與作者上按 `v` 沒有反應。
- **該怎麼改**：lib 讓每個 viewer 預設就有選取模式與複製（複製後 Success toast）；demou 升級後拿掉兩個 TODO，在兩個 viewer 上驗 `v`、`y`。

## 表單的打字說明 note 上按 `q` 會離開 app（DM16）

- **違反**：tdp K9「表單跳出的打字說明 note 上也不作用：那個 note 只收 `Enter`、`Esc`」。清單 K9-2、form-95。
- **在哪裡**：terminu-go v0.1.13 `note.go` 的 key 對 `q` 回 false，交給離開流程。
- **現況**：NEW BOOK 的 Title 上打 `aq`：`a` 跳出 `Press [Enter] to edit Title.`，`q` 接著跳出 `Quit demou?`。
- **該怎麼改**：lib 讓打字說明 note 吃掉 `Enter`、`Esc` 以外可以打的字。

## 從 Space menu 刪書，兩層 popup 只有上面那層播收起動畫（F2）

- **違反**：tdp F2「好幾層一起收起時……每一層同時收起，各自照自己的大小與層色」。清單 F2-3。
- **在哪裡**：terminu-go v0.1.13 `app.go` 的 `closeAll`：最上層以外的直接拿掉，只有最上層播。
- **現況**：`P` `Space` `r` `Enter`：書拿掉之後 REMOVE confirm 收起，Space menu 直接消失（照 tdp T1，刪除讓那一項不在了，menu 清掉是對的）。
- **該怎麼改**：lib 讓每一層一起播收起動畫。

## powerline 符號佔兩格時，換 tab 字會位移（DM12）

- **違反**：components/layout/panel 的膠囊串：`[N]` 與第一個 tab 之間的空白佔一個 powerline 符號的寬度，換 tab 時字不位移。清單 panel-37、
  panel-42。
- **在哪裡**：terminu-go v0.1.13 `panel.go` 的 `capsules`：在第一個 tab 時，`[N]` 與 tab 之間固定一個空格。
- **現況**：`DEMOU__POWERLINE_WIDTH=2`，MOVE 的 `[2]` 從 BOOKS 換到 AUTHORS，`BOOKS`、`AUTHORS` 都往右移一格。
- **該怎麼改**：lib 用 powerline 符號寬的空白。

## powerline 符號佔兩格時，表單上 focus 的值多一個 `…`（DM14）

- **違反**：components/input「值太長才從尾端截、加 `…`」；表單的 focus 色塊右邊的留白跟箭頭一樣寬（components/dialog/form、
  layout/popup）。清單 input-62、popup-20。
- **在哪裡**：terminu-go v0.1.13 `form.go` 先把值補滿 `room` 格，`block` 再照箭頭的寬度截一次，箭頭兩格時把補的空白截掉、加上 `…`。
- **現況**：`DEMOU__POWERLINE_WIDTH=2`，`P` `n` 之後 focus 那一欄是 `not set…`；Title 填 `Dune` 是 `Dune…`。
- **該怎麼改**：lib 補空白時扣掉箭頭的寬度，每一列右邊留兩格。

## 隔著 ssh 帶來的 icon 寬度讀不到

- **違反**：tdp E5 的覆寫與順序：`<APP>__` → `TERMINU__` → `LC_TERMINU_` → 探測。清單 E5-5、E5-6。
- **在哪裡**：terminu-go v0.1.13 啟動時不讀 `LC_TERMINU_ICON_WIDTH`、`LC_TERMINU_POWERLINE_WIDTH`。
- **現況**：v0.1.13 的程式裡沒有 `LC_TERMINU_`：只設這兩個變數啟動，demou 照樣自己探測。
- **該怎麼改**：lib 啟動時照上面的順序讀；demou 升級後，README 的變數表加上這兩個。
