package request

// SetIte, ...
type SetItem struct {
	Category int    `json:"category"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
	Status   string `json:"status"`
}
