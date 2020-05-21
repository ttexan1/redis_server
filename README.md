## RedisServer(もどき)


### 起動方法
```
$ go run server/main.go
```

### 説明

* tcpコネクションによるclient-serverモデル
* 各々のクライアントとの通信はgoroutineでスレッドを立てて行われる。
* 処理の流れ
  - main -> parser -> usecase -> store
  - (parserよりhandlerの方がbetterな気もするが、もっと良い名前もありそうなので、いったん保留にした。)
* 大まかにはデータの型ごとに処理のやり方にまとまりがあるので、parser, usecase, storeの中のファイルはデータの型を基準として分けている。ただし、今回実装したのはシンプルなkey-value処理だけである(domainの構造体で言うとSingleに当たるものだけ)。
* command系は定数として宣言してある。
* 保存される値は、interfaceがふさわしいかとも考えたが、基本的に文字列なので、型はstring型とし、必要に応じてintやfloatへ変換をするような実装にした。

### アピールポイント

* interfaceを用いて各パッケージの処理内容を疎結合にした。
* データベースの更新はstoreパッケージでしか発生し得ないため、思わぬ変更によるバグを防げる。
* パーサーの登録は、net/httpパッケージのHandlerの登録方法を参考にした。
* resp(redisのprotocol)のエンコード処理はdomainに封じ込めて、実装がバラつかないように注意した。


  

### 主要ファイル構造
```
.
├── command
│   ├── list.go
│   ├── other.go
│   ├── single.go
│   └── static.go
├── domain
│   ├── error.go
│   ├── list.go
│   ├── response.go
│   └── single.go
├── parser
│   ├── list.go
│   ├── parser.go
│   ├── single.go
│   └── static.go
├── server
│   └── main.go
├── store
│   ├── db.go
│   ├── list.go
│   ├── single.go
└── usecase
    ├── list.go
    ├── single.go
    └── storage.go
```