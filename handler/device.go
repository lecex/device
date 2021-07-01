package handler

import (
	"context"
	"encoding/json"
	"fmt"

	eventPB "github.com/lecex/core/proto/event"
	"github.com/micro/go-micro/v2"

	pb "github.com/lecex/device/proto/device"
	service "github.com/lecex/device/service/repository"
)

// Device 盘点
type Device struct {
	Repo      service.Device
	Publisher micro.Publisher
}

// List 获取所有设备
func (srv *Device) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	devices, err := srv.Repo.List(req.ListQuery)
	total, err := srv.Repo.Total(req.ListQuery)
	if err != nil {
		return err
	}
	res.Devices = devices
	res.Total = total
	return err
}

// Get 获取设备
func (srv *Device) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	device, err := srv.Repo.Get(req.Device)
	if err != nil {
		return err
	}
	res.Device = device
	return err
}

// Create 创建设备
func (srv *Device) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req.Device)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加设备失败")
	}
	res.Valid = true
	return err
}

// Update 更新设备
func (srv *Device) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req.Device)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新设备失败:%s", err.Error())
	}
	res.Valid = valid
	if valid {
		device, err := srv.Repo.Get(req.Device)
		if err != nil {
			return err
		}
		if err := srv.publish(ctx, device, "device.Devices.Update"); err != nil {
			return err
		}
	}
	return err
}

// Delete 删除设备
func (srv *Device) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.Device)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除设备失败")
	}
	res.Valid = valid
	return err
}

// publish 消息发布
func (srv *Device) publish(ctx context.Context, device *pb.Device, topic string) (err error) {
	data, _ := json.Marshal(&device)
	event := &eventPB.Event{
		UserId:     "",
		DeviceInfo: device.Info,
		GroupId:    "",
		Topic:      topic,
		Data:       data,
	}
	return srv.Publisher.Publish(ctx, event)
}
