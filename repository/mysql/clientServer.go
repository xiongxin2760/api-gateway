package mysql

import "context"

var clientrouterID int64

var RouterPOMap = map[int64]RouterPO{}

// 暂时先存放在内存里面
type RouterPO struct {
	ID          int64  `json:"id" gorm:"column:id"` // 自增ID
	Name        string `json:"name" gorm:"column:name"`
	Discription string `json:"discription" gorm:"column:discription"`
	LocationURL string `json:"locationUrl" gorm:"column:locationUrl"` // 客户端路由
	Target      int64  `json:"target" gorm:"column:target"`           // 绑定的server的id
}

func CreatRouterPO(ctx context.Context, po RouterPO) (int64, error) {
	clientrouterID++
	newID := clientrouterID
	po.ID = newID
	RouterPOMap[newID] = po

	return newID, nil
}

func SearchRouterPO(ctx context.Context, ID int64) (*RouterPO, error) {
	if po, exist := RouterPOMap[ID]; exist {
		return &po, nil
	}
	return nil, nil
}

func UpdateRouterPO(ctx context.Context, po RouterPO) error {
	RouterPOMap[po.ID] = po
	return nil
}
