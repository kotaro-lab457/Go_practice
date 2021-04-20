# Go 言語学習

## 基本

- 「 := 」 ... コロン,リコールで変数名を宣言

- log の出力は、「 import "log" 」を使う。代表的なのは、
  - log.Println()
  - log.Fatalln()
  - log.Printf()　　など色々ある。

型定義

- int 整数の型

### fmt フォーマット

- フォーマットパッケージ

import {
"fmt"
}

### os パッケージ機能

機能：ファイルやディレクトリの操作、環境変数の操作、プロセス情報の参照

- ファイルオープン系　メソッド

  - os.Open(ファイル名 string) ... 読み取り専用ファイル
  - os.Create(ファイル名 string) ... 新規ファイルの作成
  - os.OpenFile(ファイル名 string, フラグ int,ファイルモード（パーミション）FileMode)

  フラグ＝ビットフラグ定数を指定できる。O_RDONLY(読み取り専用でオープン)、O_CREATE(ファイルが存在しない場合に新規作成) O_APPEND(ファイル末尾に追加する)などなど。

  [https://leben.mobi/go/os-file-and-directory/go-programming/](https://leben.mobi/go/os-file-and-directory/go-programming/)

### コマンド集

- 「go run ファイル名」 ... 出力

### メソッド

- time.Now() ... 現在時刻の取得
  - Hour() ... 「時」表示

### 応用

- fot 文

「continue」で要素を追加し、継続できる。また「break」で処理の中断。。
