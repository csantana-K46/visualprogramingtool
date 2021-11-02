package database

import (
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
	"log"
)

type dgraphContext struct {
	DgraphClient *dgo.Dgraph
}

func NewDBContext(endpoint string, key string) (*grpc.ClientConn, dgraphContext) {
	conn, err := dgo.DialCloud(endpoint, key)
	if err != nil {
		log.Fatal(err)
	}
	return conn, dgraphContext{DgraphClient: dgo.NewDgraphClient(api.NewDgraphClient(conn))}
}

func NewProgram(title string, desc string, data string) Program {
	return Program{Id: "_:alice", Title: title, Description: desc, DrawFlowData: data}
}
