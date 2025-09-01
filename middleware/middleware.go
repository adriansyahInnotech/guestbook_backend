package middleware

import (
	"guestbook_backend/helper"
	"guestbook_backend/middleware/apikey"
)

type Middleware struct {
	// JWT       *jwt.Jwt
	// Websocket *websocket.Websocket
	// Casbin    *casbin.CasbinHandler
	ApiKey *apikey.ApiKey
}

func NewMiddlware(helper *helper.Helper) *Middleware {
	return &Middleware{
		// JWT:       jwt.NewJwt(),
		// Websocket: websocket.NewWebsocket(),
		ApiKey: apikey.NewApiKey(helper),
		// Casbin:    casbin.NewCasbinHandler(),
	}

}
