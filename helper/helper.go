package helper

import (
	"guestbook_backend/config"
	"guestbook_backend/helper/response"
	"guestbook_backend/helper/utils"
)

type Helper struct {
	Response response.Response
	Utils    utils.Utils
}

func NewHelper(b *config.NatsBroker) *Helper {

	return &Helper{
		Response: *response.NewResponse(),
		Utils:    *utils.NewUtils(b),
	}

}
