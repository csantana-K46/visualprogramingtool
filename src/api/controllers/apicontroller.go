package controllers

import (
	c "api/config"
	d "api/database"
	"context"
	"log"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	config := c.NewAppConfig()
	dbContext := d.NewDBContext(config.Configurations.Database.Endpoint, config.Configurations.Database.Key)
	txn := dbContext.DgraphClient.NewTxn()

	const q = `query{
		  bladerunner(func: uid(0x7a871617a)) {
			Task.title
		  }
		}`
	resp, err := txn.Query(context.Background(), q)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(resp.Json)
}