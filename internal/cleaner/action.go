package cleaner

import (
	"fmt"
	"os"
)

type Action struct {
	Group    int
	Type     ActionType
	Filename string
	To       string
}

func (a *Action) Execute() error {
	switch a.Type {
	case Delete:
		return os.Remove(a.Filename)
	default:
		return fmt.Errorf("can not execute: action %s not implemented", a.Type.String())
	}
}
