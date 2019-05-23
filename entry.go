package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"triton/configuration"
	"triton/dao"
	"triton/routes"

	"github.com/gorilla/mux"
)

const (
	nameService            = "[HISTORY_CLIEN_SERVICE]: "
	configurationSetupPath = "./settings/init.json"
)

var (
	configurationService *configuration.InitPackage
	entryPoint           *routes.EntryRoute
	serviceDao           *dao.ServiceDAO
	router               *mux.Router
)

func configureHost() {
	configurationService = &configuration.InitPackage{}
	serviceConfig, err := configurationService.Parse(configurationSetupPath)
	if err != nil {
		log.Println(nameService+" error setup init configuration with error code: ", err.Error())
		return
	}

	configurationService = serviceConfig
}

func configureRouter() {
	router = mux.NewRouter()
	entryPoint = &routes.EntryRoute{}
	entryPoint.InitRoute(serviceDao)
	router = entryPoint.SettingRoute(router)
}

func configureDao() {
	serviceDao = &dao.ServiceDAO{}
}

func init() {
	configureHost()
	configureDao()
	configureRouter()
}

func main() {
	if err := http.ListenAndServe(configurationService.ServerAddress+":"+strconv.Itoa(configurationService.ServerPort), router); err != nil {
		log.Fatal(nameService+"api was not started, error code: ", err.Error())
	}
	fmt.Println(nameService + " was started")
}
