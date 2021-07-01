package handler

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	cashierPB "github.com/lecex/device/proto/cashier"
	devicePB "github.com/lecex/device/proto/device"

	db "github.com/lecex/device/providers/database"
	service "github.com/lecex/device/service/repository"
)

const topic = "event"

// Register 注册
func Register(srv micro.Service) {
	server := srv.Server()
	publisher := micro.NewPublisher(topic, srv.Client())
	// 获取 broker 实例
	devicePB.RegisterDevicesHandler(server, &Device{&service.DeviceRepository{db.DB}, publisher})
	cashierPB.RegisterCashiersHandler(server, &Cashier{&service.CashierRepository{db.DB}})
	// 事件处理
	sub := &Subscriber{&service.CashierRepository{db.DB}}
	err := micro.RegisterSubscriber(topic, server, sub)
	if err != nil {
		log.Fatal(err)
	}
}
