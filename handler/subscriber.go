package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/micro/go-micro/v2/util/log"

	pb "github.com/lecex/core/proto/event"
	cashierPB "github.com/lecex/device/proto/cashier"

	"github.com/lecex/device/service/repository"
)

// Cashier 收银员结构
type Subscriber struct {
	Repo repository.Cashier
}

// Process 事件处理
func (sub *Subscriber) Process(ctx context.Context, event *pb.Event) (err error) {
	fmt.Println(event)
	switch event.Topic {
	case "user.Users.Delete":
		sub.delete(ctx, event.Data)
	}
	return err
}

// delete 删除
func (sub *Subscriber) delete(ctx context.Context, data []byte) (err error) {
	r := make(map[string]interface{})
	err = json.Unmarshal(data, &r)
	if err != nil {
		return
	}
	// 获取设备信息
	if userID, ok := r["id"]; ok {
		cashier := &cashierPB.Cashier{
			UserId: userID.(string),
		}
		valid, err := sub.Repo.Delete(cashier)
		if err != nil {
			log.Fatal("删除收银员失败", valid, err)
		}
	}
	return
}
