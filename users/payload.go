package users

type CreateUserPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
