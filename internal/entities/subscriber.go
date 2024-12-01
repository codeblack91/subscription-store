package entities

type Subscriber struct {
	ID            string         `json:"id"`
	Email         string         `json:"email"`
	Name          string         `json:"name"`
	Subscriptions []Subscription `json:"subscriptions"`
}
