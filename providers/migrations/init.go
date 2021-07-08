package migrations

import (
	cashierPB "github.com/lecex/device/proto/cashier"
	devicePB "github.com/lecex/device/proto/device"
	db "github.com/lecex/device/providers/database"
)

func init() {
	device()
	cashier()
}

// device 数据库结构
func device() {
	device := &devicePB.Device{}
	if !db.DB.HasTable(&device) {
		db.DB.Exec(`
			CREATE TABLE devices (
				id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
				code varchar(64) DEFAULT NULL COMMENT '编码',
				status int(11) DEFAULT 0 COMMENT '状态',
				config text DEFAULT NULL COMMENT '配置',
				info varchar(32) DEFAULT NULL COMMENT '设备信息',
				user_id varchar(36) DEFAULT NULL COMMENT '用户ID',
				created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
				PRIMARY KEY (id),
				UNIQUE KEY code (code)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

// cashier 数据库结构
func cashier() {
	cashier := &cashierPB.Cashier{}
	if !db.DB.HasTable(&cashier) {
		db.DB.Exec(`
			CREATE TABLE cashiers (
				code varchar(64) NOT NULL,
				name varchar(64) DEFAULT NULL,
				password varchar(128) DEFAULT NULL,
				user_id varchar(36) DEFAULT NULL,
				PRIMARY KEY (code),
				UNIQUE KEY user_id (user_id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}
