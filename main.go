package main

import (
	"BackendTemplate/pkg/config"
	"BackendTemplate/pkg/database"
	"BackendTemplate/pkg/logger"
	"BackendTemplate/pkg/routers"
	"embed"
	"fmt"
	"os"
)

//go:embed dist
var embedFS embed.FS

func main() {
	config.InitConfig()
	database.ConnectDateBase()
	defer database.Engine.Close()
	database.Engine.Update(&database.Clients{Online: "2"})
	database.Engine.Update(&database.Listener{Status: 2})
	database.Engine.Update(&database.Socks5{Status: 2})
	database.Engine.Update(&database.WebDelivery{Status: 2})
	r := routers.NewRouter(embedFS)
	err := r.Run(fmt.Sprintf("0.0.0.0:%d", config.ListenPort))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
