package main

import (
	c "api/config"
	"net/http"
)

func main() {
	conf := c.NewAppConfig()
	http.ListenAndServe(conf.Configurations.Server.Port, GetRoute())
	print("Running in port:" + conf.Configurations.Server.Port)
}