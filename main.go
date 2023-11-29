package main

import (
	"context"

	"github.com/wildanfaz/simple-golang-echo/cmd"
)

func main() {
	cmd.InitCmd(context.Background())
}
