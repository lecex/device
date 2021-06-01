package handler

import (
	"context"
	"fmt"

	pb "github.com/lecex/device/proto/device"
	service "github.com/lecex/device/service/repository"
)

// Device 盘点
type Device struct {
	Repo service.Device
}

// List 获取所有角色
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

// Get 获取角色
func (srv *Device) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	device, err := srv.Repo.Get(req.Device)
	if err != nil {
		return err
	}
	res.Device = device
	return err
}

// Create 创建角色
func (srv *Device) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req.Device)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加角色失败")
	}
	res.Valid = true
	return err
}

// Update 更新角色
func (srv *Device) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req.Device)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新角色失败")
	}
	res.Valid = valid
	return err
}

// Delete 删除角色
func (srv *Device) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.Device)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除角色失败")
	}
	res.Valid = valid
	return err
}
