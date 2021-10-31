package controllers

import (
	c "api/config"
	d "api/database"
	ast "api/models"
	"api/scripts"
	utils "api/utilities"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func ParseCode(w http.ResponseWriter, r *http.Request) {

	var jd ast.JData
	var nodes []ast.RequestBody
	var nodesCopy []ast.RequestBody
	var result string

	err := json.NewDecoder(r.Body).Decode(&jd)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal([]byte(jd.Data), &nodes)
	genericNodes := make(map[int]ast.AstNode)
	exeNodes := map[int]*ast.ExecutionNode{}
	nodesCopy = make([]ast.RequestBody, len(nodes))
	copy(nodesCopy, nodes)
	status := ""

	for len(nodesCopy) >= 1 {
		for i := 0; i < len(nodesCopy); i++ {
			genericNodes, status = utils.NodeFillManager(nodesCopy[i], genericNodes, exeNodes)
			if status == ast.COMPLETE {
				nodesCopy = append(nodesCopy[:i], nodesCopy[i+1:]...)
				i--
			}
		}
	}

	for _, element := range exeNodes {
		var node ast.AstNode
		if val, ok := genericNodes[element.Receptor]; ok {
			node = val
		}

		if node.AstType == "Code" {
			result = node.Code
		} else {

		}
	}

	w.WriteHeader(http.StatusOK)
	response := &ast.Response{StatusCode: "200", Result: result}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	w.Write(b)
}

func RunCode(w http.ResponseWriter, r *http.Request) {

	var jd ast.JData
	var nodes []ast.RequestBody
	var nodesCopy []ast.RequestBody

	err := json.NewDecoder(r.Body).Decode(&jd)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal([]byte(jd.Data), &nodes)
	genericNodes := make(map[int]ast.AstNode)
	exeNodes := map[int]*ast.ExecutionNode{}
	nodesCopy = make([]ast.RequestBody, len(nodes))
	copy(nodesCopy, nodes)

	status := ""

	for len(nodesCopy) >= 1 {
		for i := 0; i < len(nodesCopy); i++ {
			genericNodes, status = utils.NodeFillManager(nodesCopy[i], genericNodes, exeNodes)
			if status == ast.COMPLETE {
				nodesCopy = append(nodesCopy[:i], nodesCopy[i+1:]...)
				i--
			}
		}
	}

	var result string

	for _, element := range exeNodes {
		var node ast.AstNode

		if val, ok := genericNodes[element.Receptor]; ok {
			node = val
		}

		if node.AstType == "Code" {
			result = scripts.ExecuteCode(node.Code)
		} else {

		}
	}

	w.WriteHeader(http.StatusOK)

	response := &ast.Response{StatusCode: "200", Result: result}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	w.Write(b)
}
