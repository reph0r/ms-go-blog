package main

import (
	"ms-go-blog/common"
	"ms-go-blog/server"
)

func init() {
	//模板加载
	common.LoadTemplate()
}
func main() {
	//程序入口，一个项目 只能有一个入口
	//web程序，http协议 ip port
	server.App.Start("127.0.0.1", "8080")
}
