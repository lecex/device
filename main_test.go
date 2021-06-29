package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/lecex/device/handler"
	cashierPB "github.com/lecex/device/proto/cashier"
	devicePB "github.com/lecex/device/proto/device"
	db "github.com/lecex/device/providers/database"
	service "github.com/lecex/device/service/repository"

	_ "github.com/lecex/device/providers/migrations" // 执行数据迁移
)

func TestDevicesCreate(t *testing.T) {
	req := &devicePB.Request{
		Device: &devicePB.Device{
			Code:   `0001`,
			Status: 1,
			Config: `
			{
				"barcodeReg":      "27PPPPPBBBBBC",
				"pay": {
					"orderTitle":      "测试订单标题",
					"scanStoreName":   "admin",
					"scanPayId":       "7",
					"cardPayId":       "6"
				},
				"printer":{
					"switch":   true,
					"template": "           ******超市\n************* {{stuats}} *************\n编码    商品名称   数量   合计\n{{goods(pluCode|7,name|10,number|5,total|7)}}\n--------------------------------\n收款方式    应收金额    实收金额\n{{pays(name|12,amount|9,getAmount|9)}}\n--------------------------------\n收款员: {{userId}} 收款机: {{terminal}} \n金额: {{total}}元\n订单:{{orderNo}} 打印: {{print}} 次\n时间:{{createdAt}}\n地址:五路319号\n电话:010-2120888 ",
					"accounts": "           当日交易汇总\n--------------------------------\n商户账号: {{storeId}}\n收款员: {{userId}} 收款机: {{terminal}} \n--------------------------------\n收款方式    金额\n{{pays(name|12,amount|9)}}\n--------------------------------\n实际扫码金额:{{payTotal}}\n订单: {{count}} 笔 退款: {{returns}}  笔\n未发布: {{publish}} 笔 \n时间:{{createdAt}}\n地址:五路319号\n电话:010-2120888 ",
					"device":   "USB"
				}
			}
			`,
			Info:   `TXAP12112000354ND002112`,
			UserId: `d6ef26fe-5e78-43ce-a16d-17002c7e7fcf`,
		},
	}
	res := &devicePB.Response{}
	h := handler.Device{&service.DeviceRepository{db.DB}, nil}
	err := h.Create(context.TODO(), req, res)
	fmt.Println("DeviceDevice", res, err)
	t.Log(req, res, err)
}

func TestDevicesUpdate(t *testing.T) {
	// req := &devicePB.Request{
	// 	Device: &devicePB.Device{
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
	// fmt.Println("DeviceDevice", res, err)
	// t.Log(req, res, err)
}

func TestDevicesGet(t *testing.T) {
	req := &devicePB.Request{
		Device: &devicePB.Device{
			Code: `0001`,
		},
	}
	res := &devicePB.Response{}
	h := handler.Device{&service.DeviceRepository{db.DB}, nil}
	err := h.Get(context.TODO(), req, res)
	fmt.Println("DeviceDeviceGet", res, err)
	t.Log(req, res, err)
}

func TestCashiersCreate(t *testing.T) {
	req := &cashierPB.Request{
		Cashier: &cashierPB.Cashier{
			UserId: "123456",
			// Code:     "1024",
			// Password: "132456",
			// Name:     "测试",
		},
	}
	res := &cashierPB.Response{}
	h := handler.Cashier{&service.CashierRepository{db.DB}}
	err := h.List(context.TODO(), req, res)
	fmt.Println("DeviceCashierGet", res, err)
	t.Log(req, res, err)
}
