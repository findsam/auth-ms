package util

type Config struct {
	DB_NAME string
	DB_USER string
	DB_PWD  string
}

type SignUpRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20,alphanum"`
	Password string `json:"password" validate:"required,min=6,containsany=!@#$%^&*"`
	Email    string `json:"email" validate:"required,email"`
}

type SignInRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20,alphanum"`
	Password string `json:"password" validate:"required,min=6,containsany=!@#$%^&*"`
}
