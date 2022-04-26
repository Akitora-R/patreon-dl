package main

import (
	"log"
	"os"
	"patreon-dl/internal/app/patreon-dl/config"
	"patreon-dl/internal/app/patreon-dl/service"
	"patreon-dl/internal/app/patreon-dl/util"
)

func main() {
	args := os.Args
	for i, arg := range args {
		if arg == "--proxy" && len(args)-1 >= i+1 {
			util.ProxyUrl = args[i+1]
			log.Println("使用代理", util.ProxyUrl)
		}
	}
	service.DlAll()
	log.Printf("下载完成, 成功: %d,失败: %d,跳过: %d", config.Succeed, config.Failed, config.Ignored)
}
