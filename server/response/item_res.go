package responses

type GetItem struct {
	ID       int64  `json:"id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
	Status   string `json:"status"`
}
