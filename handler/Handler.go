package handler

import (
	"github.com/micro/go-micro/v2"
	server "github.com/micro/go-micro/v2/server"

	devicePB "github.com/lecex/device/proto/device"

	db "github.com/lecex/device/providers/database"
	service "github.com/lecex/device/service/repository"
)

// Register 注册
func Register(Server server.Server, publisher micro.Publisher) {
	// 获取 broker 实例
	devicePB.RegisterDevicesHandler(Server, &Device{&service.DeviceRepository{db.DB}, publisher})
}
