package main

import (
	c "api/config"
	"api/route"
	"net/http"
)

func main() {
	conf := c.NewAppConfig()
	http.ListenAndServe(conf.Configurations.Server.Port, route.GetRoute())
	print("Running in port " + conf.Configurations.Server.Port)
}
