package routes

import (
	"net/http"
	"triton/dao"
	"triton/enhancer"

	"github.com/gorilla/mux"
)

/*EntryRoute - роутер для обработки запросов*/
type EntryRoute struct {
	EResponser *enhancer.Responser
	Dao        *dao.ServiceDAO
}

const (
	routeToHistory = "/history"
)

/*listenHistoryCall - роутер для обработки запросов на создания новой истории*/
func (route *EntryRoute) listenHistoryCall(w http.ResponseWriter, r *http.Request) {
	route.EResponser.ResponseWithJSON(w, r, http.StatusOK, map[string]string{
		"status": "ok",
		"code":   "api was starting definition",
	})
}

/*InitRoute - настройка параметров маршрута*/
func (route *EntryRoute) InitRoute(servDao *dao.ServiceDAO) *EntryRoute {
	route.EResponser = &enhancer.Responser{}
	route.Dao = servDao
	return route
}

/*SettingRoute - настройка маршрута для роутера*/
func (route *EntryRoute) SettingRoute(router *mux.Router) *mux.Router {
	router.HandleFunc(routeToHistory, route.listenHistoryCall).Methods("POST")
	return router
}
