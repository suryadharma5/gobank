package token

import "time"

type Maker interface {
	CreateToken(username string, durartion time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
