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

// Get
func (srv *Device) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	item, err := srv.Repo.Get(req.Item)
	if err != nil {
		return err
	}
	res.Item = item
	return err
}

// Create 添加新的盘点信息
func (srv *Device) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req.Item)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加盘点信息失败")
	}
	res.Valid = true
	return err
}

// Update
func (srv *Device) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	item, valid, err := srv.Repo.Update(req.Item)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新盘点信息失败")
	}
	res.Item = item
	res.Valid = valid
	return err
}

// Delete
func (srv *Device) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return err
}
