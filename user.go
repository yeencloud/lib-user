package libuser

import "time"

type User struct {
	ID        string
	CreatedAt time.Time

	Username string
	Email    string
}
