package node_types

type NodeType int

const (
	Worker NodeType = iota
	Manager
)

func FromString(arg string) NodeType {
	switch arg {
	case "worker":
		return Worker
	case "manager":
		return Manager
	default:
		panic("Sorry, you passed a wrong type for NodeType")
	}
}
