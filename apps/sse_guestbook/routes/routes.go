package routes

import (
	"fmt"
	"guestbook_backend/apps/sse_guestbook/controllers"
	"guestbook_backend/helper/utils/hub"
	"net/http"
	"os"
)

type Routes struct {
	controller *controllers.Controllers
}

func NewRoutes(hub *hub.HubSse) *Routes {
	return &Routes{
		controller: controllers.NewControllers(hub),
	}
}

func (s *Routes) Routes(mux *http.ServeMux) {
	prefix := fmt.Sprintf("%s/v1", os.Getenv("PREFIX_API"))

	fmt.Println(prefix)

	mux.Handle(prefix+"/visitor/scan", s.controller.ScanVisitorHandler())

}
