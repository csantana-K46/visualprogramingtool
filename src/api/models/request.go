package models

type DataStr struct {
	AstType     string `json:"astType"`
	Type        string `json:"type"`
	Op          string `json:"op"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Comparators string `json:"comparators"`
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

type Response struct {
	StatusCode string
	Result     string
}

type Program struct {
	Id          string `json:"uid"`
	Name        string `json:"title"`
	Description string `json:"description"`
	Data        string `json:"drawFlowData"`
}

type DgraphProgramListQuery struct {
	Program []Program `json:"program"`
}
