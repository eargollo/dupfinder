package cleaner

import "fmt"

type ActionType int64

const (
	Keep ActionType = iota
	Delete
	Rename
)

func (at ActionType) String() string {
	switch at {
	case Keep:
		return "keep"
	case Delete:
		return "delete"
	case Rename:
		return "rename"
	}
	return "unknown"
}

func StringToAction(code string) (ActionType, error) {
	switch code {
	case "D":
		return Delete, nil
	case "d":
		return Delete, nil
	case "K":
		return Keep, nil
	case "k":
		return Keep, nil
	case "":
		return Keep, nil
	default:
		return Keep, fmt.Errorf("invalid action '%s'", code)
	}
}
