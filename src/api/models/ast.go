package models

import "fmt"

const COMPLETE = "COMPLETE"
const INPROCESS = "INPROCESS"

type Name struct {
	Id string
}

func (n Name) Parse() string {
	return fmt.Sprintf("Name(id='%s', ctx=Store())", n.Id)
}

type Constant struct {
	Value string
}

func (c Constant) Parse() string {

	return fmt.Sprintf("Constant(value=%s))", c.Value)
}

type Assign struct {
	Name  Name
	Value Constant
}

func (a Assign) Parse() string {
	return fmt.Sprintf("Assign(targets=[%s], value=%s", a.Name.Parse(), a.Value.Parse())
}

type BinOp struct {
	Left  string
	Right string
	Op    string
}

func (b BinOp) Parse() string {
	return fmt.Sprintf("BinOp(left=%s, op=%s(), right=%s)", b.Left, b.Op, b.Right)
}
func (b BinOp) IsComplete() bool {
	isComplete := false

	if b.Left != "" && b.Right != "" {
		isComplete = true
	}
	return isComplete
}

type IfElse struct {
	LeftCompare string
	Ops         string
	Comparators string
	Body        string //true
	OrElse      string // false

}

func (i IfElse) Parse() string {
	test := fmt.Sprintf(
		"Compare(left=%s, ops=[%s], comparators=[%s])", i.LeftCompare, i.Ops, i.Comparators)
	body := fmt.Sprintf("[%s]", i.Body)
	orElse := fmt.Sprintf("[%s]", i.OrElse)

	return fmt.Sprintf("If(test=%s, body=%s, orelse=%s, keywords=[]))])", test, body, orElse)
}
func (i IfElse) IsComplete() bool {
	isComplete := false

	if i.LeftCompare != "" && i.Ops != "" && i.Comparators != "" && i.Body != "" {
		isComplete = true
	}
	return isComplete
}

type PrintP struct {
	Args string
}

func (p PrintP) Parse() string {
	return fmt.Sprintf("Call(func=Name(id='print', ctx=Load()), args=[%s], keywords=[])", p.Args)
}
func (p PrintP) ParseExpr() string {
	return fmt.Sprintf("Expr(value=%s)", p.Parse())
}

type ForIn struct {
	Front  string
	To     string
	Target string
	Iter   string
	Body   string
	Orelse string
}

func (f ForIn) Parse() string {
	target := "Name(id='x', ctx=Store())"
	args := fmt.Sprintf("%s, %s", f.Front, f.To)
	iter := fmt.Sprintf(
		"Call(func=Name(id='range', ctx=Load()), args=[%s], keywords=[])", args)

	return fmt.Sprintf("For(target=%s, iter=%s, body=[%s]), Orelse=[%s]", target, iter, f.Body, f.Orelse)
}
func (f ForIn) IsComplete() bool {
	isComplete := false

	if f.Front != "" && f.To != "" && f.Body != "" {
		isComplete = true
	}

	return isComplete
}

type ExecutionNode struct {
	Letf     []int
	Receptor int
}

type AstNode struct {
	Ast         string
	NodeId      int
	Status      string
	AstType     string
	ContextName string
	Code        string
}

type PythonModule struct {
	Module string
}

func (p PythonModule) Parse() string {
	return fmt.Sprintf("Module(body=[%s])", p.Module)
}
