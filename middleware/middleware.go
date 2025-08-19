package middleware

import (
	"guestbook/middleware/casbin"
	"guestbook/middleware/jwt"
	"guestbook/middleware/websocket"
)

type Middleware struct {
	JWT       *jwt.Jwt
	Websocket *websocket.Websocket
	Casbin    *casbin.CasbinHandler
}

func NewMiddlware() *Middleware {
	return &Middleware{
		JWT:       jwt.NewJwt(),
		Websocket: websocket.NewWebsocket(),
		Casbin:    casbin.NewCasbinHandler(),
	}

}
