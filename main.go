package main

import(
	"fmt"
	"./schema/node"
)


func NewReader(expr string)(func()(byte)){
	rest := expr
	return func()(byte){
		if(len(rest) > 0){
		result := rest[0]
		rest = rest[1:]
		fmt.Println(result)
		fmt.Println(len(rest))
			return result
		}else{
			return byte(0)
		}
	}
}

/*
init expr as list

( -> init first Node


*/
func genListWithReader(reader func()(byte))(*node.Node){
	first := skipSpace(reader);
	if(uint8(first) == 40){
		end,pointer := isEnd(rea)
	}
}

func isSpace(char byte)(bool){
	if(uint8(char) == 32 || uint8(char) == 10){
		return true
	}else{
		return false
	}
}

func isEnd(reader func()(byte))(end bool,pointer byte){
	pointer = reader()
	end = false
	if(isSpace(pointer) || uint8(pointer) == 41){
		end = true
	}
	return
}

func skipSpace(reader func()(byte))(byte){
	b := reader()
	for(isSpace(b)){
		b = reader()
	}
	return b
}


func main(){
	expr := "()"
	nreader := NewReader(expr)
	fmt.Println(string(nreader()))
	nreader()
	nreader()
	nreader()
	nreader()
	fmt.Println(string(byte(0)))
}
