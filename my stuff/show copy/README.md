# show

`boilerplate`

> initiate go module

    go mod init show

> get the cobra module and the cobra CLI tool

    go get -u github.com/spf13/cobra
    go get -u github.com/spf13/cobra/cobra
    go get github.com/mitchellh/go-homedir@v1.1.0

> initiate cobra project (must be the same as the go module show)

    cobra init --pkg-name show

> test the app

    go run show.go

`functionality`

> add commands

    cobra add ip
    cobra add interface
    cobra add brief