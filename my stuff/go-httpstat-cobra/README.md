# go-httpstat-cobra

> install cobra-cli

```bash
go install github.com/spf13/cobra-cli@latest
```

> initiate cobra project

```bash
mkdir go-httpstat-cobra
cd go-httpstat-cobra
go mod init go-httpstat-cobra
cobra-cli init
```

> add first command

```bash
cobra-cli add uri
go run main.go uri # check it works
```

> add second command

```bash
cobra-cli add uri-list
go run main.go uri-list # check it works
```
