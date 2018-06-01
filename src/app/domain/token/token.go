package token

import "time"

type Token struct {
	Id          string    `json:"id"`
	Cant        int       `json:"cant"`
	CreateToken time.Time `json:"create_token"`
}
