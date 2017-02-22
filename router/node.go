package router

type Node struct {
	path map[string]Handle
}

func (n *Node) addRoute(path string, handle Handle) {

	if n.path == nil {
		n.path = make(map[string]Handle)
	}	
	n.path[path] = handle
}

func (n *Node) GetHandle(path string) Handle {
	return n.path[path]
}
