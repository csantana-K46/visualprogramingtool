package database

import (
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"log"
)

type dgraphContext struct {
	 DgraphClient *dgo.Dgraph
}

func NewDBContext(endpoint string, key string) dgraphContext {
	var dgraph dgraphContext

	if e.DgraphClient == nil{
		conn, err := dgo.DialCloud(endpoint, key)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		dgraph = dgraphContext{DgraphClient: dgo.NewDgraphClient(api.NewDgraphClient(conn))}
	}
	return dgraph
}