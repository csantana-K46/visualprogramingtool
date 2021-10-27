package controllers

import (
	c "api/config"
	d "api/database"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

type DataStr struct {
	AstType string `json:"astType"`
	Type    string `json:"type"`
	Op      string `json:"op"`
	Name    string `json:"name"`
	Value   string `json:"value"`
}

type Connection struct {
	Node   string `json:"node"`
	Output string `json:"output"`
	Input  string `json:"input"`
}

type InputOutput struct {
	Connection []Connection `json:"connections"`
}

type Outputs struct {
	Output1 InputOutput `json:"output_1"`
	Output2 InputOutput `json:"output_2"`
	Output3 InputOutput `json:"output_3"`
}

type Inputs struct {
	Input1 InputOutput `json:"input_1"`
	Input2 InputOutput `json:"input_2"`
	Input3 InputOutput `json:"input_3"`
}

type Arg struct {
	Node RequestBody
	Ast  string
}

type RequestBody struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Data      DataStr
	Outputs   Outputs `json:"outputs"`
	Inputs    Inputs  `json:"inputs"`
	Ast       string
	Arg       []Arg
	BodyBuilt string
}

type JData struct {
	Data string
}

func GetNode(nodes []RequestBody, nodeId int) RequestBody {
	var node RequestBody
	for _, n := range nodes {
		if nodeId == n.Id {
			node = n
			break
		}
	}
	return node
}

func HasInputs(inputs Inputs) bool {
	return inputs.Input1.Connection != nil && inputs.Input1.Connection[0].Node != "" ||
		inputs.Input2.Connection != nil && inputs.Input2.Connection[0].Node != ""
}

func HasOutputs(outputs Outputs) bool {
	return outputs.Output1.Connection != nil && outputs.Output1.Connection[0].Node != "" ||
		outputs.Output2.Connection != nil && outputs.Output2.Connection[0].Node != ""
}

func NodeClassification(nodes []RequestBody) ([]RequestBody, []RequestBody, []RequestBody) {
	var outputs []RequestBody
	var inputsOutputs []RequestBody
	var inputs []RequestBody

	for _, n := range nodes {
		hasInputs := HasInputs(n.Inputs)
		hasOutputs := HasOutputs(n.Outputs)

		if !hasInputs && hasOutputs {
			outputs = append(outputs, n)
		} else if hasInputs && hasOutputs {
			inputsOutputs = append(inputsOutputs, n)
		} else if hasInputs && !hasOutputs {
			inputs = append(inputs, n)
		}
	}
	return outputs, inputsOutputs, inputs
}

func BuildAst(node RequestBody) string {
	var result string
	astType := node.Data.AstType

	if astType == "Assign" {
		result = Assign(node.Data.Name, node.Data.Value)
	} else if astType == "Add" {
		a := node.Arg[0].Node.Data.Name
		b := node.Arg[1].Node.Data.Name
		result = Add(a, b)

		hasOutput := HasOutputs(node.Outputs)
		if !hasOutput {
			result = fmt.Sprintf("Expr(value=%s)", result)
		}
	}
	return result
}

func BuildNodes(outputs []RequestBody, inputsOutputs []RequestBody, inputs []RequestBody) {
	for index, n := range inputsOutputs {

		for _, o := range outputs {
			if n.Inputs.Input1.Connection != nil {
				intVar, err := strconv.Atoi(n.Inputs.Input1.Connection[0].Node)
				if err != nil {
					log.Println(err)
				}
				if intVar == o.Id {
					n.Arg = append(n.Arg, Arg{Node: o, Ast: BuildAst(o)})
				}
			}
			if n.Inputs.Input2.Connection != nil {
				intVar, err := strconv.Atoi(n.Inputs.Input2.Connection[0].Node)
				if err != nil {
					log.Println(err)
				}
				if intVar == o.Id {
					n.Arg = append(n.Arg, Arg{Node: o, Ast: BuildAst(o)})
				}
			}
		}

		if len(n.Arg) > 0 {
			n.Ast = BuildAst(n)

		}
		inputsOutputs[index] = n
	}

	for _, n := range inputsOutputs {

		for _, i := range inputs {
			if n.Outputs.Output1.Connection != nil {
				intVar, err := strconv.Atoi(n.Outputs.Output1.Connection[0].Node)
				if err != nil {
					log.Println(err)
				}
				if intVar == i.Id {
					//n.Arg = append(n.Arg, Arg{Node: i, Ast: BuildAst(i)})
					fmt.Sprintf("%s-%s", n.Ast, i.Data.AstType)
				}
			}
			if n.Outputs.Output2.Connection != nil {
				intVar, err := strconv.Atoi(n.Outputs.Output2.Connection[0].Node)
				if err != nil {
					log.Println(err)
				}
				if intVar == i.Id {
					fmt.Sprintf("%s-%s", n.Ast, i.Data.AstType)
				}
			}
		}

		if len(n.Arg) > 0 {
			n.Ast = BuildAst(n)

		}
	}
}

func Assign(name string, value string) string {
	return fmt.Sprintf("Assign(targets=[Name(id='%s', ctx=Store())], value=Constant(value=%s))", name, value)
}

func Add(a string, b string) string {
	return fmt.Sprintf("BinOp(left=Name(id='%s', ctx=Load()), op=Add(), right=Name(id='%s', ctx=Load()))", a, b)
}

func Print() string {
	return "Expr(value=Call(func=Name(id='print', ctx=Load()), args=[], keywords=[]))"
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
	var jd JData
	var ast string
	err := json.NewDecoder(r.Body).Decode(&jd)
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	var nodes []RequestBody
	json.Unmarshal([]byte(jd.Data), &nodes)

	outputs, inputsOutputs, inputs := NodeClassification(nodes)
	BuildNodes(outputs, inputsOutputs, inputs)
	fmt.Println(ast)

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
