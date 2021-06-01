package repository

import (
	"fmt"

	"github.com/go-log/log"
	"github.com/jinzhu/gorm"

	pb "github.com/lecex/device/proto/inventory"
)

//Device 仓库接口
type Device interface {
	Create(user *pb.Item) (*pb.Item, error)
	Get(user *pb.Item) (*pb.Item, error)
	Update(user *pb.Item) (*pb.Item, bool, error)
	Delete(user *pb.Item) (bool, error)
}

// DeviceRepository 用户仓库
type DeviceRepository struct {
	DB *gorm.DB
}

// All 获取所有角色信息
func (repo *DeviceRepository) All(req *pb.Request) (devices []*pb.Device, err error) {
	if err := repo.DB.Find(&devices).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return devices, nil
}

// List 获取所有角色信息
func (repo *DeviceRepository) List(req *pb.ListQuery) (devices []*pb.Device, err error) {
	db := repo.DB
	limit, offset := uitl.Page(req.Limit, req.Page) // 分页
	sort := uitl.Sort(req.Sort)                     // 排序 默认 created_at desc
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&devices).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return devices, nil
}

// Total 获取所有角色查询总量
func (repo *DeviceRepository) Total(req *pb.ListQuery) (total int64, err error) {
	devices := []pb.Device{}
	db := repo.DB
	// 查询条件
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&devices).Count(&total).Error; err != nil {
		log.Log(err)
		return total, err
	}
	return total, nil
}

// Get 获取角色信息
func (repo *DeviceRepository) Get(device *pb.Device) (*pb.Device, error) {
	if err := repo.DB.Where(&device).Find(&device).Error; err != nil {
		return nil, err
	}
	return device, nil
}

// Create 创建角色
// bug 无角色名创建角色可能引起 bug
func (repo *DeviceRepository) Create(r *pb.Device) (*pb.Device, error) {
	err := repo.DB.Create(r).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return r, fmt.Errorf("添加角色失败")
	}
	return r, nil
}

// Update 更新角色
func (repo *DeviceRepository) Update(r *pb.Device) (bool, error) {
	if r.Id == 0 {
		return false, fmt.Errorf("请传入更新id")
	}
	id := &pb.Device{
		Id: r.Id,
	}
	err := repo.DB.Model(id).Updates(r).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除角色
func (repo *DeviceRepository) Delete(r *pb.Device) (bool, error) {
	err := repo.DB.Delete(r).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
