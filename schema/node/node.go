package node


import(
	"fmt"
)

type Variable struct{
	value interface{}
	flag uint8
}
func NewVariable(v interface{},f uint8)(variable *Variable){
	return &Variable{value:v,flag:f}
}


type Node struct{
	NodeValue *Variable
	next *Node
}

func NewNode(nodeValue *Variable)(node *Node){
	return &Node{NodeValue:nodeValue,next:nil}
}

func (p *Node)GetLength()(int){
	length := 1
	pointer := p 
	for(pointer.next != nil){
		length += 1
		pointer = pointer.next
	}
	return length
}

func (p *Node)String()(string){
	return sprintf(p)
}
func sprintf(pointer *Node)string{
	if(pointer != nil){
		return fmt.Sprintf("%v -> %s",pointer.NodeValue,sprintf(pointer.next))
	}else{
		return "nil"
	}
}

func (p *Node)Concat(nextNode *Node){
	pointer := p
	for(pointer.next != nil){
		pointer = pointer.next
	}
	pointer.next = nextNode
}
