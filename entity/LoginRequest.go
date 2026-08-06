package entity

type LoginRequest struct {
	Username *string
	Password *string
}

type LoginRequestV2 struct {
	Phone *string
	Code  *string
}
