package main

import (
	"os"
	"patreon-dl/internal/app/patreon-dl/util"
)

func main() {
	args := os.Args
	if len(args) == 2 && args[0] == "--proxy" {
		util.ProxyUrl = args[1]
	}
	//service.DlPost()
}
