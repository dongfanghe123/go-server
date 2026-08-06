package entity

type RegisterRequest struct {
	Username *string `json:"username"`
	Password *string
	Email    *string
	Phone    *string
}
