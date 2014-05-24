package cntr

type node struct {
	item	interface{}
	next	*node
}

func newNode(item interface{}) *node {
	if item == nil {
		panic("nil item")
	}
	return &node{item: item}
}