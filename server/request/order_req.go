package request

// SetOrder ...
type SetOrder struct {
	Item   int64  `json:"item_id"`
	User   int64  `json:"user_id"`
	Status string `json:"status"`
}
