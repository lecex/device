package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"

	"github.com/lecex/core/util"
	pb "github.com/lecex/device/proto/device"
)

//Device 仓库接口
type Device interface {
	Create(req *pb.Device) (*pb.Device, error)
	Delete(req *pb.Device) (bool, error)
	Update(req *pb.Device) (bool, error)
	Get(req *pb.Device) (*pb.Device, error)
	All(req *pb.Request) ([]*pb.Device, error)
	List(req *pb.ListQuery) ([]*pb.Device, error)
	Total(req *pb.ListQuery) (int64, error)
}

// DeviceRepository 用户仓库
type DeviceRepository struct {
	DB *gorm.DB
}

// All 获取所有设备信息
func (repo *DeviceRepository) All(req *pb.Request) (devices []*pb.Device, err error) {
	if err := repo.DB.Find(&devices).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}
	return devices, nil
}

// List 获取所有设备信息
func (repo *DeviceRepository) List(req *pb.ListQuery) (devices []*pb.Device, err error) {
	db := repo.DB
	limit, offset := util.Page(req.Limit, req.Page) // 分页
	sort := util.Sort(req.Sort)                     // 排序 默认 created_at desc
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&devices).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}
	return devices, nil
}

// Total 获取所有设备查询总量
func (repo *DeviceRepository) Total(req *pb.ListQuery) (total int64, err error) {
	devices := []pb.Device{}
	db := repo.DB
	// 查询条件
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	if err := db.Find(&devices).Count(&total).Error; err != nil {
		log.Fatal(err)
		return total, err
	}
	return total, nil
}

// Get 获取设备信息
func (repo *DeviceRepository) Get(device *pb.Device) (*pb.Device, error) {
	if err := repo.DB.Where(&device).Find(&device).Error; err != nil {
		return nil, err
	}
	return device, nil
}

// Create 创建设备
// bug 无设备名创建设备可能引起 bug
func (repo *DeviceRepository) Create(req *pb.Device) (*pb.Device, error) {
	err := repo.DB.Create(req).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Fatal(err)
		return req, fmt.Errorf("添加设备失败")
	}
	return req, nil
}

// Update 更新设备
func (repo *DeviceRepository) Update(req *pb.Device) (bool, error) {
	if req.Id == 0 {
		return false, fmt.Errorf("请传入更新id")
	}
	id := &pb.Device{
		Id: req.Id,
	}
	err := repo.DB.Model(id).Updates(req).Error
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

// Delete 删除设备
func (repo *DeviceRepository) Delete(req *pb.Device) (bool, error) {
	err := repo.DB.Delete(req).Error
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}
