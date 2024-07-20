# go-template

** プロジェクトルートで以下を実行して go.mod を生成 **

```
go mod init my_go_project
go mod tidy
```

This is a simple Go project template.

** Requirements **

- Go 1.x
- reflex

** How to run **

```
go mod init my_go_project
go get github.com/jaschaephraim/lrserver
go mod tidy
```

```sh
go run src/main.go
```

** How to run dev**

```sh
go mod init my_go_project
go get github.com/fsnotify/fsnotify
go mod tidy

export PATH=$PATH:$(go env GOPATH)/bin

go install github.com/cespare/reflex@latest

reflex -c reflex.conf
```