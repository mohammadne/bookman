package models

type Token struct {
	Token   string
	UUID    string
	Expires int64
}

type TokenDetails struct {
	AccessToken  *Token
	RefreshToken *Token
}
