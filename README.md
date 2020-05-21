## RedisServer(もどき)


### 起動方法
```
$ go run server/main.go
```

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