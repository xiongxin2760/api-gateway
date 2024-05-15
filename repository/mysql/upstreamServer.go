package mysql

import "context"

var globalID int64

var ServerPOMap = map[int64]ServerPO{}

// 暂时先存放在内存里面
type ServerPO struct {
	ID          int64  `json:"id" gorm:"column:id"` // 自增ID
	Name        string `json:"name" gorm:"column:name"`
	Discription string `json:"discription" gorm:"column:discription"`
	Timeout     int    `json:"timeout" gorm:"column:timeout"`
	Retry       int    `json:"retry" gorm:"column:retry"`
	Balance     string `json:"balance" gorm:"column:balance"`
	Service     string `json:"service" gorm:"column:service"`
}

func CreatServerPO(ctx context.Context, po ServerPO) (int64, error) {
	globalID++
	newID := globalID
	po.ID = newID
	ServerPOMap[newID] = po

	return newID, nil
}

func SearchServerPO(ctx context.Context, ID int64) (*ServerPO, error) {
	if po, exist := ServerPOMap[ID]; exist {
		return &po, nil
	}
	return nil, nil
}

func UpdateServerPO(ctx context.Context, po ServerPO) error {
	ServerPOMap[po.ID] = po
	return nil
}
