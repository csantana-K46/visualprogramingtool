package database

import "fmt"

type Program struct {
	Id           string `json:"uid,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	DrawFlowData string `json:"drawFlowData,omitempty"`
}

func (p Program) str() string {
	return fmt.Sprintf("Id: %s, Title: %s", p.Id, p.Title)
}
