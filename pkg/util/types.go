package util

type Config struct {
	DB_NAME    string
	DB_USER    string
	DB_PWD     string
	JWT_SECRET string
	MONGO_URI  string
	STRIPE_PWD string
	STRIPE_PUBLISHABLE_KEY string
}

type SignUpRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20,alphanum"`
	Password string `json:"password" validate:"required,min=6,containsany=!@#$%^&*"`
	Email    string `json:"email" validate:"required,email"`
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}
