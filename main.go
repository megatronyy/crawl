package main

import (
	"github.com/henrylee2cn/pholcus/exec"
	_ "crawl/rule"
)

func main() {
	exec.DefaultRun("web")
}
