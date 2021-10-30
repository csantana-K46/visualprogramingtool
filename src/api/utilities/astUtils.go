package utilities

import (
	ast "api/models"
	"log"
	"strconv"
)

func HasInputs(inputs ast.Inputs) bool {
	return inputs.Input1.Connection != nil && inputs.Input1.Connection[0].Node != "" ||
		inputs.Input2.Connection != nil && inputs.Input2.Connection[0].Node != ""
}

func GetNode(nodes []ast.RequestBody, nodeId int) ast.RequestBody {
	var node ast.RequestBody
	for _, n := range nodes {
		if nodeId == n.Id {
			node = n
			break
		}
	}
	return node
}

func HasOutputs(outputs ast.Outputs) bool {
	return len(outputs.Output1.Connection) > 0 && outputs.Output1.Connection[0].Node != "" ||
		len(outputs.Output2.Connection) > 0 && outputs.Output2.Connection[0].Node != ""
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func FillAssign(node ast.RequestBody) ast.AstNode {
	gNode := ast.AstNode{}
	data := node.Data

	n := ast.Assign{Name: ast.Name{Id: data.Name}, Value: ast.Constant{Value: data.Value}}
	gNode.NodeId = node.Id
	gNode.Ast = n.Parse()
	gNode.Status = ast.COMPLETE
	gNode.AstType = data.AstType

	return gNode
}

func FillBinOp(node ast.RequestBody, genericNodes map[int]ast.AstNode, exeNodes map[int]*ast.ExecutionNode) ast.AstNode {
	gNode := ast.AstNode{}
	data := node.Data

	binOp := ast.BinOp{}
	leftId := 0
	rightId := 0
	nodeIdsToAssume := []int{}

	if val, ok := genericNodes[node.Id]; ok {
		gNode = val
	}

	if len(node.Inputs.Input1.Connection) > 0 {
		intValue, err := strconv.Atoi(node.Inputs.Input1.Connection[0].Node)
		if err != nil {
			log.Println(err)
		}
		leftId = intValue
	}

	if len(node.Inputs.Input2.Connection) > 0 {
		intValue, err := strconv.Atoi(node.Inputs.Input2.Connection[0].Node)
		if err != nil {
			log.Println(err)
		}
		rightId = intValue
	}

	binOp.Op = data.AstType

	if leftId > 0 || rightId > 0 {
		for _, g := range genericNodes { //TODO: use dict access insted of iteration

			if g.NodeId == leftId {
				if g.AstType == "Assign" {
					binOp.Left = ast.Name{Id: "fix"}.Parse()
				} else if g.AstType == "Constant" {
					binOp.Left = ast.Constant{Value: "fix"}.Parse()
				} else if g.AstType == "Add" {
					binOp.Left = g.Ast
					nodeIdsToAssume = append(nodeIdsToAssume, g.NodeId)
				}
			} else if g.NodeId == rightId {
				if g.AstType == "Assign" {
					binOp.Right = ast.Name{Id: "fix"}.Parse()
				} else if g.AstType == "Constant" {
					binOp.Right = ast.Constant{Value: "fix"}.Parse()
				} else if g.AstType == "Add" {
					binOp.Right = g.Ast
					nodeIdsToAssume = append(nodeIdsToAssume, g.NodeId)
				}
			}
		}
		gNode.NodeId = node.Id
		if HasOutputs(node.Outputs) {
			gNode.Ast = binOp.Parse() // TODO: parse expression
		} else {
			gNode.Ast = binOp.Parse()
		}
		gNode.AstType = data.AstType

		if binOp.IsComplete() {
			gNode.Status = ast.COMPLETE

			if len(nodeIdsToAssume) > 0 {
				for _, nToAssume := range nodeIdsToAssume {
					var assumedExecution *ast.ExecutionNode
					var actualExecution *ast.ExecutionNode

					if val, ok := exeNodes[nToAssume]; ok {
						assumedExecution = val
					}

					if val, ok := exeNodes[gNode.NodeId]; ok {
						actualExecution = val
					} else {
						lids := []int{leftId, rightId}
						actualExecution = &ast.ExecutionNode{Letf: lids, Receptor: gNode.NodeId}
						exeNodes[gNode.NodeId] = actualExecution
					}

					if assumedExecution != nil && actualExecution != nil {
						newList := []int{}

						for _, exeN := range assumedExecution.Letf {
							newList = append(newList, exeN)
						}

						for _, exeN := range actualExecution.Letf {
							if exeN != nToAssume {
								newList = append(newList, exeN)
							}
						}
						exeNodes[gNode.NodeId] = &ast.ExecutionNode{Letf: newList, Receptor: gNode.NodeId}
						delete(exeNodes, nToAssume)
					}
				}
			} else {
				lids := []int{leftId, rightId}
				exeNodes[gNode.NodeId] = &ast.ExecutionNode{Letf: lids, Receptor: gNode.NodeId}
			}
		} else {
			gNode.Status = ast.INPROCESS
		}
	}
	return gNode
}

func FillIfElse(node ast.RequestBody, genericNodes map[int]ast.AstNode, exeNodes map[int]*ast.ExecutionNode) ast.AstNode {
	gNode := ast.AstNode{}
	data := node.Data
	inputNodeId := 0
	bodyNodeId := 0
	orElseNodeId := 0
	nodeIdsToAssume := []int{}
	ifElse := ast.IfElse{LeftCompare: "", Ops: "", Comparators: "", Body: "", OrElse: ""}

	if val, ok := genericNodes[node.Id]; ok {
		gNode = val
	}

	if len(node.Inputs.Input1.Connection) > 0 {
		intValue, err := strconv.Atoi(node.Inputs.Input1.Connection[0].Node)
		if err != nil {
			log.Println(err)
		}
		inputNodeId = intValue
	}

	if len(node.Outputs.Output1.Connection) > 0 {
		intValue, err := strconv.Atoi(node.Outputs.Output1.Connection[0].Node)
		if err != nil {
			log.Println(err)
		}
		bodyNodeId = intValue
	}

	if len(node.Outputs.Output2.Connection) > 0 {
		intValue, err := strconv.Atoi(node.Outputs.Output2.Connection[0].Node)
		if err != nil {
			log.Println(err)
		}
		orElseNodeId = intValue
	}

	if bodyNodeId > 0 || orElseNodeId > 0 || inputNodeId > 0 {
		for _, g := range genericNodes { //TODO: use dict access insted of iteration

			if g.NodeId == inputNodeId {
				if g.AstType == "Assign" {
					ifElse.LeftCompare = ast.Name{Id: "fix"}.Parse()
				} else if g.AstType == "Constant" {
					ifElse.LeftCompare = ast.Constant{Value: "fix"}.Parse()
				} else if g.AstType == "Add" {
					ifElse.LeftCompare = g.Ast
					nodeIdsToAssume = append(nodeIdsToAssume, g.NodeId)
				} else if g.AstType == "Print" {
					ifElse.LeftCompare = g.Ast
					nodeIdsToAssume = append(nodeIdsToAssume, g.NodeId)
				}
			} else if g.NodeId == bodyNodeId {
				if g.AstType == "Assign" {
					ifElse.Body = ast.Name{Id: "fix"}.Parse()
				} else if g.AstType == "Constant" {
					ifElse.Body = ast.Constant{Value: "fix"}.Parse()
				} else if g.AstType == "Add" {
					ifElse.Body = g.Ast
					nodeIdsToAssume = append(nodeIdsToAssume, g.NodeId)
				} else if g.AstType == "Print" {
					ifElse.Body = g.Ast
					nodeIdsToAssume = append(nodeIdsToAssume, g.NodeId)
				}
			} else if g.NodeId == orElseNodeId {
				if g.AstType == "Assign" {
					ifElse.OrElse = ast.Name{Id: "fix"}.Parse()
				} else if g.AstType == "Constant" {
					ifElse.OrElse = ast.Constant{Value: "fix"}.Parse()
				} else if g.AstType == "Add" {
					ifElse.OrElse = g.Ast
					nodeIdsToAssume = append(nodeIdsToAssume, g.NodeId)
				} else if g.AstType == "Print" {
					ifElse.Body = g.Ast
					nodeIdsToAssume = append(nodeIdsToAssume, g.NodeId)
				}
			}
		}

		ifElse.Ops = data.Op
		ifElse.Comparators = data.Comparators
		gNode.NodeId = node.Id
		gNode.AstType = data.AstType
		gNode.Ast = ifElse.Parse()

		if HasOutputs(node.Outputs) {
			gNode.Ast = ifElse.Parse() // TODO: parse expression
		} else {
			gNode.Ast = ifElse.Parse()
		}
		gNode.AstType = data.AstType

		if ifElse.IsComplete() {
			gNode.Status = ast.COMPLETE

			if len(nodeIdsToAssume) > 0 {
				for _, nToAssume := range nodeIdsToAssume {
					var assumedExecution *ast.ExecutionNode
					var actualExecution *ast.ExecutionNode

					if val, ok := exeNodes[nToAssume]; ok {
						assumedExecution = val
					}

					if val, ok := exeNodes[gNode.NodeId]; ok {
						actualExecution = val
					} else {
						lids := []int{bodyNodeId, orElseNodeId}
						actualExecution = &ast.ExecutionNode{Letf: lids, Receptor: gNode.NodeId}
						exeNodes[gNode.NodeId] = actualExecution
					}

					if assumedExecution != nil && actualExecution != nil {
						newList := []int{}

						for _, exeN := range assumedExecution.Letf {
							newList = append(newList, exeN)
						}

						for _, exeN := range actualExecution.Letf {
							if exeN != nToAssume {
								newList = append(newList, exeN)
							}
						}
						exeNodes[gNode.NodeId] = &ast.ExecutionNode{Letf: newList, Receptor: gNode.NodeId}
						delete(exeNodes, nToAssume)
					}
				}
			} else {
				lids := []int{bodyNodeId, orElseNodeId}
				exeNodes[gNode.NodeId] = &ast.ExecutionNode{Letf: lids, Receptor: gNode.NodeId}
			}
		} else {
			gNode.Status = ast.INPROCESS
		}
	}
	return gNode
}

func FillPrint(node ast.RequestBody, genericNodes map[int]ast.AstNode, exeNodes map[int]*ast.ExecutionNode) ast.AstNode {
	gNode := ast.AstNode{}
	data := node.Data
	hasArgsNode := false
	leftNodeId := 0
	nodeIdsToAssume := []int{}
	printP := ast.PrintP{Args: ""}

	if val, ok := genericNodes[node.Id]; ok {
		gNode = val
	}

	if len(node.Inputs.Input1.Connection) > 0 {
		intValue, err := strconv.Atoi(node.Inputs.Input1.Connection[0].Node)
		if err != nil {
			log.Println(err)
		}
		leftNodeId = intValue
	}

	for _, g := range genericNodes { //TODO: use dict access insted of iteration
		if g.NodeId == leftNodeId {
			hasArgsNode = true
			if g.Status == ast.COMPLETE {
				if g.AstType == "Assign" {
					printP.Args = ast.Name{Id: "fix"}.Parse()
				} else if g.AstType == "Add" {
					printP.Args = g.Ast
					nodeIdsToAssume = append(nodeIdsToAssume, g.NodeId)
				} else if g.AstType == "Sub" {
					printP.Args = g.Ast
					nodeIdsToAssume = append(nodeIdsToAssume, g.NodeId)
				}
			}
		}
	}

	gNode.NodeId = node.Id
	gNode.Ast = printP.Parse()
	if !hasArgsNode || printP.Args != "" {
		gNode.Status = ast.COMPLETE
	} else {
		gNode.Status = ast.INPROCESS
	}
	gNode.AstType = data.AstType
	lids := []int{leftNodeId}
	nodeAbsorbProcess(nodeIdsToAssume, exeNodes, gNode, lids)
	return gNode
}

func nodeAbsorbProcess(nodeIdsToAssume []int, exeNodes map[int]*ast.ExecutionNode, gNode ast.AstNode, lids []int) {
	if len(nodeIdsToAssume) > 0 {
		for _, nToAssume := range nodeIdsToAssume {
			var assumedExecution *ast.ExecutionNode
			var actualExecution *ast.ExecutionNode

			if val, ok := exeNodes[nToAssume]; ok {
				assumedExecution = val
			}

			if val, ok := exeNodes[gNode.NodeId]; ok {
				actualExecution = val
			} else {
				actualExecution = &ast.ExecutionNode{Letf: lids, Receptor: gNode.NodeId}
				exeNodes[gNode.NodeId] = actualExecution
			}

			if assumedExecution != nil && actualExecution != nil {
				newList := []int{}

				for _, exeN := range assumedExecution.Letf {
					newList = append(newList, exeN)
				}

				for _, exeN := range actualExecution.Letf {
					if exeN != nToAssume {
						newList = append(newList, exeN)
					}
				}
				exeNodes[gNode.NodeId] = &ast.ExecutionNode{Letf: newList, Receptor: gNode.NodeId}
				delete(exeNodes, nToAssume)
			}
		}
	} else {
		exeNodes[gNode.NodeId] = &ast.ExecutionNode{Letf: lids, Receptor: gNode.NodeId}
	}
}

func NodeFillManager(node ast.RequestBody, genericNodes map[int]ast.AstNode, execNodes map[int]*ast.ExecutionNode) (map[int]ast.AstNode, string) {
	var status string
	var nodeBuilt ast.AstNode
	astType := node.Data.AstType

	if astType == "Assign" {
		genericNodes[node.Id] = FillAssign(node)
		status = ast.COMPLETE
	} else if astType == "Add" {
		nodeBuilt = FillBinOp(node, genericNodes, execNodes)
		if nodeBuilt.Status == ast.COMPLETE || nodeBuilt.Status == ast.INPROCESS {
			genericNodes[node.Id] = nodeBuilt
			status = nodeBuilt.Status
		}
	} else if astType == "Sub" {
		nodeBuilt = FillBinOp(node, genericNodes, execNodes)
		if nodeBuilt.Status == ast.COMPLETE || nodeBuilt.Status == ast.INPROCESS {
			genericNodes[node.Id] = nodeBuilt
			status = nodeBuilt.Status
		}
	} else if astType == "Mult" {
		nodeBuilt = FillBinOp(node, genericNodes, execNodes)
		if nodeBuilt.Status == ast.COMPLETE || nodeBuilt.Status == ast.INPROCESS {
			genericNodes[node.Id] = nodeBuilt
			status = nodeBuilt.Status
		}
	} else if astType == "Sub" {
		nodeBuilt = FillBinOp(node, genericNodes, execNodes)
		if nodeBuilt.Status == ast.COMPLETE || nodeBuilt.Status == ast.INPROCESS {
			status = nodeBuilt.Status
			genericNodes[node.Id] = nodeBuilt
		}
	} else if astType == "IfElse" {
		nodeBuilt = FillIfElse(node, genericNodes, execNodes)
		if nodeBuilt.Status == ast.COMPLETE || nodeBuilt.Status == ast.INPROCESS {
			status = nodeBuilt.Status
			genericNodes[node.Id] = nodeBuilt
		}
	} else if astType == "Print" {
		nodeBuilt = FillPrint(node, genericNodes, execNodes)
		if nodeBuilt.Status == ast.COMPLETE || nodeBuilt.Status == ast.INPROCESS {
			status = nodeBuilt.Status
			genericNodes[node.Id] = nodeBuilt
		}
	}
	return genericNodes, status
}
