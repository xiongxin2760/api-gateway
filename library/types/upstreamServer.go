package types

// api的地址
type ServerAPI struct {
	URL    string `json:"url"`
	Weight int    `json:"weight"`
}

type ServerAPISearch struct {
	ID string `json:"id"`
}
