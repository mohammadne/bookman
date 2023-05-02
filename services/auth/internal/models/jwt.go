package models

type Jwt struct {
	AccessToken  *Token
	RefreshToken *Token
}
