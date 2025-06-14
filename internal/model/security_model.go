package model

type Security struct {
	EmailVerified bool `bson:"email_verified" json:"email_verified"`
}

func NewSecurity() *Security {
	return &Security{
		EmailVerified: false,
	}
}
