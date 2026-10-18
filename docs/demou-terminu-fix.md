# demou 尚未符合 tdp 的地方

repo 照 tdp v0.7.0 的 project 建好了，app 本身還沒做。逐條待修：

- **畫面**：`cmd/demou/main.go` 現在只有一個空的 panel，讓程式跑得起來；換成 demou 真正的畫面。
- **S5、E6 家族 mark 與 splash**：還沒有 `docs/icon.svg`。有了以後用它建 splash（`terminu.NewSplash`），加測試守住兩者一致。
- **E6 L4 的畫面測試**：還沒有。多種終端機尺寸下，每一列剛好等於終端機寬度。
- **E6、E7 demo gif**：還沒有。VHS 的 tape 放 `.local/demos/demo.tape`，`make gif` 錄。
- **E7 README**：「它能做什麼」、特色、第一次啟動、快速開始、按鍵、設定與資料、限制還沒寫（兩份 README 裡標著註解的地方）。
