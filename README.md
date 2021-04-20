# Go 言語学習

特徴としては、コンパイル型の言語・コンパイル実行の高速・並行プログラミング・Google 開発

## 基本知識

- `.go`の拡張子
- `package`に属している必要があり、そのうちの一つである`main`が必要である
- `main`パッケージの中で`main`関数を実行する決まりになっている

## 基本

- 「 := 」 ... コロン,イコールで変数名を宣言

- log の出力は、「 import "log" 」を使う。代表的なのは、

  - log.Println()
    - Println() ... デフォルトのフォーマットを使用し出力する
  - log.Fatalln()
  - log.Printf()
    - Printf() ... フォーマットを指定して出力できる

  など色々ある。

### データ型

- int 整数の型
- string 文字
- bool 判定
- nil
- float64

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

### ポインタ

- ポインタ変数 ... メモリ上のアドレスを値として入れられる変数のこと。
- ポインタ型 ... 型を元にアドレスを格納する準備

```go
var n int = 100

var p *int = &n
```

`p`がポインタ変数、`*int`がポインタ型（データ値）といったところ。そしてアドレスを生成するのが`&`になる。

#### 応用

- new() ... 初期化せず、０の値になる。また、返り値がアドレス表示になる。（ポインタ型）
- make() ... 初期化し、空の値を返す。
