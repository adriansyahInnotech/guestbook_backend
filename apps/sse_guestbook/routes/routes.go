package routes

import (
	"fmt"
	"guestbook_backend/apps/sse_guestbook/controllers"
	natshelper "guestbook_backend/helper/utils/nats_helper"
	"net/http"
	"os"
)

type Routes struct {
	controller *controllers.Controllers
}

func NewRoutes(natsHelper *natshelper.NatsHelper) *Routes {
	return &Routes{
		controller: controllers.NewControllers(natsHelper),
	}
}

func (s *Routes) Routes(mux *http.ServeMux) {
	prefix := fmt.Sprintf("%s/v1", os.Getenv("PREFIX_API"))

	fmt.Println(prefix)

	mux.Handle(prefix+"/visitor/scan", s.controller.ScanVisitorHandler())

}
