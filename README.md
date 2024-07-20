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

```sh
go run src/main.go
```

** How to run dev**

```sh
go install github.com/cespare/reflex@latest

reflex -c reflex.conf
```