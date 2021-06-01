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

// Create 创建盘点信息
func (repo *DeviceRepository) Create(item *pb.Item) (*pb.Item, error) {
	err := repo.DB.Create(item).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return item, fmt.Errorf("新增盘点信息失败")
	}
	return item, nil
}

// Get 获取盘点信息
func (repo *DeviceRepository) Get(item *pb.Item) (*pb.Item, error) {
	infos := []*pb.Info{}
	if err := repo.DB.Where("plu_code = ?", item.PluCode).Where("handle = ?", item.Handle).Find(&item).Error; err != nil {
		return nil, err
	}
	if err := repo.DB.Model(&item).Related(&infos).Error; err != nil {
		if err.Error() != "record not found" {
			return nil, err
		}
	}
	item.Infos = infos
	return item, nil
}

// Update 更新用户
func (repo *DeviceRepository) Update(item *pb.Item) (*pb.Item, bool, error) {
	if item.Id == 0 {
		return item, false, fmt.Errorf("请传入更新id")
	}
	id := &pb.Item{
		Id: item.Id,
	}
	err := repo.DB.Model(id).Updates(item).Error
	if err != nil {
		log.Log(err)
		return item, false, err
	}
	return item, true, nil
}

// Delete 删除盘点信息
func (repo *DeviceRepository) Delete(item *pb.Item) (bool, error) {
	err := repo.DB.Delete(item).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
