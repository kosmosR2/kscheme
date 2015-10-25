package main

import(
	"fmt"
	_"./schema/node"
)


func reader(expr string)(func()(byte)){
	rest := expr
	return func()(byte){
		fmt.Println(rest[0])
		return rest[0]
	}
}

func main(){
	expr := "(head (body (p 233)) 233)"
	nreader := reader(expr)
	nreader()
	fmt.Println(233)
}
