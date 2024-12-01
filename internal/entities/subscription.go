package entities

type Subscription struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	PaymentsCount string  `json:"paymentsCount"`
	IsActive      bool    `json:"isActive"`
}
