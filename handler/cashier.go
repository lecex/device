package handler

import (
	"context"
	"fmt"

	pb "github.com/lecex/device/proto/cashier"

	"github.com/lecex/device/service/repository"
)

// Cashier 收银员结构
type Cashier struct {
	Repo repository.Cashier
}

// Exist 用户是否存在
func (srv *Cashier) Exist(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	res.Valid = srv.Repo.Exist(req.Cashier)
	return err
}

// All 获取所有收银员
func (srv *Cashier) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	cashiers, err := srv.Repo.All(req)
	if err != nil {
		return err
	}
	res.Cashiers = cashiers
	return err
}

// List 获取所有收银员
func (srv *Cashier) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	cashiers, err := srv.Repo.List(req.ListQuery)
	total, err := srv.Repo.Total(req.ListQuery)
	if err != nil {
		return err
	}
	res.Cashiers = cashiers
	res.Total = total
	return err
}

// Get 获取收银员
func (srv *Cashier) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	cashier, err := srv.Repo.Get(req.Cashier)
	if err != nil {
		return err
	}
	res.Cashier = cashier
	return err
}

// Create 创建收银员
func (srv *Cashier) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req.Cashier)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加收银员失败")
	}
	res.Valid = true
	return err
}

// Update 更新收银员
func (srv *Cashier) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req.Cashier)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新收银员失败")
	}
	res.Valid = valid
	return err
}

// Delete 删除收银员
func (srv *Cashier) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.Cashier)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除收银员失败")
	}
	res.Valid = valid
	return err
}
