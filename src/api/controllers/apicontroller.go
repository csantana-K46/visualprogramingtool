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
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func Response(resp string) []byte {
	response := &ast.Response{StatusCode: "200", Result: resp}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return []byte{}
	}
	return b
}

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

	isCode := false
	generatedCode := ""

	for _, element := range exeNodes {
		var node ast.AstNode
		if val, ok := genericNodes[element.Receptor]; ok {
			node = val
		}

		if node.AstType == "Code" {
			result = node.Code
			isCode = true
		} else {
			result = utils.JoinModule(node, genericNodes, *element, result)
		}
	}
	if !isCode {
		result = utils.BuildPythonModule(result)
		generatedCode = scripts.EvalCode(result)
	} else {
		generatedCode = result
	}

	w.WriteHeader(http.StatusOK)
	w.Write(Response(generatedCode))
}

func RunCode(w http.ResponseWriter, r *http.Request) {
	var data ast.JData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}

	result := scripts.ExecuteCode(data.Data)
	w.WriteHeader(http.StatusOK)
	w.Write(Response(result))
}

func Add(w http.ResponseWriter, r *http.Request) {
	var jd ast.Program

	err := json.NewDecoder(r.Body).Decode(&jd)
	if err != nil {
		log.Println(err)
	}

	program := d.NewProgram(jd.Name, jd.Description, jd.Data)
	jSchema, err := json.Marshal(program)
	if err != nil {
		fmt.Println(err)
		return
	}
	config := c.NewAppConfig()
	conn, dbContext := d.NewDBContext(config.Configurations.Database.Endpoint, config.Configurations.Database.Key)
	defer conn.Close()
	txn := dbContext.DgraphClient.NewTxn()

	mu := &api.Mutation{
		CommitNow: true,
	}
	mu.SetJson = jSchema

	ctx := context.Background()
	resp, err := txn.Mutate(ctx, mu)
	if err != nil {
		fmt.Println("Add error")
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(Response(resp.String()))
}

func List(w http.ResponseWriter, r *http.Request) {
	//var programs ast.DgraphProgramListQuery
	const query = `query{
		programs(func: has(drawFlowData)){
		uid
		title
		description
		drawFlowData
	  }
	}`

	config := c.NewAppConfig()
	conn, dbContext := d.NewDBContext(config.Configurations.Database.Endpoint, config.Configurations.Database.Key)
	defer conn.Close()
	txn := dbContext.DgraphClient.NewTxn()

	resp, err := txn.Query(context.Background(), query)
	if err != nil {
		fmt.Println("List error")
		log.Fatal(err)
	}
	w.Write(Response(string(resp.Json)))
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	variables := make(map[string]string)
	variables["$id"] = chi.URLParam(r, "id")
	const query = `query ProgramByUid($id: string){
    	programByUid(func: uid($id)){
        	uid
    		title
    		description
    		drawFlowData
    }
}`
	config := c.NewAppConfig()
	conn, dbContext := d.NewDBContext(config.Configurations.Database.Endpoint, config.Configurations.Database.Key)
	defer conn.Close()
	txn := dbContext.DgraphClient.NewTxn()

	resp, err := txn.QueryWithVars(context.Background(), query, variables)
	if err != nil {
		fmt.Println("List error")
		log.Fatal(err)
	}
	w.Write(Response(string(resp.Json)))

	/*puid := resp.Uids["alice"]
		const q = `query Me($id: string){
			me(func: uid($id)){
				title
				description
				drawFlowData
			}
	} `
		variables := make(map[string]string)
		variables["$id"] = puid
		resp2, err2 := txn.QueryWithVars(ctx, q, variables)
		if err2 != nil {
			log.Fatal(err2)
		}
		type Root struct {
			Me d.Program `json:"me"`
		}

		var root  Root
		err = json.Unmarshal(resp2.Json, &root)
		if err != nil {
			log.Fatal(err)
		}
		out, _ := json.MarshalIndent(root, "", "\t")
		fmt.Printf("%s\n", out)*/
}
