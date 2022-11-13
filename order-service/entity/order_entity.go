package entity

type Order struct {
	ID        string `json:"id"`
	Price     int64  `json:"price"`
	ProductID string `json:"product_id"`
	UserID    string `json:"user_id"`
	Quantity  int64  `json:"quantity"`
}
