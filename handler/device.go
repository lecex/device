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

// List
func (srv *Device) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	device, err := srv.Repo.List(req.Device)
	if err != nil {
		return err
	}
	res.Device = device
	return err
}

// Get
func (srv *Device) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	device, err := srv.Repo.Get(req.Device)
	if err != nil {
		return err
	}
	res.Device = device
	return err
}

// Create 添加新的盘点信息
func (srv *Device) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req.Device)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加盘点信息失败")
	}
	res.Valid = true
	return err
}

// Update
func (srv *Device) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	device, valid, err := srv.Repo.Update(req.Device)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新盘点信息失败")
	}
	res.Device = device
	res.Valid = valid
	return err
}

// Delete
func (srv *Device) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return err
}
