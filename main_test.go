package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/lecex/device/handler"
	devicePB "github.com/lecex/device/proto/device"
	db "github.com/lecex/device/providers/database"
	service "github.com/lecex/device/service/device"

	_ "github.com/lecex/device/providers/migrations" // 执行数据迁移
)

func TestItemsCreate(t *testing.T) {
	// req := &devicePB.Request{
	// 	Item: &devicePB.Item{
	// 		PluCode:     `6923450662007`,
	// 		Number:      49000,
	// 		Position:    `测试备注1`,
	// 		No:          `20140422001`,
	// 		UserId:      `d6ef26fe-5e78-43ce-a16d-17002c7e7fcf`,
	// 		CheckUserId: `d6ef26fe-5e78-43ce-a16d-17002c7e7fcf`,
	// 		Infos: []*devicePB.Info{
	// 			{
	// 				Input:    `30*5+20*6`,
	// 				Number:   27000,
	// 				Position: `货架1`,
	// 			},
	// 			{
	// 				Input:    `20*5+20*6`,
	// 				Number:   22000,
	// 				Position: `货架1`,
	// 			},
	// 		},
	// 	},
	// }
	// res := &devicePB.Response{}
	// h := handler.Device{&service.DeviceRepository{db.DB}}
	// err := h.Create(context.TODO(), req, res)
	// fmt.Println("ItemDevice", res, err)
	// t.Log(req, res, err)
}

func TestItemsUpdate(t *testing.T) {
	// req := &devicePB.Request{
	// 	Item: &devicePB.Item{
	// 		Id:          6,
	// 		PluCode:     `7923450662007`,
	// 		Number:      49000,
	// 		Position:    `测试备注3`,
	// 		No:          `20140422001`,
	// 		UserId:      `d6ef26fe-5e78-43ce-a16d-17002c7e7fcf`,
	// 		CheckUserId: `d6ef26fe-5e78-43ce-a16d-17002c7e7fcf`,
	// 		Infos: []*devicePB.Info{
	// 			{
	// 				Id:       3,
	// 				Input:    `30*5+20*6`,
	// 				Number:   27000,
	// 				Position: `货架3`,
	// 			},
	// 			{
	// 				Id:       4,
	// 				Input:    `20*5+20*6`,
	// 				Number:   22000,
	// 				Position: `货架3`,
	// 			},
	// 		},
	// 	},
	// }
	// res := &devicePB.Response{}
	// h := handler.Device{&service.DeviceRepository{db.DB}}
	// err := h.Update(context.TODO(), req, res)
	// fmt.Println("ItemDevice", res, err)
	// t.Log(req, res, err)
}

func TestItemsGet(t *testing.T) {
	req := &devicePB.Request{
		Item: &devicePB.Item{
			// Id: 6,
			PluCode: `6923450662007`,
			Handle:  false,
		},
	}
	res := &devicePB.Response{}
	h := handler.Device{&service.DeviceRepository{db.DB}}
	err := h.Get(context.TODO(), req, res)
	fmt.Println("ItemDeviceGet", res, err)
	t.Log(req, res, err)
}
