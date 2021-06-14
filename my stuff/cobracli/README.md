# cobracli

`boilerplate`

> initiate go module

    go mod init cobracli

> get the cobra module and the cobra CLI tool

    go get -u github.com/spf13/cobra
    go get -u github.com/spf13/cobra/cobra

`functionality`

> initiate cobra project (must be the same as the go module cobracli)

    cobra init --pkg-name cobracli

> add first command (hostname)

    go get github.com/mitchellh/go-homedir@v1.1.0
    cobra add hostname