package main

import (
	// 公共引入

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	_ "github.com/lecex/core/plugins" // 插件在后面执行
	"github.com/lecex/device/config"
	"github.com/lecex/device/handler"
	_ "github.com/lecex/device/providers/migrations" // 执行数据迁移
)

const topic = "websocket"

func main() {
	var Conf = config.Conf
	service := micro.NewService(
		micro.Name(Conf.Name),
		micro.Version(Conf.Version),
	)
	service.Init()
	publisher := micro.NewPublisher(topic, service.Client())
	// 注册服务
	handler.Register(service.Server(), publisher)
	// Run the server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	log.Log("serviser run ...")
}
