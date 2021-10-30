package controllers

import (
	c "api/config"
	d "api/database"
	ast "api/models"
	utils "api/utilities"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Get(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	config := c.NewAppConfig()

	conn, dbContext := d.NewDBContext(config.Configurations.Database.Endpoint, config.Configurations.Database.Key)
	defer conn.Close()
	txn := dbContext.DgraphClient.NewTxn()

	const q = `query{
		  bladerunner(func: uid(0x7a871617a)) {
			Task.title
		  }
		}`
	resp, err := txn.Query(context.Background(), q)

	if err != nil {
		fmt.Println("Hello, log fatal")
		log.Fatal(err)
	}
	w.Write(resp.Json)
}

func WriteAstScript(ast string) {
	file, err := os.OpenFile("scripts/script.py", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString(ast); err != nil {
		log.Fatal(err)
	}
}

func ParseCode(w http.ResponseWriter, r *http.Request) {
	var jd ast.JData
	err := json.NewDecoder(r.Body).Decode(&jd)
	var nodes []ast.RequestBody

	json.Unmarshal([]byte(jd.Data), &nodes)

	genericNodes := make(map[int]ast.AstNode)
	exeNodes := map[int]*ast.ExecutionNode{}
	nodesCopy := nodes
	status := ""

	for len(nodesCopy) >= 1 {
		status = ""
		for i := 0; i < len(nodesCopy); i++ {
			genericNodes, status = utils.NodeFillManager(nodesCopy[i], genericNodes, exeNodes)
			if status == ast.COMPLETE {
				nodesCopy = append(nodesCopy[:i], nodesCopy[i+1:]...)
				i--
			}
		}
	}

	if len(genericNodes) > 0 {
		print("sii")
	}

	w.WriteHeader(http.StatusOK)
	config := c.NewAppConfig()

	conn, dbContext := d.NewDBContext(config.Configurations.Database.Endpoint, config.Configurations.Database.Key)
	defer conn.Close()
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
