package repository

import (
	"fmt"
	// 公共引入
	"github.com/jinzhu/gorm"
	"github.com/lecex/core/util"
	"github.com/micro/go-micro/v2/util/log"

	pb "github.com/lecex/device/proto/cashier"
)

//Cashier 仓库接口
type Cashier interface {
	Create(cashier *pb.Cashier) (*pb.Cashier, error)
	Delete(cashier *pb.Cashier) (bool, error)
	Update(cashier *pb.Cashier) (bool, error)
	Get(cashier *pb.Cashier) (*pb.Cashier, error)
	All(req *pb.Request) ([]*pb.Cashier, error)
	List(req *pb.ListQuery) ([]*pb.Cashier, error)
	Total(req *pb.ListQuery) (int64, error)
}

// CashierRepository 收银员仓库
type CashierRepository struct {
	DB *gorm.DB
}

// All 获取所有角色信息
func (repo *CashierRepository) All(req *pb.Request) (cashiers []*pb.Cashier, err error) {
	if err := repo.DB.Find(&cashiers).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return cashiers, nil
}

// List 获取所有收银员信息
func (repo *CashierRepository) List(req *pb.ListQuery) (cashiers []*pb.Cashier, err error) {
	db := repo.DB
	limit, offset := util.Page(req.Limit, req.Page) // 分页
	sort := util.Sort(req.Sort)                     // 排序 默认 created_at desc
	// 查询条件
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&cashiers).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return cashiers, nil
}

// Total 获取所有收银员查询总量
func (repo *CashierRepository) Total(req *pb.ListQuery) (total int64, err error) {
	cashiers := []pb.Cashier{}
	db := repo.DB
	// 查询条件
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&cashiers).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
}

// Get 获取收银员信息
func (repo *CashierRepository) Get(cashier *pb.Cashier) (*pb.Cashier, error) {
	if err := repo.DB.Where(&cashier).Find(&cashier).Error; err != nil {
		return nil, err
	}
	return cashier, nil
}

// Create 创建收银员
// debug 无收银员名创建收银员可能引起 bug
func (repo *CashierRepository) Create(p *pb.Cashier) (*pb.Cashier, error) {
	err := repo.DB.Create(p).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return p, fmt.Errorf("添加收银员失败")
	}
	return p, nil
}

// Update 更新收银员
func (repo *CashierRepository) Update(p *pb.Cashier) (bool, error) {
	if p.UserId == "" {
		return false, fmt.Errorf("请传入更新id")
	}
	err := repo.DB.Where("user_id = ?", p.UserId).Updates(p).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除收银员
func (repo *CashierRepository) Delete(p *pb.Cashier) (bool, error) {
	if p.UserId == "" {
		return false, fmt.Errorf("请传入更新id")
	}
	err := repo.DB.Where("user_id = ?", p.UserId).Delete(p).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
