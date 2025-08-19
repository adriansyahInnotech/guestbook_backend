package helper

import (
	"guestbook/helper/response"
	"guestbook/helper/utils"
)

type Helper struct {
	Response response.Response
	Utils    utils.Utils
}

func NewHelper() *Helper {

	return &Helper{
		Response: *response.NewResponse(),
		Utils:    *utils.NewUtils(),
	}

}
