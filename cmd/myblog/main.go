package main

import (
	"myblog/internal/myblog"
	"os"

	_ "go.uber.org/automaxprocs"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := myblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
