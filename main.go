package main

import(
	"fmt"
)

type value interface{}

type variable struct{
	variableValue value
	typeFlag uint8
}

type listNode struct{
	nodeValue *variable
	nextNode *listNode
}

type list struct{
	fristNode *listNode
}

func (p *list)getLength()int{
	len := 0
	if(p.fristNode != nil){
		len += 1
		pointer := p.fristNode
		for(pointer.nextNode != nil){
			len += 1
		}
	}
	return len
}

func (p *list)getLastNode()*listNode{
	pointer := p.fristNode
	for(pointer.nextNode != nil){
		pointer = pointer.nextNode
	}
	return pointer
}

func (p *list)getHeadNode()*listNode{
	return p.fristNode;
}
func (p *listNode)getHeadNode()*listNode{
	return p;
}

type listOrNode interface{
	getHeadNode()*listNode
}
func (p *list)concat(lon listOrNode){
	p.getLastNode().nextNode = lon.getHeadNode();
	
	fmt.Println(p);
}
func (p *list)String()string{
	if(p.getLength() == 0){
		return fmt.Sprintf("nil")
	}else{
		pointer := p.fristNode
		return sprintf(pointer)
	}
}

func sprintf(pointer *listNode)string{
	if(pointer != nil){
		return fmt.Sprintf("%v -> %s",pointer.nodeValue,sprintf(pointer.nextNode))
	}else{
		return "nil"
	}
}



func main(){
	n1 := &listNode{nodeValue:&variable{233,0}}
	l := &list{n1}

	fmt.Println(l.getLastNode().nodeValue)
	l.concat(&listNode{nodeValue:&variable{133,0}});
	fmt.Println(l.getLength());
	fmt.Println(l);
	fmt.Println(l.fristNode.nodeValue.variableValue)
	fmt.Println(233)
}

