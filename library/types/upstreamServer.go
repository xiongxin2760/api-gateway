package types

// type ServerRegisterReq struct {
// 	Name        string         `json:"name"`
// 	Discription string         `json:"discription"`
// 	Timeout     int            `json:"timeout"`
// 	Retry       int            `json:"retry"`
// 	Balance     string         `json:"balance"`
// 	Service     []ServerAPI    `json:"service"` // TODO：升级为服务发现
// 	Plugins     map[string]any `json:"plugins"` // 待定
// }

// api的地址
type ServerAPI struct {
	URL    string `json:"url"`
	Weight int    `json:"weight"`
}

type ServerAPISearch struct {
	ID int64 `json:"id"`
}
