package responses

type GetOrder struct {
	ID        int64  `json:"id"`
	ItemId    int64  `json:"item_id"`
	ItemName  string `json:"item_name"`
	UserId    int64  `json:"user_id"`
	UserName  string `json:"user_name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
